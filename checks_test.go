package cool_obj

import (
	"testing"
)

type checkerTest[t Checker] struct {
	checker t
	isOk    bool
}

var gmailTests = []checkerTest[Gmail]{
	{"rr@gmail.com", true},
	{"@gmail.com", true},
	{"gmail.com", false},
	{"hohod@gmail.com", true},
	{"ttefrww@hotmail.com", false},
}

func TestGmail_Check(t *testing.T) {
	for _, test := range gmailTests {
		err := test.checker.Check()
		x := err != nil
		if x == test.isOk {
			t.Errorf(`g = %s | g.Check() = %v. g.Check != nil must be %v but it is not`, test.checker, test.checker.Check(), test.isOk)
		}
	}
}

var dayTests = []checkerTest[Day]{
	{2, true},
	{1, true},
	{30, true},
	{15, true},
	{31, true},
	{29, true},
	//
	{-1, false},
	{0, false},
	{32, false},
	{44, false},
	{99, false},
}

func TestDay_Check(t *testing.T) {
	for _, test := range dayTests {
		err := test.checker.Check()
		x := err != nil
		if x == test.isOk {
			t.Errorf(`day = %v | day.Check() = %v. day.Check != nil must be %v but it is not`, test.checker, test.checker.Check(), test.isOk)
		}
	}
}

var monthTests = []checkerTest[Month]{
	{1, true},
	{12, true},
	{6, true},
	{2, true},
	//
	{-1, false},
	{0, false},
	{99, false},
	{100, false},
}

func TestMonth_Check(t *testing.T) {
	for _, test := range monthTests {
		err := test.checker.Check()
		x := err != nil
		if x == test.isOk {
			t.Errorf(`month = %v | month.Check() = %v. month.Check != nil must be %v but it is not`, test.checker, test.checker.Check(), test.isOk)
		}
	}
}

var yearTest = []checkerTest[Year]{
	{2024, true},
	{2000, true},
	{2, true},
	//
	{-1, false},
	{1, false},
	{2025, false},
	{9999, false},
}

func TestYear_Check(t *testing.T) {
	for _, test := range yearTest {
		err := test.checker.Check()
		x := err != nil
		if x == test.isOk {
			t.Errorf(`year = %v | year.Check() = %v. year.Check != nil must be %v but it is not`, test.checker, test.checker.Check(), test.isOk)
		}
	}
}

var gendreCheck = []checkerTest[Gendre]{
	{"male", true},
	{"Male", true},
	{"mAlE", true},
	{"FeMalE", true},
	{"feMaLe", true},
	{"FEMALE", true},
	{"female", true},
	{"prefer not to say", true},
	//
	{"dog", false},
	{"cat", false},
	{"transformed", false},
	{"gay", false},
	{"lgbtq+", false},
}

func TestGendre_Check(t *testing.T) {
	for _, test := range gendreCheck {
		err := test.checker.Check()
		x := err != nil
		if x == test.isOk {
			t.Errorf(`gendre = %v | gendre.Check() = %v. gendre.Check != nil must be %v but it is not`, test.checker, test.checker.Check(), test.isOk)
		}
	}
}

var nameTests = []checkerTest[Name]{
	{"waed", true},
	{"promiselb", true},
	{"potato of earth", true},
	{"not a name", true},
	{"empty name", true},
	//
	{"", false},
}

func TestName_Check(t *testing.T) {
	for _, test := range nameTests {
		err := test.checker.Check()
		x := err != nil
		if x == test.isOk {
			t.Errorf(`name = %v | name.Check() = %v. name.Check != nil must be %v but it is not`, test.checker, test.checker.Check(), test.isOk)
		}
	}
}

var passwordTests = []checkerTest[Password]{
	{"en taro adun", true},
	{"131431", true},
	{"example@gmail.com", true},
	{"battle cruiser operational!!!", true},
	//
	{"", false},
	{"131431(){}", false},
	{":)", false},
}

func TestPassword_Check(t *testing.T) {
	for _, test := range passwordTests {
		err := test.checker.Check()
		x := err != nil
		if x == test.isOk {
			t.Errorf(`password = %v | password.Check() = %v. password.Check != nil must be %v but it is not`, test.checker, test.checker.Check(), test.isOk)
		}
	}
}

var birthDayTests = []checkerTest[Birthday]{
	{Birthday{1, 1, 1990}, true},
	{Birthday{31, 12, 2000}, true},
	{Birthday{20, 7, 1985}, true},
	{Birthday{15, 5, 2023}, true},
	{Birthday{28, 2, 1990}, true},
	{Birthday{29, 2, 1904}, true},
	{Birthday{31, 1, 2022}, true},
	{Birthday{1, 1, -1}, false},    // Invalid year
	{Birthday{1, 0, 1990}, false},  // Invalid month
	{Birthday{0, 1, 1990}, false},  // Invalid day
	{Birthday{30, 2, 1990}, false}, // Invalid day for February
	{Birthday{29, 2, 1990}, false}, // Invalid day for February
	{Birthday{29, 2, 1994}, false},
	{Birthday{29, 2, 1994}, false},
	{Birthday{31, 11, 2022}, false},
	{Birthday{32, 1, 2022}, false},
	{Birthday{31, 2, 1904}, false},
}

func TestBirthday_Check(t *testing.T) {
	for _, test := range birthDayTests {
		err := test.checker.Check()
		x := err != nil
		if x == test.isOk {
			t.Errorf(`birthday = %v | birthday.Check() = %v. birthday.Check != nil must be %v but it is not`, test.checker, test.checker.Check(), test.isOk)
		}
	}
}
