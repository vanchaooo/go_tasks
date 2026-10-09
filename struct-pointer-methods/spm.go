package spm

import (

)

type User struct {
	ID int
	Name string
	Age int
	Email string
	City string
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

type Employees struct {
	Name string
	Salary int
}

type Task struct {
	Status string
}

type OrderItem struct {
	Name string 
	Price int
	Quantity int
}

type Order struct {
	ID int;
	Items []OrderItem
}

type Cart struct {
	Items []OrderItem
}

type Address struct {
	City string
	Street string
	House int
}

type Person struct {
	Name string
	Address Address
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

// 16
func FindUserByID(users []User, id int) (User, bool) {
	if len(users) == 0 || id == 0 {
		return User{}, false
	}

	for _, v := range users {
		if v.ID == id {
			return v, true
		}
	}

	return User{}, false
}

// 17
func Adults(users []User) []User {
	if len(users) == 0 {
		return []User{}
	}

	result := []User{}
	for _, v := range users {
		if v.Age >= 18 {
			result = append(result, v)
		}
	}
	return result
}

// 18
func IndexUsers(users []User) map[int]User {
	if len(users) == 0 {
		return map[int]User{}
	}

	result := make(map[int]User)
	for _, v := range users {
		result[v.ID] = v
	}
	return result
}

// 19
func GroupUsersByCity(users []User) map[string][]User {
	if len(users) == 0 {
		return map[string][]User{}
	}

	result := make(map[string][]User)
    for _, u := range users {
        result[u.City] = append(result[u.City], u)
    }
    return result
}

// 20
func HighestPaid(items []Employees) (Employees, bool) {
	if len(items) == 0 {
		return Employees{}, false
	}

	result := items[0]
	for _, v := range items {
		if v.Salary > result.Salary {
			result = v
		}
	}
	return result, true
}

// 21
func CountByStatus(tasks []Task) map[string]int {
	result := make(map[string]int)

	for _, v := range tasks {
		result[v.Status]++
	}
	return result
}

// 22
func OrderTotal(o Order) int {
	result := 0

	for _, v := range o.Items {
		result += v.Price * v.Quantity
	}
	return result
}

// 23
func CartItemsCount(c Cart) int {
	result := 0

	for _, v := range c.Items {
		result += v.Quantity
	}
	return result
}

// 24
func BestStudent(students []Student) (Student, bool) {
	if len(students) == 0 {
		return Student{}, false
	}

	var bestStudent Student
	maxAVG := 0.0
	for _, v := range students {
		var avg float64
		if len(v.Grades) > 0 {
			sum := 0
			for _, j := range v.Grades {
				sum += j
			}
			avg = float64(sum) / float64(len(v.Grades))
		}

		if avg > maxAVG {
			maxAVG = avg
			bestStudent = v
		}
	}
	return bestStudent, true
}

// 25
func SameCity(a, b Person) bool {
	if a.Address.City == b.Address.City {
		return true
	}
	return false
}