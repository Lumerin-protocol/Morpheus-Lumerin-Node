package config

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-playground/validator/v10"
)

var hexadecimalRegex = regexp.MustCompile("^(0[xX])?[0-9a-fA-F]+$")

func NewValidator() (*validator.Validate, error) {
	valid := validator.New()

	err := RegisterHex32(valid)
	if err != nil {
		return nil, err
	}

	err = RegisterDuration(valid)
	if err != nil {
		return nil, err
	}

	err = RegisterEthAddr(valid)
	if err != nil {
		return nil, err
	}

	err = RegisterHexadecimal(valid)
	if err != nil {
		return nil, err
	}

	return valid, nil
}

// isHexadecimal is the validation function for validating if the current field's value is a valid hexadecimal.
func isHexadecimal(fl string) bool {
	return hexadecimalRegex.MatchString(fl)
}

func RegisterHexadecimal(v *validator.Validate) error {
	return v.RegisterValidation("hexadecimal", func(fl validator.FieldLevel) bool {
		kind := fl.Field().Kind()

		if kind == reflect.String {
			s := fl.Field().String()
			return isHexadecimal(s)
		}

		if kind == reflect.Slice {
			_, ok := fl.Field().Interface().(lib.HexString)
			return ok
		}

		return false
	})
}

func RegisterHex32(v *validator.Validate) error {
	return v.RegisterValidation("hex32", func(fl validator.FieldLevel) bool {
		kind := fl.Field().Kind()

		if kind == reflect.String {
			s := fl.Field().String()

			if !isHexadecimal(s) {
				return false
			}

			trimmed := strings.TrimPrefix(strings.TrimPrefix(s, "0x"), "0X")
			return len(trimmed) == 32
		}

		if kind == reflect.Array {
			arr, ok := fl.Field().Interface().(common.Hash)
			return ok && len(arr) == 32
		}

		return false
	})
}

func RegisterDuration(v *validator.Validate) error {
	return v.RegisterValidation("duration", func(fl validator.FieldLevel) bool {
		if kind := fl.Field().Kind(); kind != reflect.Int64 {
			return false
		}

		return fl.Field().Int() != 0
	})
}

func RegisterEthAddr(v *validator.Validate) error {
	return v.RegisterValidation("eth_addr", func(fl validator.FieldLevel) bool {
		kind := fl.Field().Kind()
		if kind == reflect.String {
			s := fl.Field().String()

			errs := v.Var(s, "eth_addr")
			if errs != nil {
				return false
			}
			return true
		}

		// if stored as common.Address
		if kind == reflect.Array {
			arr, ok := fl.Field().Interface().(common.Address)
			return ok && len(arr) == 20
		}

		return false
	})
}
