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

type jobsResponse struct {
	Jobs []types.Job `json:"jobs"`
}

func TestGetJobs(t *testing.T) {

	jobStoreMock := test.JobStoreMock{}
	jobStoreMock.Success = true

	r := gin.Default()
	r.GET("/jobs", api.GetJobs(&jobStoreMock))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/jobs", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response jobsResponse
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 2, len(response.Jobs))

	jobStoreMock.Success = false
	req, _ = http.NewRequest("GET", "/jobs", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 2, len(response.Jobs))
}
