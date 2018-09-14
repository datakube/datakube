package api_test

import (
	"encoding/json"
	"github.com/SantoDE/datahamster/internal/test"
	"github.com/SantoDE/datahamster/server/http/handlers/api"
	"github.com/SantoDE/datahamster/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type targetsResponse struct {
	Targets []types.Target `json:"targets"`
}

func TestGetTargets(t *testing.T) {

	targetStoreMock := test.TargetStoreMock{}
	targetStoreMock.Success = true

	r := gin.Default()
	r.GET("/targets", api.GetTargets(&targetStoreMock))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/targets", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response targetsResponse
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 2, len(response.Targets))

	targetStoreMock.Success = false
	req, _ = http.NewRequest("GET", "/targets", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 2, len(response.Targets))
}
