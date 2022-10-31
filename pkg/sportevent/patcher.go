package sportevent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/market"
	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

type Patcher interface {
	ApplyPatches(rawPatches json.RawMessage) error
}

type PatcherSimdJSON struct {
	sportEvent *SportEvent
	parsedJson *simdjson.ParsedJson
	rootObj    *simdjson.Object
	rootIter   *simdjson.Iter
	tmpIter    *simdjson.Iter
	tmpObj     *simdjson.Object
	//
	competitorObj *simdjson.Object
	competitor    *fixture.Competitor
	scoresObj     *simdjson.Object
	scoreObj      *simdjson.Object
	score         *fixture.Score
	//
	marketObj *simdjson.Object
	oddsObj   *simdjson.Object
	oddObj    *simdjson.Object

	tmpOdd *market.Odd
}

func NewPatcherSimdJSON(sportEvent *SportEvent) *PatcherSimdJSON {
	return &PatcherSimdJSON{
		sportEvent: sportEvent,
		parsedJson: new(simdjson.ParsedJson),
		rootObj:    new(simdjson.Object),
		rootIter:   new(simdjson.Iter),
		tmpIter:    new(simdjson.Iter),
		tmpObj:     new(simdjson.Object),
		marketObj:  new(simdjson.Object),
		oddsObj:    new(simdjson.Object),
		oddObj:     new(simdjson.Object),
		tmpOdd:     new(market.Odd),
	}
}

func (p *PatcherSimdJSON) ApplyPatches(rawPatches json.RawMessage) error {
	parsedJson, err := simdjson.Parse(rawPatches, p.parsedJson, simdjson.WithCopyStrings(true))
	if err != nil {
		return fmt.Errorf("simdjson parse: %w", err)
	}

	rootIter, err := simdutil.CreateRootIter(parsedJson)
	if err != nil {
		return fmt.Errorf("json to root iter: %w", err)
	}

	rootObj, err := rootIter.Object(p.rootObj)
	if err != nil {
		return err
	}

	for {
		path, elementType, err := rootObj.NextElement(p.rootIter)
		if err != nil {
			return fmt.Errorf("next element: %w", err)
		}

		if elementType == simdjson.TypeNone {
			// Done
			break
		}

		key, rest, partialPatch := strings.Cut(path, "/")

		switch key {
		case "bet_stop":
			value, err := p.rootIter.Bool()
			if err != nil {
				return fmt.Errorf("parse bet_stop: %w", err)
			}

			p.sportEvent.BetStop = value

		case "updated_at":
			p.sportEvent.UpdatedAt, err = simdutil.TimeFromIter(p.rootIter)
			if err != nil {
				return fmt.Errorf("parse updated_at: %w", err)
			}

		case "fixture":
			if partialPatch {
				err = p.applyFixturePatch(rest, p.rootIter)
				if err != nil {
					return err
				}

				continue
			}

			obj, err := p.rootIter.Object(p.tmpObj)
			if err != nil {
				return err
			}

			err = p.sportEvent.Fixture.UnmarshalSimdJSON(
				obj,
				p.tmpIter,
				p.tmpObj,
				p.competitorObj,
				p.competitor,
				p.scoresObj,
				p.scoreObj,
				p.score,
			)
			if err != nil {
				return fmt.Errorf("unmarshal fixture: %w", err)
			}

		case "markets":
			if partialPatch {
				err = p.applyMarketsPatch(rest, p.rootIter)
				if err != nil {
					return err
				}

				continue
			}

			p.sportEvent.Markets = make(market.Markets, 128)

			bb, err := p.rootIter.MarshalJSON()
			if err != nil {
				return fmt.Errorf("marshal markets: %w", err)
			}

			err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(bb), &p.sportEvent.Markets)
			if err != nil {
				return fmt.Errorf("sonic unmarshal markets: %w", err)
			}
		}
	}

	return nil
}

func (p *PatcherSimdJSON) applyFixturePatch(path string, iter *simdjson.Iter) error {
	var (
		err                     error
		key, rest, partialPatch = strings.Cut(path, "/")
	)

	switch key {
	case "status":
		p.sportEvent.Fixture.Status, err = simdutil.IntFromIter(iter)
	case "type":
		p.sportEvent.Fixture.Type, err = simdutil.IntFromIter(iter)
	case "start_time":
		p.sportEvent.Fixture.StartTime, err = simdutil.TimeFromIter(iter)
	case "live_coverage":
		p.sportEvent.Fixture.LiveCoverage, err = iter.Bool()
	case "competitors":
		if partialPatch {
			if p.sportEvent.Fixture.Competitors == nil {
				return fmt.Errorf("patch nil competitors")
			}

			return p.applyCompetitorsPatch(rest, iter)
		}

		obj, err := iter.Object(p.tmpObj)
		if err != nil {
			return fmt.Errorf("create competitors obj: %w", err)
		}

		p.sportEvent.Fixture.Competitors = make(fixture.Competitors, 4)

		return p.sportEvent.Fixture.Competitors.UnmarshalSimdJSON(
			obj,
			p.tmpIter,
			p.competitorObj,
			p.competitor,
			p.scoresObj,
			p.scoreObj,
			p.score,
		)

	case "tournament":
		if partialPatch {
			return p.applyTournamentPatch(rest, iter)
		}

		tournamentObj, err := iter.Object(p.tmpObj)
		if err != nil {
			return fmt.Errorf("create tournament obj: %w", err)
		}

		return p.sportEvent.Fixture.Tournament.UnmarshalSimdJSON(tournamentObj, p.tmpIter)

	case "streams":
		if partialPatch {
			if p.sportEvent.Fixture.Streams == nil {
				return fmt.Errorf("patch nil streams")
			}

			return p.applyStreamsPatch(rest, iter)
		}

		p.sportEvent.Fixture.Streams = make(fixture.Streams)
		return p.sportEvent.Fixture.Streams.FromIter(iter, p.tmpObj)

	case "venue":
		if partialPatch {
			if rest != "id" {
				return fmt.Errorf("invalid venue patch: %s", rest)
			}

			p.sportEvent.Fixture.Venue.ID, err = simdutil.UnsafeStrFromIter(iter)
			return err
		}

		return p.sportEvent.Fixture.Venue.FromIter(iter, p.tmpObj)
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
	}

	return nil
}

func (p *PatcherSimdJSON) applyCompetitorsPatch(path string, iter *simdjson.Iter) error {
	key, rest, partialPatch := strings.Cut(path, "/")
	competitor, ok := p.sportEvent.Fixture.Competitors[key]

	if !partialPatch {
		obj, err := iter.Object(p.tmpObj)
		if err != nil {
			return err
		}

		err = competitor.UnmarshalSimdJSON(obj, p.tmpIter, p.scoresObj, p.scoreObj, p.score)
		if err != nil {
			return fmt.Errorf("competitor %q unmarshal simdjson: %w", key, err)
		}

		p.sportEvent.Fixture.Competitors[key] = competitor
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent competitor: %q", key)
	}

	err := p.applyCompetitorPatch(&competitor, rest, iter)
	if err != nil {
		return fmt.Errorf("apply competitor %q patch: %w", path, err)
	}

	p.sportEvent.Fixture.Competitors[key] = competitor
	return nil
}

func (p *PatcherSimdJSON) applyCompetitorPatch(c *fixture.Competitor, path string, iter *simdjson.Iter) error {
	var (
		err                     error
		key, rest, partialPatch = strings.Cut(path, "/")
	)

	switch key {
	case "id":
		c.ID, err = simdutil.UnsafeStrFromIter(iter)
	case "type":
		c.Type, err = simdutil.IntFromIter(iter)
	case "home_away":
		c.HomeAway, err = simdutil.IntFromIter(iter)
	case "template_position":
		c.TemplatePosition, err = simdutil.IntFromIter(iter)
	case "name":
		c.Name, err = simdutil.UnsafeStrFromIter(iter)
	case "master_id":
		c.MasterID, err = simdutil.UnsafeStrFromIter(iter)
	case "country_code":
		c.CountryCode, err = simdutil.UnsafeStrFromIter(iter)
	case "scores":
		if partialPatch {
			if c.Scores == nil {
				return fmt.Errorf("partial patch nil scores")
			}

			return p.applyScoresPatch(c.Scores, rest, iter)
		}

		obj, err := iter.Object(p.tmpObj)
		if err != nil {
			return err
		}

		c.Scores = make(fixture.Scores, 4)
		return c.Scores.UnmarshalSimdJSON(obj, p.tmpIter, p.scoreObj, p.score)

	case "score":
		if c.Scores == nil {
			return fmt.Errorf("partial patch nil scores")
		}

		return p.applyScoresPatch(c.Scores, rest, iter)
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
	}

	return nil
}

func (p *PatcherSimdJSON) applyScoresPatch(s fixture.Scores, path string, iter *simdjson.Iter) error {
	key, rest, partialPatch := strings.Cut(path, "/")
	score, ok := s[key]

	if !partialPatch {
		obj, err := iter.Object(p.tmpObj)
		if err != nil {
			return err
		}

		err = score.UnmarshalSimdJSON(obj, p.tmpIter)
		if err != nil {
			return fmt.Errorf("score %q unmarshal simdjson: %w", key, err)
		}

		s[key] = score
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent score: %q", key)
	}

	err := p.applyScorePatch(&score, rest, iter)
	if err != nil {
		return fmt.Errorf("apply score %q patch: %w", path, err)
	}

	s[key] = score
	return nil
}

func (p *PatcherSimdJSON) applyScorePatch(s *fixture.Score, path string, iter *simdjson.Iter) error {
	var err error

	switch path {
	case "id":
		s.ID, err = simdutil.UnsafeStrFromIter(iter)
	case "type":
		s.Type, err = simdutil.UnsafeStrFromIter(iter)
	case "points":
		s.Points, err = simdutil.UnsafeStrFromIter(iter)
	case "number":
		s.Number, err = simdutil.IntFromIter(iter)
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", path, err)
	}

	return nil
}

func (p *PatcherSimdJSON) applyTournamentPatch(path string, iter *simdjson.Iter) error {
	var err error

	switch path {
	case "id":
		p.sportEvent.Fixture.Tournament.ID, err = simdutil.UnsafeStrFromIter(iter)
	case "name":
		p.sportEvent.Fixture.Tournament.Name, err = simdutil.UnsafeStrFromIter(iter)
	case "master_id":
		p.sportEvent.Fixture.Tournament.MasterID, err = simdutil.UnsafeStrFromIter(iter)
	case "country_code":
		p.sportEvent.Fixture.Tournament.CountryCode, err = simdutil.UnsafeStrFromIter(iter)
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", path, err)
	}

	return nil
}

func (p *PatcherSimdJSON) applyStreamsPatch(path string, iter *simdjson.Iter) error {
	key, rest, partialPatch := strings.Cut(path, "/")
	stream, ok := p.sportEvent.Fixture.Streams[key]

	if !partialPatch {
		obj, err := iter.Object(p.tmpObj)
		if err != nil {
			return err
		}

		err = stream.UnmarshalSimdJSON(obj, p.tmpIter)
		if err != nil {
			return fmt.Errorf("stream %q unmarshal simdjson: %w", key, err)
		}

		p.sportEvent.Fixture.Streams[key] = stream
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent stream: %q", key)
	}

	err := p.applyStreamPatch(&stream, rest, iter)
	if err != nil {
		return fmt.Errorf("apply stream patch: %w", err)
	}

	p.sportEvent.Fixture.Streams[key] = stream
	return nil
}

func (p *PatcherSimdJSON) applyStreamPatch(s *fixture.Stream, path string, iter *simdjson.Iter) error {
	var (
		err                     error
		key, rest, partialPatch = strings.Cut(path, "/")
	)

	switch path {
	case "id":
		s.ID, err = simdutil.UnsafeStrFromIter(iter)
	case "locale":
		s.Locale, err = simdutil.UnsafeStrFromIter(iter)
	case "url":
		s.URL, err = simdutil.UnsafeStrFromIter(iter)
	case "platforms":
		if partialPatch {
			if s.Platforms == nil {
				return fmt.Errorf("partial patch nil platforms")
			}

			return p.applyPlatformsPatch(s.Platforms, rest, iter)
		}

		platformsObj, err := iter.Object(p.tmpObj)
		if err != nil {
			return fmt.Errorf("create platforms obj: %w", err)
		}

		s.Platforms = make(fixture.Platforms)
		return s.Platforms.UnmarshalSimdJSON(platformsObj, p.tmpIter)
	case "priority":
		s.Priority, err = simdutil.IntFromIter(iter)
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
	}

	return nil
}

func (p *PatcherSimdJSON) applyPlatformsPatch(s fixture.Platforms, path string, iter *simdjson.Iter) error {
	key, rest, partialPatch := strings.Cut(path, "/")
	platform, ok := s[key]

	if !partialPatch {
		obj, err := iter.Object(p.tmpObj)
		if err != nil {
			return err
		}

		err = platform.UnmarshalSimdJSON(obj, p.tmpIter)
		if err != nil {
			return fmt.Errorf("platform %q unmarshal simdjson: %w", key, err)
		}

		s[key] = platform
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent platform: %q", key)
	}

	err := p.applyPlatformPatch(&platform, rest, iter)
	if err != nil {
		return fmt.Errorf("apply platform patch: %w", err)
	}

	s[key] = platform
	return nil
}

func (p *PatcherSimdJSON) applyPlatformPatch(platform *fixture.Platform, path string, iter *simdjson.Iter) error {
	var err error

	switch path {
	case "type":
		platform.Type, err = simdutil.UnsafeStrFromIter(iter)
	case "allowed_countries":
		array, err := iter.Array(nil)
		if err != nil {
			return err
		}

		strArray, err := array.AsString()
		if err != nil {
			return err
		}

		platform.AllowedCountries = strArray

	case "enabled":
		platform.Enabled, err = iter.Bool()
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", path, err)
	}

	return nil
}

func (p *PatcherSimdJSON) applyMarketsPatch(path string, iter *simdjson.Iter) error {
	key, rest, partialPatch := strings.Cut(path, "/")
	m, ok := p.sportEvent.Markets[key]

	if !partialPatch {
		obj, err := iter.Object(p.marketObj)
		if err != nil {
			return err
		}

		err = m.UnmarshalSimdJSON(obj, p.tmpIter, p.oddsObj, p.oddObj, p.tmpOdd)
		if err != nil {
			return fmt.Errorf("market %q unmarshal simdjson: %w", key, err)
		}

		p.sportEvent.Markets[key] = m

		return nil
	}

	if !ok {
		return fmt.Errorf("partial patch non-existent market: %q", key)
	}

	err := p.applyMarketPatch(&m, rest, iter)
	if err != nil {
		return fmt.Errorf("apply market patch: %w", err)
	}

	p.sportEvent.Markets[key] = m
	return nil
}

func (p *PatcherSimdJSON) applyMarketPatch(m *market.Market, path string, iter *simdjson.Iter) error {
	var (
		err                     error
		key, rest, partialPatch = strings.Cut(path, "/")
	)

	switch key {
	case "name":
		m.Template, err = simdutil.UnsafeStrFromIter(iter)
	case "status":
		var value int64

		value, err = iter.Int()
		m.Status = market.Status(value)
	case "type_id":
		m.TypeID, err = simdutil.IntFromIter(iter)
	case "odds":
		if !partialPatch {
			obj, err := iter.Object(p.oddsObj)
			if err != nil {
				return fmt.Errorf("create %q object: %w", key, err)
			}

			m.Odds = make(market.Odds, 4)

			return m.Odds.UnmarshalSimdJSON(obj, p.tmpIter, p.oddObj, p.tmpOdd)
		}

		if m.Odds == nil {
			return fmt.Errorf("patch nil odds")
		}

		return p.applyOddsPatch(m.Odds, rest, iter)
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
	}

	return nil
}

func (p *PatcherSimdJSON) applyOddsPatch(odds market.Odds, path string, iter *simdjson.Iter) error {
	key, rest, partialPatch := strings.Cut(path, "/")
	odd, ok := odds[key]

	if !partialPatch {
		obj, err := iter.Object(p.oddObj)
		if err != nil {
			return err
		}

		err = odd.UnmarshalSimdJSON(obj, p.tmpIter)
		if err != nil {
			return fmt.Errorf("odd %q unmarshal simdjson: %w", key, err)
		}

		odds[key] = odd
		return nil
	}

	if !ok {
		return fmt.Errorf("partial patchnon-existent odd: %q", path)
	}

	err := p.applyOddPatch(&odd, rest, iter)
	if err != nil {
		return fmt.Errorf("apply odd patch: %w", err)
	}

	odds[key] = odd
	return nil
}

func (p *PatcherSimdJSON) applyOddPatch(odd *market.Odd, path string, iter *simdjson.Iter) error {
	var err error

	switch path {
	case "name":
		odd.Template, err = simdutil.UnsafeStrFromIter(iter)
	case "value":
		odd.Value, err = simdutil.UnsafeStrFromIter(iter)
	case "is_active":
		odd.IsActive, err = iter.Bool()
	case "status":
		var value int64

		value, err = iter.Int()
		odd.Status = market.OddStatus(value)
	case "status_reason":
		odd.StatusReason, err = simdutil.UnsafeStrFromIter(iter)
	default:
		return nil
	}

	if err != nil {
		return fmt.Errorf("patch %q: %w", path, err)
	}

	return nil
}
