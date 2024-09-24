package user

import (
	"kursus-online/Config"
	"kursus-online/Controller/Dto/Request"
	"kursus-online/Controller/Dto/Response"
	"time"
)

// Controller struct, menyimpan koneksi ke database
type Controller struct {
	DB *sql.DB
}

// CreateCourse: Menangani pembuatan kursus baru
func (c *Controller) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Input tidak valid", http.StatusBadRequest)
		return
	}

	// Menggunakan repository untuk menambahkan kursus baru
	err = repository.NewRepository(c.DB).CreateCourse(context.Background(), course)
	if err != nil {
		http.Error(w, "Gagal membuat kursus", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Kursus berhasil dibuat")
}

// GetCourses: Mendapatkan semua kursus
func (c *Controller) GetCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := repository.NewRepository(c.DB).GetCourses(context.Background())
	if err != nil {
		http.Error(w, "Gagal mendapatkan kursus", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// RegisterUser: Mendaftarkan pengguna ke kursus
func (c *Controller) RegisterUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseID, err := strconv.Atoi(vars["courseID"])
	if err != nil {
		http.Error(w, "ID kursus tidak valid", http.StatusBadRequest)
		return
	}

	var registration models.Registration
	err = json.NewDecoder(r.Body).Decode(&registration)
	if err != nil {
		http.Error(w, "Input tidak valid", http.StatusBadRequest)
		return
	}

	err = repository.NewRepository(c.DB).RegisterUser(context.Background(), courseID, registration)
	if err != nil {
		http.Error(w, "Gagal mendaftarkan pengguna ke kursus", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Pengguna berhasil didaftarkan")
}

// GetCourseDetails: Mendapatkan detail kursus
func (c *Controller) GetCourseDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseID, err := strconv.Atoi(vars["courseID"])
	if err != nil {
		http.Error(w, "ID kursus tidak valid", http.StatusBadRequest)
		return
	}

	course, err := repository.NewRepository(c.DB).GetCourseDetails(context.Background(), courseID)
	if err != nil {
		http.Error(w, "Gagal mendapatkan detail kursus", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}
