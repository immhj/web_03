package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"sync"
	"time"
)

var (
	me1   chan string
	me2   chan string
	me3   chan string
	s     string
	t     string
	flag1 bool
	flag2 bool
	flag3 bool
	flag4 bool
	wg    sync.WaitGroup
)

func method1() {
	wg.Add(1)
	for flag3 == false {
		if len(me1) < 10 {
			me1 <- s
		}
	}
	wg.Done()
}
func method2() {
	for flag3 == false {
		if len(me1) > 0 && len(me2) < 5 && flag1 == false {
			t = <-me1
			fmt.Print("装弹->")
			flag1 = true
			me2 <- t
			time.Sleep(2 * time.Second)
		}
	}
}
func method3() {
	wg.Add(1)
	for flag3 == false {
		if len(me2) > 0 && len(me3) < 3 && flag2 == false {
			t = <-me2
			fmt.Print("瞄准->")
			flag2 = true
			me3 <- t
			time.Sleep(4 * time.Second)

		}
	}
	wg.Done()
}

func method4() {
	wg.Add(1)
	for flag3 == false {
		if len(me3) > 0 {
			t = <-me3
			fmt.Println("发射")
			flag1 = false
			flag2 = false
			time.Sleep(6 * time.Second)
		}
	}
	wg.Done()
}
func tion() {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		if event.Rune == 'q' {
			flag3 = true
			wg.Wait()
			flag4 = true
			return
		}
	}
}
func main() {
	s = "炮弹"
	me1 = make(chan string, 10)
	me2 = make(chan string, 5)
	me3 = make(chan string, 3)
	keyboard.Close()
	go method1()
	go method2()
	go method3()
	go method4()
	tion() //键盘读入
	if flag4 == true {
		return
	}
}
