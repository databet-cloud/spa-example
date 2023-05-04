package testutils

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func NewReqMatcher(req *http.Request, t *testing.T) ReqMatcher {
	return ReqMatcher{
		t:           t,
		expectedReq: req,
	}
}

type ReqMatcher struct {
	t           *testing.T
	expectedReq *http.Request
}

func (r ReqMatcher) Matches(x any) bool {
	actualReq, ok := x.(*http.Request)
	if !ok {
		return false
	}

	if !assert.Equal(r.t, r.expectedReq.Method, actualReq.Method) {
		return false
	}

	if !assert.Equal(r.t, r.expectedReq.URL.Query(), actualReq.URL.Query()) {
		return false
	}

	if !assert.Equal(r.t, r.expectedReq.URL.Path, actualReq.URL.Path) {
		return false
	}

	expectedBody, err := io.ReadAll(r.expectedReq.Body)
	require.NoError(r.t, err)

	actualBody, err := io.ReadAll(actualReq.Body)
	require.NoError(r.t, err)

	if len(expectedBody) > 0 && !assert.JSONEq(r.t, string(expectedBody), string(actualBody)) {
		return false
	}

	return true
}

func (r ReqMatcher) String() string {
	return fmt.Sprintf("%v", r.expectedReq)
}
