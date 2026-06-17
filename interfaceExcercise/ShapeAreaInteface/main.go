package main

import "fmt"

type AreaFinder interface {
	Area() string 
}

type Circle struct {
	radius int 
}

func (c *Circle) Area() string {
	return fmt.Sprintf("Area of circle os tjisad")
}


type Rectangle struct {
	length int
	breadth int  
}

func (c *Rectangle) Area() string  {
	return fmt.Sprintf("Area of Rectanggle os tjisad")
}





func main() {
	c:=&Circle{radius: 23}
	r:=&Rectangle{length: 2, breadth: 1}
	
	var a = []AreaFinder{c,r}

	for _, v:=range a {
		fmt.Println(v.Area())
	}


	

}