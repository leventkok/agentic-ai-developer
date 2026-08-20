package worker

type Job struct {
	Drink string
}

type Result struct {
	Drink string
	Price int
}

var prices = map[string]int{
	"latte":    45,
	"espresso": 35,
	"filter":   30,
}

func Brew(jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		price := prices[job.Drink]
		results <- Result{Drink: job.Drink, Price: price}
	}
}
