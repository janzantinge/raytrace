package core

import "math"

type Vector struct {
    Values [3]float64
}

func (v *Vector) X() float64 {
    return v.Values[0]
}

func (v *Vector) Y() float64 {
    return v.Values[1]
}

func (v *Vector) Z() float64 {
    return v.Values[2]
}

func (v *Vector) Length() float64 {
    var lengthSquared = Dot(v, v)
    if lengthSquared == 0 {
        return 0;
    }

    return math.Sqrt(lengthSquared)
}

func NewVector(x, y, z float64) *Vector {
    return &Vector {[3]float64{x, y, z}}
}

func ScaleVector(vector *Vector, factor float64) *Vector {
    return NewVector(vector.X() * factor, vector.Y() * factor, vector.Z() * factor)
}

func AddVector(a, b *Vector) *Vector {
    return NewVector(a.X() + b.X(), a.Y() + b.Y(), a.Z() + b.Z())
}

func SubVector(a, b *Vector) *Vector {
    return NewVector(a.X() - b.X(), a.Y() - b.Y(), a.Z() - b.Z())
}

func Dot(a, b *Vector) float64 {
    return a.X() * b.X() + a.Y() * b.Y() + a.Z() * b.Z()
}

func Cross(a, b *Vector) *Vector {
    return NewVector(
        a.Y() * b.Z() - a.Z() * b.Y(),
        a.Z() * b.X() - a.X() * b.Z(),
        a.X() * b.Y() - a.Y() * b.X())
}

func Tensor(a, b *Vector) *Matrix {
    return &Matrix {
        [4][4]float64{
            {a.X() * b.X(), a.X() * b.Y(), a.X() * b.Z(), 0.0},
            {a.Y() * b.X(), a.Y() * b.Y(), a.Y() * b.Z(), 0.0},
            {a.Z() * b.X(), a.Z() * b.Y(), a.Z() * b.Z(), 0.0},
            {          0.0,           0.0,           0.0, 1.0}}}
}

func Normalize(v *Vector) *Vector {
    var lengthDenominator = 1.0 / v.Length()
    return ScaleVector(v, lengthDenominator);
}

func Up() *Vector {
    return NewVector (0, 1, 0)
}

func Down() *Vector {
    return NewVector (0, -1, 0)
}

func Right() *Vector {
    return NewVector (1, 0, 0)
}

func Left() *Vector {
    return NewVector (-1, 0, 0)
}
