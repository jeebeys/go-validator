package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	_zh_lang "github.com/go-playground/validator/v10/translations/zh"
	"github.com/jeebeys/go-validator/example"
	_validator "github.com/jeebeys/go-validator/validator"
	"github.com/jeebeys/go-validator/web"
	"reflect"
	"regexp"
)

// go mod init github.com/jeebeys/go-validator
// -gcflags=-l
func main() {
	scanPath := `D:\src\workspace.golang.project\go-validator\example`
	_validatorManager := _validator.NewValidatorManager(_validator.ValidatorConfig{ScanPath: scanPath})

	_ = _validatorManager.GetValidate().RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		uuidRegex := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)
		return uuidRegex.MatchString(fl.Field().String())
	})

	_validatorManager.GetValidate().RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("label")
	})

	translator, _ := ut.New(zh.New()).GetTranslator("zh")
	_ = _zh_lang.RegisterDefaultTranslations(_validatorManager.GetValidate(), translator)
	_validatorManager.RegisterTranslator(translator)

	_validatorManager.Register(new(example.Example))

	_example := new(example.Example)

	err1 := _example.Error1(example.Param{Age: 150, UUID: "6ba7b810-9dad-11d1-80b4-00c04fd430c"})
	fmt.Println("error1: \n", err1.Error())
	//
	err2 := _example.Error2(example.Param{Name: "just4it", Age: 150, UUID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8"})
	fmt.Println("error2: ", err2.GetMessage())

	err3 := _example.Error3(example.Param{Name: "just4it", Age: 100, UUID: "6ba7b810-9dad-11d1-80b4-00c04fd430c81"})
	fmt.Println("error3: ", err3.GetMessage())

	err4 := _example.Error4(&example.Param{Name: "just4it", Age: 100, UUID: "6ba7b810-9dad-11d1-80b4-00c04fd430c81"})
	fmt.Println("error4: ", err4.GetMessage())

	//res2, err2, res := _example.Check(example.Param{Name: "just4it", Age: 100, UUID: "2"})
	//fmt.Println("result2: ", res2, err2)
	//
	//_validate := validator.New()
	//_err := _validate.Struct(example.Param{Age: 150})
	//if _err != nil {
	//	for _, err := range _err.(validator.ValidationErrors) {
	//		fmt.Println(err.Error()) // Key: 'User.Uid' Error:Field validation for 'Uid' failed on the 'eqcsfield' tag
	//	}
	//}

	fmt.Println("==========")
	jsonObj1 := web.SUCCESS.Message("abc")
	byteObj1, _ := json.Marshal(jsonObj1)
	fmt.Println(string(byteObj1))

	jsonObj2 := web.FAILURE.Message("abc")
	byteObj2, _ := json.Marshal(jsonObj2)
	fmt.Println(string(byteObj2))

	jsonArr := web.ResultSet{}.Append("abc", "ef", "eft", jsonObj1, jsonObj2)
	byteArr, _ := json.Marshal(jsonArr)
	fmt.Println(string(byteArr))
}
