package dir1

import (
	"fmt"
)

/**
map
	map的声明：map[keyType]valueType
	map的初始化：make(map[keyType]valueType, [cap])
	声明并初始化
 		mapName := map[string]int {
			"key1" : value1,
			"key2" : value2,
		}
	map中是否存在键 value : ok := mapName["key"] 存在这个键 ok = true，否则为false
	map遍历 for range
		for k , v := range mapName {
			//  k , v
		}
		for k := range mapName {
			// k , mapName[k]
		}
	map中删除键值对
		delete函数删除键值对
*/

type numFromDir1 int
type NumFromDir1 int

type City struct {
	Name string
}

func MapTest() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 80
	scoreMap["李四"] = 90
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	fmt.Printf("type of a:%T\n", scoreMap)

	scoreMap2 := map[string]int{
		"k1": 50,
		"k2": 60,
	}
	fmt.Println(scoreMap2)

	// 判断map中键是否存在
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在")
	}

	// map遍历
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	for k := range scoreMap {
		fmt.Println(k, ":", scoreMap[k])
	}

	// delete
	delete(scoreMap, "张三")
	fmt.Println(scoreMap)

}
