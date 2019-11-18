package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("====== unsort =====")
	for k, v := range barVal {
		fmt.Printf("key : %v,value : %v\n", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("===== sorted =====")
	newVal := make(map[string]int)
	for _, k := range keys {
		fmt.Printf("key: %v,value:%v\n", k, barVal[k])
		newVal[k] = barVal[k]
	}
	fmt.Println(newVal)

	invMap := key_value_piar(barVal)
	fmt.Println(invMap)
}

func key_value_piar(param map[string]int) map[int]string {
	invMap := make(map[int]string)
	for k,v:=range param{
		invMap[v] = k
	}
	return invMap
}
