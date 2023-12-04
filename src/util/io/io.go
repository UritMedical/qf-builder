package io

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// FindFiles 返回指定目录下满足条件的文件名列表 如*xxx.xxx
func FindFiles(path string, pattern string) ([]string, error) {

	output := make([]string, 0)
	// 定义正则表达式匹配文件名
	filePattern := regexp.MustCompile(pattern)
	// 遍历指定路径下的所有文件
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 判断是否为文件且文件名匹配正则表达式
		if !info.IsDir() && filePattern.MatchString(info.Name()) {
			output = append(output, filePath)
		}
		return nil
	})
	return output, err
}

// SaveToFile 保存到文件 且会自动创建文件夹
func SaveToFile(filePath string, data []byte) error {
	// 获取文件所在目录
	dir := filepath.Dir(filePath)

	// 创建目录，如果不存在的话
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("无法创建目录：%s", err)
	}

	// 将字节数组写入文件
	if err := os.WriteFile(filePath, data, os.ModePerm); err != nil {
		return fmt.Errorf("无法写入文件：%s", err)
	}

	//fmt.Printf("文件保存成功：%s\n", filePath)
	return nil
}
