package Request

type CreateMapel struct {
	IdMapel   int    `json:"id_mapel" validate:"required"`   // ID mata pelajaran
	NamaMapel string `json:"nama_mapel" validate:"required"` // Nama mata pelajaran
	Materi    string `json:"materi" validate:"required"`     // Materi pelajaran
	NamaDosen string `json:"nama_dosen" validate:"required"` // Nama dosen pengajar
}

type UpdateMapel struct {
	IdMapel   int    `json:"id_mapel" validate:"required"`   // ID mata pelajaran yang akan diupdate
	NamaMapel string `json:"nama_mapel" validate:"required"` // Nama mata pelajaran
	Materi    string `json:"materi" validate:"required"`     // Materi pelajaran
	NamaDosen string `json:"nama_dosen" validate:"required"` // Nama dosen pengajar
	UpdatedAt string `json:"updated_at"`                     // Waktu update
}
