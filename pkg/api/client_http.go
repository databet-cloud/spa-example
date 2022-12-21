package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bytedance/sonic"

	"github.com/databet-cloud/databet-go-sdk/pkg/simdutil"
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
		c.apiURL+"/markets/"+strconv.Itoa(marketID),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		rawBody, _ := io.ReadAll(httpResp.Body)
		return nil, fmt.Errorf("status code: %s, response body: %q", httpResp.Status, string(rawBody))
	}

	var resp struct {
		Market *Market `json:"market"`
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Market, nil
}

func (c *ClientHTTP) FindMarketsByIDs(ctx context.Context, marketIDs []int) ([]Market, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/markets",
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

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		rawBody, _ := io.ReadAll(httpResp.Body)
		return nil, fmt.Errorf("status code: %s, response body: %q", httpResp.Status, string(rawBody))
	}

	var resp struct {
		Markets []Market `json:"markets"`
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Markets, nil
}

func (c *ClientHTTP) FindLocalizedMarketsByIDs(ctx context.Context, locale Locale, marketIDs []int) ([]MarketLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/localized/markets/"+locale.String(),
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

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		rawBody, _ := io.ReadAll(httpResp.Body)
		return nil, fmt.Errorf("status code: %s, response body: %q", httpResp.Status, string(rawBody))
	}

	var resp struct {
		Markets []MarketLocalized `json:"markets"`
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Markets, nil
}

func (c *ClientHTTP) FindMarketsByFilters(ctx context.Context, filters *MarketFilters) ([]Market, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/markets",
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = filters.ToQueryParams().Encode()

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		rawBody, _ := io.ReadAll(httpResp.Body)
		return nil, fmt.Errorf("status code: %s, response body: %q", httpResp.Status, string(rawBody))
	}

	var resp struct {
		Markets []Market `json:"markets"`
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Markets, nil
}

func (c *ClientHTTP) FindLocalizedMarketsByFilters(ctx context.Context, locale Locale, filters *MarketFilters) ([]MarketLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/localized/markets/"+locale.String(),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = filters.ToQueryParams().Encode()

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		rawBody, _ := io.ReadAll(httpResp.Body)
		return nil, fmt.Errorf("status code: %s, response body: %q", httpResp.Status, string(rawBody))
	}

	var resp struct {
		Markets []MarketLocalized `json:"markets"`
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Markets, nil
}

func (c *ClientHTTP) FindSportByID(ctx context.Context, sportID string) (*Sport, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, c.apiURL+"/sports/"+sportID, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		rawBody, _ := io.ReadAll(httpResp.Body)
		return nil, fmt.Errorf("status code: %s, response body: %q", httpResp.Status, string(rawBody))
	}

	var resp struct {
		Sport *Sport `json:"sport"`
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
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

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		rawBody, _ := io.ReadAll(httpResp.Body)
		return nil, fmt.Errorf("status code: %s, response body: %q", httpResp.Status, string(rawBody))
	}

	var resp struct {
		Sports []Sport `json:"sports"`
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
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
		c.apiURL+"/localized/sports/"+locale.String(),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	if len(ids) > 0 {
		queryParams := url.Values{"ids[]": ids}
		httpReq.URL.RawQuery = queryParams.Encode()
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		rawBody, _ := io.ReadAll(httpResp.Body)
		return nil, fmt.Errorf("status code: %s, response body: %q", httpResp.Status, string(rawBody))
	}

	var resp struct {
		Sports []SportLocalized `json:"sports"`
	}

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return resp.Sports, nil
}

func (c *ClientHTTP) FindTournamentByID(ctx context.Context, tournamentID string) (*Tournament, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/tournaments/"+tournamentID,
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

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

	var resp struct {
		Tournament *Tournament `json:"tournament"`
		Error      *apiError   `json:"error"`
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

func (c *ClientHTTP) FindTournamentsByIDs(ctx context.Context, tournamentIDs []string) ([]Tournament, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/tournaments/by-ids",
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = url.Values{"ids[]": tournamentIDs}.Encode()

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

	var resp struct {
		Tournament []Tournament `json:"tournaments"`
		Error      *apiError    `json:"error"`
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

func (c *ClientHTTP) SearchTournaments(ctx context.Context, req *SearchTournamentsRequest) (*SearchTournamentsResponse, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/search/tournaments",
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	if query := req.ToQueryParams().Encode(); query != "" {
		httpReq.URL.RawQuery = query
	}

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

	var resp struct {
		SearchTournamentsResponse
		Error *apiError `json:"error"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if resp.Error != nil {
		return nil, c.convertApiError(resp.Error)
	}

	return &resp.SearchTournamentsResponse, nil
}

func (c *ClientHTTP) FindLocalizedTournamentByID(ctx context.Context, locale Locale, tournamentID string) (*TournamentLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/tournaments/localized/"+tournamentID+"/"+locale.String(),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

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
		c.apiURL+"/tournaments/localized/"+locale.String(),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = url.Values{"ids[]": tournamentIDs}.Encode()

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

func (c *ClientHTTP) SearchLocalizedTournaments(
	ctx context.Context,
	locale Locale,
	req *SearchTournamentsRequest,
) (*SearchLocalizedTournamentsResponse, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/search/tournaments/localized/"+locale.String(),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	if query := req.ToQueryParams().Encode(); query != "" {
		httpReq.URL.RawQuery = query
	}

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

	var resp struct {
		SearchLocalizedTournamentsResponse
		Error *apiError `json:"error"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if resp.Error != nil {
		return nil, c.convertApiError(resp.Error)
	}

	return &resp.SearchLocalizedTournamentsResponse, nil
}

func (c *ClientHTTP) FindPlayerByID(ctx context.Context, playerID string) (*Player, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/players/"+playerID,
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

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

	var resp struct {
		Player Player    `json:"player"`
		Error  *apiError `json:"error"`
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

func (c *ClientHTTP) FindPlayersByIDs(ctx context.Context, playerIDs []string) ([]Player, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/players/by-ids",
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = url.Values{"ids[]": playerIDs}.Encode()

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

	var resp struct {
		Players []Player  `json:"players"`
		Error   *apiError `json:"error"`
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

func (c *ClientHTTP) SearchPlayers(ctx context.Context, req *SearchPlayersRequest) (*SearchPlayersResponse, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/search/players",
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = req.ToQueryParams().Encode()

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

	var resp struct {
		SearchPlayersResponse
		Error *apiError `json:"error"`
	}

	err = sonic.UnmarshalString(simdutil.UnsafeStrFromBytes(rawBody), &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if resp.Error != nil {
		return nil, c.convertApiError(resp.Error)
	}

	return &resp.SearchPlayersResponse, nil
}

func (c *ClientHTTP) FindLocalizedPlayerByID(ctx context.Context, locale Locale, playerID string) (*PlayerLocalized, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/players/localized/"+playerID+"/"+locale.String(),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

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
		c.apiURL+"/players/localized/"+locale.String(),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	httpReq.URL.RawQuery = url.Values{"ids[]": playerIDs}.Encode()

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

func (c *ClientHTTP) convertApiError(err *apiError) error {
	switch err.Code {
	case errCodeTournamentNotFound:
		return fmt.Errorf("tournament: %w", ErrNotFound)
	case errCodePlayerNotFound:
		return fmt.Errorf("player: %w", ErrNotFound)
	default:
		return fmt.Errorf("%w, extra data: %v", ErrUnknown, err.Data)
	}
}
