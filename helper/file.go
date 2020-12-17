/*
@Time : 2020/12/17 9:30 上午
@Author : jia
@File : file.go
@Software : GoLand
*/

package helper

import (
	"bufio"
	"os"

	log "github.com/sirupsen/logrus"
)

//readLine 读取文件第一行
func ReadLine(path string) string {
	if !PathExists(path) {
		return ""
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer file.Close()

	var lineText string
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	lineText = scanner.Text()
	return lineText[:len(lineText)-1]
}

// PathExists 检查文件或目录是否存在
func PathExists(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
