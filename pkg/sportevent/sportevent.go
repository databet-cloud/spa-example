//go:generate go run github.com/mailru/easyjson/easyjson sportevent.go
package sportevent

import (
	"time"

	"github.com/databet-cloud/databet-go-sdk/pkg/fixture"
	"github.com/databet-cloud/databet-go-sdk/pkg/market"
)

//easyjson:json
type SportEvent struct {
	ID        string                 `json:"id"`
	Meta      map[string]interface{} `json:"meta"`
	Fixture   fixture.Fixture        `json:"fixture"`
	Markets   market.Markets         `json:"markets"`
	BetStop   bool                   `json:"bet_stop"`
	Sources   []Source               `json:"sources"`
	UpdatedAt time.Time              `json:"updated_at"`
}
