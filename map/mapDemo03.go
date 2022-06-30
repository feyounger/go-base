package main

import "fmt"

func main() {
	sliceMap := make(map[string][]string, 3)
	fmt.Println(sliceMap)
	key := "中国"
	v, ok := sliceMap[key]
	if  !ok{
		v = make([]string,0,2)
	}
	v = append(v,"武汉","杭州")
	sliceMap[key] = v

	for key, slices := range sliceMap {
		fmt.Println("key: ", key)
		for _, value := range slices {
			fmt.Println(value)
		}
	}

}
