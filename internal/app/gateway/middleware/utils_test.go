package middleware

import (
	"testing"
)

func Test_matchRoute(t *testing.T) {
	type args struct {
		route string
		path  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"GET/accounts/{id}", "GET/accounts/1"}, true},
		{"", args{"GET/accounts/{id}", "GET/accounts/ab-a"}, true},
		{"", args{"GET/accounts/{id}", "GET/accounts/ab_a"}, true},
		{"", args{"GET/accounts/{id}", "GET/accounts/1"}, true},
		{"", args{"GET/accounts/{id}", "GET/accounts/1/info"}, false},
		{"", args{"GET/accounts/{id}", "GET/accounts/1/info-a"}, false},
		{"", args{"GET/accounts/{id}", "GET/accounts/1/info_a"}, false},

		{"", args{"GET/accounts/{account_id}/info", "GET/accounts/1/info"}, true},
		{"", args{"GET/accounts/{account_id}/info", "GET/accounts/1/info_a"}, false},

		{"", args{"GET/accounts/{account-id}/info", "GET/accounts/1/info"}, true},
		{"", args{"GET/accounts/{account-id}/info", "GET/accounts/1/info_a"}, false},

		{"", args{"GET/accounts/{id}/info", "GET/accounts/1/info"}, true},
		{"", args{"GET/accounts/{id}/info", "GET/accounts/a/info"}, true},
		{"", args{"GET/accounts/{id}/info", "GET/accounts/a_a/info"}, true},
		{"", args{"GET/accounts/{id}/info", "GET/accounts/a-a/info"}, true},
		{"", args{"GET/accounts/{id}/info", "GET/accounts/1/info-a"}, false},

		{"", args{"GET/accounts/{id}/messages/{messageId}", "GET/accounts/1/messages/msg-123"}, true},
		{"", args{"GET/accounts/{id}/messages/{messageId}", "GET/accounts/1/messages"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchRoute(tt.args.route, tt.args.path); got != tt.want {
				t.Errorf("matchRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
