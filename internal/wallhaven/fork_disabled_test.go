package wallhaven

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

// Fork: Wallhaven is disabled; the client must refuse to dial wallhaven hosts.
func TestForkBlocksWallhavenDial(t *testing.T) {
	c := NewClient()
	dial := c.http.Transport.(*http.Transport).DialContext
	for _, addr := range []string{"wallhaven.cc:443", "w.wallhaven.cc:443", "WALLHAVEN.CC.:443"} {
		if _, err := dial(context.Background(), "tcp", addr); !errors.Is(err, ErrDisabled) {
			t.Errorf("dial %q: got %v, want ErrDisabled", addr, err)
		}
	}
	if _, err := dial(context.Background(), "tcp", "wallhaven.cc.example.com:443"); errors.Is(err, ErrDisabled) {
		t.Error("blocked a non-wallhaven host")
	}
}
