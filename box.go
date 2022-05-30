package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return fmt.Errorf("out of the shapesCapacity range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= 0 && i < len(b.shapes) {
		return b.shapes[i], nil
	}
	return nil, fmt.Errorf("shape by index doesn't exist or index went out of the range")
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= 0 && i < len(b.shapes) {
		result := b.shapes[i]
		newShapes := make([]Shape, 0)
		for ind, s := range b.shapes {
			if ind != i {
				newShapes = append(newShapes, s)
			}
		}
		b.shapes = newShapes
		return result, nil
	}
	return nil, fmt.Errorf("shape by index doesn't exist or index went out of the range")

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= 0 && i < len(b.shapes) {
		result := b.shapes[i]
		b.shapes[i] = shape
		return result, nil
	}
	return nil, fmt.Errorf("shape by index doesn't exist or index went out of the range")

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var p float64
	for _, s := range b.shapes {
		p += s.CalcPerimeter()
	}
	return p
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var a float64
	for _, s := range b.shapes {
		a += s.CalcArea()
	}
	return a
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	newShapes := make([]Shape, 0)
	removed := 0
	for _, s := range b.shapes {
		if _, ok := s.(Circle); !ok {
			newShapes = append(newShapes, s)
		} else {
			removed++
		}
	}
	b.shapes = newShapes
	if removed == 0 {
		return fmt.Errorf("circles are not exist in the list")
	}
	return nil
}
