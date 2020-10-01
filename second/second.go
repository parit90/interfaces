package second

import "fmt"

type GetName interface {
	getName() string
}

type Student struct {
	Name string
	Age  int
	ID   string
}

type Teacher struct {
	Name string
}

func (s Student) getName() string {
	return s.Name + "@gmail.com"
}

func (t Teacher) getName() string {
	return t.Name + "@gmail.com"
}

//export secondIntf
func SecondIntf() {
	s := Student{
		Name: "Parit",
		Age:  29,
		ID:   "123ty",
	}

	t := Teacher{
		Name: "Sharma",
	}

	str := GetName.getName(s)
	fmt.Println("Student :", str)
	str = GetName.getName(t)
	fmt.Println("Teacher :", str)
}
