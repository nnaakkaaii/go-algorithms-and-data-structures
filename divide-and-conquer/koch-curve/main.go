package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const th = math.Pi * 60.0 / 180.0

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	point := new(Point)
	point.x = x
	point.y = y
	return point
}

func koch(n int, a *Point, b *Point) {
	if n == 0 {
		return
	}

	s := NewPoint(
		(2.0*a.x+1.0*b.x)/3.0,
		(2.0*a.y+1.0*b.y)/3.0,
	)
	t := NewPoint(
		(1.0*a.x+2.0*b.x)/3.0,
		(1.0*a.y+2.0*b.y)/3.0,
	)
	u := NewPoint(
		(t.x-s.x)*math.Cos(th)-(t.y-s.y)*math.Sin(th)+s.x,
		(t.x-s.x)*math.Sin(th)+(t.y-s.y)*math.Cos(th)+s.y,
	)

	koch(n-1, a, s)
	fmt.Printf("%.8f %.8f\n", s.x, s.y)
	koch(n-1, s, u)
	fmt.Printf("%.8f %.8f\n", u.x, u.y)
	koch(n-1, u, t)
	fmt.Printf("%.8f %.8f\n", t.x, t.y)
	koch(n-1, t, b)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())

	a := NewPoint(0, 0)
	b := NewPoint(100, 0)

	fmt.Printf("%.8f %.8f\n", a.x, a.y)
	koch(n, a, b)
	fmt.Printf("%.8f %.8f\n", b.x, b.y)
}
