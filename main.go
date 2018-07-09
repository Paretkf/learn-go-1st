// package main

// import (
// 	"fmt"
// )

// func init () {
// 	println("โหล --> init ()")
// }

// func main() {
// 	var name string
// 	print("ใส่ชื่อก่อนนะ : ")
// 	fmt.Scanf("%s", &name)
// 	fmt.Printf("โหล : %s\n", name)

// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("เลข  : %d\n", i)
// 	}

// 	// var number int32
// 	// var bigNumber int64
// 	// var str string
// 	// var arr = []string { // this is array
// 	// 	"a",
// 	// 	"b"
// 	// }
// 	// var slice1 []string // this is slice
// 	slice2 := []string{}
// 	slice2 = append(slice2, "aa")
// 	fmt.Printf("สไลซ์ : %v\n", slice2)

// 	mapVal := map[string]string{
// 		"key": "valuess",
// 	}
	
// 	mapVals := map[string]interface{}{
// 		"key": "value",
// 		"number": 80,
// 	}

// 	mapValss := map[string]interface{}{
// 		"key": "value",
// 		"number": 80,
// 		"nestedMap": map[string]interface{}{
// 			"childKey": "childValues",
// 		},
// 	}

// 	fmt.Printf("mapVal : %v\n", mapVal)

// 	fmt.Printf("mapVals : %v\n", mapVals)

// 	fmt.Printf("mapVals : %v\n", mapValss)

// 	hello("Rachata", "Tongpakdee")

// 	print(" ", ifCondition(3))

// 	x,_ := Prinln ("XXX", "YYY", "ZZZ")

// 	y,z := Prinln ("XXX", "YYY", "ZZZ")

// 	// PrintIntln (10, 20, "XXX", "ZZZZ")
// }

// func hello (name, lastName string) {
// 	fmt.Println("สวัสดี : ", name, lastName)
// }

// func Prinln(a ...string) (string, string){
// 	for _, e := range a {
// 		print(e + " ")
// 	}
// 	return "Prinln", "xxxx"
// }

// func ifCondition (i int) {
// 	if i % 2 == 0 {
// 		return "เลขคู่นะ"
// 	} else if i := i+2; i % 2 == 0{
// 		return "เลขคี่นะ"
// 	}
// }


package main

import (
	"fmt"
	"math"
)

func init () {
	fmt.Println("โหล --> init ()")
}

func main() { 
	sum,avg := cal (1,2,3,4,5,6,7)
	fmt.Println(sum, avg)
}

func cal (number ...int) (int, float32) {
	sum := 0
	count := 0
	for _, e := range number {
		if IsPrime (e) {
			sum += e
			count++
		}
	}
	return sum, (float32(sum) / float32(count))
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
			if value%i == 0 {
				return false
			}
	}
	return value > 1
}