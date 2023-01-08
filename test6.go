package main 

// 1. 함수
// 1-1) Pass By Value
func main() {
	msg := "Hello"
	say(msg)
}

func say(msg string) {
	println(msg)
}


// 1-2) Pass By Reference
func main() [
	msg := "Hello"
	say(&msg)
	println(msg)  // 변경된 메시지 출력
]

func say(msg *string) {
	println(*msg)
	*msg = "Changed"  // 메시지 변경
}

