package bo

// ConcreteProductA is a specific implementation of the Product interface
type UsersBO struct{}

func (p *UsersBO) Name() string {
	return "USERS_BO"
}

func (p *UsersBO) CreateUser() {
}
