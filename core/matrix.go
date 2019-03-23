package core

import "math"

type Matrix struct {
    v [4][4]float64
}

func ScaleMatrix(m *Matrix, scale float64) *Matrix {
    var result *Matrix = new(Matrix)

    for i, row := range m.v {
        for j := range row {
            result.v[i][j] = scale * row[j]
        }
    }

    return result
}

func (m *Matrix) Transpose() *Matrix {
    var result *Matrix = new(Matrix)

    for i, row := range m.v {
        for j := range row {
            result.v[i][j] = m.v[j][i]
        }
    }

    return result
}

func (m *Matrix) Invert() *Matrix {
    var result *Matrix = new(Matrix)

    result.v[0][0] = (
        m.v[1][1] * (m.v[2][2] * m.v[3][3] - m.v[2][3] * m.v[3][2]) -
        m.v[1][2] * (m.v[2][1] * m.v[3][3] - m.v[2][3] * m.v[3][1]) +
        m.v[1][3] * (m.v[2][1] * m.v[3][2] - m.v[2][2] * m.v[3][1]))

    result.v[1][0] = (
        -m.v[0][1] * (m.v[2][2] * m.v[3][3] - m.v[2][3] * m.v[3][2]) +
        m.v[0][2] * (m.v[2][1] * m.v[3][3] - m.v[2][3] * m.v[3][1]) -
        m.v[0][3] * (m.v[2][1] * m.v[3][2] - m.v[2][2] * m.v[3][1]))

    result.v[2][0] = (
        m.v[0][1] * (m.v[1][2] * m.v[3][3] - m.v[1][3] * m.v[3][2]) -
        m.v[0][2] * (m.v[1][1] * m.v[3][3] - m.v[1][3] * m.v[3][1]) +
        m.v[0][3] * (m.v[1][1] * m.v[3][2] - m.v[1][2] * m.v[3][1]))

    result.v[3][0] = (
        -m.v[0][1] * (m.v[1][2] * m.v[2][3] - m.v[1][3] * m.v[2][2]) +
        m.v[0][2] * (m.v[1][1] * m.v[2][3] - m.v[1][3] * m.v[2][1]) -
        m.v[0][3] * (m.v[1][1] * m.v[2][2] - m.v[1][2] * m.v[2][1]))

    var determinant = (
        m.v[0][0] * result.v[0][0] +
        m.v[1][0] * result.v[1][0] +
        m.v[2][0] * result.v[2][0] +
        m.v[3][0] * result.v[3][0])

    if determinant == 0 {
        return nil
    }

    result.v[0][1] = (
        -m.v[1][0] * (m.v[2][2] * m.v[3][3] - m.v[2][3] * m.v[3][2]) +
        m.v[1][2] * (m.v[2][0] * m.v[3][3] - m.v[2][3] * m.v[3][0]) -
        m.v[1][3] * (m.v[2][0] * m.v[3][2] - m.v[2][2] * m.v[3][0]))

    result.v[0][2] = (
        m.v[1][0] * (m.v[2][1] * m.v[3][3] - m.v[2][3] * m.v[3][1]) -
        m.v[1][1] * (m.v[2][0] * m.v[3][3] - m.v[2][3] * m.v[3][0]) +
        m.v[1][3] * (m.v[2][0] * m.v[3][1] - m.v[2][1] * m.v[3][0]))

    result.v[0][3] = (
        -m.v[1][0] * (m.v[2][1] * m.v[3][2] - m.v[2][2] * m.v[3][1]) +
        m.v[1][1] * (m.v[2][0] * m.v[3][2] - m.v[2][2] * m.v[3][0]) -
        m.v[1][2] * (m.v[2][0] * m.v[3][1] - m.v[2][1] * m.v[3][0]))

    result.v[1][1] = (
        m.v[0][0] * (m.v[2][2] * m.v[3][3] - m.v[2][3] * m.v[3][2]) -
        m.v[0][2] * (m.v[2][0] * m.v[3][3] - m.v[2][3] * m.v[3][0]) +
        m.v[0][3] * (m.v[2][0] * m.v[3][2] - m.v[2][2] * m.v[3][0]))

    result.v[1][2] = (
        -m.v[0][0] * (m.v[2][1] * m.v[3][3] - m.v[2][3] * m.v[3][1]) +
        m.v[0][1] * (m.v[2][0] * m.v[3][3] - m.v[2][3] * m.v[3][0]) -
        m.v[0][3] * (m.v[2][0] * m.v[3][1] - m.v[2][1] * m.v[3][0]))

    result.v[1][3] = (
        m.v[0][0] * (m.v[2][1] * m.v[3][2] - m.v[2][2] * m.v[3][1]) -
        m.v[0][1] * (m.v[2][0] * m.v[3][2] - m.v[2][2] * m.v[3][0]) +
        m.v[0][2] * (m.v[2][0] * m.v[3][1] - m.v[2][1] * m.v[3][0]))

    result.v[2][1] = (
        -m.v[0][0] * (m.v[1][2] * m.v[3][3] - m.v[1][3] * m.v[3][2]) +
        m.v[0][2] * (m.v[1][0] * m.v[3][3] - m.v[1][3] * m.v[3][0]) -
        m.v[0][3] * (m.v[1][0] * m.v[3][2] - m.v[1][2] * m.v[3][0]))

    result.v[2][2] = (
        m.v[0][0] * (m.v[1][1] * m.v[3][3] - m.v[1][3] * m.v[3][1]) -
        m.v[0][1] * (m.v[1][0] * m.v[3][3] - m.v[1][3] * m.v[3][0]) +
        m.v[0][3] * (m.v[1][0] * m.v[3][1] - m.v[1][1] * m.v[3][0]))

    result.v[2][3] = (
        -m.v[0][0] * (m.v[1][1] * m.v[3][2] - m.v[1][2] * m.v[3][1]) +
        m.v[0][1] * (m.v[1][0] * m.v[3][2] - m.v[1][2] * m.v[3][0]) -
        m.v[0][2] * (m.v[1][0] * m.v[3][1] - m.v[1][1] * m.v[3][0]))

    result.v[3][1] = (
        m.v[0][0] * (m.v[1][2] * m.v[2][3] - m.v[1][3] * m.v[2][2]) -
        m.v[0][2] * (m.v[1][0] * m.v[2][3] - m.v[1][3] * m.v[2][0]) +
        m.v[0][3] * (m.v[1][0] * m.v[2][2] - m.v[1][2] * m.v[2][0]))

    result.v[3][2] = (
        -m.v[0][0] * (m.v[1][1] * m.v[2][3] - m.v[1][3] * m.v[2][1]) +
        m.v[0][1] * (m.v[1][0] * m.v[2][3] - m.v[1][3] * m.v[2][0]) -
        m.v[0][3] * (m.v[1][0] * m.v[2][1] - m.v[1][1] * m.v[2][0]))

    result.v[3][3] = (
        m.v[0][0] * (m.v[1][1] * m.v[2][2] - m.v[1][2] * m.v[2][1]) -
        m.v[0][1] * (m.v[1][0] * m.v[2][2] - m.v[1][2] * m.v[2][0]) +
        m.v[0][2] * (m.v[1][0] * m.v[2][1] - m.v[1][1] * m.v[2][0]))

    var determinantFactor = 1.0 / determinant

    return ScaleMatrix(result, determinantFactor).Transpose()
}

func AddMatrix(a, b *Matrix) *Matrix {
    return &Matrix {
        [4][4]float64{
            {
                a.v[0][0] + b.v[0][0],
                a.v[0][1] + b.v[0][1],
                a.v[0][2] + b.v[0][2],
                a.v[0][3] + b.v[0][3]},
            {
                a.v[1][0] + b.v[1][0],
                a.v[1][1] + b.v[1][1],
                a.v[1][2] + b.v[1][2],
                a.v[1][3] + b.v[1][3]},
            {
                a.v[2][0] + b.v[2][0],
                a.v[2][1] + b.v[2][1],
                a.v[2][2] + b.v[2][2],
                a.v[2][3] + b.v[2][3]},
            {
                a.v[3][0] + b.v[3][0],
                a.v[3][1] + b.v[3][1],
                a.v[3][2] + b.v[3][2],
                a.v[3][3] + b.v[3][3]}}}
}

func MultMatrix(a, b *Matrix) *Matrix {
    return &Matrix {
        [4][4]float64{
            {
                a.v[0][0] * b.v[0][0] + a.v[0][1] * b.v[1][0] + a.v[0][2] * b.v[2][0] + a.v[0][3] * b.v[3][0],
                a.v[0][0] * b.v[0][1] + a.v[0][1] * b.v[1][1] + a.v[0][2] * b.v[2][1] + a.v[0][3] * b.v[3][1],
                a.v[0][0] * b.v[0][2] + a.v[0][1] * b.v[1][2] + a.v[0][2] * b.v[2][2] + a.v[0][3] * b.v[3][2],
                a.v[0][0] * b.v[0][3] + a.v[0][1] * b.v[1][3] + a.v[0][2] * b.v[2][3] + a.v[0][3] * b.v[3][3]},
            {
                a.v[1][0] * b.v[0][0] + a.v[1][1] * b.v[1][0] + a.v[1][2] * b.v[2][0] + a.v[1][3] * b.v[3][0],
                a.v[1][0] * b.v[0][1] + a.v[1][1] * b.v[1][1] + a.v[1][2] * b.v[2][1] + a.v[1][3] * b.v[3][1],
                a.v[1][0] * b.v[0][2] + a.v[1][1] * b.v[1][2] + a.v[1][2] * b.v[2][2] + a.v[1][3] * b.v[3][2],
                a.v[1][0] * b.v[0][3] + a.v[1][1] * b.v[1][3] + a.v[1][2] * b.v[2][3] + a.v[1][3] * b.v[3][3]},
            {
                a.v[2][0] * b.v[0][0] + a.v[2][1] * b.v[1][0] + a.v[2][2] * b.v[2][0] + a.v[2][3] * b.v[3][0],
                a.v[2][0] * b.v[0][1] + a.v[2][1] * b.v[1][1] + a.v[2][2] * b.v[2][1] + a.v[2][3] * b.v[3][1],
                a.v[2][0] * b.v[0][2] + a.v[2][1] * b.v[1][2] + a.v[2][2] * b.v[2][2] + a.v[2][3] * b.v[3][2],
                a.v[2][0] * b.v[0][3] + a.v[2][1] * b.v[1][3] + a.v[2][2] * b.v[2][3] + a.v[2][3] * b.v[3][3]},
            {
                a.v[3][0] * b.v[0][0] + a.v[3][1] * b.v[1][0] + a.v[3][2] * b.v[2][0] + a.v[3][3] * b.v[3][0],
                a.v[3][0] * b.v[0][1] + a.v[3][1] * b.v[1][1] + a.v[3][2] * b.v[2][1] + a.v[3][3] * b.v[3][1],
                a.v[3][0] * b.v[0][2] + a.v[3][1] * b.v[1][2] + a.v[3][2] * b.v[2][2] + a.v[3][3] * b.v[3][2],
                a.v[3][0] * b.v[0][3] + a.v[3][1] * b.v[1][3] + a.v[3][2] * b.v[2][3] + a.v[3][3] * b.v[3][3]}}}
}

func MultMatrixVector(m *Matrix, v *Vector) *Vector {
    return NewVector(
        m.v[0][0] * v.X() + m.v[0][1] * v.Y() + m.v[0][2] * v.Z() + m.v[0][3],
        m.v[1][0] * v.X() + m.v[1][1] * v.Y() + m.v[1][2] * v.Z() + m.v[1][3],
        m.v[2][0] * v.X() + m.v[2][1] * v.Y() + m.v[2][2] * v.Z() + m.v[2][3])
}

func MultMatrixPoint(m *Matrix, p *Point) *Point {
    return NewPoint(
        m.v[0][0] * p.X() + m.v[0][1] * p.Y() + m.v[0][2] * p.Z() + m.v[0][3],
        m.v[1][0] * p.X() + m.v[1][1] * p.Y() + m.v[1][2] * p.Z() + m.v[1][3],
        m.v[2][0] * p.X() + m.v[2][1] * p.Y() + m.v[2][2] * p.Z() + m.v[2][3])
}

func NewTranslation(x, y, z float64) *Matrix {
    return &Matrix {
        [4][4]float64{
            {1.0, 0.0, 0.0,   x},
            {0.0, 1.0, 0.0,   y},
            {0.0, 0.0, 1.0,   z},
            {0.0, 0.0, 0.0, 1.0}}}
}

func NewRotation(x, y, z float64) *Matrix {
    var cosX = math.Cos(x)
    var cosY = math.Cos(y)
    var cosZ = math.Cos(z)

    var sinX = math.Sin(x)
    var sinY = math.Sin(y)
    var sinZ = math.Sin(z)

    return &Matrix {
        [4][4]float64{
            {                      cosY * cosZ,                      -cosY * sinZ,         sinY, 0.0},
            { sinX * sinY * cosZ + cosX * sinZ, -sinX * sinY * sinZ + cosX * cosZ, -sinX * cosY, 0.0},
            {-cosX * sinY * cosZ + sinX * sinZ,  cosX * sinY * sinZ + sinX * cosZ,  cosX * cosY, 0.0},
            {                              0.0,                               0.0,          0.0, 1.0}}}
}

func NewRotationAround(axis *Vector, theta float64) *Matrix {
    var cos = math.Cos(theta)
    var sin = math.Sin(theta)
    var ncos = 1.0 - cos

    var sinX = axis.X() * sin
    var sinY = axis.Y() * sin
    var sinZ = axis.Z() * sin

    var ncosXY = axis.X() * axis.Y() * ncos
    var ncosYZ = axis.Y() * axis.Z() * ncos
    var ncosXZ = axis.X() * axis.Z() * ncos

    return &Matrix {
        [4][4]float64{
            {ncos * axis.X() * axis.X() + cos,                    ncosXY - sinZ,                    ncosXZ + sinY, 0.0},
            {                   ncosXY - sinZ, ncos * axis.Y() * axis.Y() + cos,                    ncosYZ - sinX, 0.0},
            {                   ncosXZ - sinY,                    ncosYZ + sinX, ncos * axis.Z() * axis.Z() + cos, 0.0},
            {                             0.0,                              0.0,                              0.0, 1.0}}}
}

func NewScale(x, y, z float64) *Matrix {
    return &Matrix {
        [4][4]float64{
            {  x, 0.0, 0.0, 0.0},
            {0.0,   y, 0.0, 0.0},
            {0.0, 0.0,   z, 0.0},
            {0.0, 0.0, 0.0, 1.0}}}
}

func NewShear(x, y, z float64, normal *Vector) *Matrix {
    var shear = NewVector(x, y, z)
    return AddMatrix(Identity(), Tensor(shear, normal))
}

func Identity() *Matrix {
    return &Matrix {
        [4][4]float64{
            {1.0, 0.0, 0.0, 0.0},
            {0.0, 1.0, 0.0, 0.0},
            {0.0, 0.0, 1.0, 0.0},
            {0.0, 0.0, 0.0, 1.0}}}
}
