package main

import (
	"fmt"
	"sync"
)

type PermutationEntry []int

func getRestSlice(input PermutationEntry, usedIndex int) []int {
	ret := []int{}
	for i, v := range input {
		if i != usedIndex {
			ret = append(ret, v)
		}
	}
	return ret
}

func permutationsAsync(input PermutationEntry, waitGroup *sync.WaitGroup, out chan<- []PermutationEntry) {
	defer waitGroup.Done()
	inputLen := len(input)
	ret := []PermutationEntry{}
	if inputLen == 2 {
		ret = append(ret, input)
		ret = append(ret, make([]int, 2))
		ret[1][0] = input[1]
		ret[1][1] = input[0]
		out <- ret
	} else {
		ret := createPermutationsArray(inputLen)
		permCount_0 := fakultaet(inputLen - 1)
		for i := 0; i < inputLen; i++ {
			restSlice := getRestSlice(input, i)
			retChannel := make(chan []PermutationEntry, 1)
			waitGroup.Add(1)
			go permutationsAsync(restSlice, waitGroup, retChannel)
			select {
			case restPermutations := <-retChannel:
				for j := 0; j < permCount_0; j++ {
					index := (i * permCount_0) + j
					ret[index][0] = input[i]
					for k := 1; k < inputLen; k++ {
						ret[index][k] = restPermutations[j][k-1]
					}
				}
			}
		}
		out <- ret
	}
}

func printPermutionsAsync(numberCount int) {
	input := PermutationEntry{}
	for i := 0; i < numberCount; i++ {
		input = append(input, i)
	}
	var waitGroup sync.WaitGroup
	retChannel := make(chan []PermutationEntry, 1)
	waitGroup.Add(1)
	go permutationsAsync(input, &waitGroup, retChannel)
	select {
	case permutationSlice := <-retChannel:
		/*
			for j, v := range permutationSlice {
				fmt.Printf("%d: %v\n", j, v)
			}
		*/
		waitGroup.Wait()
		fmt.Printf("Number of permutations: %d\n\n", len(permutationSlice))
	}
}

/*
func printPermutionsAsync2(numberCount int) {
	input := PermutationEntry{}
	for i := 0; i < numberCount; i++ {
		input = append(input, i)
	}
	var waitGroup sync.WaitGroup
	permutations := createPermutationsArray(numberCount)
	retChannel := make(chan []PermutationEntry, 1)
	waitGroup.Add(1)
	go permutationsAsync2(input, &waitGroup, retChannel, permutations)
	select {
	case permutationSlice := <-retChannel:
		waitGroup.Wait()
		for j, v := range permutationSlice {
			fmt.Printf("%d: %v\n", j, v)
		}
		fmt.Printf("Number of permutations: %d\n\n", len(permutationSlice))
	}
}
*/

func testFakultaet() {
	f1 := fakultaet(1)
	f2 := fakultaet(2)
	f3 := fakultaet(3)
	f4 := fakultaet(4)
	f5 := fakultaet(5)
	fmt.Printf("f1: %v\n", f1)
	fmt.Printf("f2: %v\n", f2)
	fmt.Printf("f3: %v\n", f3)
	fmt.Printf("f4: %v\n", f4)
	fmt.Printf("f5: %v\n", f5)
}

// start

func fakultaet(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	return result
}

func createPermutationsArray(numberCount int) []PermutationEntry {
	permutationsCount := fakultaet(numberCount)
	permutations := []PermutationEntry{}
	for i := 0; i < permutationsCount; i++ {
		permutations = append(permutations, make([]int, numberCount))
	}
	return permutations
}

func permutations(input PermutationEntry) []PermutationEntry {
	inputLen := len(input)
	if inputLen == 2 {
		ret := []PermutationEntry{}
		ret = append(ret, input)
		ret = append(ret, make([]int, 2))
		ret[0] = input
		ret[1] = make([]int, 2)
		ret[1][0] = input[1]
		ret[1][1] = input[0]
		return ret
	} else {
		ret := createPermutationsArray(inputLen)
		permCount_0 := fakultaet(inputLen - 1)
		for i := 0; i < inputLen; i++ {
			restSlice := getRestSlice(input, i)
			restPermutations := permutations(restSlice)
			for j := 0; j < permCount_0; j++ {
				index := (i * permCount_0) + j
				ret[index][0] = input[i]
				for k := 1; k < inputLen; k++ {
					ret[index][k] = restPermutations[j][k-1]
				}
			}
		}
		return ret
	}
}

func printPermutions(numberCount int) {
	input := PermutationEntry{}
	for i := 0; i < numberCount; i++ {
		input = append(input, i)
	}
	permutationSlice := permutations(input)
	fmt.Printf("Number of permutations: %d\n\n", len(permutationSlice))
	// for j, v := range permutationSlice {
	// 	fmt.Printf("%d: %v\n", j, v)
	// }
}

func main() {
	printPermutions(11)
	//printPermutionsAsync(5)
}
