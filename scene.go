package main

import (
    "math"
    "image/color"

    "github.com/janzantinge/raytrace/core"
    "github.com/janzantinge/raytrace/object"
    "github.com/janzantinge/raytrace/light"
)

type Scene struct {
    objects []object.Object
    lights []light.Light
}

func (s *Scene) Hit(ray *core.Ray) (*core.Point, *core.Vector, object.Object, *core.UV) {
    var nearest = math.MaxFloat64
    var nearestPoint *core.Point = nil
    var nearestNormal *core.Vector = nil
    var nearestObject object.Object = nil
    var nearestUV *core.UV = nil

    for _, o := range s.objects {
        var factor, point, normal, uv = o.Hit(ray)
        if factor < nearest {
            nearest = factor
            nearestPoint = point
            nearestNormal = normal
            nearestObject = o
            nearestUV = uv
        }
    }

    return nearestPoint, nearestNormal, nearestObject, nearestUV
}

func (scene *Scene) ApplyLight(
    l light.Light,
    point *core.Point,
    viewDirection *core.Vector,
    normal *core.Vector,
    material *core.Material,
    uv *core.UV) *color.RGBA {

    var intensity = l.CalculateIntensity(point)
    var lightDirection = l.CalculateDirection(point)

    // Whether an object is hit is the only thing that is important when
    // determining whether a point exists in a shadow
    var _, _, object, _ = scene.Hit(&core.Ray {*core.AddPoint(point, normal), *lightDirection})

    if object != nil {
        return light.Ambient(l, intensity, material.AmbientTextureColor(uv))
    }

    return core.AddColor(
        core.AddColor(
            light.Ambient(l, intensity, material.AmbientTextureColor(uv)),
            light.Diffuse(l, intensity, point, normal, material.TextureColor(uv))),
        light.Specular(l, intensity, viewDirection, point, normal, material))
}

func (s *Scene) Trace(ray *core.Ray) *color.RGBA {
    var point, normal, object, uv = s.Hit(ray)

    if object == nil {
        return nil
    }

    var viewVector = core.Normalize(core.SubPoint(&ray.Origin, point))
    var result = &color.RGBA{0, 0, 0, 0};
    for _, l := range s.lights {
        result = core.AddColor(
            result,
            s.ApplyLight(
                l,
                point,
                viewVector,
                normal,
                object.Material(),
                uv))
    }

    return result
}
