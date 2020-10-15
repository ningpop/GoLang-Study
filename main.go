package main // 사용하는 패키지명

import (
	"fmt"
	"strings"

	"github.com/ningpop/learngo/something"
)

func multifly(a int, b int) int {
	/*
		함수는 아래처럼 작성할 수도 있다.
		func multifly (a, b int) int {
			...
		}
	*/
	return a * b
}

/*
func lenAndUpper(name string) (int, string) { // 반환값을 두개 갖는 함수
	return len(name), strings.ToUpper(name)
}
*/

/* #1.4 Functions part Two */
func lenAndUpper(name string) (length int, uppercase string) { // 'naked' return 사용
	defer fmt.Println("I'm done") // defer 문 사용: 함수의 실행이 끝난 뒤 최종적으로 실행되는 문장
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) { // 여러개의 입력값을 한번에 받는 함수
	fmt.Println(words)
}

/* #1.5 for, range, ...args */
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers { // 매개변수의 numbers의 범위만큼 순회
		total += number
	}
	return total
}

/* #1.6 If with a Twist */
/*
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // variable expression
		return false
	}
	return true
}
*/

/* #1.7 Switch */
func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

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

	/* #1.3 Functions part One */
	fmt.Println(multifly(2, 2))

	totalLength, upperName := lenAndUpper("ningpop") // 두개의 값을 반환하는 함수
	fmt.Println(totalLength, upperName)

	repeatMe("ningpop", "lynn", "dal", "marl", "flynn") // 여러개의 입력값을 한번에 전달

	/* #1.5 for, range, ...args */
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	/* #1.6 If with a Twist */
	fmt.Println(canIDrink(16))
}
