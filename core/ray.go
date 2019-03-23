package core

type Ray struct {
    Origin Point
    Direction Vector
}

func (ray *Ray) Point(factor float64) *Point {
    return AddPoint(&ray.Origin, ScaleVector(&ray.Direction, factor))
}

func (ray *Ray) Transform(matrix *Matrix) *Ray {
    var newOrigin = MultMatrixPoint(matrix, &ray.Origin)
    var newDirection = Normalize(
        SubPoint(
            MultMatrixPoint(matrix, AddPoint(&ray.Origin, &ray.Direction)),
            newOrigin))

    return &Ray {*newOrigin, *newDirection}
}
