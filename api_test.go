package dpipe

import "testing"

func TestGetPosts(t *testing.T) {
	want := "get post"
	if got := GetPosts(); got != want {
		t.Errorf("GetPosts() = %q, want %q", got, want)
	}
}
