package main

import (
	"fmt"
	"time"
)

func main() {
	commend := InputCommend()

	start := time.Now()

	FindKeyword(commend.path, commend.keyword)

	end := time.Now()

	fmt.Println("실행 시간:", end.Sub(start))
}
