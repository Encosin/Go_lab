// Go도 여타 프로그래밍 언어들처럼 Data type이 존재한다. 
// 1) bool 
// 2) string 
// 3) int, uint 
// 4) float, complex  
// 5) byte rune 


// 문자열
// 리터럴은 ' ', " " 이 두가지는 사실 파이썬에서는 차이가 별로 없는데, 여기서는 따로 구분하네??
// ' ' 로 둘러싸인 문자열은 Raw String Literal 이라 하는데, 안에 있는 문자열에 \n 이 있을 경우
// 이를 들여쓰기로 해석하지 않는다. 복수 라인의 문자열을 표현할 때 자주 사용한다.

// " " 로 둘러싸인 문자열은 Interpreted String Literal 이라 부르는데, 인용부호 안의 Escape 문자열들은 특별한 의미로 해석된다.
// 이 문자열 안에 \n 이 있을 경우 이는 NewLine으로 해석된다. 
// 이중인용부호를 이용해 문자열을 여러 라인에 걸쳐 쓰기 위해서는 + 연산자를 이용해 결합하여 사용한다.

package main 

import "fmt"

func main() {
	// Raw String Literal 복수라인.
	rawLiteral := '아리랑\n
	아리랑\n
	아라리요'
	// ㅋ 이건 안될것 같은 느낌이 팍드네

	// Interpreted String Literal 
	interLitetal := "아리랑아리랑\n아라리요"
	// 아래와 같이 +를 사용하여 두 라인에 걸쳐 사용할 수도 있다.
	interLiteral := "아리랑아리랑\n" +
					"아라리요"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
}


// 데이터 타입 변환 
// : 다른 데이터 타입으로 변환하기 위해서는 T(v) 와 같이 표현하고 이를 Type Conversion 이라 부르는데,
//  T는 변환하고자 하는 타입, v는 변환될 값을 지정한 것.
 
// func main() {
// 	var i int = 100
// 	var u uint = uint(i)
// 	var f float32 = float32(i)
// 	println(f, u)

// 	str := "ABC"
// 	bytes := []byte(str)
// 	str2 := string(bytes)
// 	println(bytes, str2)
// }





