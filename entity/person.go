package entity

// Person ...
type Person struct {
	Name  string `fake:"{name}"`
	Phone string `fake:"{phone}"`
}
