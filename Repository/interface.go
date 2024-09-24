package Repository

import "kursus-online/Repository/Product/"

type Repository struct {
	Product Product.RepositoryProduct
}

var ApplicationRepository = Repository{
	Product: Product.NewRepository(),
}
