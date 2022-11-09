package sharedresource

import "net/url"

type SportFilters struct {
	Tags []string
}

func NewSportFilters() *SportFilters {
	return &SportFilters{}
}

func (f *SportFilters) WithTags(tags ...string) *SportFilters {
	newFilters := *f
	newFilters.Tags = append(newFilters.Tags, tags...)
	return &newFilters
}

func (f *SportFilters) ToQueryParams() url.Values {
	queryParams := url.Values{}

	for i := range f.Tags {
		queryParams.Add("tags[]", f.Tags[i])
	}

	return queryParams
}
