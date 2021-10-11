//go:generate mockgen -source=me.go -destination ../mock/mock_me.go -package mock
package lagoon_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/uselagoon/lagoon-cli/internal/mock"
	"github.com/uselagoon/lagoon-cli/internal/schema"
)

type meCalls struct {
	Users []schema.User
}

func TestGetMeInfo(t *testing.T) {
	var testCases = map[string]struct {
		expect *meCalls
	}{
		"simple": {expect: &meCalls{
			Users: []schema.User{},
		}},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			ctx := context.Background()
			// set up the mock me
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()
			me := mock.NewMockMe(ctrl)
			// use the provided meCalls to set the expectations
			for i := range tc.expect.Users {
				me.EXPECT().Me(ctx, &tc.expect.Users[i])
			}
		})
	}
}
