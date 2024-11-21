package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"strconv"
)

// defalut size
var size = 1000

const iterator = 254

// 从z=0 开始迭代。 c, c^2+c,(c^2+c)^2+c,.... 所有满足条件的c的集合
type Mandelbrot struct {
}

func (f Mandelbrot) DrawWithParam(writer io.Writer, param map[string]string) error {
	if v, ok := param["size"]; ok {
		s, _ := strconv.Atoi(v)
		size = s
	}

	// init rec
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	for i := 0; i < size; i++ {
		x := 4 * (float64(i)/float64(size) - 0.5)
		for j := 0; j < size; j++ {
			y := 4 * (float64(j)/float64(size) - 0.5)
			img.Set(i, j, mandelbrot(complex(x, y)))
		}
	}

	return png.Encode(writer, img)
}

// z := z^2 + c
// 特点，如果 c in M，则 |c| <= 2; 反过来不一定成立
// 如果  c in M，则 |z| <= 2. 这个特性可以用来发现 c 是否属于 M
func mandelbrot(c complex128) color.Color {
	var z complex128

	// 如果迭代 200 次发现 z 还是小于 2，则认为 c 属于 M
	for i := uint8(0); i < iterator; i++ {
		if cmplx.Abs(z) > 2 {
			return getColor(i)
		}
		z = z*z + c
	}

	return color.Black
}

// 根据迭代次数计算一个合适的像素值
func getColor(n uint8) color.Color {
    paletted := [16]color.Color{
        color.RGBA{66, 30, 15, 255},    // # brown 3
        color.RGBA{25, 7, 26, 255},     // # dark violett
        color.RGBA{9, 1, 47, 255},      //# darkest blue
        color.RGBA{4, 4, 73, 255},      //# blue 5
        color.RGBA{0, 7, 100, 255},     //# blue 4
        color.RGBA{12, 44, 138, 255},   //# blue 3
        color.RGBA{24, 82, 177, 255},   //# blue 2
        color.RGBA{57, 125, 209, 255},  //# blue 1
        color.RGBA{134, 181, 229, 255}, // # blue 0
        color.RGBA{211, 236, 248, 255}, // # lightest blue
        color.RGBA{241, 233, 191, 255}, // # lightest yellow
        color.RGBA{248, 201, 95, 255},  // # light yellow
        color.RGBA{255, 170, 0, 255},   // # dirty yellow
        color.RGBA{204, 128, 0, 255},   // # brown 0
        color.RGBA{153, 87, 0, 255},    // # brown 1
        color.RGBA{106, 52, 3, 255},    // # brown 2
    }
    return paletted[n%16]
}

func (f Mandelbrot) Name() string {
	return "mandelbrot"
}
