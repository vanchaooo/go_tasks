package spm

import (

)

type User struct {
	ID int
	Name string
	Age int
	Email string
}

type Point struct {
	X int
	Y int
}

type Rectangle struct {
	Width int
	Height int
}

type Pair struct {
	A int
	B int
}

type Range struct {
	From int
	To int
}

type Student struct {
	Name string
	Grades []int
}

type Product struct {
	Name string
	Price int
}

type Employee struct {
	Name string
	Rate int
	Hours int
}

type ClockTime struct {
	Hour int
	Minute int
}

type Date struct {
	Year int
	Month int
	Day int
}



// 1
func NewUser(id int, name string, age int) *User {
	if id == 0 || age == 0 {
		return nil
	}
	if name == "" {
		return nil
	}

	user := &User {
		ID: id,
		Name: name,
		Age: age,
	}
	return user
}

// 2
func NewPoint(x, y int) *Point {
	if x == 0 || y == 0 {
		return nil
	}

	point := &Point{
		X: x,
		Y: y,
	}
	return point
}

// 3
func Area(r Rectangle) int {
	result := r.Height * r.Width
	return result
}

// 4
func IsOrigin(p Point) bool {
	if p.X == 0 && p.Y == 0 {
		return true
	}
	return false
}

// 5
func Older(a, b User) User {
	if a.Age > b.Age {
		return a
	}
	return b
}

// 6
func WithEmail(u User, email string) User {
	if email == "" {
		return u
	}
	u.Email = email
	return u
}

// 7
func SwapPair(p Pair) Pair {
	p.A, p.B = p.B, p.A
	return p
}

// 8
func ManhattanDistance(a, b Point) int {
	result := (a.X - b.X) + (a.Y - b.Y)
	return result
}

// 9
func NormalizeRange(r Range) Range {
	if r.From > r.To {
		r.From, r.To = r.To, r.From
	}
	return r
}

// 10
func MovePoint(p Point, dx, dy int) Point {
	p.X = dx
	p.Y = dy
	return p
}

// 11
func AverageGrade(s Student) float64 {
	if len(s.Grades) == 0 {
		return 0
	}

	sum := 0
	for _, v := range s.Grades {
		sum += v
	}
	result := sum / len(s.Grades)
	return float64(result)
}

// 12
func Discounted(p Product, percent int) Product {
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}

	p.Price = p.Price * (100 - percent) / 100
	return p
}

// 13
func Salary(e Employee) int {
	return e.Hours * e.Rate
}

// 14
func NormalizeTime(t ClockTime) ClockTime {
	total := t.Hour*60 + t.Minute

    total = total % 1440
    if total < 0 {
        total += 1440
    }

    t.Hour = total / 60
    t.Minute = total % 60
    return t
}

// 15
func CompareDate(a, b Date) int {
	if a.Year != b.Year {
        if a.Year < b.Year {
            return -1
        }
        return 1
    }

    if a.Month != b.Month {
        if a.Month < b.Month {
            return -1
        }
        return 1
    }

    if a.Day != b.Day {
        if a.Day < b.Day {
            return -1
        }
        return 1
    }

    return 0 
}