package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

func isFileOrDirectory(path string) string {
	if strings.Contains(path, ".") {
		return "file"
	} else {
		return "directory"
	}
}

func textFileCheck(path string) bool {
	if strings.Contains(path, ".txt") {
		return true
	} else {
		return false
	}
}

func ReadPathType(path string, keyword string) {
	if isFileOrDirectory(path) == "file" {
		ReadFile(path, keyword)
	} else {
		ReadFiles(path, keyword)
	}
}

func ReadFilesPath(path string) ([]string, error) {
	var filesPath []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf(err.Error())
		}
		if !info.IsDir() {
			if textFileCheck(path) {
				filesPath = append(filesPath, path)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return filesPath, nil
}

func ReadFiles(path string, keyword string) {
	paths, _ := ReadFilesPath(path)

	var wg sync.WaitGroup
	results := make(chan []string, len(paths))

	for _, path := range paths {
		wg.Add(1)
		go func(p string, k string) {
			defer wg.Done()
			content := ReadFile(p, k)
			results <- content
		}(path, keyword)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for content := range results {
		for _, str := range content {
			fmt.Println(str)
		}
	}
}

func ReadFilesNo(path string, keyword string) {
	paths, err := ReadFilesPath(path)
	if err != nil {
		fmt.Println("Error reading paths:", err)
		return
	}

	for _, path := range paths {
		ReadFile(path, keyword)
	}
}

func ReadFile(path string, keyword string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s 파일을 읽는데 실패했습니다.\n", path)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	keywordLog := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			keywordLog = append(keywordLog, "File:"+path+"/ Line: "+strconv.Itoa(lineNumber))
		}
		lineNumber++
	}
	return keywordLog
}
