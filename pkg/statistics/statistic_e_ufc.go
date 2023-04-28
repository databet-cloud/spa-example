package statistics

type EUFCReasonType = string

const (
	EUFCReasonKo    EUFCReasonType = "ko"
	EUFCReasonTKO   EUFCReasonType = "tko"
	EUFCReasonDQ    EUFCReasonType = "dq"
	EUFCReasonSUB   EUFCReasonType = "sub"
	EUFCReasonD     EUFCReasonType = "d"
	EUFCReasonTD    EUFCReasonType = "td"
	EUFCReasonOther EUFCReasonType = "o"
)

func (s EUFCStatistic) Typ() string {
	return s.Type
}

type EUFCStatistic struct {
	Type   string      `json:"type"`
	Rounds []EUFCRound `json:"rounds"`
	Result EUFCResult  `json:"result"`
}

type EUFCRound struct {
	Number int   `json:"number"`
	Timer  Timer `json:"timer"`
}

type EUFCResult struct {
	RoundNumber *int            `json:"round_number"`
	Winner      *Team           `json:"winner"`
	Time        *int            `json:"time"`
	Reason      *EUFCReasonType `json:"reason"`
}
