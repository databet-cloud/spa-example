package patch

type Patch map[string]any

type Number interface {
	~int | ~uint | ~float64
}
