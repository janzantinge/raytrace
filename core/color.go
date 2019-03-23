package core

import "math"
import "image/color"

const MULT_SCALE = 1.0 / 255.0

func ClampColorComponent(value float64) uint8 {
    return uint8(math.Max(0.0, math.Min(255.0, value)))
}

func ScaleColor(c *color.RGBA, scale float64) *color.RGBA {
    return &color.RGBA {
        ClampColorComponent(float64(c.R) * scale),
        ClampColorComponent(float64(c.G) * scale),
        ClampColorComponent(float64(c.B) * scale),
        ClampColorComponent(float64(c.A) * scale)}
}

func AddColor(first, second *color.RGBA) *color.RGBA {
    return &color.RGBA {
        ClampColorComponent(float64(first.R) + float64(second.R)),
        ClampColorComponent(float64(first.G) + float64(second.G)),
        ClampColorComponent(float64(first.B) + float64(second.B)),
        ClampColorComponent(float64(first.A) + float64(second.A))}
}

func MultColor(first, second *color.RGBA) *color.RGBA {
    return &color.RGBA {
        ClampColorComponent((float64(first.R) * float64(second.R)) * MULT_SCALE),
        ClampColorComponent((float64(first.G) * float64(second.G)) * MULT_SCALE),
        ClampColorComponent((float64(first.B) * float64(second.B)) * MULT_SCALE),
        ClampColorComponent((float64(first.A) * float64(second.A)) * MULT_SCALE)}
}
