package api_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/databet-cloud/databet-go-sdk/pkg/api"
	"github.com/databet-cloud/databet-go-sdk/pkg/api/mocks"
)

const testURL = "http://api-test"

type ClientHTTPTestSuite struct {
	suite.Suite

	client    *http.Client
	apiClient *api.ClientHTTP
}

func TestClientHTTP(t *testing.T) {
	suite.Run(t, new(ClientHTTPTestSuite))
}

func (s *ClientHTTPTestSuite) SetupSuite() {
	s.client = http.DefaultClient
	s.apiClient = api.NewClientHTTP(s.client, testURL)
}

func (s *ClientHTTPTestSuite) TestFindTournamentByID() {
	testCases := []struct {
		name        string
		httpResp    *http.Response
		expected    *api.Tournament
		expectedErr error
	}{
		{
			name:     "succeed",
			httpResp: s.makeResponse(http.StatusOK, "find-tournament-by-id/response-success.json"),
			expected: &DefaultTournament,
		},
		{
			name:        "not found",
			httpResp:    s.makeResponse(http.StatusOK, "find-tournament-by-id/response-not-found.json"),
			expectedErr: api.ErrNotFound,
		},
		{
			name:        "unauthorized",
			httpResp:    s.makeResponse(http.StatusUnauthorized, ""),
			expectedErr: api.ErrInvalidCertificate,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			req := s.makeRequest(
				http.MethodGet,
				fmt.Sprintf("%s/tournaments/tournamentID", testURL),
				http.NoBody,
			)
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(req)).Return(tc.httpResp, nil)

			actual, actualErr := s.apiClient.FindTournamentByID(context.Background(), "tournamentID")

			s.Equal(tc.expected, actual)
			s.ErrorIs(actualErr, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestFindTournamentsByIDs() {
	testCases := []struct {
		name        string
		httpResp    *http.Response
		expected    []api.Tournament
		expectedErr error
	}{
		{
			name:     "succeed",
			httpResp: s.makeResponse(http.StatusOK, "find-tournaments-by-ids/response-success.json"),
			expected: []api.Tournament{DefaultTournament},
		},
		{
			name:        "unauthorized",
			httpResp:    s.makeResponse(http.StatusUnauthorized, ""),
			expectedErr: api.ErrInvalidCertificate,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			req := s.makeRequest(
				http.MethodGet,
				fmt.Sprintf("%s/tournaments/by-ids?ids[]=id1&ids[]=id2", testURL),
				http.NoBody,
			)
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(req)).Return(tc.httpResp, nil)

			actual, actualErr := s.apiClient.FindTournamentsByIDs(context.Background(), []string{"id1", "id2"})

			s.Equal(tc.expected, actual)
			s.ErrorIs(actualErr, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestSearchTournaments() {
	testCases := []struct {
		name                string
		req                 *api.SearchTournamentsRequest
		httpResp            *http.Response
		expected            *api.SearchTournamentsResponse
		expectedQueryParams url.Values
		expectedErr         error
	}{
		{
			name: "succeed",
			req: api.NewSearchTournamentsRequest().
				SetSearchNames("name1", "name2").
				SetSearchNameStrategy(api.SearchByOccurrence).
				SetCountryCode("UA").
				SetDateStartFrom(time.Date(2022, 12, 12, 0, 0, 0, 0, mustLoadLocation(s.T(), "Europe/Kyiv"))).
				SetDateStartTo(time.Date(2022, 12, 13, 0, 0, 0, 0, time.UTC)).
				SetDateEndFrom(time.Date(2022, 12, 12, 20, 0, 0, 0, time.UTC)).
				SetDateEndTo(time.Date(2022, 12, 12, 20, 0, 0, 0, time.UTC)).
				SetKeywords("keyword").
				SetDuplicated(true).
				SetDeactivated(false).
				SetLimit(10).
				SetOffset(5),
			expectedQueryParams: url.Values{
				"country_code":           []string{"UA"},
				"date_start_from":        []string{"2022-12-12T00:00:00+02:00"},
				"date_start_to":          []string{"2022-12-13T00:00:00Z"},
				"date_end_from":          []string{"2022-12-12T20:00:00Z"},
				"date_end_to":            []string{"2022-12-12T20:00:00Z"},
				"deactivated":            []string{"false"},
				"duplicated":             []string{"true"},
				"keywords[]":             []string{"keyword"},
				"limit":                  []string{"10"},
				"offset":                 []string{"5"},
				"search_string_strategy": []string{"by_occurrence"},
				"search_strings[]":       []string{"name1", "name2"},
			},
			httpResp: s.makeResponse(http.StatusOK, "search-tournaments/response-success.json"),
			expected: &api.SearchTournamentsResponse{
				Tournaments: []api.Tournament{DefaultTournament},
				Count:       5,
			},
		},
		{
			name:        "unauthorized",
			req:         api.NewSearchTournamentsRequest(),
			httpResp:    s.makeResponse(http.StatusUnauthorized, ""),
			expectedErr: api.ErrInvalidCertificate,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			req := s.makeRequest(
				http.MethodGet,
				fmt.Sprintf("%s/search/tournaments", testURL),
				http.NoBody,
			)
			req.URL.RawQuery = tc.expectedQueryParams.Encode()
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(req)).Return(tc.httpResp, nil)

			actual, actualErr := s.apiClient.SearchTournaments(context.Background(), tc.req)

			s.Equal(tc.expected, actual)
			s.ErrorIs(actualErr, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestFindLocalizedTournamentByID() {
	testCases := []struct {
		name        string
		httpResp    *http.Response
		expected    *api.TournamentLocalized
		expectedErr error
	}{
		{
			name:     "succeed",
			httpResp: s.makeResponse(http.StatusOK, "find-localized-tournament-by-id/response-success.json"),
			expected: &DefaultLocalizedTournament,
		},
		{
			name:        "not found",
			httpResp:    s.makeResponse(http.StatusOK, "find-localized-tournament-by-id/response-not-found.json"),
			expectedErr: api.ErrNotFound,
		},
		{
			name:        "unauthorized",
			httpResp:    s.makeResponse(http.StatusUnauthorized, ""),
			expectedErr: api.ErrInvalidCertificate,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			req := s.makeRequest(
				http.MethodGet,
				fmt.Sprintf("%s/tournaments/localized/tournamentID/en", testURL),
				http.NoBody,
			)
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(req)).Return(tc.httpResp, nil)

			actual, actualErr := s.apiClient.FindLocalizedTournamentByID(context.Background(), api.LocaleEnglish, "tournamentID")

			s.Equal(tc.expected, actual)
			s.ErrorIs(actualErr, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestFindLocalizedTournamentsByIDs() {
	testCases := []struct {
		name        string
		httpResp    *http.Response
		expected    []api.TournamentLocalized
		expectedErr error
	}{
		{
			name:     "succeed",
			httpResp: s.makeResponse(http.StatusOK, "find-localized-tournaments-by-ids/response-success.json"),
			expected: []api.TournamentLocalized{DefaultLocalizedTournament},
		},
		{
			name:        "unauthorized",
			httpResp:    s.makeResponse(http.StatusUnauthorized, ""),
			expectedErr: api.ErrInvalidCertificate,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			req := s.makeRequest(
				http.MethodGet,
				fmt.Sprintf("%s/tournaments/localized/en?ids[]=id1&ids[]=id2", testURL),
				http.NoBody,
			)
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(req)).Return(tc.httpResp, nil)

			actual, actualErr := s.apiClient.FindLocalizedTournamentsByIDs(
				context.Background(),
				api.LocaleEnglish,
				[]string{"id1", "id2"},
			)

			s.Equal(tc.expected, actual)
			s.ErrorIs(actualErr, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestSearchLocalizedTournaments() {
	testCases := []struct {
		name                string
		req                 *api.SearchTournamentsRequest
		httpResp            *http.Response
		expected            *api.SearchLocalizedTournamentsResponse
		expectedQueryParams url.Values
		expectedErr         error
	}{
		{
			name: "succeed",
			req: api.NewSearchTournamentsRequest().
				SetSearchNames("name1", "name2").
				SetSearchNameStrategy(api.SearchByOccurrence).
				SetCountryCode("UA").
				SetDateStartFrom(time.Date(2022, 12, 12, 0, 0, 0, 0, mustLoadLocation(s.T(), "Europe/Kyiv"))).
				SetDateStartTo(time.Date(2022, 12, 13, 0, 0, 0, 0, time.UTC)).
				SetDateEndFrom(time.Date(2022, 12, 12, 20, 0, 0, 0, time.UTC)).
				SetDateEndTo(time.Date(2022, 12, 12, 20, 0, 0, 0, time.UTC)).
				SetKeywords("keyword").
				SetDuplicated(true).
				SetDeactivated(false).
				SetLimit(10).
				SetOffset(5),
			expectedQueryParams: url.Values{
				"country_code":           []string{"UA"},
				"date_start_from":        []string{"2022-12-12T00:00:00+02:00"},
				"date_start_to":          []string{"2022-12-13T00:00:00Z"},
				"date_end_from":          []string{"2022-12-12T20:00:00Z"},
				"date_end_to":            []string{"2022-12-12T20:00:00Z"},
				"deactivated":            []string{"false"},
				"duplicated":             []string{"true"},
				"keywords[]":             []string{"keyword"},
				"limit":                  []string{"10"},
				"offset":                 []string{"5"},
				"search_string_strategy": []string{"by_occurrence"},
				"search_strings[]":       []string{"name1", "name2"},
			},
			httpResp: s.makeResponse(http.StatusOK, "search-localized-tournaments/response-success.json"),
			expected: &api.SearchLocalizedTournamentsResponse{
				Tournaments: []api.TournamentLocalized{DefaultLocalizedTournament},
				Count:       5,
			},
		},
		{
			name:        "unauthorized",
			req:         api.NewSearchTournamentsRequest(),
			httpResp:    s.makeResponse(http.StatusUnauthorized, ""),
			expectedErr: api.ErrInvalidCertificate,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			req := s.makeRequest(
				http.MethodGet,
				fmt.Sprintf("%s/search/tournaments/localized/en", testURL),
				http.NoBody,
			)
			req.URL.RawQuery = tc.expectedQueryParams.Encode()
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(req)).Return(tc.httpResp, nil)

			actual, actualErr := s.apiClient.SearchLocalizedTournaments(context.Background(), "en", tc.req)

			s.Equal(tc.expected, actual)
			s.ErrorIs(actualErr, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestFindPlayerByID() {
	testCases := []struct {
		name        string
		httpResp    *http.Response
		expected    *api.Player
		expectedErr error
	}{
		{
			name:     "succeed",
			httpResp: s.makeResponse(http.StatusOK, "find-player-by-id/response-success.json"),
			expected: &DefaultPlayer,
		},
		{
			name:        "not found",
			httpResp:    s.makeResponse(http.StatusOK, "find-player-by-id/response-not-found.json"),
			expectedErr: api.ErrNotFound,
		},
		{
			name:        "unauthorized",
			httpResp:    s.makeResponse(http.StatusUnauthorized, ""),
			expectedErr: api.ErrInvalidCertificate,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			req := s.makeRequest(
				http.MethodGet,
				fmt.Sprintf("%s/players/playerID", testURL),
				http.NoBody,
			)
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(req)).Return(tc.httpResp, nil)

			actual, actualErr := s.apiClient.FindPlayerByID(context.Background(), "playerID")

			s.Equal(tc.expected, actual)
			s.ErrorIs(actualErr, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestFindPlayersByIDs() {
	testCases := []struct {
		name        string
		httpResp    *http.Response
		expected    []api.Player
		expectedErr error
	}{
		{
			name:     "succeed",
			httpResp: s.makeResponse(http.StatusOK, "find-players-by-ids/response-success.json"),
			expected: []api.Player{DefaultPlayer},
		},
		{
			name:        "unauthorized",
			httpResp:    s.makeResponse(http.StatusUnauthorized, ""),
			expectedErr: api.ErrInvalidCertificate,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			req := s.makeRequest(
				http.MethodGet,
				fmt.Sprintf("%s/players/by-ids?ids[]=id1&ids[]=id2", testURL),
				http.NoBody,
			)
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(req)).Return(tc.httpResp, nil)

			actual, actualErr := s.apiClient.FindPlayersByIDs(context.Background(), []string{"id1", "id2"})

			s.Equal(tc.expected, actual)
			s.ErrorIs(actualErr, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestFindLocalizedPlayerByID() {
	testCases := []struct {
		name        string
		httpResp    *http.Response
		expected    *api.PlayerLocalized
		expectedErr error
	}{
		{
			name:     "succeed",
			httpResp: s.makeResponse(http.StatusOK, "find-localized-player-by-id/response-success.json"),
			expected: &DefaultLocalizedPlayer,
		},
		{
			name:        "not found",
			httpResp:    s.makeResponse(http.StatusOK, "find-localized-player-by-id/response-not-found.json"),
			expectedErr: api.ErrNotFound,
		},
		{
			name:        "unauthorized",
			httpResp:    s.makeResponse(http.StatusUnauthorized, ""),
			expectedErr: api.ErrInvalidCertificate,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			req := s.makeRequest(
				http.MethodGet,
				fmt.Sprintf("%s/players/localized/playerID/en", testURL),
				http.NoBody,
			)
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(req)).Return(tc.httpResp, nil)

			actual, actualErr := s.apiClient.FindLocalizedPlayerByID(context.Background(), api.LocaleEnglish, "playerID")

			s.Equal(tc.expected, actual)
			s.ErrorIs(actualErr, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestFindLocalizedPlayersByIDs() {
	testCases := []struct {
		name        string
		httpResp    *http.Response
		expected    []api.PlayerLocalized
		expectedErr error
	}{
		{
			name:     "succeed",
			httpResp: s.makeResponse(http.StatusOK, "find-localized-players-by-ids/response-success.json"),
			expected: []api.PlayerLocalized{DefaultLocalizedPlayer},
		},
		{
			name:        "unauthorized",
			httpResp:    s.makeResponse(http.StatusUnauthorized, ""),
			expectedErr: api.ErrInvalidCertificate,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			req := s.makeRequest(
				http.MethodGet,
				fmt.Sprintf("%s/players/localized/en?ids[]=id1&ids[]=id2", testURL),
				http.NoBody,
			)
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(req)).Return(tc.httpResp, nil)

			actual, actualErr := s.apiClient.FindLocalizedPlayersByIDs(
				context.Background(),
				api.LocaleEnglish,
				[]string{"id1", "id2"},
			)

			s.Equal(tc.expected, actual)
			s.ErrorIs(actualErr, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) newReqMatcher(req *http.Request) reqMatcher {
	return reqMatcher{
		s:           &s.Suite,
		expectedReq: req,
	}
}

func (s *ClientHTTPTestSuite) makeResponse(status int, filePath string) *http.Response {
	var body io.ReadCloser = http.NoBody

	if filePath != "" {
		fileContent, err := os.ReadFile("testdata/" + filePath)
		s.Require().NoError(err)
		body = io.NopCloser(bytes.NewReader(fileContent))
	}

	return &http.Response{
		Status:     http.StatusText(status),
		StatusCode: status,
		Body:       body,
	}
}

func (s *ClientHTTPTestSuite) makeRequest(method string, url string, body io.Reader) *http.Request {
	httpReq, err := http.NewRequest(method, url, body)
	s.Require().NoError(err)

	return httpReq
}

type reqMatcher struct {
	s           *suite.Suite
	expectedReq *http.Request
}

func (r reqMatcher) Matches(x any) bool {
	actualReq, ok := x.(*http.Request)
	if !ok {
		return false
	}

	if !r.s.Equal(r.expectedReq.Method, actualReq.Method) {
		return false
	}

	if !r.s.Equal(r.expectedReq.URL.Query(), actualReq.URL.Query()) {
		return false
	}

	if !r.s.Equal(r.expectedReq.URL.Path, actualReq.URL.Path) {
		return false
	}

	expectedBody, err := io.ReadAll(r.expectedReq.Body)
	r.s.Require().NoError(err)

	actualBody, err := io.ReadAll(actualReq.Body)
	r.s.Require().NoError(err)

	if len(expectedBody) > 0 && !r.s.JSONEq(string(expectedBody), string(actualBody)) {
		return false
	}

	return true
}

func (r reqMatcher) String() string {
	return fmt.Sprintf("%v", r.expectedReq)
}

func pointTo[T any](v T) *T {
	return &v
}

func mustParseTime(rawTime string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, rawTime)
	if err != nil {
		panic(err)
	}

	return parsedTime
}

func mustLoadLocation(t *testing.T, rawLocation string) *time.Location {
	t.Helper()

	loc, err := time.LoadLocation(rawLocation)
	require.NoError(t, err)

	return loc
}
