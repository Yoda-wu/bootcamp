package main

import (
	"testing"
)

func TestMain(m *testing.M) {

}

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expect := "Tom"
	if expect != output {
		t.Error("output is not equal to expect")
	}
}
