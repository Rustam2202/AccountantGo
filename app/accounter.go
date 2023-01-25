package acc

import (
	"go/doc/comment"
	"time"
)

type Record struct {
	Id      uint
	Date    time.Time
	Income  float32
	Spend   float32
	Comment string
}

type Records struct {
	Records []Record
}

func (r *Record) AddIncome() {

}
