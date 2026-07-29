package modules

var Employees []Employee

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
	Email    string  `json:"email"`
	IsActive bool    `json:"is_active"`
}
