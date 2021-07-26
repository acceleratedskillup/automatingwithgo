package main

import "fmt"

func main() {
	//example1()
	//example2()
	example3()
	//example4()
	//example5()
}

//assigning maps to variables
func example1() {
	nameToScoreMap := map[string]int{
		"john": 64,
		"may":  84,
	}
	fmt.Printf("nameToScoreMap) m-addr:%p\n", &nameToScoreMap)
	fmt.Printf("%v\n", nameToScoreMap)
	copiedMap := nameToScoreMap
	fmt.Printf("copiedMap) m-addr:%p\n", &copiedMap)
	fmt.Printf("%v\n", copiedMap)
	copiedMap["john"] = 5
	fmt.Println("after new assignment")
	fmt.Printf("copiedMap) m-addr:%p\n", &copiedMap)
	fmt.Printf("%v\n", copiedMap)
	fmt.Printf("nameToScoreMap) m-addr:%p\n", &nameToScoreMap)
	fmt.Printf("%v\n", nameToScoreMap)
}

//passing maps to functions
func example2() {
	nameToScoreMap := map[string]int{
		"john": 64,
		"may":  84,
	}
	fmt.Printf("nameToScoreMap) m-addr:%p\n", &nameToScoreMap)
	fmt.Printf("%v\n", nameToScoreMap)
	mutateMap(nameToScoreMap)
	fmt.Println("after mutation")
	fmt.Printf("nameToScoreMap) m-addr:%p\n", &nameToScoreMap)
	fmt.Printf("%v\n", nameToScoreMap)
}

func mutateMap(input map[string]int) {
	input["peter"] = 50
	fmt.Printf("input) m-addr:%p\n", &input)
	fmt.Printf("%v\n", input)
}

//returning maps from functions
func example3() {
	nameToScoreMap := map[string]int{
		"john": 64,
		"may":  84,
	}
	fmt.Printf("nameToScoreMap) m-addr:%p\n", &nameToScoreMap)
	fmt.Printf("%v\n", nameToScoreMap)
	nameToScoreMap = mutateMapWithReturn(nameToScoreMap)
	fmt.Println("after mutation with return")
	fmt.Printf("nameToScoreMap) m-addr:%p\n", &nameToScoreMap)
	fmt.Printf("%v\n", nameToScoreMap)
}
func mutateMapWithReturn(input map[string]int) map[string]int {
	input["peter"] = 50
	fmt.Printf("input) m-addr:%p\n", &input)
	fmt.Printf("%v\n", input)
	return input
}

//testing existings of a key in a map
func example4() {
	nameToScoreMap := map[string]int{
		"john": 64,
		"may":  84,
	}

	if score, ok := nameToScoreMap["john"]; ok {
		fmt.Println("score:", score)
	} else {
		fmt.Println("key not found")
	}

	//To test for a key without retrieving the value,
	//use an underscore in place of the first value
	if _, ok := nameToScoreMap["john"]; ok {
		fmt.Println("key found")
	} else {
		fmt.Println("key not found")
	}
}

//how to delete keys
func example5() {
	nameToScoreMap := map[string]int{
		"john": 64,
		"may":  84,
	}
	fmt.Printf("before %v\n", nameToScoreMap)
	delete(nameToScoreMap, "john")
	delete(nameToScoreMap, "john")
	delete(nameToScoreMap, "john")
	fmt.Printf("after %v\n", nameToScoreMap)
}
