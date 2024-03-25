package example

import (
	"errors"
	"github.com/jeebeys/go-validator/web"
)

type Example struct {
}

type Param struct {
	Name string `validate:"required" label:"名字"`
	UUID string `validate:"uuid" label:"主键"`
	Age  uint8  `validate:"gte=0,lte=130" label:"年龄"`
}

// @Validate
func (d *Example) Error1(par Param) error {
	return errors.New("error")
}

// @Validate
func (t *Example) Error2(par Param) web.ResultObj {
	return web.ResultObj{}
}

// @Validate
func (t *Example) Error3(par Param) *web.ResultObj {
	return &web.ResultObj{}
}
