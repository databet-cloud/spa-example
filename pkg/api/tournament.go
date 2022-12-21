package api

type TournamentLocalized struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SportID     string `json:"sport_id"`
	// CountryCode ISO 3166-1 alpha-2
	CountryCode    string         `json:"country_code"`
	OrganizationID string         `json:"organization_id"`
	Logo           Logo           `json:"logo"`
	Locale         Locale         `json:"locale"`
	Keywords       []string       `json:"keywords"`
	Meta           map[string]any `json:"meta"`
}
