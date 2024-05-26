package monitor

import (
	"context"
	"testing"
)

func TestPing(t *testing.T) {
	tests := []struct {
		URL string
		Up bool
	}{
		{"https://google.com", true},
		{"https://intercity.ng", true},
		{"https://resend.com", true},
		{"https://some-random-non-existing.ng", false},
	}

	for _, test := range tests {
		res, err := Ping(context.TODO(), test.URL)
		if err != nil {
			t.Errorf("unexpected error while pinging [%s] - %v", test.URL, err)
			return
		}

		if res.Up != test.Up {
			t.Errorf("ping [%s], want up=%v, got up=%v", test.URL, test.Up, res.Up)
		}
	}
}
