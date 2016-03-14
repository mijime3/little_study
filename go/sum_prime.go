package main

import (
//	"fmt"
	"flag"	// コマンドラインオプションのパーサー
//	"time"
	"sync"
)

var num = flag.Int("num", 1000, "")
var cnt = flag.Int("cnt", 1, "")

func main() {
	flag.Parse()	// パラメータリストを調べてflagに設定

	var wg sync.WaitGroup

	for i:=0; i< *cnt; i++ {
		wg.Add(1)
		go func(targetNum int){
			defer wg.Done()
			sum_prime(targetNum)
		}(*num+i*100)
	}
	wg.Wait()
}

func sum_prime(num int) {
	//var start,end time.Time
	//start = time.Now()
	//fmt.Printf("num=%d, start=%s\n", num, start)

	var sum int = 0
	for i := int(1); i <= num ; i++ {
		if is_prime(i) {
			sum += i
		}
	}
	//end = time.Now()

	//fmt.Printf("num=%d, sum=%d, start=%s, end=%s\n", num, sum, start, end)
	
	//return sum
}


func is_prime(n int) bool{                                                                                                                                                      
        for i := 2; i < n; i++ {
                if n % i == 0 {
                        return false;
                }                                                                                                                                                               
        }
        return true;
}                                                                                                                                                                               
