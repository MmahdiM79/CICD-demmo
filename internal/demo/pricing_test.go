package demo

import "testing"

func TestNoDiscount(t *testing.T) {
	if got := CalculateTotal([]Item{{Price: 10, Qty: 2}, {Price: 5, Qty: 1}}, ""); got != 25 {
		t.Fatalf("got %v, want 25", got)
	}
}

func TestStudentDiscount(t *testing.T) {
	if got := CalculateTotal([]Item{{Price: 10, Qty: 2}, {Price: 5, Qty: 1}}, "STUDENT20"); got != 20 {
		t.Fatalf("got %v, want 20", got)
	}
}

func TestDiscountLowersPrice(t *testing.T) {
	items := []Item{{Price: 10, Qty: 2}, {Price: 5, Qty: 1}}
	if CalculateTotal(items, "WELCOME10") >= CalculateTotal(items, "") {
		t.Fatal("discount must lower the price")
	}
}

func TestUnknownDiscount(t *testing.T) {
	if got := CalculateTotal([]Item{{Price: 10, Qty: 2}, {Price: 5, Qty: 1}}, "FREE-STUFF"); got != 25 {
		t.Fatalf("got %v, want 25", got)
	}
}
