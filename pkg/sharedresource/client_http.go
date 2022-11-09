package sharedresource

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type ClientHTTP struct {
	httpClient        *http.Client
	sharedResourceURL string
}

var _ Client = (*ClientHTTP)(nil)

func NewClientHTTP(httpClient *http.Client, sharedResourceURL string) *ClientHTTP {
	return &ClientHTTP{
		httpClient:        httpClient,
		sharedResourceURL: sharedResourceURL,
	}
}

func (c *ClientHTTP) FindMarketByID(ctx context.Context, marketID int) (*Market, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.sharedResourceURL+"/markets/"+strconv.Itoa(marketID),
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
		return nil, fmt.Errorf("status code: %s, response body: %s", httpResp.Status, string(rawBody))
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
		c.sharedResourceURL+"/markets",
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
		return nil, fmt.Errorf("status code: %s, response body: %s", httpResp.Status, string(rawBody))
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
		c.sharedResourceURL+"/localized/markets/"+locale.String(),
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
		return nil, fmt.Errorf("status code: %s, response body: %s", httpResp.Status, string(rawBody))
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
		c.sharedResourceURL+"/markets",
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
		return nil, fmt.Errorf("status code: %s, response body: %s", httpResp.Status, string(rawBody))
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
		c.sharedResourceURL+"/localized/markets/"+locale.String(),
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
		return nil, fmt.Errorf("status code: %s, response body: %s", httpResp.Status, string(rawBody))
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
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, c.sharedResourceURL+"/sports/"+sportID, http.NoBody)
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
		return nil, fmt.Errorf("status code: %s, response body: %s", httpResp.Status, string(rawBody))
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
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, c.sharedResourceURL+"/sports", http.NoBody)
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
		return nil, fmt.Errorf("status code: %s, response body: %s", httpResp.Status, string(rawBody))
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
		c.sharedResourceURL+"/localized/sports/"+locale.String(),
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
		return nil, fmt.Errorf("status code: %s, response body: %s", httpResp.Status, string(rawBody))
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
