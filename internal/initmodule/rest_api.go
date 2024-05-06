package initmodule

import (
	"coincap/internal/rest_api/controller"
	"coincap/internal/usecase"
)

type (
	RestAPIWrapper struct {
		Controller *controller.Schema
		Usecases   *usecase.Schema
	}
)

func NewRestAPIWrapper(
	ctrl *controller.Schema,
	usecases *usecase.Schema,
) *RestAPIWrapper {
	return &RestAPIWrapper{
		Controller: ctrl,
		Usecases:   usecases,
	}
}
