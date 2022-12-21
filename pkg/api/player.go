package api

import "time"

type PlayerLocalized struct {
	ID              string     `json:"id"`
	OriginalID      string     `json:"original_id"`
	Type            PlayerType `json:"type"`
	SourceID        string     `json:"source_id"`
	Locale          Locale     `json:"locale"`
	Name            string     `json:"name"`
	Nickname        string     `json:"nickname"`
	Description     string     `json:"description"`
	Keywords        []string   `json:"keywords"`
	Version         string     `json:"version"`
	ProviderID      string     `json:"provider_id"`
	ProviderVersion string     `json:"provider_version"`
	SportID         string     `json:"sport_id"`
	Logo            struct {
		File string `json:"file"`
		URL  string `json:"url"`
	} `json:"logo"`
	LogoAutoUpdate bool      `json:"logo_auto_update"`
	CountryCode    string    `json:"country_code"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	// Statistics ???
	Statistics any            `json:"statistics"`
	Duplicated bool           `json:"duplicated"`
	UpdatedAt  time.Time      `json:"updated_at"`
	Meta       map[string]any `json:"meta"`
}

type PlayerType int

type PlayerLocalization struct {
	Locale      Locale   `json:"locale"`
	Name        string   `json:"name"`
	Nickname    string   `json:"nickname"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}
