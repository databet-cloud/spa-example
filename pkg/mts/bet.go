package mts

import (
	"net"
	"time"

	"github.com/databet-cloud/databet-go-sdk/pkg/restriction"
)

type Bet struct {
	Version       string                  `json:"version"`
	ID            string                  `json:"id"`
	ForeignID     string                  `json:"foreign_id"`
	BookmakerID   string                  `json:"bookmaker_id"`
	Status        BetStatus               `json:"status"`
	Type          BetType                 `json:"type"`
	Stake         MultiMoney              `json:"stake"`
	Refund        MultiMoney              `json:"refund"`
	RefundBase    MultiMoney              `json:"refund_base"`
	Selections    []*Selection            `json:"selections"`
	PlayerInfo    PlayerInfo              `json:"player_info"`
	CreatedAt     time.Time               `json:"created_at"`
	Events        []Event                 `json:"events"`
	Markers       []Marker                `json:"markers"`
	PlaceContext  *PlaceContext           `json:"place_context,omitempty"`
	CashOutOrders map[string]CashOutOrder `json:"cash_out_orders"`
}

type BetTypeCode int

const (
	BetTypeCodeSingle  BetTypeCode = 0
	BetTypeCodeExpress BetTypeCode = 1
	BetTypeCodeSystem  BetTypeCode = 2
)

type BetType struct {
	Code BetTypeCode `json:"code"`
	Size []int       `json:"size"`
}

func NewBetType(code BetTypeCode, size []int) BetType {
	return BetType{Code: code, Size: size}
}

type BetStatusCode int

const (
	BetStatusCodePlaced BetStatusCode = iota
	BetStatusCodeAccepted
	BetStatusCodeDeclined
	BetStatusCodeSettled
	BetStatusCodeUnsettled
	BetStatusCodeRefunded
	BetStatusCodeRefundedManually
)

type BetStatus struct {
	Code   BetStatusCode `json:"code"`
	Reason string        `json:"reason"`
}

type Money struct {
	Value        string `json:"value"`
	CurrencyCode string `json:"currency_code"`
}

func NewMoney(value string, currencyCode string) Money {
	return Money{Value: value, CurrencyCode: currencyCode}
}

type MultiMoney struct {
	Base   Money `json:"base"`
	Origin Money `json:"origin"`
}

func NewMultiMoney(base Money, origin Money) MultiMoney {
	return MultiMoney{Base: base, Origin: origin}
}

type Selection struct {
	SportEventIsLive bool            `json:"sport_event_is_live"`
	SportEventID     string          `json:"sport_event_id"`
	MarketID         string          `json:"market_id"`
	OddID            string          `json:"odd_id"`
	Value            string          `json:"value"`
	Marge            string          `json:"marge"`
	Status           SelectionStatus `json:"status"`
	MarketType       int             `json:"market_type"`
	SportID          string          `json:"sport_id"`
	TournamentID     string          `json:"tournament_id"`
}

type SelectionStatus int

const (
	SelectionNotResulted SelectionStatus = iota
	SelectionWin
	SelectionLoss
	SelectionHalfWin
	SelectionHalfLoss
	SelectionRefunded
	SelectionRefundedManually
)

type PlayerInfo struct {
	PlayerID    string `json:"id"`
	RiskGroupID string `json:"risk_group_id"`
	ClientIP    net.IP `json:"client_ip"`
	CountryCode string `json:"country_code"`
}

func NewPlayerInfo(playerID string, riskGroupID string, clientIP net.IP, countryCode string) PlayerInfo {
	return PlayerInfo{PlayerID: playerID, RiskGroupID: riskGroupID, ClientIP: clientIP, CountryCode: countryCode}
}

type EventName string

const (
	BetPlaced           EventName = "placed"
	BetAccepted         EventName = "accepted"
	BetDeclined         EventName = "declined"
	BetSettled          EventName = "settled"
	BetUnsettled        EventName = "unsettled"
	BetRefunded         EventName = "refunded"
	BetRefundedManually EventName = "refunded_manually"
	MarkerAdded         EventName = "marker_added"
	MarkerRemoved       EventName = "marker_removed"
	BetSelectionChanged EventName = "selection.changed"
)

type Event struct {
	Name    string    `json:"name"`
	Time    time.Time `json:"time"`
	Payload any       `json:"payload"`
}

type MarkerType string

const (
	MarkerTypeMatchFixing MarkerType = "match_fixing"
	MarkerTypeAfterGoal   MarkerType = "after_goal"
	MarkerTypeSureBet     MarkerType = "sure_bet"
	MarkerTypeSureBetAuto MarkerType = "sure_bet_auto"
	MarkerTypeForkAuto    MarkerType = "fork_auto"
	MarkerTypeEmpty       MarkerType = ""
)

type Marker struct {
	Type      MarkerType `json:"type"`
	Removed   bool       `json:"removed"`
	RemovedBy string     `json:"removed_by"`
	RemovedAt time.Time  `json:"removed_at"`
	CreatedBy string     `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
}

type PlaceContext struct {
	Restrictions     []restriction.Restriction `json:"restrictions"`
	PlayerLimit      *PlayerLimit              `json:"player_limit,omitempty"`
	Fixtures         []Fixture                 `json:"fixtures"`
	SportEventRisks  []SportEventRisk          `json:"sport_event_risk"`
	SportEventLimits []SportEventLimit         `json:"sport_event_limits"`
	MarketLimits     []MarketLimit             `json:"market_limits"`
	MaxBet           string                    `json:"max_bet"`
	DelayMs          uint64                    `json:"delay_ms"`
	Forks            []Fork                    `json:"forks"`
	BetStopSettings  []BetStopSettings         `json:"bet_stop_settings"`
}

func NewPlaceContext(
	restrictions []restriction.Restriction,
	playerLimit *PlayerLimit,
	fixtures []Fixture,
	sportEventRisks []SportEventRisk,
	sportEventLimits []SportEventLimit,
	marketLimits []MarketLimit,
	maxBet string,
	delayMs uint64,
	forks []Fork,
	betStopSettings []BetStopSettings,
) *PlaceContext {
	return &PlaceContext{
		Restrictions:     restrictions,
		PlayerLimit:      playerLimit,
		Fixtures:         fixtures,
		SportEventRisks:  sportEventRisks,
		SportEventLimits: sportEventLimits,
		MarketLimits:     marketLimits,
		MaxBet:           maxBet,
		DelayMs:          delayMs,
		Forks:            forks,
		BetStopSettings:  betStopSettings,
	}
}

type Fork struct {
	SportEventID string  `json:"sport_event_id"`
	OddID        string  `json:"odd_id"`
	MarketID     string  `json:"market_id"`
	Percent      float32 `json:"percent"`
	ParserID     string  `json:"parser_id"`
}

func NewFork(
	sportEventID string,
	oddID string,
	marketID string,
	percent float32,
	parserID string,
) Fork {
	return Fork{
		SportEventID: sportEventID,
		OddID:        oddID,
		MarketID:     marketID,
		Percent:      percent,
		ParserID:     parserID,
	}
}

type PlayerLimit struct {
	PlayerID    string               `json:"player_id"`
	RiskGroup   RiskGroup            `json:"risk_group"`
	SportEvents map[string]RiskGroup `json:"sport_events"`
}

func NewPlayerLimit(playerID string, riskGroup RiskGroup, sportEvents map[string]RiskGroup) *PlayerLimit {
	return &PlayerLimit{PlayerID: playerID, RiskGroup: riskGroup, SportEvents: sportEvents}
}

type RiskGroup struct {
	ID                    string          `json:"id"`
	Title                 string          `json:"title"`
	StakeMultiplier       StakeMultiplier `json:"stake_multiplier"`
	BetProcessingDelayMs  int             `json:"bet_processing_delay_ms"`
	BetIntervalMultiplier string          `json:"bet_interval_multiplier"`
}

func NewRiskGroup(
	id string,
	title string,
	stakeMultiplier StakeMultiplier,
	betProcessingDelayMs int,
	betIntervalMultiplier string,
) RiskGroup {
	return RiskGroup{
		ID:                    id,
		Title:                 title,
		StakeMultiplier:       stakeMultiplier,
		BetProcessingDelayMs:  betProcessingDelayMs,
		BetIntervalMultiplier: betIntervalMultiplier,
	}
}

type StakeMultiplier struct {
	Live     string `json:"live"`
	Prematch string `json:"prematch"`
}

func NewStakeMultiplier(live string, prematch string) StakeMultiplier {
	return StakeMultiplier{Live: live, Prematch: prematch}
}

type Fixture struct {
	ID     string `json:"id"`
	IsLive bool   `json:"is_live"`
}

func NewFixture(id string, isLive bool) Fixture {
	return Fixture{ID: id, IsLive: isLive}
}

type SportEventRisk struct {
	SportEventID string `json:"sport_event_id"`
	Total        string `json:"total"`
	Ordinar      string `json:"ordinar"`
}

func NewSportEventRisk(sportEventID string, total string, ordinar string) SportEventRisk {
	return SportEventRisk{SportEventID: sportEventID, Total: total, Ordinar: ordinar}
}

type SportEventLimit struct {
	SportEventID      string    `json:"sport_event_id"`
	Group             Group     `json:"group,omitempty"`
	LiveBetLimits     BetLimits `json:"live_bet_limits"`
	PrematchBetLimits BetLimits `json:"prematch_bet_limits"`
	ForbiddenBetTypes []int     `json:"forbidden_bet_types"`
}

func NewSportEventLimit(
	sportEventID string,
	group Group,
	liveBetLimits BetLimits,
	prematchBetLimits BetLimits,
	forbiddenBetTypes []int,
) SportEventLimit {
	return SportEventLimit{
		SportEventID:      sportEventID,
		Group:             group,
		LiveBetLimits:     liveBetLimits,
		PrematchBetLimits: prematchBetLimits,
		ForbiddenBetTypes: forbiddenBetTypes,
	}
}

type BetLimits struct {
	BetIntervalMs         int `json:"interval_ms"`
	BetIntervalMultiplier int `json:"interval_multiplier"`
	BetDelayMs            int `json:"delay_ms"`
}

func NewBetLimits(betIntervalMs int, betIntervalMultiplier int, betDelayMs int) BetLimits {
	return BetLimits{BetIntervalMs: betIntervalMs, BetIntervalMultiplier: betIntervalMultiplier, BetDelayMs: betDelayMs}
}

type Group struct {
	MaxSportEventRisk string `json:"max_sport_event_risk"`
	MaxBetRisk        string `json:"max_bet_risk"`
}

func NewGroup(maxSportEventRisk string, maxBetRisk string) Group {
	return Group{MaxSportEventRisk: maxSportEventRisk, MaxBetRisk: maxBetRisk}
}

type MarketLimit struct {
	Version              string    `json:"version"`
	ID                   string    `json:"id"`
	SportID              string    `json:"sport_id"`
	MarketID             string    `json:"market_id"`
	BetMaxRiskMultiplier float64   `json:"bet_max_risk_multiplier"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func NewMarketLimit(
	version string,
	id string,
	sportID string,
	marketID string,
	betMaxRiskMultiplier float64,
	createdAt time.Time,
	updatedAt time.Time,
) MarketLimit {
	return MarketLimit{
		Version:              version,
		ID:                   id,
		SportID:              sportID,
		MarketID:             marketID,
		BetMaxRiskMultiplier: betMaxRiskMultiplier,
		CreatedAt:            createdAt,
		UpdatedAt:            updatedAt,
	}
}

type BetStopStatus string

const (
	BetStopStatusInitial BetStopStatus = ""
	BetStopStatusStopped BetStopStatus = "stopped"
	BetStopStatusStarted BetStopStatus = "started"
)

type BetStopSettingsManual struct {
	Status     BetStopStatus `json:"status" bson:"status"`
	AutoCancel bool          `json:"auto_cancel" bson:"auto_cancel"`
}

type BetStopSettings struct {
	SportEventID string                `json:"sport_event_id" bson:"sport_event_id"`
	BetStop      bool                  `json:"bet_stop" bson:"bet_stop"`
	System       BetStopStatus         `json:"system" bson:"system"`
	Manual       BetStopSettingsManual `json:"manual" bson:"manual"`
}

func NewBetStopSettings(
	sportEventID string,
	betStop bool,
	system BetStopStatus,
	manual BetStopSettingsManual,
) BetStopSettings {
	return BetStopSettings{
		SportEventID: sportEventID,
		BetStop:      betStop,
		System:       system,
		Manual:       manual,
	}
}
