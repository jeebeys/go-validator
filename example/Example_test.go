package example

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	_validator "github.com/jeebeys/go-validator/validator"
	"regexp"
	"testing"
)

// -gcflags=-l
// @Validate
func TestExample(t *testing.T) {
	scanPath := `D:\src\workspace.golang.project\go-validator\example`
	_validateManager := _validator.NewValidatorManager(_validator.ValidatorConfig{ScanPath: scanPath})
	_ = _validateManager.GetValidate().RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		uuidRegex := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)
		return uuidRegex.MatchString(fl.Field().String())
	})
	_validateManager.Register(new(Example))
	_example := new(Example)

	res1 := _example.Error1(Param{Name: "just4it", Age: 150, UUID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8"})
	fmt.Println("result1: ", res1)

	res2 := _example.Error1(Param{Name: "just4it", Age: 100, UUID: "2"})
	fmt.Println("result2: ", res2)
	//_validate := validator.New()
	//_err := _validate.Struct(example.Param{Age: 150})
	//if _err != nil {
	//	for _, err := range _err.(validator.ValidationErrors) {
	//		fmt.Println(err.Error()) // Key: 'User.Uid' Error:Field validation for 'Uid' failed on the 'eqcsfield' tag
	//	}
	//}
}
