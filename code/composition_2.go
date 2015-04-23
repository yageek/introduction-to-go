package main

import "fmt"

type Employee struct {
	Name string
}
type Boss struct {
	Employee
}

func (e *Employee) Work() {
	fmt.Println("Work Hard")
}
func (e *Employee) FillTimeSheet() {
	fmt.Println("I fill my time sheet as", e.Name)
}
func (b *Boss) Work() {
	b.Employee.Work()
	fmt.Println("Chilling :)")
}
func main() {
	boss := &Boss{Employee: Employee{"God"}}
	boss.FillTimeSheet()
	boss.Work()
}
