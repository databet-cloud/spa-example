package mts

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/bytedance/sonic"

	apierror "github.com/databet-cloud/databet-go-sdk/pkg/error"
	"github.com/databet-cloud/databet-go-sdk/pkg/restriction"
)

type ClientHTTP struct {
	httpClient *http.Client
	mtsURL     string
}

var _ Client = (*ClientHTTP)(nil)

func NewClientHTTP(httpClient *http.Client, mtsURL string) *ClientHTTP {
	return &ClientHTTP{
		httpClient: httpClient,
		mtsURL:     mtsURL,
	}
}

func (c *ClientHTTP) PlaceBet(ctx context.Context, req *PlaceBetRequest) (*Bet, []restriction.Restriction, error) {
	rawBody, err := sonic.Marshal(req)
	if err != nil {
		return nil, nil, fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/bets", c.mtsURL),
		bytes.NewReader(rawBody),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("create http request: %w", err)
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode == http.StatusForbidden {
		return nil, nil, apierror.NewUser(ErrCodeAccessForIPDenied, nil)
	}

	var resp placeBetResponse

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, nil, fmt.Errorf("unmarshal response: %w, status code: %s", err, httpResp.Status)
	}

	if resp.Error != nil {
		return nil, nil, resp.Error
	}

	return resp.Bet, resp.Restrictions, nil
}

func (c *ClientHTTP) DeclineBet(ctx context.Context, req *DeclineBetRequest) error {
	rawBody, err := sonic.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/bets", c.mtsURL),
		bytes.NewReader(rawBody),
	)
	if err != nil {
		return fmt.Errorf("create http request: %w", err)
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode == http.StatusForbidden {
		return apierror.NewUser(ErrCodeAccessForIPDenied, nil)
	}

	if httpResp.StatusCode == http.StatusNoContent {
		return nil
	}

	var resp errorResponse

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return fmt.Errorf("unmarshal response: %w, status code: %s", err, httpResp.Status)
	}

	if resp.Error != nil {
		return resp.Error
	}

	return nil
}

func (c *ClientHTTP) CalculateCashOut(ctx context.Context, req *CalculateCashOutRequest) (*CashOutAmount, []restriction.Restriction, error) {
	rawBody, err := sonic.Marshal(req)
	if err != nil {
		return nil, nil, fmt.Errorf("marshal request: %s", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/bets/%s/cash-out-orders/calculate", c.mtsURL, req.BetID),
		bytes.NewReader(rawBody),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("create http request: %w", err)
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode == http.StatusForbidden {
		return nil, nil, apierror.NewUser(ErrCodeAccessForIPDenied, nil)
	}

	var resp calculateCashOutResponse

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, nil, fmt.Errorf("unmarshal response: %w, status code: %s", err, httpResp.Status)
	}

	if resp.Error != nil {
		return nil, nil, resp.Error
	}

	return resp.Amount, resp.Restrictions, nil
}

func (c *ClientHTTP) PlaceCashOutOrder(
	ctx context.Context,
	req *PlaceCashOutOrderRequest,
) (*Bet, *CashOutOrder, []restriction.Restriction, error) {
	rawBody, err := sonic.Marshal(req)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("marshal request: %s", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/bets/%s/cash-out-orders/place", c.mtsURL, req.BetID),
		bytes.NewReader(rawBody),
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("create http request: %w", err)
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode == http.StatusForbidden {
		return nil, nil, nil, apierror.NewUser(ErrCodeAccessForIPDenied, nil)
	}

	var resp placeCashOutOrderResponse

	rawBody, err = io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("read response body: %w", err)
	}

	err = sonic.Unmarshal(rawBody, &resp)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("unmarshal response: %w, status code: %s", err, httpResp.Status)
	}

	if resp.Error != nil {
		return nil, nil, nil, resp.Error
	}

	return resp.Bet, resp.CashOutOrder, resp.Restrictions, nil
}

func (c *ClientHTTP) CancelCashOutOrder(ctx context.Context, req *CancelCashOutOrderRequest) error {
	rawBody, err := sonic.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal request: %s", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPatch,
		fmt.Sprintf("%s/bets/%s/cash-out-orders/%s/cancel", c.mtsURL, req.BetID, req.CashOutOrderID),
		bytes.NewReader(rawBody),
	)
	if err != nil {
		return fmt.Errorf("create http request: %w", err)
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode == http.StatusForbidden {
		return apierror.NewUser(ErrCodeAccessForIPDenied, nil)
	}

	if httpResp.StatusCode == http.StatusNoContent {
		return nil
	}

	var resp errorResponse

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return fmt.Errorf("unmarshal response: %w, status code: %s", err, httpResp.Status)
	}

	if resp.Error != nil {
		return resp.Error
	}

	return nil
}

func (c *ClientHTTP) GetRestrictions(ctx context.Context, req *GetRestrictionsRequest) ([]restriction.Restriction, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/restrictions", c.mtsURL), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("create http request: %w", err)
	}

	query := httpReq.URL.Query()

	query.Set("player_id", req.PlayerID)
	query.Set("bet_type", strconv.Itoa(req.BetType))
	query.Set("currency_code", req.CurrencyCode)
	query.Set("odd_accept_strategy", strconv.Itoa(int(req.OddAcceptStrategy)))

	for i, sel := range req.Selections {
		query.Set(fmt.Sprintf("selections[%d][sport_event_id]", i), sel.SportEventID)
		query.Set(fmt.Sprintf("selections[%d][market_id]", i), sel.MarketID)
		query.Set(fmt.Sprintf("selections[%d][odd_id]", i), sel.OddID)
		query.Set(fmt.Sprintf("selections[%d][value]", i), sel.Value)
	}

	for i, systemSize := range req.SystemSizes {
		query.Set(fmt.Sprintf("system_sizes[%d]", i), strconv.Itoa(systemSize))
	}

	httpReq.URL.RawQuery = query.Encode()

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode == http.StatusForbidden {
		return nil, apierror.NewUser(ErrCodeAccessForIPDenied, nil)
	}

	var resp getRestrictionsResponse

	rawBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	err = sonic.Unmarshal(rawBody, &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w, status code: %s", err, httpResp.Status)
	}

	if resp.Error != nil {
		return nil, resp.Error
	}

	return resp.Restrictions, nil
}

func (c *ClientHTTP) GetMaxBet(ctx context.Context, req *GetMaxBetRequest) (string, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/max-bet", c.mtsURL), http.NoBody)
	if err != nil {
		return "", fmt.Errorf("create http request: %w", err)
	}

	query := httpReq.URL.Query()

	query.Set("player_id", req.PlayerID)

	for i, sel := range req.Selections {
		query.Set(fmt.Sprintf("selections[%d][sport_event_id]", i), sel.SportEventID)
		query.Set(fmt.Sprintf("selections[%d][market_id]", i), sel.MarketID)
		query.Set(fmt.Sprintf("selections[%d][odd_id]", i), sel.OddID)
		query.Set(fmt.Sprintf("selections[%d][value]", i), sel.Value)
	}

	httpReq.URL.RawQuery = query.Encode()

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("do http request: %w", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode == http.StatusForbidden {
		return "", apierror.NewUser(ErrCodeAccessForIPDenied, nil)
	}

	var resp getMaxBetResponse

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return "", fmt.Errorf("unmarshal response: %w, status code: %s", err, httpResp.Status)
	}

	if resp.Error != nil {
		return "", resp.Error
	}

	return resp.MaxBet, nil
}
