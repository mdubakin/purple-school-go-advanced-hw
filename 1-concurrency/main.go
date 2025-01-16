package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// G1 генерирует 10 чисел от 0 до 100 и отправляет по одному числу в канал numCh
	// G2 получает числа из numCh, считает квадрат числа и отправить в канал squareCh
	// Gmain читает из канала squareCh и направляет число в stdout

	var wg sync.WaitGroup
	numCh := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		r := rand.New(rand.NewSource(time.Now().UnixMilli()))
		for range 10 {
			numCh <- r.Intn(100 + 1)
		}
		close(numCh)
	}()

	squareCh := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range numCh {
			squareCh <- num * num
		}
		close(squareCh)
	}()

	for square := range squareCh {
		fmt.Printf("%v ", square)
	}
	// wg.Wait() - кажется, что в этом случае необходимости в дополнительной синхронизации нет,
	// так как горутины синхронизированы каналами
}
