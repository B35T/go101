package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page int
	Skills []string
}

type response2 struct {
	Page int	`json:"page"`
	Skills []string	`json:"skills"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB,_ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB,_ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB,_ := json.Marshal("gofer")
	fmt.Println(string(strB))

	slcB := &response1{1,[]string{"go", "swift", "R", "python"}}
	sclBJson,_ := json.Marshal(slcB)
	fmt.Println(string(sclBJson))

	slcC := &response2{2,[]string{"no","one","data"}}
	sclCJson , _ := json.Marshal(slcC)
	fmt.Println(string(sclCJson))


	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt , &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)


	str := `{"page":1,"skills":["name","apple"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Skills)
	fmt.Println(res.Skills[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple":5,"banana":10}
	enc.Encode(d)

}
