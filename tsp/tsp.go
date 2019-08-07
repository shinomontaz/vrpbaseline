package tsp

import (
	"fmt"
	"math/rand"

	"github.com/shinomontaz/vrpbaseline/types"
)

func solve(points []types.Order, dm [][]types.Distance) []int {

	solution := rand.Perm(len(points))
	fitness := routeDistance(solution, dm)

	fmt.Println("start", solution, fitness)

	for i := 0; i < 10000; i++ {
		chunk := rand.Intn(len(points) - 2)
		for j := 0; j < chunk; j++ {
			for k := j + 1; k < chunk; k++ {
				new_solution := step(solution, j, k)
				new_distance := routeDistance(new_solution, dm)
				if new_distance < fitness {
					fitness = new_distance
					solution = new_solution
					fmt.Println(fitness)

					break
				}
			}
		}
	}
	fmt.Println("end", solution, fitness)

	return solution
}

func step(in []int, i, j int) []int {
	return in
}

func routeDistance(route []int, dm [][]types.Distance) float64 {
	d := 0.0
	for i := 1; i < len(route); i++ {
		d += dm[route[i-1]][route[i]].Road
	}

	return d
}
