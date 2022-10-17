package market

type FilterOption struct {
	IDs      []string
	TypeIDs  []int
	Statuses []int
}

type normalizedFilters struct {
	ids      map[string]struct{}
	typeIDs  map[int]struct{}
	statuses map[Status]struct{}
}

type FilterFn func(m Market) bool

func FilterOptionsToFilterFn(options []FilterOption) FilterFn {
	res := normalizedFilters{
		ids:      make(map[string]struct{}),
		typeIDs:  make(map[int]struct{}),
		statuses: make(map[Status]struct{}),
	}

	for _, fo := range options {
		for _, id := range fo.IDs {
			res.ids[id] = struct{}{}
		}

		for _, typeID := range fo.TypeIDs {
			res.typeIDs[typeID] = struct{}{}
		}

		for _, status := range fo.Statuses {
			res.statuses[Status(status)] = struct{}{}
		}
	}

	return res.ToFilter()
}

func (mf normalizedFilters) ToFilter() func(m Market) bool {
	filterFuns := make([]FilterFn, 0, 4)

	if len(mf.ids) > 0 {
		filterFuns = append(filterFuns, func(m Market) bool {
			_, ok := mf.ids[m.ID]
			return ok
		})
	}

	if len(mf.typeIDs) > 0 {
		filterFuns = append(filterFuns, func(m Market) bool {
			_, ok := mf.typeIDs[m.TypeID]
			return ok
		})
	}

	if len(mf.statuses) > 0 {
		filterFuns = append(filterFuns, func(m Market) bool {
			_, ok := mf.statuses[m.Status]
			return ok
		})
	}

	if len(filterFuns) > 0 {
		return func(m Market) bool {
			for _, f := range filterFuns {
				if !f(m) {
					return false
				}
			}

			return true
		}
	}

	return func(m Market) bool {
		return true
	}
}
