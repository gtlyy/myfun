package myfun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试 int ---> string
func TestIntToString(t *testing.T) {
	var a int = 123456789
	s := "123456789"
	ss := IntToString(a)
	assert.True(t, s == ss)
}

// 测试 int64 ---> string
func TestInt64ToString(t *testing.T) {
	var a int64 = 123456789
	s := "123456789"
	ss := Int64ToString(a)
	assert.True(t, s == ss)
}

// 测试 string ---> int64
func TestStringToInt64(t *testing.T) {
	var a int64
	a = 123456789
	s := "123456789"
	aa := StringToInt64(s)
	assert.True(t, a == aa)
}

// 测试 string ---> float64
func TestStringToFloat64(t *testing.T) {
	var rf float64
	rf = 0.00782
	s := "0.00782"
	ff := StringToFloat64(s)
	assert.True(t, ff == rf)
}

// 测试 float64 ---> string
func TestFloat64ToString(t *testing.T) {
	// 测试 取整
	f := 111.23459
	ss := Float64ToString(f, 0)
	assert.True(t, ss == "111")
	// 测试 指定小数位数
	ss2 := Float64ToString(f, 5)
	assert.True(t, ss2 == "111.23459")
}

// 测试 计算小数点后有几位小数
func TestCountFloat(t *testing.T) {
	f := 111.23459
	n := CountFloat(f)
	assert.True(t, n == 5)
}

// 测试：float64 ---> int64
func TestFloat64ToInt64(t *testing.T) {
	f := 222.48093
	a := Float64ToInt64(f)
	assert.True(t, a == 222)
}

// 测试：float64 ---> int
func TestFloat64ToInt(t *testing.T) {
	f := 222.48093
	a := Float64ToInt(f)
	assert.True(t, a == 222)
}

// 测试函数 In() ：查询字符串是否在数组中
func TestIn(t *testing.T) {
	strArray := []string{"apple", "banana", "cherry", "date"}
	target := "banana"
	assert.True(t, In(target, strArray))
}
