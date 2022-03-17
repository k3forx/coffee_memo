package reader_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/reader"
)

func TestUser_GetByID(t *testing.T) {
	t.Parallel()

	userReader := reader.NewUserReader(testClient)

	cases := map[string]struct {
		userID   int
		expected model.User
	}{
		"has_rows": {
			userID: 1,
			expected: model.User{
				ID: 1,
			},
		},
		"no_rows": {
			userID:   -100,
			expected: model.User{},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := userReader.GetByID(context.Background(), c.userID)
			fmt.Printf("actual: %+v\n", actual)

			if err != nil {
				t.Errorf("err should be nil, but got %v", err)
			}
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("GetByID() mismatch (-want +got): %s\n", diff)
			}
		})
	}
}
