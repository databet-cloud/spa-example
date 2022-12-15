package api

import "time"

type Sport struct {
	ID            string             `json:"id"`
	Version       string             `json:"version"`
	Localizations SportLocalizations `json:"localizations"`
	Tags          []string           `json:"tags"`
	UpdatedAt     time.Time          `json:"updatedAt"`
}

type SportLocalizations map[Locale]SportLocalization

type SportLocalization struct {
	Locale Locale `json:"locale"`
	Name   string `json:"name"`
}

type SportLocalized struct {
	ID        string    `json:"id"`
	Version   string    `json:"version"`
	Locale    Locale    `json:"locale"`
	Name      string    `json:"name"`
	Tags      []string  `json:"tags"`
	UpdatedAt time.Time `json:"updated_at"`
}
