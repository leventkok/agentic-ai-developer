package pricing

import "testing"



func TestTotal(t *testing.T){
	got := Total(45, 2);
	want := 90.0;

	if got != want{
		t.Errorf("Total(45, 2) = %v, want %v", got, want);
	}
}


func TestApplyDiscount(t *testing.T){
	got := ApplyDiscount(100, 10);
	want := 90.0;


	if got != want{
		t.Fatalf("ApplyDiscount(100, 10) = %v, want %v", got, want);
	}
}


func TestTotalSingleItem(t *testing.T){
	got := Total(30, 1);
	want := 30.0;

	if got != want{
		t.Fatalf("got %v, want %v", got, want);
	}
}