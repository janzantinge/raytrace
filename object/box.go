package object

import (
    "math"

    "github.com/janzantinge/raytrace/core"
)

const EPSILON = 0.0000001

type Box struct {
    data ObjectData

    min, max core.Point
    uvBounds [3][2]core.UVBounds
}

func NewBox(transform *core.Matrix, material *core.Material, min, max *core.Point) *Box {
    return &Box {
        ObjectData {
            *transform,
            *transform.Invert(),
            *material,
            nil},
        *min,
        *max,
        [3][2]core.UVBounds{
            {
                core.UVBounds {core.UV {0.0, 1.0 / 3.0}, core.UV {0.25, 2.0 / 3.0}},
                core.UVBounds {core.UV {0.5, 1.0 / 3.0}, core.UV {0.75, 2.0 / 3.0}}},
            {
                core.UVBounds {core.UV {0.5, 2.0 / 3.0}, core.UV {0.75, 1.0}},
                core.UVBounds {core.UV {0.5, 0.0}, core.UV {0.75, 1.0 / 3.0}}},
            {
                core.UVBounds {core.UV {0.25, 1.0 / 3.0}, core.UV {0.5, 2.0 / 3.0}},
                core.UVBounds {core.UV {0.75, 1.0 / 3.0}, core.UV {1.0, 2.0 / 3.0}}}}}
}

func (box *Box) Material() *core.Material {
    return &box.data.material
}

func (box *Box) UV(dimension, back int, hitPoint *core.Point) *core.UV {
    var bounds = box.uvBounds[dimension][back]

    var uFactor, vFactor float64 = 0, 0

    switch dimension {
    case 0:
        uFactor = (hitPoint.Z() - box.min.Z()) /
                  (box.max.Z() - box.min.Z())
        vFactor = (hitPoint.Y() - box.min.Y()) /
                  (box.max.Y() - box.min.Y())

        if back == 0 {
            uFactor = 1.0 - uFactor
        }
    case 1:
        uFactor = (hitPoint.Z() - box.min.Z()) /
                  (box.max.Z() - box.min.Z())
        vFactor = (hitPoint.X() - box.min.X()) /
                  (box.max.X() - box.min.X())

        if back == 1 {
            vFactor = 1.0 - vFactor
        }
    case 2:
        uFactor = (hitPoint.X() - box.min.X()) /
                  (box.max.X() - box.min.X())
        vFactor = (hitPoint.Y() - box.min.Y()) /
                  (box.max.Y() - box.min.Y())

        if back == 1 {
            uFactor = 1.0 - uFactor
        }
    }

    return &core.UV {
        (bounds.Max.U - bounds.Min.U) * uFactor + bounds.Min.U,
        (bounds.Max.V - bounds.Min.V) * (1.0 - vFactor) + bounds.Min.V}
}

func (box *Box) Hit(initialRay *core.Ray) (float64, *core.Point, *core.Vector, *core.UV) {
    var ray = initialRay.Transform(&box.data.inverseTransform)

    var hitMin = 0.0
    var hitMax = math.MaxFloat64

    var normal = core.NewVector(0, 0, 0)
    var back = 0
    var dimension = 0

    for d := 0; d < 3; d += 1 {
        if math.Abs(ray.Direction.Values[d]) < EPSILON {
            if ray.Origin.Values[d] < box.min.Values[d] ||
               ray.Origin.Values[d] > box.max.Values[d] {
                return math.MaxFloat64, nil, nil, nil
            }
        } else {
            var ood = 1.0 / ray.Direction.Values[d]
            var hit1 = (box.min.Values[d] - ray.Origin.Values[d]) * ood
            var hit2 = (box.max.Values[d] - ray.Origin.Values[d]) * ood

            var normalFactor = -1.0
            var b = 0
            if hit1 > hit2 {
                hit1, hit2 = hit2, hit1
                normalFactor = 1
                b = 1
            }

            hitMin = math.Max(hitMin, hit1)
            hitMax = math.Min(hitMax, hit2)

            if hitMin > hitMax {
                return math.MaxFloat64, nil, nil, nil
            }

            if hitMin == hit1 {
                normal = core.NewVector(0, 0, 0)
                normal.Values[d] = normalFactor

                dimension = d
                back = b
            }
        }
    }

    var hitPoint = ray.Point(hitMin)
    var worldHitPoint = core.MultMatrixPoint(&box.data.transform, hitPoint)

    return core.SubPoint(worldHitPoint, &initialRay.Origin).Length(),
           worldHitPoint,
           core.Normalize(
               core.SubPoint(
                   core.MultMatrixPoint(&box.data.transform, core.AddPoint(hitPoint, normal)),
                   worldHitPoint)),
           box.UV(dimension, back, hitPoint)
}
