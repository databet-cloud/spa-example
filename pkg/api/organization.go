package api

type OrganizationLocalized struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// SportIDs for which organization could be used
	SportIDs []string `json:"sport_ids"`
	Logo     Logo     `json:"logo"`
	Locale   Locale   `json:"locale"`
}
