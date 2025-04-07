// 功能：提供一些数据转化函数，以及一些常用函数。
package myfun

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 数据转换 =========================================== Start
func IntToString(arg int) string {
	return strconv.Itoa(arg)
}

func Int64ToString(arg int64) string {
	return strconv.FormatInt(arg, 10)
}

func StringToInt64(arg string) int64 {
	value, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		return 0
	} else {
		return value
	}
}

func StringToInt(arg string) int {
	value, err := strconv.Atoi(arg)
	if err != nil {
		return 0
	} else {
		return value
	}
}

func StringToFloat64(arg string) float64 {
	value, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return 0
	} else {
		return value
	}
}

func Float64ToString(arg float64, bit int) string {
	if bit == 0 {
		a := Float64ToInt64(arg)
		return Int64ToString(a)
	}
	format1 := "%." + IntToString(bit) + "f"
	return fmt.Sprintf(format1, arg)
	// return fmt.Sprint(arg)
}

// 计算小数点后有几位小数
func CountFloat(f float64) int {
	numstr := fmt.Sprint(f)
	tmp := strings.Split(numstr, ".")
	if len(tmp) <= 1 {
		return 0
	}
	return len(tmp[1])
}

// 计算小数点后有几位小数
func CountStrFloat(sf string) int {
	tmp := strings.Split(sf, ".")
	if len(tmp) <= 1 {
		return 0
	}
	return len(tmp[1])
}

// float64 to int64
func Float64ToInt64(f float64) int64 {
	s := fmt.Sprintf("%1.0f", f)
	return StringToInt64(s)
}

// float64 to int
func Float64ToInt(f float64) int {
	s := fmt.Sprintf("%1.0f", f)
	return StringToInt(s)
}

// 数据转换 =========================================== End

// 字符串函数 =========================================== Start
// 将[]string{"aaa", "bbb", "ccc"}  --->  "aaa,bbb,ccc"
func ArrayToString(a []string) string {
	// var s string
	// for i, x := range a {
	// 	s = s + x
	// 	if i != len(a)-1 {
	// 		s = s + ","
	// 	}
	// }
	// return s
	return strings.Join(a, ",")
}

// 查询字符串是否在数组中的函数
func In(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

// 字符串函数 =========================================== End

// 其他，如下：

// go 要写太多的错误检测函数了
func IfError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}

// 实现三目运算：a == b ? c : d
func T3O(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}
