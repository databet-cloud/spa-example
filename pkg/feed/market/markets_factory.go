package market

import "fmt"

func MarketsFromMarketIter(iter *Iterator) (Markets, error) {
	var (
		err    error
		market Market
		res    = make(Markets, 128)
	)

	err = iter.Rewind()
	if err != nil {
		return nil, fmt.Errorf("rewind: %w", err)
	}

	for {
		m, err := iter.Next(&market)
		if err != nil {
			return nil, fmt.Errorf("next: %w", err)
		}

		if m == nil {
			break
		}

		res[m.ID] = *m
	}

	return res, nil
}
