package handlers_test

import (
	"errors"
	"github.com/datakube/datakube/internal/test"
	"github.com/datakube/datakube/server/http/handlers"
	"github.com/datakube/datakube/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLatestFile(t *testing.T) {

	dfsMock := test.DumpFileStoreMock{}
	storageMock := test.StorageMock{}

	dfOk := types.DumpFile{
		Target: "testTarget",
		File: types.File{
			Path: "/tmp/test_file",
			Name: "testfile.sql",
		},
		ID: 1337,
	}

	dfNoFile := types.DumpFile{
		Target: "noFile",
		File: types.File{
			Path: "",
		},
		ID: 1338,
	}

	dfsMock.On("LoadOneDumpFileByTarget", "testTarget").Return(dfOk, nil)
	dfsMock.On("LoadOneDumpFileByTarget", "noDumpFile").Return(types.DumpFile{}, nil)
	dfsMock.On("LoadOneDumpFileByTarget", "error").Return(dfOk, errors.New("Test Error"))
	dfsMock.On("LoadOneDumpFileByTarget", "noFile").Return(dfNoFile, nil)

	storageMock.On("ReadFile", "/tmp/test_file").Return([]byte("TEST DATA"), nil)
	storageMock.On("ReadFile", "").Return([]byte(""), errors.New("no path"))

	r := gin.Default()
	fileRoutes := r.Group("/files/download/")
	fileRoutes.GET("/:name/latest", handlers.GetLatestFile(dfsMock, storageMock))

	wOk := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/files/download/testTarget/latest", nil)
	r.ServeHTTP(wOk, req)
	assert.Equal(t, 200, wOk.Code)
	assert.Equal(t, "TEST DATA", wOk.Body.String())
	assert.Equal(t, "File Transfer", wOk.Header().Get("Content-Description"))
	assert.Equal(t, "binary", wOk.Header().Get("Content-Transfer-Encoding"))
	assert.Equal(t, "attachment; filename=testfile.sql", wOk.Header().Get("Content-Disposition"))
	assert.Equal(t, "application/octet-stream", wOk.Header().Get("Content-Type"))

	wBadRequest := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/files/download/noDumpFile/latest", nil)
	r.ServeHTTP(wBadRequest, req)
	assert.Equal(t, 400, wBadRequest.Code)
	assert.Equal(t, "{\"message\":\"No Dumps for target noDumpFile found\"}", wBadRequest.Body.String())

	wNoFile := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/files/download/noFile/latest", nil)
	r.ServeHTTP(wNoFile, req)
	assert.Equal(t, "{\"message\":\"no path\"}", wNoFile.Body.String())
	assert.Equal(t, 500, wNoFile.Code)

	wError := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/files/download/error/latest", nil)
	r.ServeHTTP(wError, req)
	assert.Equal(t, "{\"message\":\"Test Error\"}", wError.Body.String())
	assert.Equal(t, 500, wError.Code)
}

func TestGetFile(t *testing.T) {
	dfsMock := test.DumpFileStoreMock{}
	storageMock := test.StorageMock{}

	dfOk := types.DumpFile{
		Target: "testTarget",
		File: types.File{
			Path: "/tmp/test_file",
			Name: "testfile.sql",
		},
		ID: 1337,
	}


	dfsMock.On("LoadOneDumpFileByName", "testfile.sql").Return(dfOk, nil)
	dfsMock.On("LoadOneDumpFileByName", "noDumpFile").Return(types.DumpFile{}, nil)
	dfsMock.On("LoadOneDumpFileByName", "error").Return(dfOk, errors.New("Test Error"))

	storageMock.On("ReadFile", "/tmp/test_file").Return([]byte("TEST DATA"), nil)
	storageMock.On("ReadFile", "").Return([]byte(""), errors.New("no path"))

	r := gin.Default()
	fileRoutes := r.Group("/files/download/")
	fileRoutes.GET("/:name", handlers.GetFile(dfsMock, storageMock))

	wOk := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/files/download/testfile.sql", nil)
	r.ServeHTTP(wOk, req)
	assert.Equal(t, 200, wOk.Code)
	assert.Equal(t, "TEST DATA", wOk.Body.String())
	assert.Equal(t, "File Transfer", wOk.Header().Get("Content-Description"))
	assert.Equal(t, "binary", wOk.Header().Get("Content-Transfer-Encoding"))
	assert.Equal(t, "attachment; filename=testfile.sql", wOk.Header().Get("Content-Disposition"))
	assert.Equal(t, "application/octet-stream", wOk.Header().Get("Content-Type"))

	wBadRequest := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/files/download/noDumpFile", nil)
	r.ServeHTTP(wBadRequest, req)
	assert.Equal(t, 400, wBadRequest.Code)
	assert.Equal(t, "{\"message\":\"No Filename by that name noDumpFile found\"}", wBadRequest.Body.String())

	wError := httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/files/download/error", nil)
	r.ServeHTTP(wError, req)
	assert.Equal(t, "{\"message\":\"Test Error\"}", wError.Body.String())
	assert.Equal(t, 500, wError.Code)
}
