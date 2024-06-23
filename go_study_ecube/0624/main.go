package main

import (
	"fmt"
	"sync"
	"time"
)

type Calculator interface {
	plus() int
	minus() int
}

type Parameter struct {
	n1, n2 int
}

func (p Parameter) plus() int {
	return p.n1 + p.n2
}

func (p Parameter) minus() int {
	return p.n1 - p.n2
}

func say(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s, "***", i+1)
	}
}

func main() {
	// interface, struct, method
	var p Calculator = Parameter{10, 20}
	fmt.Println(p.plus())
	fmt.Println(p.minus())

	// goroutine
	// 함수를 동기적으로 실행
	say("Sync")

	// 함수를 비동기적으로 실행
	go say("Async1")
	go say("Async2")
	go say("Async3")

	// 3초 대기
	time.Sleep(time.Second * 1)

	var wait sync.WaitGroup
	wait.Add(2) // 몇개의 goroutine을 기다릴지 설정 만약 숫자보다 goroutine이 적으면 deadlock 발생, 많으면 2개까지만 기다리고 끝나는 경우가 생김

	go func(a, b int) {
		defer wait.Done()
		fmt.Println(a, "+", b, "=", a+b)
	}(1, 2)

	go func(a, b int) {
		defer wait.Done()
		fmt.Println(a, "-", b, "=", a-b)
	}(3, 2)

	wait.Wait()

	// go channel
	// 정수형 채널을 생성한다
	ch := make(chan bool)

	go func() {
		ch <- true //채널에 값을 보낸다
	}()

	i := <-ch // 채널로부터 값을 받는다
	// 채널로부터 값을 받을 때까지 대기한다 -> 락을 구현할 수 있을 것 같은데?
	fmt.Println(i)

	// // unbuffered channel
	// // 채널에 버퍼를 설정하지 않으면 채널에 값을 보내거나 받을 때까지 대기한다
	// // 채널에 값을 보내거나 받을 때까지 대기한다
	// unbufferedCh := make(chan bool)
	// unbufferedCh <- true 수신자가 없기 때문에 deadlock 발생
	// println(<-unbufferedCh)

	// buffered channel
	// 채널에 버퍼를 설정하면 채널에 값을 보내거나 받을 때까지 대기하지 않는다
	bufferedCh := make(chan bool, 2)
	bufferedCh <- true
	bufferedCh <- false
	close(bufferedCh) // 채널을 닫는다. 채널을 받으면 수신은 가능하지만 송신은 불가능하다.
	fmt.Println("1", <-bufferedCh)
	fmt.Println("2", <-bufferedCh)
	fmt.Println("2", <-bufferedCh) // 마지막 값으로 false가 들어가 있기 때문에 false 출력

	ch2 := make(chan string, 1)
	sendChan(ch2)
	receiveChan(ch2)

	// channel range

	ch4 := make(chan int, 2)

	// 채널에 송신
	ch4 <- 1
	ch4 <- 2

	// 채널을 닫는다
	close(ch4)

	// 방법1
	// 채널이 닫힌 것을 감지할 때까지 계속 수신
	/*
	   for {
	       if i, success := <-ch4; success {
	           println(i)
	       } else {
	           break
	       }
	   }
	*/

	// 방법2
	// 위 표현과 동일한 채널 range 문
	for i := range ch4 {
		println(i)
	}

	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

	// for루프 안에 select문을 쓰면서 goroutine이 모두 끝나길 기다린 후 Exit 레이블로 이동
EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")

		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}
}

func sendChan(ch chan<- string) {
	ch <- "Data"
	// x := <-ch // 송신전용 채널에서 수신을 하기때문에 에러발생
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
