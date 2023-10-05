package bo

import (
	casbin "github.com/casbin/casbin/v2"
	"github.com/dirkarnez/stemexapi/utils"
	"gorm.io/gorm"
)

// // Product is the interface that all products implement
type IBO interface {
	Name() string
}

// // ConcreteProductB is another specific implementation of the Product interface
// type ConcreteProductB struct{}

// func (p *ConcreteProductB) Use() string {
// 	return "Using ConcreteProductB"
// }

// // Creator is the interface for the factory
// type Creator interface {
// 	CreateProduct() Product
// }

// // ConcreteCreatorA is a specific implementation of the Creator interface for creating ConcreteProductA
// type ConcreteCreatorA struct{}

type Factory struct {
	db      *gorm.DB
	e       *casbin.SyncedEnforcer
	usersBO *UsersBO
}

func (f *Factory) GetUsersBO() *UsersBO {
	return f.usersBO
}

// // ConcreteCreatorB is a specific implementation of the Creator interface for creating ConcreteProductB
// type ConcreteCreatorB struct{}

// func (c *ConcreteCreatorB) CreateProduct() Product {
// 	return &ConcreteProductB{}
// }

func NewFactory(db *gorm.DB, e *casbin.SyncedEnforcer) *Factory {
	factory := &Factory{db, e, &UsersBO{}}
	_, enforcerErr := e.AddPolicy("admin", factory.usersBO.Name(), "read")
	utils.CheckError(enforcerErr)

	_, enforcerErr = e.AddPolicy("admin", factory.usersBO.Name(), "write")
	utils.CheckError(enforcerErr)

	return factory
}

// func s() {
// 	// Create a ConcreteCreatorA
// 	creatorA :=
// 	productA := creatorA.CreateUsersBO()
// 	fmt.Println(productA.s())

// 	// Create a ConcreteCreatorB
// 	creatorB := &ConcreteCreatorB{}
// 	productB := creatorB.CreateProduct()
// 	fmt.Println(productB.Use())
// }
