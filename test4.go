// 조건문

// if 문은 해당조건이 맞으면 {} 블럭 안의 내용을 실행한다. 
// Go 는 조건식을 괄호 안에 x 반드시 { 

// } 이런식으로 써야한다. 
// 이게 원래 교수님이 꼴값이라고 하셨는데... 
// C나 C++ 에서는 이렇게 쓰는게 가독성에 좋다하셨다.
// 교수님말이 다맞음..ㅋ
// {

// }

// if 문의 조건식은 반드시 Boolean 식으로 표현되어야 한다. 

if k == 1 {  // 같은 라인
	println("One")
}

// Go 에서는 if 문을 사용하기 전에 간단한 문장을 함께 실행할 수 있다. 
// 주의할 점은 이렇게 실행하는 문장은 if 문 안에서만 동작하고 펑한다는 것이다. 
// 예로

if val := i * 2; val < max {
	println(val)
}

val++ // if 문을 벗어난 문장은 에러, 근데 이건 미리 선언해둔 변수를 이용하면 문제 없을듯??



// ----------------------------
// Switch 
여러 값을 비교 or 다수의 조건식 체크 
Switch문 뒤에 하나의 변수를 지정하고, case문에 해당 변수가 가질 수 있는 값들을 지정하여, 
각 경우에 다른 문장 블럭들을 실행할 수 있다. 역시 C랑 비슷하군..

package main 
func main() {
	var name string 
	var category = 1

	switch category {
	case 1:
		name = "Tiger"
	case 2:
		name = "Lion"
	case 3, 4:
		name = "Cat"
	default:
		name = "Dog"
	}
	println(name)

	// Expression을 사용한 경우
	switch X := category << 2; x - 1 {
		// ...
	}
}

/*
C++, C#, Java 등과 다른 Go의 switch 문의 특징은?
1) Switch 뒤에 expression 이 없을 수 있다. 
: 다른 언어는 switch 키워드 뒤에 변수나 expression을 반드시 두지만,
  Go는 이를 쓰지 않아도 된다. 이 경우 Go는 switch expression을 true 로 생각하고,
  첫번째 case 문으로 이동하여 검사한다.

2) case 문에 expression을 쓸 수 있음. 
: 다른 언어의 case 문은 일반적으로 리터럴 값만을 갖지만, Go는 case 문에 복잡한 expression을 가질 수 있다. 
예로 C에서는 
switch 1: or switch 0: 등만 사용하였다. 

3) No default fall through
: 다른 언어의 case문은 break 를 쓰지 않는 한 다음 case 로 이동하지만 
  Go는 다른 case 로 가지 않는다. 

4) Type Switch 
: 다른 언어의 switch 는 일반적으로 변수의 값을 기준으로 case 로 분기하지만,
Go 는 그 변수의 Type에 따라 case 로 분기할 수 있다. */


func grade(score int) {
	switch {

	case score >= 90:
		println("A")

	case score >= 80:
		println("B")

	case score >= 70:
		println("C")

	case score >= 60:
		println("D")

	default:
		println("F")
	}
}


switch 변수의 타입을 검사하는 type switch 

switch v.(type) {
case int:
	println("int")

case bool:
	println("bool")

case string:
	println("string")

default:
	println("unknown")
}

Go 는 case 문 마지막에 break 가 없어도 알아서 빠져나온다.
Go 가 만약 C나 C#처럼 계속 밑의 문장들을 실행하게 하려면 
fallthrough 문을 명시해주면 된다. 

package main 

import "fmt"

func main() {
	check(2)
}

func check(val int) {
	switch val {
	case 1:
		fmt.println("1 이하")
		fallthrough
	
	case 2:
		fmt.println("2 이하")
		fallthrough
	
	case 3:
		fmt.println("3 이하")
		fallthrough

	default:
		fmt.println("default 도달")
	
	}
}