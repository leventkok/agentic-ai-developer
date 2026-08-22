package pricing

import "testing"





func TestDiscountedTotal(t *testing.T){
	test := []struct{
		name string;
		price float64;
		qty int;
		percent int;
		want float64;
		wantErr bool;

	}{
		{name: "latte x2 ten percent off", price: 45, qty: 2, percent: 10, want: 81},
		{name: "single item no discount", price: 30, qty: 1, percent: 0, want: 30},
		{name: "full discount", price: 50, qty: 1, percent: 100, want: 0},
		{name: "zero qty", price: 45, qty: 0, percent: 10, wantErr: true},
		{name: "negative percent", price: 45, qty: 1, percent: -5, wantErr: true},
		{name: "percent over 100", price: 45, qty: 1, percent: 101, wantErr: true},
	}


	for _, tt := range test{
		t.Run(tt.name, func(t *testing.T){
			got, err := DiscountedTotal(tt.price, tt.qty, tt.percent);
			if tt.wantErr{
				if err == nil{
					t.Fatalf("expected error, got nil");
				}
				return;
			}
			if err != nil{
				t.Fatalf("unexpected error: %v", err);
			}
			if got != tt.want{
				t.Errorf("got %v, want %v", got, tt.want);
			}
		});
	}

}


func assertDiscount(t *testing.T, price float64, qty, percent int, want float64) {
	t.Helper()
	got, err := DiscountedTotal(price, qty, percent)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}