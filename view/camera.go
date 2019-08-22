package view

import "github.com/janzantinge/raytrace/core"

type Resolution struct {
    x, y uint32
}

func NewResolution(x, y uint32) *Resolution{
    return &Resolution {x, y}
}

type Camera struct {
    position core.Point
    screen Rectangle
    resolution Resolution
}

func NewCamera(position *core.Point, rectangle *Rectangle, resolution *Resolution) *Camera {
    return &Camera {*position, *rectangle, *resolution}
}

func (c *Camera) ScreenWidth() uint32 {
    return c.resolution.x
}

func (c *Camera) ScreenHeight() uint32 {
    return c.resolution.y
}

func (c *Camera) Ray(x, y uint32) *core.Ray {
    if x > c.resolution.x || y > c.resolution.y {
        return nil
    }

    var target = c.screen.point (
        core.LerpFloat(
            0.0,
            c.screen.width,
            core.Linear(0.0, float64(c.resolution.x), float64(x))),
        core.LerpFloat(
            0.0,
            c.screen.height,
            1.0 - core.Linear(0.0, float64(c.resolution.y), float64(y))))

    return &core.Ray {c.position, *core.Normalize(core.SubPoint(target, &c.position))}
}
