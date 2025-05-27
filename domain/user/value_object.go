package user

import "golang.org/x/crypto/bcrypt"

type UserRole string

const (
	AdminRole    UserRole = "admin"
	ManagerRole  UserRole = "manager"
	EmployeeRole UserRole = "employee"
	CustomerRole UserRole = "customer"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 12)
}

func ComparePassword(hashedPwd, plainPwd string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
}