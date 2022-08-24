package mytime

import (
	"encoding/json"
	"fmt"
	"testing"
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
