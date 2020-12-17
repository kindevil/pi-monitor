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
	"strings"

	log "github.com/sirupsen/logrus"
)

func ReadLines(filename string) ([]string, error) {
	return ReadLinesOffsetN(filename, 0, -1)
}

func ReadLinesOffsetN(filename string, offset uint, n int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer f.Close()

	var ret []string

	r := bufio.NewReader(f)
	for i := 0; i < n+int(offset) || n < 0; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if i < int(offset) {
			continue
		}
		ret = append(ret, strings.Trim(line, "\n"))
	}

	return ret, nil
}

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
