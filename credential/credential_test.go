package credential

import (
	"reflect"
	"testing"

	"golang.org/x/net/context"
)

func TestCredentials(t *testing.T) {
	tests := []struct {
		username string
		password string
		secure   bool
	}{{
		username: "testuser",
		password: "testpassword",
		secure:   false,
	}}
	for _, tt := range tests {
		c := NewPassCred(tt.username, tt.password, tt.secure)
		if got, want := c.RequireTransportSecurity(), tt.secure; got != want {
			t.Errorf("RequireTransportSecurity() on %+v failed: got %v, want %v", tt, got, want)
		}
		got, err := c.GetRequestMetadata(context.Background(), "foo")
		if err != nil {
			t.Errorf("c() on %v failed: %s", tt, err)
		}
		want := map[string]string{"username": tt.username, "password": tt.password}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetRequestMetadata() on %+v failed: got: %v, want %v", tt, got, want)
		}
	}
}
