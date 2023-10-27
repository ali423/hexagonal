package util

import "errors"

const (
	INVALID_INPUT_DATA_FORMAT         = "فرمت اطلاعات ورودی اشتباه است."
	VALIDATION_ERROR_REQUIRED         = "فیلد '%s' الزامی است."
	VALIDATION_ERROR_GREATER_THAN     = `مقدار %q باید بزرگتر %v باشد.`
	VALIDATION_ERROR_REQUIRED_WITHOUT = "فیلد %q در صورت نبود فیلد %q ضروری است."
	VALIDATION_ERROR_DEFAULT          = "مقدار فیلد %q معتبر نیست."
	TRANSACTION_FAILED                = "عملیات با خطا روبرو شد."
	INVALID_INPUT                     = "داده های ورودی معتبر نیست."
	FORBIDDEN_ACTION                  = "مجوز دسترسی به این بخش را ندارید."
)

var FIELDS = map[string]string{}

var ErrGeneral = errors.New(TRANSACTION_FAILED)
var ErrInvalidInput = errors.New(INVALID_INPUT)
var ErrForbiddenAction = errors.New(FORBIDDEN_ACTION)
