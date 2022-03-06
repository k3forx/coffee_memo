package inject

import (
	"github.com/golang/mock/gomock"
	"github.com/k3forx/coffee_memo/pkg/reader"
)

func NewMockInjector(ctrl *gomock.Controller) Injector {
	return Injector{
		Reader: Reader{
			User: reader.NewMockUser(ctrl),
		},
	}
}
