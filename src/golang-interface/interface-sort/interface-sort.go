package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []Person{
		{"Guowei", 26},
		{"Yuyingying", 23},
		{"Jiangchunjie", 25},
		{"Xiongruiqi", 24},
	}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (PersonList ByAge) Len() int {
	return len(PersonList)
}

func (PersonList ByAge) Swap(i, j int) {
	PersonList[i], PersonList[j] = PersonList[j], PersonList[i]
}

func (PersonList ByAge) Less(i, j int) bool {
	return PersonList[i].Age < PersonList[j].Age
}
