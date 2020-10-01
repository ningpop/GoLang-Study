package main // 사용하는 패키지명

import (
	"fmt"

	"github.com/ningpop/learngo/something"
)

// name := "ningpop" // 사용 불가
func main() {
	/* #1.1 Packages and Imports */
	fmt.Println("Hello World!") // Hello World 출력
	something.SayHello()        // 대문자로 시작하므로 export해서 사용 가능
	// something.sayBye() 		// 소문자로 시작하므로 export해서 사용 불가

	/* #1.2 Variables and Constants */
	const nameConst string = "ningpop" // name이라는 이름을 가진 string형 상수에 "ningpop"이라는 문자열을 저장
	// nameConst = "Lisa"              // name은 상수이기때문에 값 변경이 불가능
	fmt.Println(nameConst)

	var nameVar string = "ningpop" // name이라는 이름을 가진 string형 변수에 "ningpop"이라는 문자열을 저장
	nameVar = "Lisa"               // name은 변수이기때문에 값 변경 가능
	fmt.Println(nameVar)

	name := "ningpop" // 사용 가능
	fmt.Println(name)
}
