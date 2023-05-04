package statistics_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/databet-cloud/databet-go-sdk/internal/testutils"
	"github.com/databet-cloud/databet-go-sdk/internal/testutils/mocks"
	"github.com/databet-cloud/databet-go-sdk/pkg/statistics"
	"github.com/databet-cloud/databet-go-sdk/pkg/statistics/testdata"
)

const testURL = "http://api-test"

func TestClientHTTP(t *testing.T) {
	client := http.DefaultClient
	sapiClient := statistics.NewClientHTTP(client, testURL)

	type req struct {
		fixtureID string
		version   string
	}

	testCases := []struct {
		name        string
		httpResp    *http.Response
		expected    statistics.Statistics
		expectedErr error
		req         req
	}{
		{
			name:     "successfully find statistics by id",
			httpResp: testutils.MustMakeResponse(http.StatusOK, "testdata/response_find_statistics_by_id_success.json"),
			expected: testdata.ExpectedStatisticsResponse,
			req: req{
				fixtureID: "85d130b1-6c8f-4170-bb4f-f80343b55c01",
				version:   "1",
			},
		},
		{
			name:        "custom error",
			httpResp:    testutils.MustMakeResponse(http.StatusInternalServerError, "testdata/response_find_statistics_by_id_fail.json"),
			expected:    nil,
			expectedErr: statistics.ErrUnknown,
			req: req{
				fixtureID: "85d130b1-6c8f-4170-bb4f-f80343b55c01",
				version:   "1",
			},
		},
		{
			name:        "unauthorized",
			httpResp:    testutils.MustMakeResponse(http.StatusUnauthorized, ""),
			expectedErr: statistics.ErrInvalidCertificate,
		},
		{
			name:        "ip forbidden",
			httpResp:    testutils.MustMakeResponse(http.StatusForbidden, ""),
			expectedErr: statistics.ErrForbidden,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			roundTripper := mocks.NewRoundTripper(ctrl)
			client.Transport = roundTripper

			httpReq := testutils.MustMakeRequest(
				http.MethodGet,
				fmt.Sprintf("%s/statistics/%s/%s", testURL, tc.req.fixtureID, tc.req.version),
				http.NoBody,
			)
			roundTripper.EXPECT().RoundTrip(testutils.NewReqMatcher(httpReq, t)).Return(tc.httpResp, nil)

			actual, actualErr := sapiClient.FindFixtureStatisticsByID(context.Background(), tc.req.fixtureID, tc.req.version)

			assert.Equal(t, tc.expected, actual)
			assert.ErrorIs(t, actualErr, tc.expectedErr)
		})
	}

}
