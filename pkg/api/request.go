package api

import (
	"net/url"
	"strconv"
	"time"
)

type SearchStrategy string

func (s SearchStrategy) String() string {
	return string(s)
}

const (
	SearchByOccurrence SearchStrategy = "by_occurrence"
	SearchByExact      SearchStrategy = "by_exact"
)

type SearchTournamentsRequest struct {
	// SearchNames specifies name patterns (en locale) to search.
	SearchNames []string
	// SearchNameStrategy specifies the strategy to search tournaments by their names (en locale).
	// If not specified, SearchByOccurrence strategy is applied.
	SearchNameStrategy *SearchStrategy
	SportIDs           []string
	// CountryCode ISO 3166-1 alpha-2
	CountryCode   *string
	DateStartFrom time.Time
	DateStartTo   time.Time
	DateEndFrom   time.Time
	DateEndTo     time.Time
	Keywords      []string
	Duplicated    *bool
	Deactivated   *bool
	Limit         *int
	Offset        *int
}

func NewSearchTournamentsRequest() *SearchTournamentsRequest {
	return &SearchTournamentsRequest{}
}

func (r *SearchTournamentsRequest) SetSearchNames(searchNames ...string) *SearchTournamentsRequest {
	r.SearchNames = searchNames
	return r
}

func (r *SearchTournamentsRequest) SetSearchNameStrategy(searchNameStrategy SearchStrategy) *SearchTournamentsRequest {
	r.SearchNameStrategy = &searchNameStrategy
	return r
}

func (r *SearchTournamentsRequest) SetSportIDs(sportIDs ...string) *SearchTournamentsRequest {
	r.SportIDs = sportIDs
	return r
}

func (r *SearchTournamentsRequest) SetCountryCode(countryCode string) *SearchTournamentsRequest {
	r.CountryCode = &countryCode
	return r
}

func (r *SearchTournamentsRequest) SetDateStartFrom(dateStartFrom time.Time) *SearchTournamentsRequest {
	r.DateStartFrom = dateStartFrom
	return r
}

func (r *SearchTournamentsRequest) SetDateStartTo(dateStartTo time.Time) *SearchTournamentsRequest {
	r.DateStartTo = dateStartTo
	return r
}

func (r *SearchTournamentsRequest) SetDateEndFrom(dateEndFrom time.Time) *SearchTournamentsRequest {
	r.DateEndFrom = dateEndFrom
	return r
}

func (r *SearchTournamentsRequest) SetDateEndTo(dateEndTo time.Time) *SearchTournamentsRequest {
	r.DateEndTo = dateEndTo
	return r
}

func (r *SearchTournamentsRequest) SetKeywords(keywords ...string) *SearchTournamentsRequest {
	r.Keywords = keywords
	return r
}

func (r *SearchTournamentsRequest) SetDuplicated(duplicated bool) *SearchTournamentsRequest {
	r.Duplicated = &duplicated
	return r
}

func (r *SearchTournamentsRequest) SetDeactivated(deactivated bool) *SearchTournamentsRequest {
	r.Deactivated = &deactivated
	return r
}

func (r *SearchTournamentsRequest) SetLimit(limit int) *SearchTournamentsRequest {
	r.Limit = &limit
	return r
}

func (r *SearchTournamentsRequest) SetOffset(offset int) *SearchTournamentsRequest {
	r.Offset = &offset
	return r
}

func (r *SearchTournamentsRequest) ToQueryParams() url.Values {
	values := url.Values{}

	if r.SearchNames != nil {
		values["search_strings[]"] = r.SearchNames
	}

	if r.SearchNameStrategy != nil {
		values.Set("search_string_strategy", r.SearchNameStrategy.String())
	}

	if r.SportIDs != nil {
		values["sport_ids[]"] = r.SportIDs
	}

	if r.CountryCode != nil {
		values.Set("country_code", *r.CountryCode)
	}

	if !r.DateStartFrom.IsZero() {
		values.Set("date_start_from", r.DateStartFrom.Format(time.RFC3339))
	}
	if !r.DateStartTo.IsZero() {
		values.Set("date_start_to", r.DateStartTo.Format(time.RFC3339))
	}

	if !r.DateEndFrom.IsZero() {
		values.Set("date_end_from", r.DateEndFrom.Format(time.RFC3339))
	}

	if !r.DateEndTo.IsZero() {
		values.Set("date_end_to", r.DateEndTo.Format(time.RFC3339))
	}

	if r.Keywords != nil {
		values["keywords[]"] = r.Keywords
	}

	if r.Duplicated != nil {
		values.Set("duplicated", strconv.FormatBool(*r.Duplicated))
	}

	if r.Deactivated != nil {
		values.Set("deactivated", strconv.FormatBool(*r.Deactivated))
	}

	if r.Limit != nil {
		values.Set("limit", strconv.Itoa(*r.Limit))
	}

	if r.Offset != nil {
		values.Set("offset", strconv.Itoa(*r.Offset))
	}

	return values
}

type SearchPlayersRequest struct {
	IDs                []string
	SearchNames        []string
	SportIDs           []string
	Keywords           []string
	SearchNameStrategy *SearchStrategy
	CountryCode        *string
	Duplicated         *bool
	Limit              *int
	Offset             *int
}

func (r *SearchPlayersRequest) SetIDs(ids ...string) *SearchPlayersRequest {
	r.IDs = ids
	return r
}

func (r *SearchPlayersRequest) SetSearchNames(searchNames ...string) *SearchPlayersRequest {
	r.SearchNames = searchNames
	return r
}

func (r *SearchPlayersRequest) SetSportIDs(sportIDs ...string) *SearchPlayersRequest {
	r.SportIDs = sportIDs
	return r
}

func (r *SearchPlayersRequest) SetKeywords(keywords ...string) *SearchPlayersRequest {
	r.Keywords = keywords
	return r
}

func (r *SearchPlayersRequest) SetSearchNameStrategy(searchNameStrategy SearchStrategy) *SearchPlayersRequest {
	r.SearchNameStrategy = &searchNameStrategy
	return r
}

func (r *SearchPlayersRequest) SetCountryCode(countryCode string) *SearchPlayersRequest {
	r.CountryCode = &countryCode
	return r
}

func (r *SearchPlayersRequest) SetDuplicated(duplicated bool) *SearchPlayersRequest {
	r.Duplicated = &duplicated
	return r
}

func (r *SearchPlayersRequest) SetLimit(limit int) *SearchPlayersRequest {
	r.Limit = &limit
	return r
}

func (r *SearchPlayersRequest) SetOffset(offset int) *SearchPlayersRequest {
	r.Offset = &offset
	return r
}

func (r *SearchPlayersRequest) ToQueryParams() url.Values {
	values := url.Values{}

	if len(r.IDs) > 0 {
		values["ids[]"] = r.IDs
	}

	if r.SearchNames != nil {
		values["search_strings[]"] = r.SearchNames
	}

	if r.SearchNameStrategy != nil {
		values.Set("search_string_strategy", r.SearchNameStrategy.String())
	}

	if r.SportIDs != nil {
		values["sport_ids[]"] = r.SportIDs
	}

	if r.CountryCode != nil {
		values.Set("country_code", *r.CountryCode)
	}

	if r.Keywords != nil {
		values["keywords[]"] = r.Keywords
	}

	if r.Duplicated != nil {
		values.Set("duplicated", strconv.FormatBool(*r.Duplicated))
	}

	if r.Limit != nil {
		values.Set("limit", strconv.Itoa(*r.Limit))
	}

	if r.Offset != nil {
		values.Set("offset", strconv.Itoa(*r.Offset))
	}

	return values
}
