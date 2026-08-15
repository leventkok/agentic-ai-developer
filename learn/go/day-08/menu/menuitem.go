package menu

type MenuItem interface {
	String() string
	PriceTL() int
}
