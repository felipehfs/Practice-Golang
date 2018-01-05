package main 

import "fmt"

type Human struct {
 	name string
 	age int
 }

 func (h *Human) Walk() string {
 	return fmt.Sprintf("%s is walking!", h.name)
 }
 func main() {
 	h := Human{"Felipe", 23}
 	fmt.Println(h.Walk())
 }