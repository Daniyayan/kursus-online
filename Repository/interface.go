package Repository

import (
	"kursus-online/Repository/Coursus"
	"kursus-online/Repository/Register"
	Mapel "kursus-online/Repository/mapel"
)

type Repository struct {
	Coursus  Coursus.RepositoryCoursus
	Mapel    Mapel.RepositoryMapel
	Register Register.RepositoryRegister
	//Product Product.RepositoryProduct
}

var ApplicationRepository = Repository{
	Coursus:  Coursus.NewRepository(),
	Mapel:    Mapel.NewRepository(),
	Register: Register.NewRepository(),
	//Product: Product.NewRepository(),
}
