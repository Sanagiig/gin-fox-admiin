package system

import (
	"fmt"
	"gin-one/model/system/request"
)

type FormService struct{}

func (f *FormService) CreateForm(form request.CreateForm) error {
	fmt.Println("form", form)
	return nil
}
