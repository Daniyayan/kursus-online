package Response

type Mapel struct {
	IdMapel   int     `json:"id_mapel"`   // ID mata pelajaran
	NamaMapel string  `json:"nama_mapel"` // Nama mata pelajaran
	Materi    string  `json:"materi"`     // Materi pelajaran
	NamaDosen string  `json:"nama_dosen"` // Nama dosen pengajar
	CreatedAt *string `json:"created_at"` // Timestamp untuk saat mata pelajaran dibuat
	UpdatedAt *string `json:"updated_at"` // Timestamp untuk saat mata pelajaran diupdate
}

type MapelList struct {
	IdMapel   int    `json:"id_mapel"`   // ID mata pelajaran
	NamaMapel string `json:"nama_mapel"` // Nama mata pelajaran
	Materi    string `json:"materi"`     // Materi pelajaran
	NamaDosen string `json:"nama_dosen"` // Nama dosen pengajar
}

type MapelDetail struct {
	IdMapel   int     `json:"id_mapel"`   // ID mata pelajaran
	NamaMapel string  `json:"nama_mapel"` // Nama mata pelajaran
	Materi    string  `json:"materi"`     // Materi pelajaran
	NamaDosen string  `json:"nama_dosen"` // Nama dosen pengajar
	CreatedAt *string `json:"created_at"` // Timestamp untuk saat mata pelajaran dibuat
	UpdatedAt *string `json:"updated_at"` // Timestamp untuk saat mata pelajaran diupdate
}
