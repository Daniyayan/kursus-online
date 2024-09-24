package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/labstack/echo"
)

// Controller struct
type userInterface interface {
	CreateCourse(etx echo.Context) (err error)
	GettCourse(ctx echo.Context) (err error)
	RegisterUser(ctx echo.Context) (err error)
	GetCourseDetails(ctx echo.Context) (err error)
}

// CreateCourse: Handle creating a new course
func (c *Controller) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO courses (name, description, duration, price) VALUES (?, ?, ?, ?)`
	_, err = c.DB.Exec(query, course.Name, course.Description, course.Duration, course.Price)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error creating course", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Course created successfully")
}

// GetCourses: Get all courses
func (c *Controller) GetCourses(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, name, description, duration, price FROM courses`
	rows, err := c.DB.Query(query)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching courses", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var course models.Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.Duration, &course.Price)
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, course)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// RegisterUser: Register a user to a course
func (c *Controller) RegisterUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseID, err := strconv.Atoi(vars["courseID"])
	if err != nil {
		http.Error(w, "Invalid course ID", http.StatusBadRequest)
		return
	}

	var registration models.Registration
	err = json.NewDecoder(r.Body).Decode(&registration)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO registrations (user_id, course_id, date_registered, status) VALUES (?, ?, ?, ?)`
	_, err = c.DB.Exec(query, registration.UserID, courseID, registration.DateRegistered, registration.Status)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User registered successfully")
}

// GetCourseDetails: Get details of a specific course
func (c *Controller) GetCourseDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseID, err := strconv.Atoi(vars["courseID"])
	if err != nil {
		http.Error(w, "Invalid course ID", http.StatusBadRequest)
		return
	}

	var course models.Course
	query := `SELECT id, name, description, duration, price FROM courses WHERE id = ?`
	err = c.DB.QueryRow(query, courseID).Scan(&course.ID, &course.Name, &course.Description, &course.Duration, &course.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Course not found", http.StatusNotFound)
		} else {
			log.Fatal(err)
			http.Error(w, "Error fetching course details", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}
