package myfun

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试 ReadJSONFile(): 通用函数，用于读取 JSON 文件并解析数据到指定的结构体类型
func TestReadJSONFile(t *testing.T) {
	// Create a temporary JSON file for testing
	tempFile, err := ioutil.TempFile("", "test.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write test JSON data to the temporary file
	testData := []byte(`{"name": "John", "age": 30}`)
	if _, err := tempFile.Write(testData); err != nil {
		t.Fatalf("Failed to write test data to file: %v", err)
	}

	// Call the ReadJSONFile function with the temporary file
	var result struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	err = ReadJSONFile(tempFile.Name(), &result)
	if err != nil {
		t.Fatalf("ReadJSONFile failed: %v", err)
	}

	// Check if the data was parsed correctly
	expectedResult := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "John",
		Age:  30,
	}
	assert.True(t, reflect.DeepEqual(result, expectedResult))
}

// TestWriteJSONFile tests the WriteJSONFile function.
func TestWriteJSONFile(t *testing.T) {
	// Define a sample struct for testing
	type Sample struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "test_*.json")
	if err != nil {
		t.Fatalf("failed to create temporary file: %s", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up the temporary file after the test

	// Sample data to write to the file
	data := Sample{Name: "Alice", Age: 30}

	// Call the WriteJSONFile function
	err = WriteJSONFile(tmpFile.Name(), data)
	assert.NoError(t, err, "expected no error while writing JSON to file")

	// Read back the content of the file to verify
	fileData, err := os.ReadFile(tmpFile.Name())
	assert.NoError(t, err, "expected no error while reading the JSON file")

	// Expected JSON output
	expectedJson := `{
  "name": "Alice",
  "age": 30
}`
	assert.JSONEq(t, expectedJson, string(fileData), "the written JSON file content does not match the expected value")
}

// 测试 JsonBytes ---> Struct
func TestJsonBytes2Struct(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	jsonBytes := []byte(`{"name": "Alice", "age": 30}`)
	var result Person

	err := JsonBytes2Struct(jsonBytes, &result)

	assert.NoError(t, err, "expected no error during unmarshalling")
	assert.Equal(t, Person{Name: "Alice", Age: 30}, result, "expected result does not match")
}

// 测试 Struct ---> JsonString
func TestStruct2JsonString(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	person := Person{Name: "Alice", Age: 30}
	expectedJson := `{"name":"Alice","age":30}`

	jsonString, err := Struct2JsonString(person)

	assert.NoError(t, err, "expected no error during marshalling")
	assert.JSONEq(t, expectedJson, jsonString, "expected JSON string does not match")
}
