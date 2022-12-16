// 반복문
// C 언어와 다른 점은 while문이 없고, for문만 있다는 것이다.

package main

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)
}

// 조건식만 쓰는 for 루프
func main() {
	n := 1
	for n < 100 {
		n *= 2
		// if n > 90 {
		// 	break
		// }
		println(n)
	}
}

func main() {
	for {
		println("Infinite loop")
	}
}

// 4. for range 문
names := []string{"홍길동", "이순신", "강감찬"}
for index, name := range names {
	println(index, name)
}

// 5. break, continue, goto 
// package main
func main() {
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue  // for 루프 시작으로
		}
		a++
		if a > 10 {
			break
		}
	}
	if a == 11 {
		goto END  // goto 사용 예(for루프와 관계없이 사용가능)
	}
	println(a)

END:
	println("End")
}