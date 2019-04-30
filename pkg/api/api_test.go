package api

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testPostId = 231192908
)

func TestGetComments(t *testing.T) {
	comments, err := GetComments(testPostId)
	assert.NoError(t, err)

	r, _ := json.MarshalIndent(comments, "", "    ")
	log.Printf("%s", r)
	log.Printf("len: %d", len(comments))
}

func TestGetPostsUntil(t *testing.T) {
	posts, err := GetPostsUntil("dressup", 230955909)
	assert.NoError(t, err)

	r, _ := json.MarshalIndent(posts, "", "    ")
	log.Printf("%s", r)
	log.Printf("len: %d", len(posts))
}

func TestContains(t *testing.T) {
	testJson := `[{"id":444},{"id":333},{"id":222},{"id":111}]]`

	contained, _ := contains(testJson, 999)
	assert.False(t, contained)
}
