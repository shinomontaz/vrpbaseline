package tsp

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/shinomontaz/vrpbaseline/types"
)

func TestTsp(t *testing.T) {
	points := createPoints(15)

	dm := make([][]types.Distance, 0, 15)

	// заполнить матрицу

	result := solve(points, dm)

	fmt.Println(result)
}

func createPoints(n int) []types.Order {
	res := make([]types.Order, 0)
	for i := 0; i < n; i++ {
		res = append(res, types.Order{
			Lat:  rand.Float64() * 100,
			Long: rand.Float64() * 100,
		})
	}

	return res
}
