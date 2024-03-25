package validator

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/jeebeys/go-validator/aop"
	"github.com/jeebeys/go-validator/ast"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"unicode"
)

var config ValidatorConfig
var methodLocationMap = make(map[string]*struct{})

type joinPointStructInfo struct {
	StructIndex int // orm会话对象在方法入参列表的位置
}

type ValidatorConfig struct {
	ScanPath string
}

func (c *ValidatorConfig) check() error {
	_, err := os.Stat(c.ScanPath)
	return err
}

func (c *ValidatorConfig) Reload() error {
	return nil
}

type ValidatorManager struct {
	validate  *validator.Validate
	validator *Validator
}

func NewValidatorManager(cfg ValidatorConfig) *ValidatorManager {
	if err := cfg.check(); err != nil {
		panic(err)
	}
	config = cfg
	scanGoFile()
	_validator := new(Validator)
	_validator.validate = validator.New()
	aop.RegisterAspect(_validator)
	_validateManager := new(ValidatorManager)
	_validateManager.validate = _validator.validate
	_validateManager.validator = _validator
	return _validateManager
}

func (t *ValidatorManager) GetValidate() *validator.Validate {
	return t.validate
}

func (t *ValidatorManager) RegisterTranslator(translator ut.Translator) (tm *ValidatorManager) {
	t.validator.translator = translator
	return t
}

func (t *ValidatorManager) Register(objects ...interface{}) (tm *ValidatorManager) {
	for _, v := range objects {
		aop.RegisterPoint(reflect.TypeOf(v))
	}
	return t
}

func scanGoFile() {
	err := filepath.Walk(config.ScanPath, walkFunc)
	if err != nil {
		return
	}
}

func walkFunc(fullPath string, info os.FileInfo, err error) error {
	if info == nil {
		return err
	}
	if info.IsDir() {
		return nil
	} else {
		if GO_FILE_SUFIX == path.Ext(fullPath) {
			cacheMethodLocationMap(fullPath)
		}
		return nil
	}
}

func cacheMethodLocationMap(fullPath string) {
	bt, err := os.ReadFile(fullPath)
	if err != nil {
		return
	}
	result := ast.ScanFuncDeclByComment("", string(bt), COMMENT_NAME)
	if result == nil {
		return
	}
	if result.RecvMethods != nil {
		for _, l := range result.RecvMethods {
			for _, v := range l {
				for i, s := range v.MethodName {
					// skip the private method
					if i == 0 && unicode.IsUpper(s) {
						methodLocationMap[fmt.Sprintf("%s.%s.%s", v.PkgName, v.RecvName, v.MethodName)] = new(struct{})
					}
					break
				}
			}
		}
	}
}
