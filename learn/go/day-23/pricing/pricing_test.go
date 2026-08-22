package pricing

import (
	"fmt"
	"testing"
)



func BenchmarkDiscountedTotal(b *testing.B){
	for i := 0; i < b.N; i++{
		_, _ = DiscountedTotal(45, 2, 10);
	}
}

func BenchmarkDiscountedTotalWithSetup(b *testing.B){
	price := 45.0;
	qty := 2;
	percent := 10;

	b.ResetTimer();

	for i := 0; i < b.N; i++{
		_, _ = DiscountedTotal(price, qty, percent);
	}
}


func ExampleDiscountedTotal() {
	total, err := DiscountedTotal(45, 2, 10);
	if err != nil{
		fmt.Println("Error:", err);
		return;
	}
	fmt.Println("Total:", total);
}

func ExampleTotal() {
	fmt.Println("Total:", Total(30, 2));
}