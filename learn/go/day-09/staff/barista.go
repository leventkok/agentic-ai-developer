package staff

import "fmt"

type Barista struct {
	Staff
	Shift string
}

func (b Barista) String() string {
	return fmt.Sprintf("Barista: %s - shift %s", b.Name, b.Shift)

}
