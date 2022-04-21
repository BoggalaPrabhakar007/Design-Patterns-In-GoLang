/*Liskov substitution principle*/
/*In an object-oriented class-based language, the concept of the Liskov substitution principle
is that a user of a base class should be able to function properly of all derived classes.*/

package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	Width, Height int
}

func (r *Rectangle) GetWidth() int {
	return r.Width
}

func (r *Rectangle) SetWidth(width int) {
	r.Width = width
}

func (r *Rectangle) GetHeight() int {
	return r.Height
}

func (r *Rectangle) SetHeight(height int) {
	r.Height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Printf("expexted area:%d,actual area:%d", expectedArea, actualArea)
}

type Square struct{
	Rectangle
}

func NewSquare(size int)*Square{
	sq:=Square{}
	sq.Width=size
	sq.Height=size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.Height = width
	s.Width = width
}

func (s *Square) SetHeight(height int) {
	s.Height = height
	s.Width = height
}


func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)
	sq:=NewSquare(3)
	UseIt(sq)
}
