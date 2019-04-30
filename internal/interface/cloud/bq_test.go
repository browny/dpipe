package cloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testProjectID = "cyberwind"
)

func TestQuery(t *testing.T) {
	err := Query(testProjectID)
	assert.NoError(t, err)
}

func TestCreateTable(t *testing.T) {
	err := CreateTable(testProjectID)
	assert.NoError(t, err)
}

func TestWritePosts(t *testing.T) {
	err := WritePosts(testProjectID)
	assert.NoError(t, err)
}
