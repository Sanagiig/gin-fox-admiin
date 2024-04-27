package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var compareMap = map[string]bool{
	"lt":    true,
	"le":    true,
	"eq":    true,
	"ne":    true,
	"ge":    true,
	"gt":    true,
	"in":    true,
	"notIn": true,
}

func SetValidatorComppareMap(k string, open bool) (bool, error) {
	_, ok := compareMap[k]
	if !ok {
		return false, errors.New(k + " not exist in compareMap")
	}

	compareMap[k] = open
	return true, nil
}

func Verify(struc interface{}, rules Rules) (err error, innerErr error) {
	typ := reflect.TypeOf(struc)
	val := reflect.ValueOf(struc)

	kd := typ.Kind()
	if kd != reflect.Struct {
		return nil, errors.New("Verify expect struct")
	}

	num := val.NumField()
	for i := 0; i < num; i++ {
		fieldTyp := typ.Field(i)
		fieldVal := val.Field(i)
		if fieldTyp.Type.Kind() == reflect.Struct {
			err, innerErr = Verify(fieldVal.Interface(), rules)
			if err != nil || innerErr != nil {
				return
			}
		}

		rule, ok := rules[fieldTyp.Name]

		if ok && len(rule) > 0 {
			for _, r := range rule {
				ruleParts := strings.SplitN(r, "=", 2)
				switch {
				case ruleParts[0] == NotEmptyStr:
					if isEmpty(fieldVal) {
						return errors.New(fieldTyp.Name + " 缺少数据"), nil
					}
				case ruleParts[0] == RegexpStr:
					if !regexpMatch(ruleParts[1], fieldVal.String()) {
						return errors.New(fieldTyp.Name + " 格式校验不通过, 必须满足 [ " + ruleParts[1] + " ]"), nil
					}
				case compareMap[ruleParts[0]]:
					err, innerErr = compareVerify(fieldVal, ruleParts[0], ruleParts[1])
					if err != nil || innerErr != nil {
						err = errors.Join(fmt.Errorf("field [%s] error.\n", fieldTyp.Name), err)
						return
					}
				default:
					rulesChecker := CustomizeMap[ruleParts[0]]
					if rulesChecker != nil {
						err, innerErr = rulesChecker.Verify(fieldVal, ruleParts[1])
						if err != nil || innerErr != nil {
							err = errors.Join(fmt.Errorf("field [%s] error.\n", fieldTyp.Name), err)
							return
						}
					}
				}

			}
		}
	}
	return
}

func isEmpty(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String, reflect.Slice:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func regexpMatch(rule, matchStr string) bool {
	return regexp.MustCompile(rule).MatchString(matchStr)
}

//@function: compareVerify
//@description: 长度和数字的校验方法 根据类型自动校验
//@param: value reflect.Value, verifyStr string
//@return: bool

func compareVerify(value reflect.Value, verifyName string, verifyStr string) (err error, innerErr error) {
	switch value.Kind() {
	case reflect.String:
		switch verifyName {
		case "in":
			return strInOrNot(verifyStr, value.String(), true)
		case "notIn":
			return strInOrNot(verifyStr, value.String(), false)
		default:
			return compare(len([]rune(value.String())), verifyName, verifyStr)
		}
	case reflect.Slice, reflect.Array:
		return compare(value.Len(), verifyName, verifyStr)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), verifyName, verifyStr)
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), verifyName, verifyStr)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), verifyName, verifyStr)
	default:
		return nil, errors.New("缺少数据校验的类型")
	}
}

func compare(value interface{}, verifyName string, verifyStr string) (err error, innerErr error) {
	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compareNum(verifyName, verifyStr, val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compareNum(verifyName, verifyStr, val.Uint())
	case reflect.Float32, reflect.Float64:
		return compareNum(verifyName, verifyStr, val.Float())
	default:
		return nil, errors.New("value kind must be number (int , float , uint)")
	}
	return
}

func compareNum[T int64 | uint64 | float64](verifyName string, verifyStr string, val T) (error, error) {
	var pv T

	if verifyName != "in" && verifyName != "notIn" {
		vKind := reflect.TypeOf(val).Kind()
		v := reflect.ValueOf(&pv)
		switch vKind {
		case reflect.Int64:
			tmp, innerErr := strconv.ParseInt(verifyStr, 10, 64)
			if innerErr != nil {
				return nil, innerErr
			}
			v.Elem().SetInt(tmp)
		case reflect.Uint64:
			tmp, innerErr := strconv.ParseUint(verifyStr, 10, 64)
			if innerErr != nil {
				return nil, innerErr
			}
			v.Elem().SetUint(tmp)
		case reflect.Float64:
			tmp, innerErr := strconv.ParseFloat(verifyStr, 64)
			if innerErr != nil {
				return nil, innerErr
			}
			v.Elem().SetFloat(tmp)
		default:
			return nil, errors.New("value kind must be number (int , float , uint)")
		}
	}

	switch verifyName {
	case "lt":
		if val >= pv {
			return fmt.Errorf("must less then %s\n", verifyStr), nil
		}
	case "le":
		if val > pv {
			return fmt.Errorf("must less or equals %s\n", verifyStr), nil
		}
	case "eq":
		if val != pv {
			return fmt.Errorf("must equals %s\n", verifyStr), nil
		}
	case "ne":
		if val == pv {
			return fmt.Errorf("must not equals %s\n", verifyStr), nil
		}
	case "ge":
		if val < pv {
			return fmt.Errorf("must greater or equals %s\n", verifyStr), nil
		}
	case "gt":
		if val <= pv {
			return fmt.Errorf("must greater then %s\n", verifyStr), nil
		}
	case "in":
		list := strings.Split(verifyStr[1:len(verifyStr)-1], ",")
		for _, str := range list {
			v, vErr := strconv.ParseInt(str, 10, 64)
			if vErr != nil {
				return nil, vErr
			}
			if v == int64(val) {
				return nil, nil
			}
		}
		return fmt.Errorf("must in %s", verifyStr), nil
	case "notIn":
		list := strings.Split(verifyStr[1:len(verifyStr)-1], ",")
		for _, str := range list {
			v, vErr := strconv.ParseInt(str, 10, 64)
			if vErr != nil {
				return nil, vErr
			}
			if v == int64(val) {
				return fmt.Errorf("must not in %s", verifyStr), nil
			}
		}
	}
	return nil, nil
}

func strInOrNot(verifyStr string, val string, isIn bool) (error, error) {
	strList := strings.Split(verifyStr[1:len(verifyStr)], ",")
	has := false
	for i := 0; i < len(strList); i++ {
		if val == strList[i] {
			has = true
			break
		}
	}

	if isIn {
		if has {
			return nil, nil
		}
		return fmt.Errorf("must in %s", verifyStr), nil
	}

	if has {
		return fmt.Errorf("must not in %s", verifyStr), nil
	}
	return nil, nil
}
