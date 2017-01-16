package concerto

import (
	"testing"
)

func TestMin(t *testing.T) {
	expected := 10
	min := min(10, 30)

	if min != expected {
		t.Errorf("Expected %d, got %d", expected, min)
	}
}

func TestIsSong(t *testing.T) {
	song1 := "Basket Case"
	song2 := "BasketCase"
	expected := true

	isSong := isSong(song1, song2)

	if isSong != expected {
		t.Errorf("Expected %t, got %t", expected, isSong)
	}
}
