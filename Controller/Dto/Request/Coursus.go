package Request

type CreateCoursus struct {
	IdCoursus   int     `json:"id" validate:"required"`
	Nama        string  `json:"nama" validate:"required"`        // Nama kursus
	Description string  `json:"description" validate:"required"` // Deskripsi kursus
	Duration    string  `json:"duration" validate:"required"`    // Durasi kursus
	Period      string  `json:"periode" validate:"required"`     // Periode kursus
	Price       float64 `json:"price" validate:"required"`       // Harga kursus
}

type UpdateCoursus struct {
	IdCoursus   int     `json:"idCoursus" validate:"required"`   // ID kursus yang akan diupdate
	Nama        string  `json:"nama" validate:"required"`        // Nama kursus
	Description string  `json:"description" validate:"required"` // Deskripsi kursus
	Duration    string  `json:"duration" validate:"required"`    // Durasi kursus
	Period      string  `json:"periode" validate:"required"`     // Periode kursus
	Price       float64 `json:"price" validate:"required"`       // Harga kursus
	UpdatedAt   string  `json:"updated_at"`                      // Waktu update
}
