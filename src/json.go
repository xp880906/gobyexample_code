package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page	int
	Fruits	[]string
}

type Response2 struct {
	Page	int		 `json:"page"`
	Fruits	[]string `json:"fruits"`
}

func _showMarshal(v interface{}) {
	b, _ := json.Marshal(v)
	fmt.Println(string(b))
}

func main() {
	_showMarshal(true)
	_showMarshal(1)
	_showMarshal(2.34)
	_showMarshal("gopher")
	_showMarshal([]string{"apple", "pear", "peach"})
	_showMarshal(map[string]int{"apple": 5, "lettuce": 7})
	_showMarshal(&Response1{Page: 1, Fruits: []string{"apple", "pear", "peach"}})
	_showMarshal(&Response2{Page: 1, Fruits: []string{"apple", "pear", "peach"}})

	byt := []byte(`{"num":6.13,"strs":["a","b","c"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(strs)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	rep := Response2{}
	if err := json.Unmarshal([]byte(str), &rep); err != nil {
		panic(rep)
	}
	fmt.Println(rep)
	fmt.Printf("%#v\n", rep.Fruits[0])

	encoder := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 1, "lettuce": 2}
	err := encoder.Encode(d)
	if err != nil {
		panic(err)
	}
}
