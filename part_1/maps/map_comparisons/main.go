package main

import (
	"fmt"
	"reflect"
)

func main() {
	//example1()
	//example2()
	example3()
}
func example1() {
	map1 := map[string]int{
		"john": 64,
		"may":  84,
	}

	map2 := map[string]int{
		"john": 64,
		"may":  84,
	}
	//fmt.Println("isEqual?", map1 == map2)
	fmt.Println("isNil?", map1 == nil)
	fmt.Println("isNil?", map2 == nil)
}

func example2() {
	map1 := map[string]int{
		"john": 64,
		"may":  84,
	}

	map2 := map[string]int{
		"john": 64,
		"may":  84,
	}
	equal := reflect.DeepEqual(map1, map2)
	fmt.Println(equal)
}

func example3() {
	map1 := map[string]int{
		"john": 64,
		"may":  84,
	}

	map2 := map[string]int{
		"john": 64,
		"may":  84,
	}
	map1String := fmt.Sprintf("%v", map1)
	map2String := fmt.Sprintf("%v", map2)

	fmt.Println(map1String)
	fmt.Println(map2String)
	fmt.Println(map1String == map2String)
}
