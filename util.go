package utils

import (
	"reflect"
	"strconv"
	"strings"
)

/**
 * 类型检测 要检测的变量  期望变量类型
 */
func TypeCheck(params interface{}, t string) bool {
	//数据初始化
	var (
		return_val bool = false
	)
	v := reflect.ValueOf(params)
	//获取传递参数类型
	v_t := v.Type()

	//类型名称对比
	if v_t.String() == t {
		return_val = true
	}
	return return_val
}

/**
 * 获取数据类型 (string, int, float64, []int, []string, map[string]int ...)
 */
func GetType(params interface{}) string {
	//数据初始化
	v := reflect.ValueOf(params)
	//获取传递参数类型
	v_t := v.Type()

	//类型名称对比
	return v_t.String()
}

/**
 * 给定元素值 是否在 指定的数组中
 */
func InArray(needle interface{}, hystack []interface{}) bool {
	for _, item := range hystack {
		if GetType(needle) == GetType(item) {
			if needle == item {
				return true
			}
		}
	}

	return false
}

/**
 * int 数组转 interface
 */
func Aitoi(arr []int) []interface{} {
	var tmp []interface{}

	for _, item := range arr {
		tmp = append(tmp, item)
	}

	return tmp
}

/**
 * string 数组转 interface
 */
func Astoi(arr []string) []interface{} {
	var tmp []interface{}

	for _, item := range arr {
		tmp = append(tmp, item)
	}

	return tmp
}

func ParseInt(data interface{}) int {
	dataType := GetType(data)
	var res int
	switch dataType {
	case "string":
		res, _ = strconv.Atoi(data.(string))
	case "int":
		res = data.(int)
	default:
		panic("只能转换字符串类型")
	}

	return res
}

func ParseStr(data interface{}) string {
	dataType := GetType(data)
	var res string
	switch dataType {
	case "int":
		res = strconv.Itoa(data.(int))
	case "string":
		res = data.(string)
	default:
		panic("只能转换数字类型")
	}

	return res
}

/**
 * 三元运算
 */
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

/**
 * 添加单引号
 */
func AddSingleQuotes(data interface{}) string {
	return "'" + strings.Trim(ParseStr(data), " ") + "'"
}

/**
 * 字符串转数组, 接受混合类型, 最终输出的是字符串类型
 */
func Implode(data interface{}, glue string) string {
	var tmp []string
	for _, item := range data.([]interface{}) {
		tmp = append(tmp, ParseStr(item))
	}

	return strings.Join(tmp, glue)
}
