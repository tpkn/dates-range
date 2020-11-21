package datesrange

import (
	"fmt"
	"testing"
	"time"
)

type ranges struct {
	title string
	from string
	to string
}

var today = time.Now().Format("2006-01-02")
var yesterday = time.Now().AddDate(0, 0, -1).Format("2006-01-02")

var input = []ranges{
	ranges {"Normal range", "2020-06-29", "2020-07-01"},
	ranges {"From this day", today, ""},
	ranges {"From this day to yesterday", today, yesterday},
	ranges {"Empty dates range", "", ""},
	ranges {"Empty start date", "", today},
	ranges {"Empty end date", today, ""},
}

func TestNew(t *testing.T) {
	for _, l := range input {
		list, err := New(l.from, l.to)
		
		fmt.Println("----------------------------------------")
		fmt.Println(l.title)
		
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println(list)
		}
	}
}
