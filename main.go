package main

import (
	"fmt"
	"time"
)

func main() {
	keyword, path := InputCommend()
	fmt.Printf("%s %s\n", keyword, path)
	start := time.Now()
	ReadPathType(path, keyword)
	end := time.Now()
	fmt.Println("실행 시간:", end.Sub(start))
}
