package goenv

import (
	"os"
	"testing"
	"time"
)

func TestExists(t *testing.T) {
	k := "TESTVAR"
	e := os.Setenv(k, "poop")

	if e != nil {
		t.Error(e)
	}

	if !Exists(k) {
		t.Error("should exist")
	}

	e = os.Unsetenv(k)

	if e != nil {
		t.Error(e)
	}

	if Exists(k) {
		t.Error("should not exist")
	}
}

func TestBool(t *testing.T) {
	k := "BOOLVAR"
	e := os.Setenv(k, "1")

	if e != nil {
		t.Error(e)
	}

	if !Bool(k, false) {
		t.Error("should be true")
	}

	e = os.Setenv(k, "0")

	if e != nil {
		t.Error(e)
	}

	if Bool(k, true) {
		t.Error("should be false")
	}

	e = os.Unsetenv(k)

	if e != nil {
		t.Error(e)
	}

	if Bool(k, false) {
		t.Error("should be false")
	}
}

func TestInt(t *testing.T) {
	k := "INTVAR"
	e := os.Setenv(k, "5")

	if e != nil {
		t.Error(e)
	}

	i, e := Int(k, 0)

	if e != nil {
		t.Error(e)
	}

	if i != 5 {
		t.Errorf("expected 5, got %d", i)
	}

	e = os.Setenv(k, "not a number")

	if e != nil {
		t.Error(e)
	}

	_, e = Int(k, 0)

	if e == nil {
		t.Error("expected error")
	}

	e = os.Unsetenv(k)

	if e != nil {
		t.Error(e)
	}

	i, e = Int(k, 69)

	if e != nil {
		t.Error(e)
	}

	if i != 69 {
		t.Errorf("expected 69, got %d", i)
	}
}

func TestSecond(t *testing.T) {
	k := "SECVAR"
	x := 5 * time.Second
	e := os.Setenv(k, "5")

	if e != nil {
		t.Error(e)
	}

	s, e := Second(k, 0)

	if e != nil {
		t.Error(e)
	}

	if s != x {
		t.Errorf("expected %+v, got %+v", x, s)
	}
}
