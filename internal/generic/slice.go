package generic

func CastSlice[T any](s []any) []T {
	res := make([]T, 0, len(s))

	for i := range s {
		if v, ok := s[i].(T); ok {
			res = append(res, v)
		}
	}

	return res
}
