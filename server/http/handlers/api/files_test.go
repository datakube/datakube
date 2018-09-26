package api_test

import (
	"encoding/json"
	"github.com/datakube/datakube/internal/test"
	"github.com/datakube/datakube/server/http/handlers/api"
	"github.com/datakube/datakube/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type filesResponse struct {
	Dumps  []types.DumpFile `json:"dumps"`
	Error string      `json:"error"`
}

func TestGetFiles(t *testing.T) {

	dumpStoreMock := test.DumpFileStoreMock{}
	dumpStoreMock.Success = true

	r := gin.Default()
	r.GET("/dumps", api.GetFiles(&dumpStoreMock))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/dumps", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response filesResponse
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, 2, len(response.Dumps))
	assert.Empty(t, response.Error)

	dumpStoreMock.Success = false
	errorRec := httptest.NewRecorder()
	r.ServeHTTP(errorRec, req)

	assert.Equal(t, 500, errorRec.Code)
	json.Unmarshal([]byte(errorRec.Body.String()), &response)
	assert.Equal(t, 0, len(response.Dumps))
	assert.NotNil(t, response.Error)
}

