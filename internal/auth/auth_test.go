package auth_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		"no header": {
			expectErr: "no authorization header",
		},
		"missing key": {
			key:       "Authorization",
			expectErr: "no authorization header",
		},
		"bad value": {
			key:       "Authorization",
			value:     "-",
			expectErr: "malformed authorization",
		},
		"using beaerer instead of api token": {
			key:       "Authorization",
			value:     "Bearer xxxxxx",
			expectErr: "malformed authorization",
		},
		"using correcct structure": {
			key:       "Authorization",
			value:     "ApiKey xxxxxx",
			expect:    "xxxxxx",
			expectErr: "not expecting an error",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetApiKey: %v ", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			output, err := auth.GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Unexpected TestGetApiKey:%v\n", err)
				return
			}
			if output != test.expect {
				t.Errorf("Unexpected: TestGetApiKey:%v\n", output)
				return
			}
		})
	}
}
