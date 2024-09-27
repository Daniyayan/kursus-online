package Register

import (
	"context"
	"kursus-online/Controller/Dto/Request"
	//"kursus-online/Controller/Dto/Response"
)

// RepositoryRegister interface defines the methods for the 'register' repository
type RepositoryRegister interface {
	CreateRegister(ctx context.Context, param Request.CreateRegister) (err error)
	//CreateRegister(ctx context.Context, param Request.CreateRegister) (err error)
	//CheckExistsRegisterEmail(ctx context.Context, email string) (exists bool, err error)
	//CheckExistsRegisterId(ctx context.Context, id string) (exists bool, err error)
	//UpdateRegister(ctx context.Context, param Request.UpdateRegister) (err error)
	//DeleteRegister(ctx context.Context, id string) (err error)
	//ListRegister(ctx context.Context, sortBy string) (res []Response.RegisterList, err error)
	//DetailRegister(ctx context.Context, id string) (res Response.RegisterDetail, err error)
}

// Register struct implements the RepositoryRegister interface
type register struct{}

// NewRepository creates a new repository instance for the 'register' table
func NewRepository() RepositoryRegister {
	return &register{}
}
