package dependencies

import (
	"app/pkg/app/dependencies"
	"app/pkg/logic"
	"app/pkg/logic/sales"
	"app/pkg/logic/users"
	"errors"
	"fmt"
)

func Initialize(deps *dependencies.Dependencies) (err error) {
	defer func() {
		r := recover()
		switch r := r.(type) {
		case error:
			err = r
		case string:
			err = errors.New(r)
		default:
			err = fmt.Errorf("unexpected panic: %v", r)
		}
	}()
	logic.Dependencies().SetUserLogics(users.NewLogic(deps))
	logic.Dependencies().SetSalesLogics(sales.NewLogic(deps))

	return
}
