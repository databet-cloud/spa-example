package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bytedance/sonic"

	"github.com/databet-cloud/databet-go-sdk/internal/simdutil"
)

type ClientHTTP struct {
	httpClient *http.Client
	apiURL     string
}

var _ Client = (*ClientHTTP)(nil)

func NewClientHTTP(httpClient *http.Client, apiURL string) *ClientHTTP {
	return &ClientHTTP{
		httpClient: httpClient,
		apiURL:     apiURL,
	}
}

func (c *ClientHTTP) FindMarketByID(ctx context.Context, marketID int) (*Market, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/markets/%d", c.apiURL, marketID),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Market *Market `json:"market"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Market, nil
}

func (c *ClientHTTP) FindMarketsByIDs(ctx context.Context, marketIDs []int) ([]Market, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/markets", c.apiURL),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	queryParams := make(url.Values)
	queryParams["ids[]"] = make([]string, len(marketIDs))
	for i := range marketIDs {
		queryParams.Add("ids[]", strconv.Itoa(marketIDs[i]))
	}

	httpReq.URL.RawQuery = queryParams.Encode()

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Markets []Market `json:"markets"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Markets, nil
}

func (c *ClientHTTP) FindLocalizedMarketsByIDs(ctx context.Context, locale Locale, marketIDs []int) ([]MarketLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/localized/markets/%s", c.apiURL, locale),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	queryParams := make(url.Values)
	queryParams["ids[]"] = make([]string, len(marketIDs))
	for i := range marketIDs {
		queryParams.Add("ids[]", strconv.Itoa(marketIDs[i]))
	}

	httpReq.URL.RawQuery = queryParams.Encode()

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Markets []MarketLocalized `json:"markets"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Markets, nil
}

func (c *ClientHTTP) FindMarketsByFilters(ctx context.Context, filters *MarketFilters) ([]Market, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/markets", c.apiURL),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = filters.ToQueryParams().Encode()

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Markets []Market `json:"markets"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Markets, nil
}

func (c *ClientHTTP) FindLocalizedMarketsByFilters(ctx context.Context, locale Locale, filters *MarketFilters) ([]MarketLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/localized/markets/%s", c.apiURL, locale),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = filters.ToQueryParams().Encode()

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Markets []MarketLocalized `json:"markets"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Markets, nil
}

func (c *ClientHTTP) FindSportByID(ctx context.Context, sportID string) (*Sport, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/sports/%s", c.apiURL, sportID),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Sport *Sport `json:"sport"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Sport, nil
}

func (c *ClientHTTP) FindSportsByFilters(ctx context.Context, filters *SportFilters) ([]Sport, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, c.apiURL+"/sports", http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = filters.ToQueryParams().Encode()

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Sports []Sport `json:"sports"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Sports, nil
}

// GetAllLocalizedSports finds sports by ids (optional argument, by default all sports are returned)
// and translates them to a given locale.
func (c *ClientHTTP) GetAllLocalizedSports(ctx context.Context, locale Locale, ids ...string) ([]SportLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/localized/sports/%s", c.apiURL, locale),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	if len(ids) > 0 {
		queryParams := url.Values{"ids[]": ids}
		httpReq.URL.RawQuery = queryParams.Encode()
	}

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Sports []SportLocalized `json:"sports"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Sports, nil
}

func (c *ClientHTTP) FindLocalizedTournamentByID(ctx context.Context, locale Locale, tournamentID string) (*TournamentLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/tournaments/localized/%s/%s", c.apiURL, tournamentID, locale),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Tournament *TournamentLocalized `json:"tournament"`
		Error      *apiError            `json:"error"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if resp.Error != nil {
		return nil, c.convertApiError(resp.Error)
	}

	return resp.Tournament, nil
}

func (c *ClientHTTP) FindLocalizedTournamentsByIDs(
	ctx context.Context,
	locale Locale,
	tournamentIDs []string,
) ([]TournamentLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/tournaments/localized/%s", c.apiURL, locale),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = url.Values{"ids[]": tournamentIDs}.Encode()

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Tournament []TournamentLocalized `json:"tournaments"`
		Error      *apiError             `json:"error"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if resp.Error != nil {
		return nil, c.convertApiError(resp.Error)
	}

	return resp.Tournament, nil
}

func (c *ClientHTTP) FindLocalizedPlayerByID(ctx context.Context, locale Locale, playerID string) (*PlayerLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/players/localized/%s/%s", c.apiURL, playerID, locale),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Player PlayerLocalized `json:"player"`
		Error  *apiError       `json:"error"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if resp.Error != nil {
		return nil, c.convertApiError(resp.Error)
	}

	return &resp.Player, nil
}

func (c *ClientHTTP) FindLocalizedPlayersByIDs(ctx context.Context, locale Locale, playerIDs []string) ([]PlayerLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/players/localized/%s", c.apiURL, locale),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = url.Values{"ids[]": playerIDs}.Encode()

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Players []PlayerLocalized `json:"players"`
		Error   *apiError         `json:"error"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if resp.Error != nil {
		return nil, c.convertApiError(resp.Error)
	}

	return resp.Players, nil
}

func (c *ClientHTTP) FindLocalizedTeamByID(ctx context.Context, locale Locale, teamID string) (*TeamLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/teams/localized/%s/%s", c.apiURL, teamID, locale),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Team  TeamLocalized `json:"team"`
		Error *apiError     `json:"error"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if resp.Error != nil {
		return nil, c.convertApiError(resp.Error)
	}

	return &resp.Team, nil
}

func (c *ClientHTTP) FindLocalizedTeamsByIDs(ctx context.Context, locale Locale, teamIDs []string) ([]TeamLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/teams/localized/%s", c.apiURL, locale),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = url.Values{"ids[]": teamIDs}.Encode()

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Teams []TeamLocalized `json:"teams"`
		Error *apiError       `json:"error"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if resp.Error != nil {
		return nil, c.convertApiError(resp.Error)
	}

	return resp.Teams, nil
}

func (c *ClientHTTP) FindLocalizedOrganizationsByIDs(ctx context.Context, locale Locale, organizationIDs []string) ([]OrganizationLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/organizations/localized/%s/by-ids", c.apiURL, locale),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = url.Values{"ids[]": organizationIDs}.Encode()

	rawBody, err := c.doAPIRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Organizations []OrganizationLocalized `json:"organizations"`
		Error         *apiError               `json:"error"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if resp.Error != nil {
		return nil, c.convertApiError(resp.Error)
	}

	return resp.Organizations, nil
}

func (c *ClientHTTP) doAPIRequest(httpReq *http.Request) ([]byte, error) {
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	rawBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	if httpResp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("%w, response body: %q", ErrInvalidCertificate, string(rawBody))
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %s, response body: %q", httpResp.Status, string(rawBody))
	}

	return rawBody, nil
}

func (c *ClientHTTP) convertApiError(err *apiError) error {
	switch err.Code {
	case errCodeTournamentNotFound:
		return fmt.Errorf("tournament: %w", ErrNotFound)
	case errCodePlayerNotFound:
		return fmt.Errorf("player: %w", ErrNotFound)
	case errCodeTeamNotFound:
		return fmt.Errorf("team: %w", ErrNotFound)
	default:
		return fmt.Errorf("%w, extra data: %v", ErrUnknown, err.Data)
	}
}
