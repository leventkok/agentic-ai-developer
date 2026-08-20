

package kitchen

import(
	"fmt"
	"time"

)

func BrewUntilDone(orders <-chan string, done <-chan struct{}) {
    for {
        select {
        case <-done:
            fmt.Println("Kitchen closing — stop brewing")
            return
        case drink := <-orders:
            fmt.Println("Brewing:", drink)
            time.Sleep(200 * time.Millisecond)
        }
    }
}