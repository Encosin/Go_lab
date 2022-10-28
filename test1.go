package main

func main() {
	println("Test")
}

// 변수는 Go 키워드 var를 사용하여 선언한다. 
var a int

// 이렇게 변수 선언문에서 초기값을 할당할 수도 있다. 
var f float32 = 11.

// 일단 선언된 변수는 그 뒤 문장에서 해당 타입의 값을 할당할 수 있다.
a = 10
f = 12.0

// 변수를 여러개 지정할 시 마지막에만 타입을 붙여준다.
var i, j, k int 

// 복수 변수들의 값 지정을 해주고 싶은 경우 타입 뒤에 숫자를 넣는다. 
// 이때 값들은 순서에 맞게 들어간다.
var i, j, k int = 1, 2, 3


// 변수를 선언하면서  초기값을 지정하지 않으면, Zero value를 기본적으로 할당한다.
// 숫자형 0, bool 타입 false, string형 """"(빈문자열)


// Go에서는 할당되는 값을 보고 역으로 추론할 수 있는데, 
// 아래 코드에서 n는 정수형 1이 할당, s는 문자열로 Hi가 할당된다.
var n = 1 var s = "Go is so dif"


// 상수 선언 방법, 여기서 상수는 고정된 값을 의미한다.
const c int = 10
const s string = "Hi"

// 굳이 int or string을 쓰지 않아도 사실 알아서 추리하긴한다. 근데 썩 좋은 방법은 아니니 타입을 명시해주는게 좋다.
const c = 10
const s = "Hi"

// 상수는 여러개로 묶는 것도 가능하다.
const (
	enco = "Good Dev :)"
	enco_2 = "Cloud Engineer"

)

// 상수값을 0부터 순차적으로 부여하기 위해 iota라는 identifier를 사용할 수도 있다.
const (
	Apple = iota // 0 
	Grape 		 // 1
	Orange		 // 2
)


// 예약 키워드라는게 있다. 예로 파이썬의 def, init, import 등이 있는데, 
// 물론 Go 에도 있다. 
1) break      default 		func    interface   select
2) case		  defer			go 		map 		struct
3) chan 	  else			goto	package 	switch 
4) const 	  fallthrough	if 		range		type
5) continue   for			import  return		var








