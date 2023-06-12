package transport

// ДТО установки отдела пользователю.
type SetDepartmentQuery struct {
	// Идентификатор пользователя.
	EmployeeId int64 `json:"employeeId"`

	// Идентификатор отдела.
	DepartmentId int64 `json:"departmentId"`
}
