package sportevent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/minio/simdjson-go"

	"github.com/databet-cloud/databet-go-sdk/pkg/market"
	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
)

func (se *SportEvent) ApplyPatches(rawPatches json.RawMessage) error {
	rootIter, err := simdutil.JSONToRootIter(rawPatches)
	if err != nil {
		return fmt.Errorf("json to root iter: %w", err)
	}

	obj := new(simdjson.Object)
	tmpObj := new(simdjson.Object)

	if obj, err = rootIter.Object(obj); err != nil {
		return fmt.Errorf("object: %w", err)
	}

	var iter simdjson.Iter

	for {
		path, elementType, err := obj.NextElement(&iter)
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
			value, err := iter.Bool()
			if err != nil {
				return err
			}

			se.BetStop = value

		case "updated_at":
			value, err := iter.String()
			if err != nil {
				return err
			}

			t, err := time.Parse(time.RFC3339, value)
			if err != nil {
				return fmt.Errorf("parse time: %w", err)
			}

			se.UpdatedAt = t

		case "fixture":
			if partialPatch {
				err = se.Fixture.ApplyPatchSimdJSON(rest, &iter)
				if err != nil {
					return err
				}

				continue
			}

			obj, err := iter.Object(tmpObj)
			if err != nil {
				return err
			}

			err = se.Fixture.UnmarshalSimdJSON(obj)
			if err != nil {
				return fmt.Errorf("unmarshal fixture: %w", err)
			}

		case "markets":
			if partialPatch {
				err = se.Markets.ApplyPatchSimdJSON(rest, &iter)
				if err != nil {
					return err
				}

				continue
			}

			tmpIter := iter
			marketIter, err := market.NewIterator(&tmpIter)
			if err != nil {
				return fmt.Errorf("new market iter: %w", err)
			}

			se.Markets, err = market.MarketsFromMarketIter(marketIter)
			if err != nil {
				return fmt.Errorf("unmarshal markets: %w", err)
			}

		default:
			continue
		}
	}

	return err
}
