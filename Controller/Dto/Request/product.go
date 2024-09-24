package Request

type CreateCourse struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Duration    string  `json:"duration"`
	Price       float64 `json:"price"`
}

type UpdateProduct struct {
	Id          string  `json:"id"`
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float64 `json:"rating" validate:"gte=1,lte=5"`
	Image       string  `json:"image"`
	YouTubeLink string  `json:"youtube_link"` // Tambahkan properti untuk link YouTube
}
