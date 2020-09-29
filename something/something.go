package something

import "fmt"

func sayBye() { // 소문자로 시작하는 함수는 private
	fmt.Println("Bye")
}

func SayHello() { // 대문자로 시작하는 함수는 export해서 사용 가능
	fmt.Println("Hello")
}
