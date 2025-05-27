package user

type User struct {
	ID        string
	UserName  string
	FirstName string
	LastName  string
	Email     string
	Mobile    string
	Password  string // Hashed password
	Role      UserRole
	IsActive  bool
}
