package Response

type Coursus struct {
	IdCoursus   int     `json:"id_coursus"`
	Nama        string  `json:"nama"`
	Description string  `json:"description"`
	Duration    string  `json:"duration"` // Sesuaikan tipe data jika perlu
	Period      string  `json:"periode"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"` // Tipe data dapat disesuaikan sesuai kebutuhan
	UpdatedAt   string  `json:"updated_at"` // Tipe data dapat disesuaikan sesuai kebutuhan
}

type CoursusList struct {
	IdCoursus   int     `json:"id_coursus"`  // Menggunakan int untuk id_coursus
	Nama        string  `json:"nama"`        // Nama kursus
	Description string  `json:"description"` // Deskripsi kursus
	Duration    string  `json:"duration"`    // Durasi kursus
	Period      string  `json:"periode"`     // Periode kursus
	Price       float64 `json:"price"`       // Harga kursus
}

type CoursusDetail struct {
	IdCoursus   int     `json:"id_coursus"`  // Menggunakan int untuk id_coursus
	Nama        string  `json:"nama"`        // Nama kursus
	Description string  `json:"description"` // Deskripsi kursus
	Duration    string  `json:"duration"`    // Durasi kursus
	Period      string  `json:"periode"`     // Periode kursus
	Price       float64 `json:"price"`       // Harga kursus
	CreatedAt   *string `json:"createdAt"`   // Timestamp untuk saat kursus dibuat
	UpdatedAt   *string `json:"updatedAt"`   // Timestamp untuk saat kursus diupdate
}
