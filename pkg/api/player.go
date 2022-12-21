package api

type PlayerLocalized struct {
	ID          string   `json:"id"`
	Locale      Locale   `json:"locale"`
	Nickname    string   `json:"nickname"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	SportID     string   `json:"sport_id"`
	Logo        Logo     `json:"logo"`
	// CountryCode ISO 3166-1 alpha-2
	CountryCode string `json:"country_code"`
}
