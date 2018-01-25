package main

import (
	"fmt"
	"math"
	"unicode"
)

type Result interface {
	String() string
	equal(s string) bool
}

func toResult(s string) Result {
	dots := 0
	nums := 0
	dotp := 0
	for i, c := range s {
		if unicode.IsNumber(c) {
			nums++
		} else if c == ',' || c == '.' {
			dots++
			dotp = i
		}
	}

	if dots == 1 && dots+nums == len(s) {
		var res float64
		fmt.Sscan(s, &res)
		return Float{
			res,
			len(s) - 1 - dotp,
		}
	}

	if nums == len(s) {
		var res int64
		fmt.Sscan(s, &res)
		return Num(res)
	}

	return Text(s)

}

type Num int64

type Float struct {
	x         float64
	precesion int
}

type Text string

type Problem struct {
	statement string
	result    Result
	level     int
	typ       string
	imag      string
}

func (a Num) equal(b string) bool {
	return fmt.Sprint(int64(a)) == b
}

func (a Float) equal(b string) bool {
	c := float64(a.x)
	fmt.Println("Is ", math.Abs(c-a.x), "less than", math.Pow(10, -float64(a.precesion)), "and relative error is ", math.Abs(c-a.x)/c)
	if math.Abs(c-a.x) < math.Pow(10, -float64(a.precesion))/2 && math.Abs(c-a.x)/c < math.Pow(10, -float64(a.precesion)/2) {
		return true
	}
	return false
}

func (a Text) equal(b string) bool {
	return string(a) == b
}

func (a Num) String() string {
	return fmt.Sprint(int64(a))
}

func (a Text) String() string {
	return string(a)
}

func (a Float) String() string {
	return fmt.Sprint(a.x)
}

func (p Problem) solved(answer string) bool {
	fmt.Printf("Problem je %v a odpoved %v\n", p.result, answer)
	if p.result.equal(answer) {
		return true
	}
	return false
}

func (p Problem) String() string {
	return fmt.Sprintf("%v -> %v | %v of type %v", p.statement, p.result, p.level, p.typ)
}
