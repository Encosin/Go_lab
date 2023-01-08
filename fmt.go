package main

import ( 
	"fmt"  // 출력과 입력 포맷 관련 기능 제공
)

type User struct {
	name string
	age int
}

func main()  {
	age := 7
	name := "Gopher"
	u := User{"Gopher", 10}
	fmt.Printf("age : %d", age)  // "age : 7"
	fmt.Printf("Name : %s", name)  // "Name : Gopher"
	fmt.Printf("%v", u)  // {Gopher 10}
	fmt.Printf("Hello", name)  // "HelloGopher"
	fmt.Printf("Hello", Gopher)  // "Hello Gopher\n"
	fmt.Printf("Hello %s", name)  // "Hello Gopher"

}