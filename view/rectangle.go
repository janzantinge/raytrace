package view

import "github.com/janzantinge/raytrace/core"

type Rectangle struct {
    center core.Point
    right, up core.Vector
    width, height float64
}

func NewRectangle(
    center *core.Point,
    right, up *core.Vector,
    width, height float64) *Rectangle {

    return &Rectangle {*center, *right, *up, width, height}
}

func (r *Rectangle) point(x, y float64) *core.Point {
    if x > r.width || y > r.height {
        return nil
    }

    x -= r.width * 0.5
    y -= r.height * 0.5

    return core.AddPoint(
        &r.center,
        core.AddVector(
            core.ScaleVector(&r.right, x),
            core.ScaleVector(&r.up, y)))
}
