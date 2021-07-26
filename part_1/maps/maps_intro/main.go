package main

import "fmt"

func main() {
	example1()
	example2()
	example3()
}

func example1() {
	/*
		map[KeyType]ValueType
		creating a nil map below
		The zero value of a map is nil.
		A nil map has no keys, nor can keys be added.
	*/
	var nameToScoreMap map[string]int
	fmt.Printf("nameToScoreMap) len:%d\n", len(nameToScoreMap))
	fmt.Printf("%v\n", nameToScoreMap)

	/* examples of valid keys
	map[int]int
	map[byte]int
	*/

	/* examples of invalid keys
	map[[]string]bool
	map[map[string]int]string
	*/

	/* examples of valid values
	map[string][]int
	map[string]map[string]string
	*/

	fmt.Println("nameToScoreMap[angela]:", nameToScoreMap["angela"])
	nameToScoreMap["angela"] = 90 //will result in error
}

func example2() {
	//another way to create a map
	nameToScoreMap := map[string]int{
		"john": 64,
		"may":  84,
	}
	fmt.Printf("nameToScoreMap) len:%d\n", len(nameToScoreMap))
	fmt.Printf("%v\n", nameToScoreMap)

	fmt.Println("nameToScoreMap[angela]:", nameToScoreMap["angela"])
	nameToScoreMap["angela"] = 90

	fmt.Printf("nameToScoreMap) len:%d\n", len(nameToScoreMap))
	fmt.Printf("%v\n", nameToScoreMap)

	nameToScoreMap["angela"] = 100
	fmt.Printf("nameToScoreMap) len:%d\n", len(nameToScoreMap))
	fmt.Printf("%v\n", nameToScoreMap)
}

func example3() {
	/*The make function returns a map of the given type, initialized and ready for use.

	 */
	nameToScoreMap := make(map[string]int)
	fmt.Printf("nameToScoreMap) len:%d\n", len(nameToScoreMap))
	fmt.Printf("%v\n", nameToScoreMap)

}
