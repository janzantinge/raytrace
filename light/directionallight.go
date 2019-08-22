package light

import (
    "image/color"

    "github.com/janzantinge/raytrace/core"
)

type DirectionalLight struct {
    data LightData
    direction core.Vector
}

func NewDirectionalLight(data *LightData, direction *core.Vector) *DirectionalLight {
    return &DirectionalLight {*data, *direction}
}

func (light *DirectionalLight) Color() *color.RGBA {
    return &light.data.c
}

func (light *DirectionalLight) Scale() *LightScale {
    return &light.data.scale
}

func (light *DirectionalLight) Attenuation() *LightAttenuation {
    return &light.data.attenuation
}

func (light *DirectionalLight) CalculateDirection(point *core.Point) *core.Vector {
    return core.Normalize(core.ScaleVector(&light.direction, -1))
}

func (light *DirectionalLight) CalculateIntensity(point *core.Point) float64 {
    return 1.0
}
