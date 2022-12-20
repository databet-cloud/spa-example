package api

type SearchTournamentsResponse struct {
	Tournaments []Tournament `json:"tournaments"`
	Count       int          `json:"count"`
}

type SearchLocalizedTournamentsResponse struct {
	Tournaments []TournamentLocalized `json:"tournaments"`
	Count       int                   `json:"count"`
}
