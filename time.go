package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	go test2()
	go test3()
	go test4()
	go test5()
	go test6()
	go test8() //interval
	go test9()

	// formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
	// 	t.Year(), t.Month(), t.Day(),
	// 	t.Hour(), t.Minute(), t.Second())
	for i := 0; i < 3; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println("time now: ", time.Now().UnixMilli())
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Good bye")
	fmt.Println("----------------------------------test 7---------------------------------------")
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	// select {
	// // case <-time.After(2 * time.Second):
	// // 	fmt.Println("overslept")
	// case <-ctx.Done():
	// 	fmt.Println("giờ hiện tại: ", time.Now()) // prints "context deadline exceeded"
	// }

}

func test2() {
	t := time.Now()
	// formatted := fmt.Sprintf("date time: ", t.Day())
	fmt.Println("date time", t)
}
func test3() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	go func(ctx context.Context) {
		for {
			time.Sleep(time.Second * 3)
			select {
			case <-ctx.Done():
				fmt.Println("Cancelled:", ctx.Err())
				return
			default:
				fmt.Println("Sleep")
			}
		}
	}(ctx)
	time.Sleep(10 * time.Second)
}
func test4() {
	nix := time.Unix(0, 1592190294764144364).UTC()
	fmt.Println("Số phút: ", nix.Minute())
}
func test5() {
	d := time.Unix(0, 592190385).UTC()
	fmt.Println("day: ", d.Weekday())
}
func test6() {
	fmt.Println("Có 4 loại :seconds, milliseconds, microseconds and nanoseconds")
}
func test8() {
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(time.Now().Unix(), "done")
	}
}
func x(ctx context.Context) {

}
func test9() {
	step := time.Duration(100 * time.Millisecond)
	f := func() {
		fmt.Println("I'm study")
	}
	Timer1 := time.AfterFunc(step, f)
	defer Timer1.Stop()
	time.Sleep(10 * time.Second)
	/*time.AfterFunc(time.Millisecond*100, func() {
		fmt.Println("i'm study")
	})
	time.Sleep(time.Second)*/
}
