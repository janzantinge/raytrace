package core

import "fmt"
import "os"
import "image"
import "image/color"
import "image/draw"
import _ "image/jpeg"
import _ "image/png"

type UV struct {
    U, V float64
}

type UVBounds struct {
    Min, Max UV
}

type Material struct {
    Ambient, Diffuse, Specular color.RGBA
    SpecularExponent float64
    TextureMap, AmbientTextureMap, NormalMap, SpecularMap *image.RGBA
}

func LoadTexture(filepath string) *image.RGBA {
    reader, err := os.Open(filepath)
    if err != nil {
        fmt.Println("Failed to open", filepath, "-", err)

        return nil
    }
    defer reader.Close()

    texture, _, err := image.Decode(reader)
    if err != nil {
        fmt.Println("Failed to decode", filepath, "-", err)

        return nil
    }

    var rgbaTexture = image.NewRGBA(texture.Bounds())

    draw.Draw(rgbaTexture, texture.Bounds(), texture, image.ZP, draw.Src)

    return rgbaTexture
}

func NewSimpleMaterial(color *color.RGBA, specularExponent float64) *Material {
    return &Material {*color, *color, *color, specularExponent, nil, nil, nil, nil}
}

func NewSpecularMaterial(
    color, specular *color.RGBA,
    specularExponent float64) *Material {

    return &Material {*color, *color, *specular, specularExponent, nil, nil, nil, nil}
}

func NewTextureMaterial(
    color, specular *color.RGBA,
    specularExponent float64,
    textureFilepath, ambientTextureFilepath string) *Material {

    return &Material {
        *color,
        *color,
        *specular,
        specularExponent,
        LoadTexture(textureFilepath),
        LoadTexture(ambientTextureFilepath),
        nil,
        nil}
}

func (material *Material) TextureColor(uv *UV) *color.RGBA {
    if material.TextureMap == nil {
        return &material.Diffuse
    }

    var result = material.TextureMap.RGBAAt(
        LerpInt(0, material.TextureMap.Bounds().Max.X, uv.U),
        LerpInt(0, material.TextureMap.Bounds().Max.Y, uv.V))

    return &result
}

func (material *Material) AmbientTextureColor(uv *UV) *color.RGBA {
    if material.AmbientTextureMap == nil {
        return &material.Ambient
    }

    var result = material.AmbientTextureMap.RGBAAt(
        LerpInt(0, material.AmbientTextureMap.Bounds().Max.X, uv.U),
        LerpInt(0, material.AmbientTextureMap.Bounds().Max.Y, uv.V))

    return &result
}
