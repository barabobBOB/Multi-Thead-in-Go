package main

import (
	"fmt"
	"sync"
)

// readFiles 함수는 디렉토리 내 모든 파일에서 키워드를 비동기적으로 검색합니다.
func readFiles_multi(path string, keyword string) {
	paths, err := readFilesPath(path)
	if err != nil {
		fmt.Println("파일 경로를 읽는 도중 에러 발생:", err)
		return
	}

	var wg sync.WaitGroup
	for _, filePath := range paths {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			content := readFile(p, keyword)
			for _, str := range content {
				fmt.Println(str)
			}
		}(filePath)
	}

	wg.Wait() // 모든 고루틴이 완료될 때까지 대기
}
