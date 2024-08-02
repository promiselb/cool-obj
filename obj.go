package cool_obj

type Name string

type Day uint // 1 <= Day <= 31

type Month uint // 1 <= Month <= 12

type Year uint

type Gmail string

type Gendre string

type Password string

type Birthday struct {
	Day
	Month
	Year
}
