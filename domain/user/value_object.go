package user

type UserRole string

const (
	AdminRole    UserRole = "admin"
	ManagerRole  UserRole = "manager"
	EmployeeRole UserRole = "employee"
	CustomerRole UserRole = "customer"
)
