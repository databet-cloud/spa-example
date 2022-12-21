package api

type TeamLocalized struct {
	ID             string   `json:"id"`
	CountryCode    string   `json:"country_code"`
	Name           string   `json:"name"`
	SportID        string   `json:"sport_id"`
	OrganizationID string   `json:"organization_id"`
	Logo           Logo     `json:"logo"`
	Description    string   `json:"description"`
	Keywords       []string `json:"keywords"`
}
