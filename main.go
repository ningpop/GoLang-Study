package main // 사용하는 패키지명

import (
	"fmt"

	"github.com/ningpop/learngo/something"
)

func main() {
	fmt.Println("Hello World!") // Hello World 출력
	something.SayHello()        // 대문자로 시작하므로 export해서 사용 가능
	// something.sayBye() 		// 소문자로 시작하므로 export해서 사용 불가
}
