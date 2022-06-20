package main

import (
	"context"
	"fmt"
	"time"
)

func x(ctx context.Context) context.Context {
	return context.WithValue(context.Background(), "current", time.Now().UnixNano())
}
func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test9()
	test8() //interval
}

// formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
// 	t.Year(), t.Month(), t.Day(),
// 	t.Hour(), t.Minute(), t.Second())
func test1() {
	fmt.Println("----------------------------------bài 1-----------------------------------------")
	for i := 0; i < 3; i++ {
		fmt.Println("time now: ", time.Now().UnixMilli())
		time.Sleep(3 * time.Second)
	}
	fmt.Println("kết thúc")

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
	fmt.Println("----------------------------bài 2-------------------------------")
	t := time.Now().Unix()
	// formatted := fmt.Sprintf("date time: ", t.Day())
	fmt.Println("date time", t)
}

func test3() {
	fmt.Println("----------------------------bài 3-------------------------------")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	for i := 1; i <= 3; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Cancelled:", ctx.Err())
			return
		case <-time.After(3 * time.Second):
			fmt.Println("Sleep lan thu: ", i)
		}
	}
}

func test4() {
	fmt.Println("----------------------------bài 4-------------------------------")

	nix := time.Unix(0, 1592190294764144364).UTC()
	fmt.Println("Số phút: ", nix.Minute())
}
func test5() {
	fmt.Println("----------------------------bài 5-------------------------------")
	d := time.Unix(0, 592190385).UTC()
	fmt.Println("day: ", d.Weekday())
}

func test6() {
	fmt.Println("----------------------------bài 6-------------------------------")

	fmt.Println("Có 4 loại :seconds, milliseconds, microseconds and nanoseconds")
}

func test7() {
	fmt.Println("----------------------------bài 7-------------------------------")

	ctx2 := context.Background()
	ctx2 = x(ctx2)

	currentTime := ctx2.Value("current").(int64)
	fmt.Println(currentTime)

	time.Sleep(time.Second * 3)
	currentTimeAfter3s := time.Now().UnixNano()
	fmt.Println(currentTimeAfter3s)

	result := currentTimeAfter3s - currentTime
	fmt.Println(result)

}

func test8() {
	fmt.Println("----------------------------bài 8-------------------------------")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancelled:", ctx.Err())
			return
		case <-time.After(100 * time.Millisecond):
			fmt.Println(time.Now().Unix(), "done")
		}
	}
}

func test9() {
	fmt.Println("----------------------------bài 9------------------------------")
	f := func() {
		fmt.Println("I'm study")
	}
	Timer1 := time.AfterFunc(time.Duration(100*time.Millisecond), f)
	defer Timer1.Stop()
	// Calling sleep method
	time.Sleep(2 * time.Second)
	/*time.AfterFunc(time.Millisecond*100, func() {
		fmt.Println("i'm study")
	})
	time.Sleep(time.Second)*/
}
