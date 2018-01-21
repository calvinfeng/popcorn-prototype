package matfac

import (
	"gonum.org/v1/gonum/mat"
)

type Recommender struct {
	UserIdToIndex  map[int]int
	UserIndexToId  map[int]int
	MovieIdToIndex map[int]int
	MovieIndexToId map[int]int
	RatingMap      map[int]map[int]float64
	MovieMap       map[int]*Movie
}

func NewRecommender(ratingMap map[int]map[int]float64, movieMap map[int]*Movie) *Recommender {
	var i, j int
	userIdToIndex := make(map[int]int)
	userIndexToId := make(map[int]int)
	movieIdToIndex := make(map[int]int)
	movieIndexToId := make(map[int]int)

	i = 0
	for userId := range ratingMap {
		userIdToIndex[userId] = i
		userIndexToId[i] = userId
		i += 1

		for movieId := range ratingMap[userId] {
			movieMap[movieId].Ratings = append(movieMap[movieId].Ratings, ratingMap[userId][movieId])
		}
	}

	// Compute average rating for each movie
	j = 0
	for movieId := range movieMap {
		movieIdToIndex[movieId] = j
		movieIndexToId[j] = movieId
		j += 1

		movieMap[movieId].AvgRating = Average(movieMap[movieId].Ratings)
	}

	return &Recommender{
		UserIdToIndex:  userIdToIndex,
		UserIndexToId:  userIndexToId,
		MovieIdToIndex: movieIdToIndex,
		MovieIndexToId: movieIndexToId,
		RatingMap:      ratingMap,
		MovieMap:       movieMap,
	}
}

// We should have I users and J movies, with K latent features.
func (r *Recommender) GetRatingMatrix() *mat.Dense {
	I, J := len(r.RatingMap), len(r.MovieMap)
	R := mat.NewDense(I, J, nil)
	for i := 0; i < I; i += 1 {
		for j := 0; j < J; j += 1 {
			userId := r.UserIndexToId[i]
			movieId := r.MovieIndexToId[j]
			if _, ok := r.RatingMap[userId][movieId]; ok {
				R.Set(i, j, r.RatingMap[userId][movieId])
			} else {
				R.Set(i, j, r.MovieMap[movieId].AvgRating)
			}
		}
	}

	return R
}
