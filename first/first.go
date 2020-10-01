package first

import "fmt"

type Person struct {
	Name string
}

// export first
func Firsteg() {
	var i interface{}
	i = 4
	fmt.Println(i)
	i = 4.5

	//interface type conversion
	j := int(i.(float64))

	fmt.Println(i, j)
	i = "Parit"
	fmt.Println(i)

	//interface can hold non-primitive type data as well
	i = Person{Name: "Parit Sharma"}
	fmt.Println(i)
}

//export Secondeg
func Secondeg() {
	var i interface{}

	i = 3
	checkType(i)

	i = 3.5
	checkType(i)

	i = "Parit"
	checkType(i)

	type MyCustomeType uint
	i = MyCustomeType(5)
	checkType(i)
}

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("i am int")
	case float64:
		fmt.Println("I am float")
	case string:
		fmt.Println("I am string")
	default:
		fmt.Println("I am otherwise")
	}
}
