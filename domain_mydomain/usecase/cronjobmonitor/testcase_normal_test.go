package cronjobmonitor

import (
	"context"
	"testing"

	"demo-cronjobmonitor/domain_mydomain/model/entity"
)

type mockOutportNormal struct {
	t *testing.T
}

// TestCaseNormal is for the case where ...
// explain the purpose of this test here with human readable naration...
func TestCaseNormal(t *testing.T) {

	ctx := context.Background()

	mockOutport := mockOutportNormal{
		t: t,
	}

	res, err := NewUsecase(&mockOutport).Execute(ctx, InportRequest{
		Shipments:     "10-10-2022:01:08:56",
		Shipments_log: "17-10-2022:03:10:06",
	})

	if err != nil {
		t.Errorf("%v", err.Error())
		t.FailNow()
	}

	t.Logf("%v", res)

}

func (r *mockOutportNormal) SaveUser(ctx context.Context, obj *entity.User) error {

	return nil
}
