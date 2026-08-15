package staff

import "fmt"

type Staff struct {
	Name  string
	Email string
}

func (s Staff) String() string {
	return fmt.Sprintf("%s: (%s)", s.Name, s.Email)
}
