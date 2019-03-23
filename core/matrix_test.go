package core

import "testing"

func TestMatrixInvertIdentity(t *testing.T) {
    var expected = identity()
    var actual = expected.invert()

    if actual == nil {
        t.Error(expected, "inverted is nil")
    }

    actual = actual.invert()

    if *expected != *actual {
        t.Error(expected, " != ", actual)
    }
}

func TestMatrixInvertTranslation(t *testing.T) {
    var expected = newTranslation(1, 2, 3)
    var actual = expected.invert()

    if actual == nil {
        t.Error(expected, "inverted is nil")
    }

    actual = actual.invert()

    if *expected != *actual {
        t.Error(expected, " != ", actual)
    }
}

func TestMatrixInvertScale(t *testing.T) {
    var expected = newScale(1, 2, 3)
    var actual = expected.invert()

    if actual == nil {
        t.Error(expected, "inverted is nil")
    }

    actual = actual.invert()

    if *expected != *actual {
        t.Error(expected, " != ", actual)
    }
}

func TestMatrixInvertRotation(t *testing.T) {
    var expected = newRotation(90, 90, 90)
    var actual = expected.invert()

    if actual == nil {
        t.Error(expected, "inverted is nil")
    }

    actual = actual.invert()

    if *expected != *actual {
        t.Error(expected, " != ", actual)
    }
}
