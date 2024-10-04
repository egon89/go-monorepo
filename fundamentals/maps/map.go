package maps

import "fmt"

func MapExample() {
	fmt.Println("> mapExample")
	myMap := map[string]int{"a": 10, "b": 20, "c": 30}
	fmt.Println(myMap)

	fmt.Printf("b: %d\n", myMap["b"])
	myMap["d"] = 40
	fmt.Println(myMap)

	delete(myMap, "b")
	fmt.Println(myMap)

	for key, value := range myMap {
		fmt.Printf("key: %v, value: %d\n", key, value)
	}

	anotherMap := make(map[string]int)
	anotherMap["e"] = 50
	fmt.Println(anotherMap)

	i, ok := anotherMap["e"]
	fmt.Printf("i: %v, ok: %v\n", i, ok)

	j, ok := anotherMap["f"]
	fmt.Printf("j: %v, ok: %v\n", j, ok)

	_, exists := anotherMap["e"]
	fmt.Printf("does 'e' exist?: %v\n", exists)

	/*
		map[a:10 b:20 c:30]
		b: 20
		map[a:10 b:20 c:30 d:40]
		map[a:10 c:30 d:40]
		key: c, value: 30
		key: d, value: 40
		key: a, value: 10
		map[e:50]
		i: 50, ok: true
		j: 0, ok: false
		does 'e' exist?: true
	*/
}
