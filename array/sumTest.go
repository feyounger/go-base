package main

import "fmt"

func sumArr(arr *[5]int) {
	fmt.Printf("arr: %p\n", &arr)
	arr[0] = 12132
}

func main() {
	arr := []string{"1", "2", "4", "4", "5"}

	fmt.Println(removeRepByMap(arr))
}

func findArr(sum int, arr [5]int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			target := arr[i] + arr[j]
			if target == sum {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

//slice去重
func removeRepByMap(slc []string) []string {
	result := []string{}         //存放返回的不重复切片
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}
