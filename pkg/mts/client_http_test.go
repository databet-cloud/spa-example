package mts_test

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/databet-cloud/databet-go-sdk/pkg/mts"
	"github.com/databet-cloud/databet-go-sdk/pkg/mts/mocks"
)

const testURL = "http://mts-test"

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

type ClientHTTPTestSuite struct {
	suite.Suite

	client    *http.Client
	mtsClient *mts.ClientHTTP
}

func TestClientHTTP(t *testing.T) {
	suite.Run(t, new(ClientHTTPTestSuite))
}

func (s *ClientHTTPTestSuite) SetupSuite() {
	s.client = http.DefaultClient
	s.mtsClient = mts.NewClientHTTP(s.client, testURL)
}

//go:embed testdata/place-bet/request.json
var rawPlaceBetReq []byte

func (s *ClientHTTPTestSuite) TestPlaceBet() {
	testCases := []struct {
		name        string
		input       []byte
		httpResp    *http.Response
		expected    *mts.PlaceBetResponse
		expectedErr error
	}{
		{
			name:     "succeed",
			input:    rawPlaceBetReq,
			httpResp: s.makeResponse(http.StatusOK, "place-bet/response-success.json"),
			expected: &mts.PlaceBetResponse{
				Bet: &mts.Bet{
					ID:          "bh5cppigcvvoqss2htfg",
					BookmakerID: "1",
					Status:      mts.BetStatus{Code: 0, Reason: ""},
					Type:        mts.BetType{Code: 0, Size: []int(nil)},
					Stake:       mts.MultiMoney{Base: mts.Money{Value: "3.554952", CurrencyCode: "USD"}, Origin: mts.Money{Value: "3.000000", CurrencyCode: "EUR"}},
					Refund:      mts.MultiMoney{Base: mts.Money{Value: "3.554952", CurrencyCode: "USD"}, Origin: mts.Money{Value: "3.000000", CurrencyCode: "EUR"}},
					RefundBase:  mts.MultiMoney{Base: mts.Money{Value: "3.554952", CurrencyCode: "USD"}, Origin: mts.Money{Value: "3.000000", CurrencyCode: "EUR"}},
					PlayerInfo:  mts.PlayerInfo{PlayerID: "", RiskGroupID: "", ClientIP: net.IP(nil), CountryCode: ""},
					CreatedAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					Selections: []*mts.Selection{
						{SportEventID: "1", MarketID: "1", OddID: "1", Value: "1.5", Marge: "0.06"},
					},
				},
			},
		},
		{
			name:     "got restrictions",
			input:    rawPlaceBetReq,
			httpResp: s.makeResponse(http.StatusBadRequest, "response-restrictions.json"),
			expected: &mts.PlaceBetResponse{
				Restrictions: []mts.Restriction{
					{Type: "test_restriction"},
				},
			},
		},
		{
			name:        "bad request error",
			input:       rawPlaceBetReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-bad-request-response.json"),
			expectedErr: mts.ErrBadRequest,
		},
		{
			name:        "certificate error",
			input:       rawPlaceBetReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-certificate-response.json"),
			expectedErr: mts.ErrInvalidCertificate,
		},
		{
			name:        "unknown api error",
			input:       rawPlaceBetReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-response.json"),
			expectedErr: mts.ErrUnknown,
		},
		{
			name:        "ip forbidden",
			input:       rawPlaceBetReq,
			httpResp:    s.makeResponse(http.StatusForbidden, ""),
			expectedErr: mts.ErrAccessForIPDenied,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			httpReq := s.makeRequest(http.MethodPost, fmt.Sprintf("%s/bets", testURL), bytes.NewReader(rawPlaceBetReq))
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(httpReq)).Return(tc.httpResp, nil)

			var req *mts.PlaceBetRequest
			s.NoError(json.Unmarshal(tc.input, &req))

			resp, err := s.mtsClient.PlaceBet(context.Background(), req)

			s.Equal(tc.expected, resp)
			s.ErrorIs(err, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestDeclineBet() {
	testCases := []struct {
		name        string
		input       []byte
		httpResp    *http.Response
		expectedErr error
	}{
		{
			name:     "succeed",
			input:    []byte(`{"bet_id": "bet1", "reason": "test"}`),
			httpResp: s.makeResponse(http.StatusNoContent, ""),
		},
		{
			name:        "bad request error",
			input:       []byte(`{"bet_id": "bet1", "reason": "test"}`),
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-bad-request-response.json"),
			expectedErr: mts.ErrBadRequest,
		},
		{
			name:        "certificate error",
			input:       []byte(`{"bet_id": "bet1", "reason": "test"}`),
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-certificate-response.json"),
			expectedErr: mts.ErrInvalidCertificate,
		},
		{
			name:        "unknown api error",
			input:       []byte(`{"bet_id": "bet1", "reason": "test"}`),
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-response.json"),
			expectedErr: mts.ErrUnknown,
		},
		{
			name:        "ip forbidden",
			input:       []byte(`{"bet_id": "bet1", "reason": "test"}`),
			httpResp:    s.makeResponse(http.StatusForbidden, ""),
			expectedErr: mts.ErrAccessForIPDenied,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			httpReq := s.makeRequest(http.MethodDelete, fmt.Sprintf("%s/bets", testURL), bytes.NewReader(tc.input))
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(httpReq)).Return(tc.httpResp, nil)

			var req *mts.DeclineBetRequest
			s.NoError(json.Unmarshal(tc.input, &req))

			err := s.mtsClient.DeclineBet(context.Background(), req)

			s.ErrorIs(err, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestCalculateCashOut() {
	defaultReq := &mts.CalculateCashOutRequest{
		BetID:  "bet1",
		Amount: mts.Money{Value: "100", CurrencyCode: "USD"},
		Selections: []mts.CashOutSelection{
			{
				SportEventID: "sportEvent1",
				MarketID:     "market1",
				OddID:        "odd1",
				Value:        "1.42",
			},
		},
	}

	testCases := []struct {
		name        string
		input       *mts.CalculateCashOutRequest
		httpResp    *http.Response
		expected    *mts.CalculateCashOutResponse
		expectedErr error
	}{
		{
			name:     "succeed",
			input:    defaultReq,
			httpResp: s.makeResponse(http.StatusOK, "calculate-cash-out/response-success.json"),
			expected: &mts.CalculateCashOutResponse{
				Amount: &mts.CashOutAmount{
					RefundAmount:    "10.01",
					MinAmount:       "10",
					MinRefundAmount: "10",
					MaxAmount:       "20.01",
					MaxRefundAmount: "20.01",
					Ranges: []mts.CashOutRange{
						{FromAmount: "10.1", ToAmount: "20.1", RefundRatio: "1.5"},
					},
				},
				Restrictions: nil,
			},
		},
		{
			name:     "got restrictions",
			input:    defaultReq,
			httpResp: s.makeResponse(http.StatusBadRequest, "response-restrictions.json"),
			expected: &mts.CalculateCashOutResponse{
				Restrictions: []mts.Restriction{{Type: "test_restriction"}},
			},
		},
		{
			name:        "bad request error",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-bad-request-response.json"),
			expectedErr: mts.ErrBadRequest,
		},
		{
			name:        "certificate error",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-certificate-response.json"),
			expectedErr: mts.ErrInvalidCertificate,
		},
		{
			name:        "unknown api error",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-response.json"),
			expectedErr: mts.ErrUnknown,
		},
		{
			name:        "ip forbidden",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusForbidden, ""),
			expectedErr: mts.ErrAccessForIPDenied,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			reqBody, err := json.Marshal(tc.input)
			s.NoError(err)

			httpReq := s.makeRequest(
				http.MethodPost,
				fmt.Sprintf("%s/bets/%s/cash-out-orders/calculate", testURL, tc.input.BetID),
				bytes.NewReader(reqBody),
			)
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(httpReq)).Return(tc.httpResp, nil)

			resp, err := s.mtsClient.CalculateCashOut(context.Background(), tc.input)

			s.Equal(tc.expected, resp)
			s.ErrorIs(err, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestPlaceCashOutOrder() {
	defaultReq := &mts.PlaceCashOutOrderRequest{
		ID:           "order1",
		BetID:        "bet1",
		Amount:       mts.Money{Value: "100", CurrencyCode: "EUR"},
		RefundAmount: mts.Money{Value: "20", CurrencyCode: "EUR"},
		CreatedAt:    time.Unix(1, 0).UTC().String(),
		Selections: []mts.CashOutSelection{
			{
				SportEventID: "sportEvent1",
				MarketID:     "market1",
				OddID:        "odd1",
				Value:        "1.42",
			},
		},
	}

	testCases := []struct {
		name        string
		input       *mts.PlaceCashOutOrderRequest
		httpResp    *http.Response
		expected    *mts.PlaceCashOutOrderResponse
		expectedErr error
	}{
		{
			name:     "succeed",
			input:    defaultReq,
			httpResp: s.makeResponse(http.StatusOK, "place-cash-out-order/response-success.json"),
			expected: &mts.PlaceCashOutOrderResponse{
				Bet: &mts.Bet{
					ID:          "bh5cppigcvvoqss2htfg",
					BookmakerID: "1",
					Status:      mts.BetStatus{Code: 0, Reason: ""},
					Type:        mts.BetType{Code: 0, Size: []int(nil)},
					Stake:       mts.MultiMoney{Base: mts.Money{Value: "3.554952", CurrencyCode: "USD"}, Origin: mts.Money{Value: "3.000000", CurrencyCode: "EUR"}},
					Refund:      mts.MultiMoney{Base: mts.Money{Value: "3.554952", CurrencyCode: "USD"}, Origin: mts.Money{Value: "3.000000", CurrencyCode: "EUR"}},
					RefundBase:  mts.MultiMoney{Base: mts.Money{Value: "3.554952", CurrencyCode: "USD"}, Origin: mts.Money{Value: "3.000000", CurrencyCode: "EUR"}},
					PlayerInfo:  mts.PlayerInfo{PlayerID: "", RiskGroupID: "", ClientIP: net.IP(nil), CountryCode: ""},
					CreatedAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					Selections: []*mts.Selection{
						{SportEventID: "1", MarketID: "1", OddID: "1", Value: "1.5", Marge: "0.06"},
					},
				},
				CashOutOrder: &mts.CashOutOrder{
					ID: "order1",
					Amount: mts.MultiMoney{
						Base:   mts.Money{Value: "100", CurrencyCode: "USD"},
						Origin: mts.Money{Value: "100", CurrencyCode: "USD"},
					},
					RefundAmount: mts.MultiMoney{
						Base:   mts.Money{Value: "113.56", CurrencyCode: "USD"},
						Origin: mts.Money{Value: "113.56", CurrencyCode: "USD"},
					},
					CreatedAt: mustParseTime(s.T(), "2006-01-02T15:04:05+07:00"),
					Selections: []mts.CashOutOrderSelection{
						{OddID: "1", MarketID: "1", SportEventID: "1", Value: "8"},
					},
				},
			},
		},
		{
			name:     "got restrictions",
			input:    defaultReq,
			httpResp: s.makeResponse(http.StatusBadRequest, "response-restrictions.json"),
			expected: &mts.PlaceCashOutOrderResponse{
				Restrictions: []mts.Restriction{{Type: "test_restriction"}},
			},
		},
		{
			name:        "bad request error",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-bad-request-response.json"),
			expectedErr: mts.ErrBadRequest,
		},
		{
			name:        "certificate error",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-certificate-response.json"),
			expectedErr: mts.ErrInvalidCertificate,
		},
		{
			name:        "bad request error",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-bad-request-response.json"),
			expectedErr: mts.ErrBadRequest,
		},
		{
			name:        "certificate error",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-certificate-response.json"),
			expectedErr: mts.ErrInvalidCertificate,
		},
		{
			name:        "unknown api error",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-response.json"),
			expectedErr: mts.ErrUnknown,
		},
		{
			name:        "ip forbidden",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusForbidden, ""),
			expectedErr: mts.ErrAccessForIPDenied,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			reqBody, err := json.Marshal(tc.input)
			s.NoError(err)

			httpReq := s.makeRequest(
				http.MethodPost,
				fmt.Sprintf("%s/bets/%s/cash-out-orders/place", testURL, tc.input.BetID),
				bytes.NewReader(reqBody),
			)
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(httpReq)).Return(tc.httpResp, nil)

			resp, err := s.mtsClient.PlaceCashOutOrder(context.Background(), tc.input)

			s.Equal(tc.expected, resp)
			s.ErrorIs(err, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestCancelCashOutOrder() {
	defaultReq := &mts.CancelCashOutOrderRequest{
		BetID:          "bet1",
		CashOutOrderID: "order1",
		Context: &mts.CashOutContext{
			Restrictions: []mts.Restriction{
				{Type: "test_restriction"},
			},
		},
	}

	testCases := []struct {
		name        string
		input       *mts.CancelCashOutOrderRequest
		httpResp    *http.Response
		expectedErr error
	}{
		{
			name:     "succeed",
			input:    defaultReq,
			httpResp: s.makeResponse(http.StatusOK, "place-cash-out-order/response-success.json"),
		},
		{
			name:        "bad request error",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-bad-request-response.json"),
			expectedErr: mts.ErrBadRequest,
		},
		{
			name:        "certificate error",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-certificate-response.json"),
			expectedErr: mts.ErrInvalidCertificate,
		},
		{
			name:        "unknown api error",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-response.json"),
			expectedErr: mts.ErrUnknown,
		},
		{
			name:        "ip forbidden",
			input:       defaultReq,
			httpResp:    s.makeResponse(http.StatusForbidden, ""),
			expectedErr: mts.ErrAccessForIPDenied,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctrl := gomock.NewController(s.T())
			defer ctrl.Finish()

			roundTripper := mocks.NewRoundTripper(ctrl)
			s.client.Transport = roundTripper

			reqBody, err := json.Marshal(tc.input)
			s.NoError(err)

			httpReq := s.makeRequest(
				http.MethodPatch,
				fmt.Sprintf("%s/bets/%s/cash-out-orders/%s/cancel", testURL, tc.input.BetID, tc.input.CashOutOrderID),
				bytes.NewReader(reqBody),
			)
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(httpReq)).Return(tc.httpResp, nil)

			err = s.mtsClient.CancelCashOutOrder(context.Background(), tc.input)

			s.ErrorIs(err, tc.expectedErr)
		})
	}
}

func (s *ClientHTTPTestSuite) TestGetRestrictions() {
	defaultReq := &mts.GetRestrictionsRequest{
		PlayerID: "player1",
		BetType:  1,
		Selections: []mts.RestrictionsSelection{{
			SportEventID: "sportEvent1",
			MarketID:     "market1",
			OddID:        "odd1",
			Value:        "1.42",
		}, {
			SportEventID: "sportEvent2",
			MarketID:     "market2",
			OddID:        "odd2",
			Value:        "2.42",
		}},
		SystemSizes:       []int{1},
		CurrencyCode:      "USD",
		OddAcceptStrategy: mts.AcceptStrategyAlwaysAllowed,
	}

	defaultQueryParams := url.Values{
		"player_id":                     []string{"player1"},
		"bet_type":                      []string{"1"},
		"currency_code":                 []string{"USD"},
		"odd_accept_strategy":           []string{"4"},
		"system_sizes[0]":               []string{"1"},
		"selections[0][sport_event_id]": []string{"sportEvent1"},
		"selections[0][market_id]":      []string{"market1"},
		"selections[0][odd_id]":         []string{"odd1"},
		"selections[0][value]":          []string{"1.42"},
		"selections[1][sport_event_id]": []string{"sportEvent2"},
		"selections[1][market_id]":      []string{"market2"},
		"selections[1][odd_id]":         []string{"odd2"},
		"selections[1][value]":          []string{"2.42"},
	}

	testCases := []struct {
		name        string
		input       *mts.GetRestrictionsRequest
		queryParams url.Values
		httpResp    *http.Response
		expected    []mts.Restriction
		expectedErr error
	}{
		{
			name:        "succeed",
			input:       defaultReq,
			queryParams: defaultQueryParams,
			httpResp:    s.makeResponse(http.StatusOK, "restrictions/response-success.json"),
			expected: []mts.Restriction{
				{
					Type: mts.RestrictionMaxBet,
					Context: map[string]any{
						mts.CtxKeyMaxBet: "240.96",
					},
				},
				{
					Type: mts.RestrictionMarketStatus,
					Context: map[string]any{
						mts.CtxKeyMarketID:     "20",
						mts.CtxKeySportEventID: "85d130b1-6c8f-4170-bb4f-f80343b55c01",
						mts.CtxKeyStatus:       float64(3),
					},
				},
				{
					Type: mts.RestrictionOddStatus,
					Context: map[string]any{
						mts.CtxKeyIsActive:     true,
						mts.CtxKeyMarketID:     "20",
						mts.CtxKeyOddID:        "1",
						mts.CtxKeySportEventID: "85d130b1-6c8f-4170-bb4f-f80343b55c01",
						mts.CtxKeyStatus:       float64(1),
					},
				},
				{
					Type: mts.RestrictionSelectionValue,
					Context: map[string]any{
						mts.CtxKeyMarketID:     "20",
						mts.CtxKeyOddID:        "1",
						mts.CtxKeySportEventID: "85d130b1-6c8f-4170-bb4f-f80343b55c01",
						mts.CtxKeyValue:        "1.00",
					},
				},
			},
		},
		{
			name:        "bad request error",
			input:       defaultReq,
			queryParams: defaultQueryParams,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-bad-request-response.json"),
			expectedErr: mts.ErrBadRequest,
		},
		{
			name:        "certificate error",
			input:       defaultReq,
			queryParams: defaultQueryParams,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-certificate-response.json"),
			expectedErr: mts.ErrInvalidCertificate,
		},

		{
			name:        "unknown api error",
			input:       defaultReq,
			queryParams: defaultQueryParams,
			httpResp:    s.makeResponse(http.StatusBadRequest, "error-response.json"),
			expectedErr: mts.ErrUnknown,
		},
		{
			name:        "ip forbidden",
			input:       defaultReq,
			queryParams: defaultQueryParams,
			httpResp:    s.makeResponse(http.StatusForbidden, ""),
			expectedErr: mts.ErrAccessForIPDenied,
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
				fmt.Sprintf("%s/restrictions?%s", testURL, tc.queryParams.Encode()),
				http.NoBody,
			)
			roundTripper.EXPECT().RoundTrip(s.newReqMatcher(req)).Return(tc.httpResp, nil)

			actual, actualErr := s.mtsClient.GetRestrictions(context.Background(), tc.input)

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

func mustParseTime(t *testing.T, rawTime string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, rawTime)
	require.NoError(t, err)

	return parsedTime
}
