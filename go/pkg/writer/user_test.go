package writer_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/k3forx/coffee_memo/pkg/entity"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/writer"
	db_helper "github.com/k3forx/coffee_memo/test/db"
)

func TestUserWriter_SignUp(t *testing.T) {
	impl := writer.NewUserWriter(testClient)

	cases := map[string]struct {
		user     model.User
		expected model.User
	}{
		"success": {
			user: model.User{
				Username: "test user",
				Email:    "test@test.jp",
				Password: "test pass",
				Flags:    model.UserFlags{model.UserFlagEmailActivated},
			},
			expected: model.User{
				Username:  "test user",
				Email:     "test@test.jp",
				Password:  "test pass",
				Flags:     model.UserFlags{model.UserFlagEmailActivated},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := impl.Create(context.Background(), &c.user)
			if err != nil {
				t.Errorf("err should be nil, but got:\n%v", err)
			}
			t.Cleanup(func() {
				db_helper.DeleteUser(context.Background(), testClient, entity.NewUser(c.user))
			})

			opts := cmp.Options{
				cmpopts.IgnoreFields(model.User{}, "ID"),
				cmpopts.EquateApproxTime(time.Minute),
			}
			if diff := cmp.Diff(c.expected, c.user, opts); diff != "" {
				t.Errorf("UserWriter.Create mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
