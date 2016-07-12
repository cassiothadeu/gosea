package services

var users map[int]*User

func init() {
	users = make(map[int]*User)
	user := &User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
		Admin:     true,
	}

	user2 := &User{
		ID:        2,
		FirstName: "Test",
		LastName:  "User",
		Admin:     false,
	}

	users[user.ID] = user
	users[user2.ID] = user2
}

// User defines a user for our application
type User struct {
	ID        int
	FirstName string
	LastName  string
	Admin     bool
}

// UserService provides a CRUD interface for Users
type UserService interface {
	Create(*User) error
	Read(int) (*User, error)
	Update(*User) error
	Delete(int) error
}

// NewUserService creates a new UserService
func NewUserService() UserService {
	return &userService{}
}

type userService struct{}

func (s *userService) Create(u *User) error {
	users[u.ID] = u
	return nil
}

func (s *userService) Read(id int) (*User, error) {
	return users[id], nil
}

func (s *userService) Update(u *User) error {
	users[u.ID] = u
	return nil
}

func (s *userService) Delete(id int) error {
	delete(users, id)
	return nil
}
