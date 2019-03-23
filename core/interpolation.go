package core

import "image/color"

func Linear(min, max, current float64) float64 {
    return (current - min) / (max - min)
}

func LerpInt(first, second int, factor float64) int {
    return int(float64(first + second) * factor)
}

func LerpFloat(first, second, factor float64) float64 {
    return (first + second) * factor
}

func LerpColor(first, second *color.RGBA, factor float64) *color.RGBA {
    return &color.RGBA {
        uint8(float64(first.R + second.R) * factor),
        uint8(float64(first.G + second.G) * factor),
        uint8(float64(first.B + second.B) * factor),
        uint8(float64(first.A + second.A) * factor)}
}
