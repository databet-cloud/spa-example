package api

type TeamLocalized struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// CountryCode ISO 3166-1 alpha-2
	CountryCode string   `json:"country_code"`
	SportID     string   `json:"sport_id"`
	Logo        Logo     `json:"logo"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}
