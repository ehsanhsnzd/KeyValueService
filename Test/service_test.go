package Test

import (
	"biges/Repository"
	"testing"
)

func TestSet(t *testing.T) {
	Repo := Repository.NewData()

	result := Repo.SetValue("foo", "bar")

	if result != true {
		t.Errorf("Set failed. expected %v, got %v", true, result)
	} else {
		t.Logf("Set success. expected %v, got %v", true, result)
	}

}

func TestGet(t *testing.T) {
	expect:="bar"
	Repo := Repository.NewData()

	Repo.SetValue("foo", expect)
	result,_ := Repo.GetValue("foo")

	if result != expect {
		t.Errorf("Get failed. expected %v, got %v", expect, result)
	} else {
		t.Logf("Get success. expected %v, got %v", expect, result)
	}

}
