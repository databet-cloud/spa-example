package api

import "time"

type Tournament struct {
	ID string `json:"id"`
	// OriginalID foreign id, maybe omit
	OriginalID string `json:"original_id"`
	// Type Replica/Standalone, maybe omit, ask frontends
	Type TournamentType `json:"type"`
	// SourceID of original tournament
	SourceID string `json:"source_id"`
	Version  string `json:"version"`
	// ProviderID maybe omit?
	ProviderID string `json:"provider_id"`
	SportID    string `json:"sport_id"`
	// CountryCode ISO 3166-1 alpha-2
	CountryCode string `json:"country_code"`
	// Organization deprecated
	Organization string `json:"organization"`
	// OrganizationID, maybe omit/add api
	OrganizationID string `json:"organization_id"`
	Duplicated     bool   `json:"duplicated"`
	Logo           struct {
		// File, ask hollowj
		File string `json:"file"`
		// URL to cdn.gin.bet, ask hollowj
		URL string `json:"url"`
	} `json:"logo"`
	LimitGroups struct {
		PrematchID *string `json:"prematch_id"`
		LiveID     *string `json:"live_id"`
	} `json:"limit_groups"`
	DateStart time.Time      `json:"date_start"`
	DateEnd   time.Time      `json:"date_end"`
	UpdatedAt time.Time      `json:"updated_at"`
	Meta      map[string]any `json:"meta"`
	// LocalizationOverridden used to not override by auto import, maybe omit
	LocalizationOverridden bool `json:"localization_overridden"`

	Localizations []TournamentLocalization `json:"localizations"`
}

type TournamentLocalization struct {
	Locale      Locale   `json:"locale"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}

type TournamentType int

type TournamentLocalized struct {
	ID string `json:"id"`
	// OriginalID foreign id, maybe omit
	OriginalID  string `json:"original_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// Type Replica/Standalone, maybe omit, ask frontends
	Type TournamentType `json:"type"`
	// SourceID of original tournament
	SourceID string `json:"source_id"`
	Version  string `json:"version"`
	// ProviderID maybe omit?
	ProviderID string `json:"provider_id"`
	SportID    string `json:"sport_id"`
	// CountryCode ISO 3166-1 alpha-2
	CountryCode string `json:"country_code"`
	// Organization deprecated
	Organization string `json:"organization"`
	// OrganizationID, maybe omit/add api
	OrganizationID string `json:"organization_id"`
	Duplicated     bool   `json:"duplicated"`
	Logo           struct {
		// File, ask hollowj
		File string `json:"file"`
		// URL to cdn.gin.bet, ask hollowj
		URL string `json:"url"`
	} `json:"logo"`
	LimitGroups struct {
		PrematchID *string `json:"prematch_id"`
		LiveID     *string `json:"live_id"`
	} `json:"limit_groups"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
	UpdatedAt time.Time `json:"updated_at"`
	Locale    Locale    `json:"locale"`
	// Keywords like tags, maybe omit
	Keywords []string       `json:"keywords"`
	Meta     map[string]any `json:"meta"`
	// LocalizationOverridden used to not override by auto import, maybe omit
	LocalizationOverridden bool `json:"localization_overridden"`
}
