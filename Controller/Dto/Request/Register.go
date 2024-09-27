package Request

type CreateRegister struct {
	IdRegis  int    `json:"id_regis" validate:"required"`    // ID registrasi
	Email    string `json:"email" validate:"required,email"` // Email pengguna
	NoTelpon int    `json:"no_telpon" validate:"required"`   // Nomor telepon pengguna
	Status   string `json:"status" validate:"required"`      // Status pengguna (misalnya: aktif/non-aktif)
}

type UpdateRegister struct {
	IdRegis   int    `json:"id_regis" validate:"required"`    // ID registrasi yang akan diupdate
	Email     string `json:"email" validate:"required,email"` // Email pengguna
	NoTelpon  int    `json:"no_telpon" validate:"required"`   // Nomor telepon pengguna
	Status    string `json:"status" validate:"required"`      // Status pengguna
	UpdatedAt string `json:"updated_at"`                      // Waktu update
}
