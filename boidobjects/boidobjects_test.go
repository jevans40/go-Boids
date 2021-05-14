package boidobjects

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkBoids(b *testing.B) {
	Squares := make(map[int]Square)
	for i := 0; i < 100000; i++ {
		Squares[i] = Square{}
	}
	fmt.Print(int((time.Second * 180) / (time.Millisecond * 16)))
	x := int((time.Second * 180) / (time.Millisecond * 16))
	_ = x
	tosend := make([]float32, 28)
	lastTime := time.Now()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			if j%120 == 0 {
				fmt.Printf("%v per frame on average\n", time.Since(lastTime)/120)
				lastTime = time.Now()
			}
			for _, v := range Squares {
				v.Update(i)
				v.Render(tosend)
			}
		}
	}
}
