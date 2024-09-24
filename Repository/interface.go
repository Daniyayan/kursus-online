package Repository

import (
	"kursus-online/Repository/Course"
)

type Repository struct {
	Course Course.RepositoryCourse
	//Product Product.RepositoryProduct
}

var ApplicationRepository = Repository{
	Course: Course.NewRepository(),
	//Product: Product.NewRepository(),
}
