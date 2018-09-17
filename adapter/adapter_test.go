package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewAdapter(t *testing.T) {
	adapter, err := CreateNewAdapter("testhost", "testport", "testdb", "testuser", "testpw", "mysql")
	assert.Nil(t, err)
	assert.NotNil(t, adapter)
}