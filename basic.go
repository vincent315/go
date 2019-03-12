package main

import (
	"fmt"
	_ "fmt"
	"math"
	"math/cmplx"
)
var(
	a = 56565
)
func variable()  {
	var a int
	var s string
	fmt.Println(a,s)
	
}
func variable2()  {
	var a int = 3;
	var s string  = "sdsd"
	fmt.Println(a,s)

}
func variableType()  {
	var a,b,c,d = 3,"aaddad",true,'a'
	fmt.Println(a,b,c,d)
	
}
func variableSimple()  {
	a,b,c,d := 3,"aaddad",true,'a'
	a = 6
	fmt.Println("------------------")
	fmt.Println(a,b,c,d)
}
func enums()  {
	const(
		java = iota //自增
		cpp
		golang
	)
	fmt.Println(java,cpp,golang)
	
}
func eular()  {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	
}
func consts()  {
	const test = "test"
	const a,b = 3,4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(test,c)
	
}

func main() {
	fmt.Print("hello word")
	variable()
	variable2()
	variableType()
	variableSimple()
	fmt.Println(a)
	fmt.Println("----------------")
	eular()
	fmt.Println("----------------")
	consts()
	enums()
}
