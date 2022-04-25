package inject

import (
	"github.com/golang/mock/gomock"
	readerMock "github.com/k3forx/coffee_memo/pkg/reader/mock"
	writerMock "github.com/k3forx/coffee_memo/pkg/writer/mock"
)

func NewMockInjector(ctrl *gomock.Controller) Injector {
	return Injector{
		Reader: Reader{
			User: readerMock.NewMockUser(ctrl),
		},
		Writer: Writer{
			CoffeeBean: writerMock.NewMockCoffeeBean(ctrl),
			User:       writerMock.NewMockUser(ctrl),
		},
	}
}
