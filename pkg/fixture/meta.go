//go:generate go run github.com/mailru/easyjson/easyjson meta.go
package fixture

// easyjson:json
type Meta map[string]interface{}

func (m Meta) Clone() Meta {
	result := make(Meta, len(m))
	for k, v := range m {
		result[k] = v
	}

	return result
}
