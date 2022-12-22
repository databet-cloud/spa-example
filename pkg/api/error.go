package api

import "fmt"

const (
	errCodeTournamentNotFound = "tournament_not_found"
	errCodePlayerNotFound     = "player_not_found"
	errCodeTeamNotFound       = "team_not_found"
)

var (
	ErrNotFound           = fmt.Errorf("not found")
	ErrUnknown            = fmt.Errorf("unknown error")
	ErrInvalidCertificate = fmt.Errorf("invalid certificate")
)

type apiError struct {
	Code string         `json:"code"`
	Data map[string]any `json:"data"`
}
