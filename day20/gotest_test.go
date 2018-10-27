package gotest

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(1, 0); i != 3 || e != nil {
		t.Error("除法函數沒通過測試")
	} else {
		t.Log("測試通過")
	}
}

func Test_Division_table(t *testing.T) {
	tables := []struct {
		x float64
		y float64
	}{
		{3, 1},
		{6, 2},
		{9, 3},
		{8, 2},
	}

	for _, table := range tables {
		if i, e := Division(table.x, table.y); i != 3 || e != nil {
			t.Error("除法函數沒通過測試")
		} else {
			t.Log("測試通過")
		}
	}
}
