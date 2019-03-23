package light

import (
    "math"
    "image/color"

    "github.com/jzantinge/raytrace/core"
)

type LightScale struct {
    ambient, diffuse, specular float64
}

func NewLightScale(ambient, diffuse, specular float64) *LightScale {
    return &LightScale {ambient, diffuse, specular}
}

type LightAttenuation struct {
    constant, linear, quadratic float64
}

func NewLightAttenuation(constant, linear, quadratic float64) *LightAttenuation {
    return &LightAttenuation {constant, linear, quadratic}
}

type Light interface {
    Color() *color.RGBA
    Scale() *LightScale
    Attenuation() *LightAttenuation
    CalculateDirection(point *core.Point) *core.Vector
    CalculateIntensity(point *core.Point) float64
}

type LightData struct {
    c color.RGBA
    scale LightScale
    attenuation LightAttenuation
}

func NewLightData(
    c *color.RGBA,
    scale *LightScale,
    attenuation *LightAttenuation) *LightData {

    return &LightData {*c, *scale, *attenuation}
}

func Ambient(light Light, intensity float64, color *color.RGBA) *color.RGBA {
    return core.ScaleColor(
        core.MultColor(
            core.ScaleColor(light.Color(), light.Scale().ambient),
            color),
        1.0)
}

func Diffuse(
    light Light,
    intensity float64,
    point *core.Point,
    normal *core.Vector,
    color *color.RGBA) *color.RGBA {

    return core.ScaleColor(
        core.MultColor(
            core.ScaleColor(
                core.ScaleColor(light.Color(), light.Scale().diffuse),
                math.Max(0.0, core.Dot(normal, light.CalculateDirection(point)))),
            color),
        intensity)
}

func Specular(
    light Light,
    intensity float64,
    viewDirection *core.Vector,
    point *core.Point,
    normal *core.Vector,
    material *core.Material) *color.RGBA {

    var lightDirection = light.CalculateDirection(point)

    if core.Dot(lightDirection, normal) < 0 {
        return &color.RGBA {0, 0, 0, 0xff}
    }

    var reflection = core.Normalize(
        core.SubVector(
            core.ScaleVector(
                core.ScaleVector(normal, core.Dot(lightDirection, normal)),
                2.0),
            lightDirection))

    var specularFactor = math.Pow(
        math.Max(
            0.0,
            core.Dot(reflection, viewDirection)),
        material.SpecularExponent)
    var specularColor = core.ScaleColor(&material.Specular, specularFactor)

    var lightColor = core.ScaleColor(
        core.ScaleColor(light.Color(), light.Scale().specular),
        intensity)

    return core.MultColor(lightColor, specularColor)
}
