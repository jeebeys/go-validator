package validator

import (
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/jeebeys/go-validator/aop"
	"github.com/jeebeys/go-validator/web"
	"reflect"
	"strings"
)

var methodStructMap = make(map[string]*joinPointStructInfo)

type Validator struct {
	validate   *validator.Validate
	translator translator.Translator
}

func (t *Validator) Before(point *aop.JoinPoint, methodLocation string) bool {

	index := len(point.Result) - 1
	types := "type-e"
	for i, v := range point.Result {
		val := v.Interface()
		switch val.(type) {
		case web.ResultObj:
			index = i
			types = "type-o"
			break
		case *web.ResultObj:
			index = i
			types = "type-p"
			break
		}
	}
	for _, v := range point.Params {
		if v.Kind() == reflect.Struct || v.Kind() == reflect.Pointer {
			_error := t.validate.Struct(v.Interface())
			if _error != nil {
				var messages []string
				for _, err := range _error.(validator.ValidationErrors) {
					messages = append(messages, err.Translate(t.translator))
				}
				_message := strings.Join(messages, ";")
				switch types {
				case "type-e":
					point.Result[index] = reflect.ValueOf(_error)
					break
				case "type-o":
					point.Result[index] = reflect.ValueOf(web.FAILURE.Message(_message))
					break
				case "type-p":
					point.Result[index] = reflect.ValueOf(web.FAILURE.Message(_message).Ptr())
					break
				}
				return false
			}
		}
	}
	return true
}

func (t *Validator) After(point *aop.JoinPoint, methodLocation string) {

}

func (t *Validator) Finally(point *aop.JoinPoint, methodLocation string) {

}

func (t *Validator) IsMatch(methodLocation string) bool {
	return true
}

func (t *Validator) Error() error {
	return nil
}
