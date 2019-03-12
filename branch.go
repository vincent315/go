package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int)  string{
	g := ""
	switch  {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d" , score))
	case score  < 60:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g

}
func fors() int {
	sum := 0
	for i := 1;i <= 100; i++{
		sum += 1
	}
	return sum
}
func forsc()  {


}
func main()  {
	const filename = "abc.txt"
	bytes, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",bytes)
	}
	fmt.Println(grade(20),
		grade(59),
		grade(61),
		grade(79),
		grade(89),
		grade(99),
		)
	fmt.Println("--------------------")
	fmt.Println(fors())
}
