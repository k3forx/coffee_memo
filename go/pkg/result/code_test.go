package result_test

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/coffee_memo/pkg/result"
)

func TestCode_String(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		code     result.Code
		expected string
	}{
		"unspecified": {
			code:     result.CodeUnspecified,
			expected: "error",
		},
		"not_found": {
			code:     result.CodeNotFound,
			expected: "not found",
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, c.code.String()); diff != "" {
				t.Errorf("Code.String() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCode_ToStatusCode(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		code     result.Code
		expected int
	}{
		"ok": {
			code:     result.CodeOK,
			expected: http.StatusOK,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, c.code.ToStatusCode()); diff != "" {
				t.Errorf("Code.ToStatusCode() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
