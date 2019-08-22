package light

import (
    "math"
    "image/color"

    "github.com/janzantinge/raytrace/core"
)

type SpotLight struct {
    data LightData

    position core.Point
    direction core.Vector
    angle float64
    exponent float64
}

func NewSpotLight(
    data *LightData,
    position *core.Point,
    direction *core.Vector,
    angle float64,
    exponent float64) *SpotLight {

    return &SpotLight {*data, *position, *direction, angle, exponent}
}

func (light *SpotLight) Color() *color.RGBA {
    return &light.data.c
}

func (light *SpotLight) Scale() *LightScale {
    return &light.data.scale
}

func (light *SpotLight) Attenuation() *LightAttenuation {
    return &light.data.attenuation
}

func (light *SpotLight) CalculateDirection(point *core.Point) *core.Vector {
    return core.Normalize(core.ScaleVector(&light.direction, -1))
}

func (light *SpotLight) CalculateIntensity(point *core.Point) float64 {
    var pointDirection = core.SubPoint(point, &light.position)
    var distanceSquared = core.Dot(pointDirection, pointDirection)
    var distance = math.Sqrt(distanceSquared)

    pointDirection = core.ScaleVector(pointDirection, 1.0 / distance)

    var lightAngle = core.Dot(&light.direction, pointDirection)

    var intensity = 0.0

    if lightAngle >= math.Cos(light.angle) {
        var numerator = math.Pow(lightAngle, light.exponent)

        var denominator = light.Attenuation().constant +
                          light.Attenuation().linear * distance +
                          light.Attenuation().quadratic * distanceSquared

        intensity = numerator / denominator
    }

    return intensity
}
