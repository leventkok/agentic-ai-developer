package staff

import "fmt"

type Manager struct {
	Staff
	Level int
}

func (m Manager) String() string {
	return fmt.Sprintf("Manager Level %d: %s", m.Level, m.Name)
}
