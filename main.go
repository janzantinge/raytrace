package main

import (
    "flag"
    "math"
    "os"
    "image/jpeg"
    "image/color"

    "github.com/janzantinge/raytrace/core"
    "github.com/janzantinge/raytrace/view"
    "github.com/janzantinge/raytrace/object"
    "github.com/janzantinge/raytrace/light"
)

func main() {
    filepath := flag.String(
        "output",
        "/home/janzantinge/projects/trace.jpeg",
        "The desired output image filepath.")

    flag.Parse()

    var camera = view.NewCamera (
        core.NewPoint(0, 50, 0),
        view.NewRectangle (
            core.NewPoint(0, 50, 10),
            core.Right(),
            core.Up(),
            32,
            18),
        view.NewResolution (3840, 2160))

    var scene = Scene {
        []object.Object {
            object.NewSphere(
                core.Identity(),
                core.NewSpecularMaterial(
                    &color.RGBA {0xff, 0x0f, 0x0f, 0xff},
                    &color.RGBA {0xff, 0xff, 0xff, 0xff},
                    10.0),
                core.NewPoint(-50, 50, 100),
                10),
            object.NewSphere(
                core.MultMatrix(
                    core.MultMatrix(
                        core.NewTranslation(0, 50, 100),
                        core.NewRotation(0, 65, 0)),
                    core.NewScale(40, 40, 40)),
                core.NewSpecularMaterial(
                    &color.RGBA {0xff, 0x0f, 0x0f, 0xff},
                    &color.RGBA {0xff, 0xff, 0xff, 0xff},
                    1.0),
                core.Origin(),
                1),
            object.NewSphere(
                core.MultMatrix(
                    core.NewTranslation(75, 50, 100),
                    core.NewScale(2, 1, 1)),
                core.NewSpecularMaterial(
                    &color.RGBA {0xff, 0x0f, 0x0f, 0xff},
                    &color.RGBA {0xff, 0xff, 0xff, 0xff},
                    1.0),
                core.Origin(),
                15),
            object.NewBox(
                core.MultMatrix(
                    core.NewTranslation(-70, 90, 80),
                    core.NewRotation(0, math.Pi, 0)),
                core.NewSpecularMaterial(
                    &color.RGBA {0xff, 0xff, 0x0f, 0xff},
                    &color.RGBA {0xff, 0xff, 0xff, 0xff},
                    1.0),
                core.NewPoint(-10, -10, -10),
                core.NewPoint(10, 10, 10)),
            object.NewBox(
                core.NewTranslation(70, 25, 80),
                core.NewSpecularMaterial(
                    &color.RGBA {0xff, 0xff, 0x0f, 0xff},
                    &color.RGBA {0xff, 0xff, 0xff, 0xff},
                    1.0),
                core.NewPoint(-10, -10, -10),
                core.NewPoint(10, 10, 10)),
            object.NewBox(
                core.MultMatrix(
                    core.NewTranslation(0, -50, 100),
                    core.NewScale(50, 1, 50)),
                core.NewSpecularMaterial(
                    &color.RGBA {0x0f, 0xff, 0x0f, 0xff},
                    &color.RGBA {0xff, 0xff, 0xff, 0xff},
                    10.0),
                core.NewPoint(-5, -5, -5),
                core.NewPoint(5, 5, 5))},
        []light.Light {
            light.NewSpotLight (
                light.NewLightData (
                    &color.RGBA {0xff, 0xff, 0xff, 0xff},
                    light.NewLightScale(0.2, 0.2, 0.001),
                    light.NewLightAttenuation(1.0, 0.0, 0.0000001)),
                core.NewPoint(0.0, 200.0, 10.0),
                core.Normalize(core.NewVector(-0.2, -1, 0.5)),
                math.Pi * 0.125,
                3.0),
            light.NewSpotLight (
                light.NewLightData (
                    &color.RGBA {0xff, 0xff, 0xff, 0xff},
                    light.NewLightScale(0.2, 0.2, 0.001),
                    light.NewLightAttenuation(1.0, 0.0, 0.0000001)),
                core.NewPoint(0.0, 200.0, 10.0),
                core.Normalize(core.NewVector(0, -1, 0.5)),
                math.Pi * 0.125,
                3.0),
            light.NewSpotLight (
                light.NewLightData (
                    &color.RGBA {0xff, 0xff, 0xff, 0xff},
                    light.NewLightScale(0.2, 0.2, 0.001),
                    light.NewLightAttenuation(1.0, 0.0, 0.0000001)),
                core.NewPoint(0.0, 200.0, 10.0),
                core.Normalize(core.NewVector(0.2, -1, 0.5)),
                math.Pi * 0.125,
                3.0)}}

    var rayTracer = RayTracer {*camera, scene}

    var imageData = rayTracer.Run()

    var file, _ = os.Create(*filepath)

    jpeg.Encode(file, imageData, &jpeg.Options {100})
}
