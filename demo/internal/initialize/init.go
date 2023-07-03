/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  initialize
 * @Date: 2023/07/03 11:46
 */

package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en1 "github.com/go-playground/validator/v10/translations/en"
	zh1 "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func TransInitialize(locale string) error {
	validate := binding.Validator.Engine().(*validator.Validate) // 修改gin框架中的validator引擎属性,实现定制
	zhT := zh.New()                                              // 中文翻译器
	enT := en.New()
	uni := ut.New(zhT, zhT, enT)
	var ok bool
	Trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s)", locale)
	}
	if locale == "en" {
		en1.RegisterDefaultTranslations(validate, Trans)
	} else {
		zh1.RegisterDefaultTranslations(validate, Trans)
	}
	return nil
}
