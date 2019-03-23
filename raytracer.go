package main

import (
    "fmt"
    "time"
    "image"
    "image/color"

    "github.com/jzantinge/raytrace/core"
    "github.com/jzantinge/raytrace/view"
)

type RayTracer struct {
    camera view.Camera
    scene Scene
}

var startColor = &color.RGBA {0, 0, 0xff, 0xff}
var endColor = &color.RGBA {0x0f, 0x0f, 0xff, 0xff}

type Block struct {
    startx, starty, endx, endy uint32
}

func (rt *RayTracer) GenerateImage(
    blockChannel chan Block,
    done chan bool,
    result *image.RGBA) {

loop:
    for {
        select {
        case block := <-blockChannel:
            for y := block.starty; y < block.endy; y += 1 {
                for x := block.startx; x < block.endx; x += 1 {
                    var ray = rt.camera.Ray(x, y)
                    var color = rt.scene.Trace(ray)
                    if color == nil {
                        color = core.LerpColor(
                            startColor,
                            endColor,
                            core.Linear(
                                0.0,
                                float64(rt.camera.ScreenHeight()),
                                float64(y)))
                    }

                    result.Set(int(x), int(y), *color)
                }
            }
        default:
            break loop
        }
    }

    done <- true
}

const GOROUTINE_COUNT = 4
const X_BLOCK_COUNT = uint32(8)
const Y_BLOCK_COUNT = uint32(8)

func (rt *RayTracer) Run() image.Image {
    var start = time.Now()

    var result = image.NewRGBA (
        image.Rectangle {
            image.Point{0, 0},
            image.Point{int(rt.camera.ScreenWidth()), int(rt.camera.ScreenHeight())}})

    blocks := make(chan Block, X_BLOCK_COUNT * Y_BLOCK_COUNT)

    var xBlockSize = rt.camera.ScreenWidth() / X_BLOCK_COUNT
    var yBlockSize = rt.camera.ScreenHeight() / Y_BLOCK_COUNT

    var xStart, yStart uint32
    xEnd, yEnd := xBlockSize, yBlockSize

    for xBlock := uint32(0); xBlock < X_BLOCK_COUNT; xBlock += 1 {
        for yBlock := uint32(0); yBlock < Y_BLOCK_COUNT; yBlock += 1 {
            blocks <- Block{xStart, yStart, xEnd, yEnd}

            yStart = yEnd
            yEnd += yBlockSize
        }

        xStart = xEnd
        xEnd += xBlockSize

        yStart = uint32(0)
        yEnd = yBlockSize
    }

    var done [GOROUTINE_COUNT]chan bool
    for index := range done {
        done[index] = make(chan bool)
        go rt.GenerateImage(blocks, done[index], result)
    }

    for index := range done {
        <-done[index]
    }

    fmt.Println("time taken", time.Since(start))

    return result
 }
