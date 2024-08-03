package cool_obj

type Name string

type Day int // 1 <= Day <= 31

type Month int // 1 <= Month <= 12

type Year int

type Gmail string

type Gendre string

type Password string

type Birthday struct {
	Day
	Month
	Year
}
