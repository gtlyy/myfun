package myfun

import (
	"reflect"
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

// 测试：实现三目运算：a == b ? c : d
func TestT3O(t *testing.T) {
	type testStruct struct{ field int }

	tests := []struct {
		name      string
		condition bool
		trueVal   interface{}
		falseVal  interface{}
		want      interface{}
	}{
		{"true returns int", true, 42, 0, 42},
		{"false returns string", false, "apple", "banana", "banana"},
		{"true with nil", true, nil, "not-nil", nil},
		{"false with nil", false, "not-nil", nil, nil},
		{"different types", true, 3.14, "pi", 3.14},
		{"struct comparison", true, testStruct{5}, testStruct{10}, testStruct{5}},
		{"pointer comparison",
			true,
			&testStruct{7},
			&testStruct{9},
			&testStruct{7}},
		{"typed nil pointer",
			false,
			(*testStruct)(nil),
			(*testStruct)(nil),
			(*testStruct)(nil)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := T3O(tt.condition, tt.trueVal, tt.falseVal)

			// 特殊处理nil比较
			if tt.want == nil {
				if got != nil {
					t.Errorf("expected nil, got %v (%T)", got, got)
				}
				return
			}

			// 使用反射处理复杂类型比较
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("T3O() = %v (%T), want %v (%T)",
					got, got, tt.want, tt.want)
			}
		})
	}
}
