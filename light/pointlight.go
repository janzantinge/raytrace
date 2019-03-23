package light

import (
    "math"
    "image/color"

    "github.com/jzantinge/raytrace/core"
)

type PointLight struct {
    data LightData
    position core.Point
}

func NewPointLight(data *LightData, position *core.Point) *PointLight {
    return &PointLight {*data, *position}
}

func (light *PointLight) Color() *color.RGBA {
    return &light.data.c
}

func (light *PointLight) Scale() *LightScale {
    return &light.data.scale
}

func (light *PointLight) Attenuation() *LightAttenuation {
    return &light.data.attenuation
}

func (light *PointLight) CalculateDirection(point *core.Point) *core.Vector {
    return core.Normalize(core.SubPoint(&light.position, point))
}

func (light *PointLight) CalculateIntensity(point *core.Point) float64 {
    var pointDirection = core.SubPoint(point, &light.position)
    var distanceSquared = core.Dot(pointDirection, pointDirection)
    var distance = math.Sqrt(distanceSquared)

    var denominator = light.Attenuation().constant +
                      light.Attenuation().linear * distance +
                      light.Attenuation().quadratic * distanceSquared

    return 1.0 / denominator
}
