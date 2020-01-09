package words

import "testing"

func TestNumberWords(t *testing.T) {
	if got := numberWords(3); got != "three" {
		t.Errorf("got %s, want %s\n", got, "three")
	}
	if got := numberWords(4); got != "four" {
		t.Errorf("got %s, want %s\n", got, "four")
	}
	if got := numberWords(2); got != "two" {
		t.Errorf("got %s, want %s\n", got, "two")
	}
	if got := numberWords(10); got != "ten" {
		t.Errorf("got %s, want %s\n", got, "ten")
	}
	if got := numberWords(22); got != "twenty two" {
		t.Errorf("got %s, want %s\n", got, "twenty two")
	}
	if got := numberWords(52); got != "fifty two" {
		t.Errorf("got %s, want %s\n", got, "fifty two")
	}
	if got := numberWords(952); got != "nine hundred fifty two" {
		t.Errorf("got %s, want %s\n", got, "nine hundred fifty two")
	}
	if got := numberWords(0); got != "zero" {
		t.Errorf("got %s, want %s\n", got, "zero")
	}
	if got := numberWords(100); got != "one hundred" {
		t.Errorf("got %s, want %s\n", got, "one hundred")
	}
	if got := numberWords(200); got != "two hundred" {
		t.Errorf("got %s, want %s\n", got, "two hundred")
	}
	if got := numberWords(210); got != "two hundred ten" {
		t.Errorf("got %s, want %s\n", got, "two hundred ten")
	}
	if got := numberWords(211); got != "two hundred eleven" {
		t.Errorf("got %s, want %s\n", got, "two hundred eleven")
	}
	if got := numberWords(901); got != "nine hundred one" {
		t.Errorf("got %s, want %s\n", got, "nine hundred one")
	}
	if got := numberWords(999); got != "nine hundred ninety nine" {
		t.Errorf("got %s, want %s\n", got, "nine hundred ninety nine")
	}
}
