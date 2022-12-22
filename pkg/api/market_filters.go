package api

import (
	"net/url"
	"strconv"
)

type MarketFilters struct {
	SportIds   []string
	Tags       []string
	Template   string
	MarketType string
	// IncludeDeprecated indicates whether you want your response to contain deprecated markets or not.
	// By default, response doesn't contain such markets.
	IncludeDeprecated *bool
}

func NewMarketFilters() *MarketFilters {
	return new(MarketFilters)
}

func (f *MarketFilters) WithSportIDs(sportIDs ...string) *MarketFilters {
	newFilters := *f
	newFilters.SportIds = append(newFilters.SportIds, sportIDs...)
	return &newFilters
}

func (f *MarketFilters) WithTags(tags ...string) *MarketFilters {
	newFilters := *f
	newFilters.Tags = append(newFilters.Tags, tags...)
	return &newFilters
}

func (f *MarketFilters) WithTemplate(template string) *MarketFilters {
	newFilters := *f
	newFilters.Template = template
	return &newFilters
}

func (f *MarketFilters) WithMarketType(marketType string) *MarketFilters {
	newFilters := *f
	newFilters.MarketType = marketType
	return &newFilters
}

func (f *MarketFilters) WithIncludeDeprecated(includeDeprecated bool) *MarketFilters {
	newFilters := *f
	newFilters.IncludeDeprecated = &includeDeprecated
	return &newFilters
}

func (f *MarketFilters) ToQueryParams() url.Values {
	queryParams := make(url.Values)

	for i := range f.SportIds {
		queryParams.Add("sport_ids[]", f.SportIds[i])
	}

	for i := range f.Tags {
		queryParams.Add("tags[]", f.Tags[i])
	}

	if f.Template != "" {
		queryParams.Set("template", f.Template)
	}

	if f.MarketType != "" {
		queryParams.Set("market_type", f.MarketType)
	}

	if f.IncludeDeprecated != nil {
		queryParams.Set("with_deprecated", strconv.FormatBool(*f.IncludeDeprecated))
	}

	return queryParams
}
