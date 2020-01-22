package main

import "fmt"

// Animal ...
type Animal struct {
	name string
	age  int
	sex  string
}

// Speak ...
func (a Animal) Speak() (bool, error) {
	fmt.Printf("Dog... bow bow\n")
	return true, nil
}

func (a Animal) String() string {
	return fmt.Sprintf("Animal:: name:'%v' Sex: '%v' age: '%v' \n\n", a.name, a.sex, a.age)
}

// Human ...
type Human struct {
	name string
	age  int
	sex  string
}

// Speak ../
func (h Human) Speak() (bool, error) {
	fmt.Printf("Hello Friend, how are you?\n")
	return true, nil
}

func (h Human) String() string {
	return fmt.Sprintf("Human: name: '%v' sex: '%v' age: '%v'\n\n", h.name, h.sex, h.age)
}

// Iface ...
type Iface interface {
	Speak() (bool, error)
	// Describe()
	String() string
}

func main() {
	var if2 Iface = Animal{"Dog", 3, "Male"}
	var if1 Iface = Human{name: "Robert", age: 45, sex: "male"}
	var if3 Iface = Animal{}
	var if4 Iface = Human{}

	if2.Speak()
	fmt.Println(if2.String())
	// if2.Describe()
	if1.Speak()
	fmt.Println(if1.String())
	// if1.Describe()

	// if3.Describe()
	if3.Speak()
	fmt.Println(if3.String())

	// if4.Describe()
	if4.Speak()
	fmt.Println(if4.String())

}
