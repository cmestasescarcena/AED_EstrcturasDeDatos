package main

import (
	"fmt"
	"math/rand"
)

// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func Permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func RandomPermutations(_cities []int, _nPopulation int) [][]int {
	var selected = rand.Perm(119)[:_nPopulation]
	//print("Seleccionados ")
	//fmt.Println(selected)
	var possibleGenes = Permutations(_cities)
	//fmt.Println(possibleGenes[4])
	//fmt.Println(Permutations(_cities))
	//fmt.Println(len(Permutations(_cities)))
	//print("Seleccionados A ")
	//fmt.Println(possibleGenes[0])
	//print("Seleccionados B ")
	//fmt.Println(possibleGenes[119])
	//print("Seleccionados C ")
	//fmt.Println(len(selected))
	//fmt.Println(possibleGenes)
	var start = 0
	//var flag = true
	var genes [][]int
	for j := 0; j < len(selected); j++ {
		//fmt.Println(selected[j])
		//fmt.Println(possibleGenes[selected[j]])
		start = possibleGenes[selected[j]][0]
		possibleGenes[selected[j]] = append(possibleGenes[selected[j]], start)
		genes = append(genes, possibleGenes[selected[j]])

	}
	return genes
}

type Population struct {
	genes        [][]int
	parents      [][]int
	score        int
	best         []int
	adjacencyMat [][]int
}

func NewPopulation(_cities []int, _adjacencyMat [][]int, _nPopulation int) Population {
	var parents [][]int
	var best []int
	return Population{RandomPermutations(_cities, _nPopulation), parents, 0, best, _adjacencyMat}
}

func distance(_cromosome []int, _adjacencyMat [][]int) int {
	var sum = 0
	for i := 0; i < (len(_cromosome) - 1); i++ {
		sum += _adjacencyMat[_cromosome[i]][_cromosome[i+1]]
	}
	return sum
}

func evaluate(_population Population, _max int) Population {
	var distances []int
	for i := 0; i < len(_population.genes); i++ {
		distances = append(distances, distance(_population.genes[i], _population.adjacencyMat))
		//		fmt.Println(_population.genes[i])
		//		fmt.Println(distances[i])
		//		fmt.Printf("\n")
	}
	/*
		for i := 0; i < len(_population.genes); i++ {
			fmt.Println(_population.genes[i])
			fmt.Println(distances[i])
			fmt.Printf("\n")
		}
	*/

	//fmt.Printf("\n")
	var a = minMax(distances)
	//fmt.Println(len(distances))
	//var tmp []int
	//tmp = _population.genes[a[1]]
	//fmt.Println(_population.genes[a[1]])
	//print("\n\n\n\n")
	//fmt.Printf("\nMejor")
	//print("\n")
	//fmt.Println(_population.score)
	//fmt.Println(_population.best)
	//fmt.Println(tmp)
	//fmt.Println(a[0])
	if _max > a[0] {
		_population.score = a[0]
		_population.best = _population.genes[a[1]]
	}
	_population.parents = append(_population.parents, _population.genes[a[1]])
	_population.genes = deleteElement(_population.genes, a[1])
	//fmt.Println(_population.best)
	return _population
}

func selection(_population Population, _nPopulation int, _flag bool) Population {
	var k = _nPopulation / 2
	i := 0
	_population.parents = nil
	for i < k {
		if i == 0 && _flag == true {
			_population = evaluate(_population, 999999)
		} else {
			_population = evaluate(_population, _population.score)
		}

		i++
	}
	return _population
}

/*
func evaluate(_genes [][]int, _adjacencyMat [][]int) {
	var distances []int
	for i := 0; i < len(_genes); i++ {
		distances = append(distances, distance(_genes[i], _adjacencyMat))
		fmt.Println(_genes[i])
		fmt.Println(distances[i])
	}
}
*/

func minMax(array []int) []int {
	var min = 99999999
	var j = -1
	//fmt.Println(array)
	for i := 0; i < len(array); i++ {
		if min > array[i] {
			min = array[i]
			j = i
		}
	}
	var values []int
	values = append(values, min)
	values = append(values, j)
	return values
}

// https://www.tutorialspoint.com/delete-elements-in-a-slice-in-golang
func deleteElement(_slice [][]int, _index int) [][]int {
	return append(_slice[:_index], _slice[_index+1:]...)
}

func swap(_gene []int) []int {
	var s = rand.Perm(5)[:2]
	//fmt.Println(s)
	var tmp = _gene[s[1]]
	_gene[s[1]] = _gene[s[0]]
	_gene[s[0]] = tmp
	_gene[len(_gene)-1] = _gene[0]
	return _gene
}

func mutate(_population Population, _nPopulation int) Population {
	_population.genes = [][]int{}
	var k = _nPopulation / 2
	i := 0
	var tmp []int
	j := 0
	for j < 5 {
		tmp = append(tmp, _population.best[j])
		j++
	}
	tmp = append(tmp, _population.best[0])
	for i < k {
		_population.genes = append(_population.genes, swap(_population.parents[i]))
		i++
	}
	_population.best = tmp
	return _population
}
func main() {
	/*
		var cities = []int{0, 1, 2, 3, 4}
		var adjacencyMat = [][]int{
			{0, 113, 147, 167, 56},
			{113, 0, 98, 142, 137},
			{147, 98, 0, 58, 135},
			{167, 142, 58, 0, 133},
			{56, 137, 135, 133, 0},
		}

		pop := NewPopulation(cities, adjacencyMat, 20)
		// Selección de varios padres
		fmt.Println(pop.genes)
		pop = selection(pop, 20)
		fmt.Println(pop.genes)
		fmt.Println(pop.parents)
		fmt.Println(pop.score)
		fmt.Println(pop.best)
		pop = mutate(pop, 20)
		fmt.Println(pop.genes)
		fmt.Println(pop.score)
		fmt.Println(pop.best)
	*/
	////////////////////////////////////////////////////////////
	//print(cities[0])
	//print(adjacencyMat[1][4])
	//fmt.Println(Permutations(cities))
	//fmt.Println(len(Permutations(cities)))
	//var cities2 = [5]int{}
	//fmt.Println(cities2)
	//fmt.Println(rand.Perm(120)[:5])
	//fmt.Println(adjacencyMat)
	//fmt.Println(RandomPermutations(cities, 5))
	//fmt.Println(pop.genes)
	//fmt.Println(distance(pop.genes[0], pop.adjacencyMat))
	// Selección inicial de un padre
	/*
		pop = evaluate(pop)
		fmt.Println(pop.genes)
		fmt.Println(pop.score)
		fmt.Println(pop.best)
	*/
	// Selección de varios padres
	var cities = []int{0, 1, 2, 3, 4}
	var adjacencyMat = [][]int{
		{0, 113, 147, 167, 56},
		{113, 0, 98, 142, 137},
		{147, 98, 0, 58, 135},
		{167, 142, 58, 0, 133},
		{56, 137, 135, 133, 0},
	}

	var n = 20
	pop := NewPopulation(cities, adjacencyMat, n)
	var flag = true
	k := 0
	for k < 3 {
		fmt.Printf("Generación ")
		fmt.Println(k + 1)
		fmt.Println(pop.genes)
		pop = selection(pop, n, flag)
		fmt.Println(pop.genes)
		pop = mutate(pop, n)
		fmt.Println(pop.parents)
		fmt.Println(pop.score)
		fmt.Println(pop.best)
		n = n / 2
		flag = false
		k++
	}

}
