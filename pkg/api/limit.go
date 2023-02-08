package api

import "time"

type LimitFixtureState int

const (
	LimitFixtureStateInProgress LimitFixtureState = iota + 1
	LimitFixtureStateFinished
)

type FindSportEventLimitsRequest struct {
	SportEventIDs      []string            `json:"sport_event_ids"`
	LimitLastUpdatedAt *time.Time          `json:"limit_last_updated-at"`
	FixtureStates      []LimitFixtureState `json:"fixture_states"`
}

func NewFindSportEventLimitRequest() *FindSportEventLimitsRequest {
	return new(FindSportEventLimitsRequest)
}

func (r *FindSportEventLimitsRequest) SetSportEventIDs(ids ...string) *FindSportEventLimitsRequest {
	r.SportEventIDs = ids
	return r
}

func (r *FindSportEventLimitsRequest) SetLimitLastUpdatedAt(t time.Time) *FindSportEventLimitsRequest {
	r.LimitLastUpdatedAt = &t
	return r
}

func (r *FindSportEventLimitsRequest) SetFixtureStates(fixtureStatus ...LimitFixtureState) *FindSportEventLimitsRequest {
	r.FixtureStates = fixtureStatus
	return r
}

type FindSportEventLimitsResponse struct {
	Limits []SportEventLimit `json:"limits"`
}

type SportEventLimit struct {
	BookmakerID  string            `json:"bookmaker_id"`
	SportEventID string            `json:"sport_event_id"`
	FixtureState LimitFixtureState `json:"fixture_state"`

	BetLimit       BetLimit      `json:"bet_limit"`
	Risks          Risks         `json:"risks"`
	MarketsLimits  []MarketLimit `json:"markets_limits"`
	LimitUpdatedAt time.Time     `json:"limit_updated_at"`

	Version   string    `json:"version"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BetLimit struct {
	BetDelay int `json:"bet_delay"`
}

type Risks struct {
	MaxBetRisk string `json:"max_bet_risk"`
}

type MarketLimit struct {
	MarketID             string `json:"market_id"`
	MaxBetRiskMultiplier string `json:"max_bet_risk_multiplier"`
}
