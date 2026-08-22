package cafe

import "testing"

func TestLineTotal(t *testing.T) {
	tests := []struct {
		name    string
		price   float64
		qty     int
		want    float64
		wantErr bool
	}{
		{name: "latte x2", price: 45, qty: 2, want: 90},
		{name: "zero qty", price: 45, qty: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LineTotal(tt.price, tt.qty)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplyDiscount(t *testing.T) {
	tests := []struct {
		name    string
		total   float64
		percent int
		want    float64
		wantErr bool
	}{
		{name: "ten percent off", total: 100, percent: 10, want: 90},
		{name: "invalid percent", total: 100, percent: 101, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ApplyDiscount(tt.total, tt.percent)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLineTotal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = LineTotal(45, 2)
	}
}

// Task 4: break a want value on purpose, run go test ./cafe -v, read the failure, then fix.
