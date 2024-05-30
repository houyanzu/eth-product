package mytime

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
	"time"
)

func Test(t *testing.T) {
	ti := NewFromNow()
	fmt.Println(ti.Time())
	var aaa struct {
		Tt DateTime
	}
	var bbb struct {
		Tt DateTime
	}
	aaa.Tt = ti
	js, _ := json.Marshal(&aaa)
	fmt.Println(string(js))
	json.Unmarshal(js, &bbb)
	fmt.Println(bbb.Tt.Time())
}

func TestDateTime_DiffDays(t *testing.T) {
	date1 := NewFromNow()
	date2 := NewFromNow().Add(-87 * time.Hour)

	fmt.Println(date2.DiffDays(date1))
}

func Test2(t *testing.T) {
	aaa := make(map[int8]decimal.Decimal)
	aaa[1] = aaa[1].Add(decimal.NewFromInt(123))
	fmt.Println(aaa)
}
