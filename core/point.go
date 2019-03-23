package core

type Point struct {
    Values [3]float64
}

func (p *Point) X() float64 {
    return p.Values[0]
}

func (p *Point) Y() float64 {
    return p.Values[1]
}

func (p *Point) Z() float64 {
    return p.Values[2]
}

func (p *Point) Sub(b *Point) *Vector {
    return NewVector(p.X() - b.X(), p.Y() - b.Y(), p.Z() - b.Z())
}

func NewPoint(x, y, z float64) *Point {
    return &Point {[3]float64{x, y, z}}
}

func SubPoint(a, b *Point) *Vector {
    return a.Sub(b)
}

func AddPoint(a *Point, v *Vector) *Point {
    return NewPoint(a.X() + v.X(), a.Y() + v.Y(), a.Z() + v.Z())
}

func Origin() *Point {
    return NewPoint(0, 0, 0)
}
