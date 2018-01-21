package lowrank

type Movie struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	AvgRating float64   `json:"avg_rating"`
	Ratings   []float64 `json:"ratings"`
}
