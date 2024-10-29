package main

import (
	"reflect"
	"testing"
)

func TestReorganize(t *testing.T) {
	data := []string{"i love music", "i love Music", "we love MuSiC", "they Love Music", "", "i do not love musiC"}
	got := reorganize(2, 1, true, data)
	want := []string{"usic", "usic", "usic", "usic", "", "ot love music"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("reorganize(2, 1, true, data) = %s; want = %s", got, want)
	}
}

func TestResult(t *testing.T) {
	data := []string{"i love music", "i love Music", "we love MuSiC", "they Love Music", "", "i do not love musiC"}
	userVision := []string{"usic", "usic", "usic", "usic", "", "ot love music"}
	got := result(true, false, false, userVision, data)
	want := []string{"1 i love music", "1 ", "1 i do not love musiC"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("result(true, false, false, userVision, data) = %s; want = %s", got, want)
	}
}
