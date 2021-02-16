package gwl

type Point struct {
    X int
    Y int
}

func P(x, y int) Point {
    return Point{ x, y }
}

func P16(x, y int16) Point {
    return Point{ int(x), int(y) }
}

type Bounds struct {
    X int
    Y int
    Width int
    Height int
}

func B(x, y, w, h int) Bounds {
    return Bounds{ x, y, w, h }
}

func B16(x, y int16, w, h uint16) Bounds {
    return Bounds{ int(x), int(y), int(w), int(h) }
}
