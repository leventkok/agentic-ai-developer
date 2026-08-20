


package pipeline

import "fmt"


func Generate(nums ...int) <-chan int{
	out := make(chan int)
	go func(){
		for _, n := range nums{
			out <- n
		}
		close(out)
	}()
	return out
}

func Square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}
func Print(in <-chan int) {
    for n := range in {
        fmt.Println("Result:", n)
    }
}