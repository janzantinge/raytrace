package object

import (
    "math"

    "github.com/janzantinge/raytrace/core"
)

const CIRCLE_PERIMITER_LENGTH_FACTOR = 1.0 / (2.0 * math.Pi)

type Sphere struct {
    data ObjectData

    position core.Point
    radius float64
}

func NewSphere(
    transform *core.Matrix,
    material *core.Material,
    position *core.Point,
    radius float64) *Sphere {

    return &Sphere {
        ObjectData {
            *transform,
            *transform.Invert(),
            *material,
            nil},
        *position,
        radius}
}

func (sphere *Sphere) Material() *core.Material {
    return &sphere.data.material
}

func (sphere *Sphere) UV(normal *core.Vector) *core.UV {
    var angle = math.Atan2(normal.X(), normal.Z())

    return &core.UV {
        1.0 - ((angle * CIRCLE_PERIMITER_LENGTH_FACTOR) + 0.5),
        1.0 - ((normal.Y() * 0.5) + 0.5)}
}

func (sphere *Sphere) Hit(initialRay *core.Ray) (float64, *core.Point, *core.Vector, *core.UV) {
    var ray = initialRay.Transform(&sphere.data.inverseTransform)

    var m = core.SubPoint(&ray.Origin, &sphere.position);

    var b = core.Dot(m, &ray.Direction)
    var c = core.Dot(m, m) - (sphere.radius * sphere.radius)

    // if c > 0 {
    //     fmt.Println("Ahead of the sphere.", ray.origin, sphere.position, c, Dot(m, m), sphere.radius * sphere.radius)
    //     return math.MaxFloat64, nil, nil, nil
    // }

    if b > 0 {
        // The ray is facing away from the sphere
        return math.MaxFloat64, nil, nil, nil
    }

    var discriminent = (b * b) - c

    if discriminent < 0 {
        // ray missed the sphere
        return math.MaxFloat64, nil, nil, nil
    }

    var sqrtDiscriminent = math.Sqrt(discriminent)
    var factor = -b - sqrtDiscriminent

    if factor < 0 {
        // The ray started ahead of or within the sphere, see if the ray exits
        // the sphere.
        factor = -b + sqrtDiscriminent

        if factor < 0 {
            // The ray started ahead of the sphere, no collision
            return math.MaxFloat64, nil, nil, nil
        }
    }

    var hitPoint = ray.Point(factor)
    var worldHitPoint = core.MultMatrixPoint(&sphere.data.transform, hitPoint)
    var worldSpherePosition = core.MultMatrixPoint(&sphere.data.transform, &sphere.position)

    return factor,
           worldHitPoint,
           core.Normalize(core.SubPoint(worldHitPoint, worldSpherePosition)),
           sphere.UV(core.Normalize(core.SubPoint(hitPoint, core.Origin())))
}
