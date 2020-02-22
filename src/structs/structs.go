package structs

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

// Shape interface has Area() method
type Shape interface {
	Area() float64
}

// Perimeter method takes returns a float64
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area method of Rectangle returns a float64
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area method of Circle returns a float64
func (c Circle) Area() float64 {
	return c.Radius * float64(3.14)
}
