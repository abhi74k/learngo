package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type ShapeOp interface {
	Area() float64
}

type Square struct {
	a int
}

type Circle struct {
	r int
}

type Rectangle struct {
	w int
	h int
}

func (s *Square) Area() float64 {
	return float64(s.a * s.a)
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(float64(c.r), 2)
}

func (r *Rectangle) Area() float64 {
	return float64(r.h * r.w)
}

func PrintArea(op ShapeOp) {
	fmt.Println("Area: ", op.Area())
}

func CountChars(s string) map[string]int {
	r := make(map[string]int, len(s))

	for _, c := range s {
		r[strings.ToLower(string(c))]++
	}

	return r
}

func DeDup(nums []int) []int {
	seen := make(map[int]struct{}, len(nums))
	ret := make([]int, 0)

	for _, n := range nums {
		if _, ok := seen[n]; ok {
			continue
		}
		seen[n] = struct{}{}
		ret = append(ret, n)
	}

	return ret
}

func main() {

	fmt.Println("Abhinav")
	fmt.Println(strings.ToUpper("Abhinav"))

	const x = 1
	var y int = 2
	var z = 3
	a := 1

	fmt.Println(x, y, z, a)

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	if a == 1 {
		fmt.Println("a == 1")
	}

	err := fmt.Errorf("This is an error %d", a)
	fmt.Println("Error : ", err)

	circle := Circle{r: 10}
	square := Square{a: 10}
	rectangle := Rectangle{w: 10, h: 2}

	fmt.Println("Calling area on circle ", circle.Area())

	PrintArea(&circle)
	PrintArea(&square)
	PrintArea(&rectangle)

	names := []string{"Abhinav", "Martina"}
	fmt.Println(names)

	var hashmap map[string]int
	hashmap = make(map[string]int, 10)
	hashmap["a"] = 1
	fmt.Println(hashmap)
	fmt.Println(hashmap["b"])

	var s []int
	s = make([]int, 0, 10)
	s = append(s, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5)
	fmt.Println(s)

	fmt.Println(CountChars("Abhinav"))

	fmt.Println(DeDup([]int{1, 1, 2, 3, 2, 3}))

	fmt.Println([]int32("hi"))

	// TIME OPERATIONS
	fmt.Println(time.Now())

	timeNow := time.Now()

	fmt.Println(time.Since(timeNow))

	t1 := timeNow.Add(2 * time.Hour)
	fmt.Println(t1)

	t2 := timeNow.Add(-2 * time.Hour)
	fmt.Println(t2)

	d1 := t1.Sub(t2)
	fmt.Println(d1)

	tz, _ := time.LoadLocation("Asia/Calcutta")
	t3 := time.Now().In(tz)
	fmt.Println(t3)

	time.Sleep(time.Second)

	fmt.Println(t1.Before(t2))
	fmt.Println(t1.After(t2))
	fmt.Println(t1.Equal(t2))

	fmt.Println(t1.Format("01-02-2006"))
	fmt.Println(t2.Format("2006-01-02 15:04:05"))
}
