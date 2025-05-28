package user

type User struct {
	ID        string
	Username  string
	Password  string
	Role      string
	IsActive  bool
	Email     *string
	FirstName *string
	LastName  *string
	Mobile    *string
}
