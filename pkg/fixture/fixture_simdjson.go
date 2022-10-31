package fixture

import (
	"fmt"
	"strings"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

func (f *Fixture) UnmarshalSimdJSON(obj *simdjson.Object) error {
	iter := new(simdjson.Iter)
	tmpObj := new(simdjson.Object)

	for {
		name, elementType, err := obj.NextElementBytes(iter)
		if err != nil {
			return err
		}

		if elementType == simdjson.TypeNone {
			break
		}

		switch string(name) {
		case "id":
			f.ID, err = simdutil.UnsafeStrFromIter(iter)
		case "version":
			f.Version, err = simdutil.IntFromIter(iter)
		case "owner_id":
			f.OwnerID, err = simdutil.UnsafeStrFromIter(iter)
		case "template":
			f.Template, err = simdutil.UnsafeStrFromIter(iter)
		case "status":
			f.Status, err = simdutil.IntFromIter(iter)
		case "type":
			f.Type, err = simdutil.IntFromIter(iter)
		case "sport_id":
			f.SportID, err = simdutil.UnsafeStrFromIter(iter)
		case "tournament":
			err = f.Tournament.FromIter(iter, tmpObj)
		case "venue":
			err = f.Venue.FromIter(iter, tmpObj)
		case "competitors":
			f.Competitors = make(Competitors)
			err = f.Competitors.FromIter(iter, tmpObj)

		case "streams":
			f.Streams = make(Streams)
			err = f.Streams.FromIter(iter, tmpObj)
		case "live_coverage":
			f.LiveCoverage, err = iter.Bool()
		case "start_time":
			f.StartTime, err = simdutil.TimeFromIter(iter)
		case "flags":
			f.Flags, err = simdutil.IntFromIter(iter)
		case "created_at":
			f.CreatedAt, err = simdutil.TimeFromIter(iter)
		case "updated_at":
			f.UpdatedAt, err = simdutil.TimeFromIter(iter)
		case "published_at":
			f.PublishedAt, err = simdutil.TimeFromIter(iter)
		}

		if err != nil {
			return fmt.Errorf("%q unmarshal: %w", name, err)
		}
	}

	return nil
}

func (f *Fixture) ApplyPatchSimdJSON(path string, iter *simdjson.Iter) error {
	var (
		err                     error
		key, rest, partialPatch = strings.Cut(path, "/")
	)

	switch key {
	case "status":
		f.Status, err = simdutil.IntFromIter(iter)
	case "type":
		f.Type, err = simdutil.IntFromIter(iter)
	case "start_time":
		f.StartTime, err = simdutil.TimeFromIter(iter)
	case "live_coverage":
		f.LiveCoverage, err = iter.Bool()
	case "competitors":
		if partialPatch {
			if f.Competitors == nil {
				return fmt.Errorf("patch nil competitors")
			}

			return f.Competitors.ApplyPatchSimdJSON(rest, iter)
		}

		f.Competitors = make(Competitors)
		return f.Competitors.FromIter(iter, nil)

	case "tournament":
		if partialPatch {
			return f.Tournament.ApplyPatchSimdJSON(rest, iter)
		}

		return f.Tournament.FromIter(iter, nil)

	case "streams":
		if partialPatch {
			if f.Streams == nil {
				return fmt.Errorf("patch nil streams")
			}

			return f.Streams.ApplyPatchSimdJSON(rest, iter)
		}

		f.Streams = make(Streams)
		return f.Streams.FromIter(iter, nil)

	case "venue":
		if partialPatch {
			return f.Venue.ApplyPatchSimdJSON(rest, iter)
		}

		return f.Venue.FromIter(iter, nil)
	}

	if err != nil {
		return fmt.Errorf("%q unmarshal: %w", key, err)
	}

	return nil
}
