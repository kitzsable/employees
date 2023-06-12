package models

// Сотрудник.
type Employee struct {
	// Идентификатор сотрудника.
	Id int64 `json:"id"`

	// Имя сотрудника.
	Name string `json:"name"`

	// Пол сотрудника.
	Sex string `json:"sex"`

	// Возраст сотрудника.
	Age int `json:"age"`

	// Зарплата сотрудника.
	Salary int `json:"salary"`

	// Отдел, в котором работает сотрудник.
	Department Department `json:"department"`
}
