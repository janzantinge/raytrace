package object

import "github.com/jzantinge/raytrace/core"

type Object interface {
    Material() *core.Material
    Hit(ray *core.Ray) (float64, *core.Point, *core.Vector, *core.UV)
}

type ObjectData struct {
    transform core.Matrix
    inverseTransform core.Matrix
    material core.Material
    children []Object
}
