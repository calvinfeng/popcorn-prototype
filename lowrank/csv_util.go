package lowrank

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func LoadRatingsByUserID(filepath string) (map[int]map[int]float64, error) {
	if csvFile, fileErr := os.Open(filepath); fileErr == nil {
		reader := csv.NewReader(bufio.NewReader(csvFile))
		ratingsByUserId := make(map[int]map[int]float64)
		for {
			if row, readerErr := reader.Read(); readerErr == nil {
				userId, idErr := strconv.ParseInt(row[0], 10, 32)
				movieId, movieIDErr := strconv.ParseInt(row[1], 10, 32)
				rating, ratingErr := strconv.ParseFloat(row[2], 64)

				if idErr == nil && movieIDErr == nil && ratingErr == nil {
					if _, ok := ratingsByUserId[int(userId)]; !ok {
						ratingsByUserId[int(userId)] = make(map[int]float64)
					}

					ratingsByUserId[int(userId)][int(movieId)] = rating
				}
			} else {
				if readerErr == io.EOF {
					break
				} else {
					fmt.Printf("Unexpected reader error: %v\n", readerErr)
				}
			}
		}

		return ratingsByUserId, nil
	} else {
		return nil, fileErr
	}
}

func LoadMovies(filepath string) (map[int]*Movie, error) {
	if csvFile, err := os.Open(filepath); err == nil {
		reader := csv.NewReader(bufio.NewReader(csvFile))

		movieById := make(map[int]*Movie)
		for {
			if row, readerErr := reader.Read(); readerErr != nil {
				if readerErr == io.EOF {
					break
				} else {
					fmt.Printf("Unexpected reader error: %v", readerErr)
				}
			} else {
				// bitSize defines range of values. If the value corresponding to s cannot be represented by a signed
				// integer of the given size, err.Err = ErrRange.
				id, parseErr := strconv.ParseInt(row[0], 10, 64)
				if parseErr == nil {
					movieById[int(id)] = &Movie{
						ID:    int(id),
						Title: row[1],
					}
				}
			}
		}

		return movieById, nil
	} else {
		return nil, err
	}
}
