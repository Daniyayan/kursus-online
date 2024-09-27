package Response

type Register struct {
	IdRegis   int     `json:"id_regis"`   // ID registrasi
	Email     string  `json:"email"`      // Email pengguna
	NoTelpon  int     `json:"no_telpon"`  // Nomor telepon pengguna
	Status    string  `json:"status"`     // Status pengguna (aktif/non-aktif)
	CreatedAt *string `json:"created_at"` // Timestamp untuk saat registrasi dibuat
	UpdatedAt *string `json:"updated_at"` // Timestamp untuk saat registrasi diupdate
}

type RegisterList struct {
	IdRegis  int    `json:"id_regis"`  // ID registrasi
	Email    string `json:"email"`     // Email pengguna
	NoTelpon int    `json:"no_telpon"` // Nomor telepon pengguna
	Status   string `json:"status"`    // Status pengguna (aktif/non-aktif)
}

type RegisterDetail struct {
	IdRegis   int     `json:"id_regis"`   // ID registrasi
	Email     string  `json:"email"`      // Email pengguna
	NoTelpon  int     `json:"no_telpon"`  // Nomor telepon pengguna
	Status    string  `json:"status"`     // Status pengguna (aktif/non-aktif)
	CreatedAt *string `json:"created_at"` // Timestamp untuk saat registrasi dibuat
	UpdatedAt *string `json:"updated_at"` // Timestamp untuk saat registrasi diupdate
}
