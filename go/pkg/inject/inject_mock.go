package inject

import (
	"github.com/golang/mock/gomock"
	readerMock "github.com/k3forx/coffee_memo/pkg/reader/mock"
	writerMock "github.com/k3forx/coffee_memo/pkg/writer/mock"
)

func NewMockInjector(ctrl *gomock.Controller) Injector {
	return Injector{
		Reader: Reader{
			UserCoffeeBean: readerMock.NewMockUserCoffeeBean(ctrl),
			User:           readerMock.NewMockUser(ctrl),
		},
		Writer: Writer{
			UserBrewRecipe: writerMock.NewMockUserBrewRecipe(ctrl),
			UserCoffeeBean: writerMock.NewMockUserCoffeeBean(ctrl),
			User:           writerMock.NewMockUser(ctrl),
		},
	}
}
