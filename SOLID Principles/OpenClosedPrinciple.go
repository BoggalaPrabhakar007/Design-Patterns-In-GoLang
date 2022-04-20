/*Open/Closed principle*/

/*Open Closed is considered one of the most important principles in an object-oriented class-based language.
The concept suggests that modules should be written in a way so that we can add new modules or new functionalities
in a module without requiring existing modules to be modified.*/

package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	Name  string
	Color Color
	Size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.Color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.Size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) {
	//Filters products based on size,color
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.Color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.Size == s.size
}

type BetterFilter struct {
}

func (bf *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"apple", green, small}
	tree := Product{"tree", green, large}
	house := Product{"house", blue, large}
	products := []Product{apple, tree, house}
	fmt.Printf("Green Products(old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf("%s is green\n", v.Name)
	}

	fmt.Printf("Green Products(new):\n")
	greenSpec:=ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf("%s is green\n", v.Name)
	}
}
