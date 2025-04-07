package myfun

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadJSONFile 通用函数，用于读取 JSON 文件并解析数据到指定的结构体类型
func ReadJSONFile(filePath string, data interface{}) error {
	// 读取文件内容
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("无法读取文件: %s", err)
	}
	// 解析 JSON 数据
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return fmt.Errorf("无法解析 JSON: %s", err)
	}
	return nil
}

// WriteJSONFile 通用函数，用于将指定的结构体数据写入到 JSON 文件中
func WriteJSONFile(filePath string, data interface{}) error {
	// 打开文件，如果文件不存在则创建
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("无法打开文件: %s", err)
	}
	defer file.Close() // 确保在函数退出时关闭文件

	// 将数据转换为 JSON 格式
	jsonData, err := json.MarshalIndent(data, "", "  ") // 使用缩进格式化输出
	if err != nil {
		return fmt.Errorf("无法转换数据为 JSON: %s", err)
	}

	// 写入 JSON 数据到文件
	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("无法写入文件: %s", err)
	}

	return nil
}

// json byte array to struct
func JsonBytes2Struct(jsonBytes []byte, result interface{}) error {
	err := json.Unmarshal(jsonBytes, result)
	return err
}

// json string to struct
func JsonString2Struct(jsonString string, result interface{}) error {
	return JsonBytes2Struct([]byte(jsonString), result)
}

// struct to json string
func Struct2JsonString(structt interface{}) (jsonString string, err error) {
	data, err := json.Marshal(structt)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
