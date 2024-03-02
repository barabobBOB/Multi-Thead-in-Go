package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FindKeyword 함수는 주어진 경로에서 키워드를 찾습니다.
// 파일 또는 디렉토리 여부에 따라 적절한 함수를 호출합니다.
func FindKeyword(path string, keyword string) {
	if isFileOrDirectory(path) == "file" {
		// 단일 파일일 경우
		results := readFile(path, keyword)
		for _, result := range results {
			fmt.Println(result)
		}
	} else {
		// 디렉토리일 경우
		readFiles(path, keyword)
	}
}

// readFilesPath 함수는 주어진 디렉토리 내 모든 텍스트 파일의 경로를 반환합니다.
func readFilesPath(path string) ([]string, error) {
	var filesPath []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && textFileCheck(path) {
			filesPath = append(filesPath, path)
		}
		return nil
	})
	return filesPath, err
}

// isFileOrDirectory 함수는 주어진 경로가 파일인지 디렉토리인지 확인합니다.
func isFileOrDirectory(path string) string {
	if strings.Contains(path, ".") {
		return "file"
	}
	return "directory"
}

// textFileCheck 함수는 주어진 경로가 텍스트 파일인지 확인합니다.
func textFileCheck(path string) bool {
	return strings.HasSuffix(path, ".txt")
}

// readFiles 함수는 디렉토리 내 모든 파일에서 키워드를 비동기적으로 검색합니다.
func readFiles(path string, keyword string) {
	paths, _ := readFilesPath(path)

	for _, path := range paths {
		content := readFile(path, keyword)
		for _, str := range content {
			fmt.Println(str)
		}
	}
}

// readFile 함수는 주어진 파일에서 키워드를 검색하고 결과를 반환합니다.
func readFile(path string, keyword string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file %s\n", path)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	var keywordLog []string
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), keyword) {
			keywordLog = append(keywordLog, fmt.Sprintf("File: %s / Line: %d", path, lineNumber))
		}
		lineNumber++
	}
	return keywordLog
}
