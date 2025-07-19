package dependencies

import (
	"app/pkg/app/dependencies"
	"app/pkg/logic"
	"app/pkg/logic/sales"
	"app/pkg/logic/users"
)

func Initialize(deps *dependencies.Dependencies) (err error) {
	logic.Dependencies().SetUserLogics(users.NewLogic(deps))
	logic.Dependencies().SetSalesLogics(sales.NewLogic(deps))

	return
}
