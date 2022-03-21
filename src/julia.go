package main

import (
  "image"
  "image/color"
  "image/png"
  "log"
  "math/cmplx"
  "os"
  "sync"
  "strconv"
  "fmt"
  "time"
)

type ComplexFunc func(complex128) complex128

var Funcs []ComplexFunc = []ComplexFunc{
  func(z complex128) complex128 { return z*z - 0.61803398875 },
  func(z complex128) complex128 { return z*z + complex(0, 1) },
  func(z complex128) complex128 { return z*z + complex(-0.835, -0.2321) },
  func(z complex128) complex128 { return z*z + complex(0.45, 0.1428) },
  func(z complex128) complex128 { return z*z*z + 0.400 },
  func(z complex128) complex128 { return cmplx.Exp(z*z*z) - 0.621 },
  func(z complex128) complex128 { return (z*z*z)/cmplx.Log(z) + complex(0.268, 0.060) },
  func(z complex128) complex128 { return cmplx.Sqrt(cmplx.Sinh(z*z)) + complex(0.065, 0.122) },
}

func main() {
  timer := time.Now()
  for n, fn := range Funcs {
    err := CreatePng("picture-"+strconv.Itoa(n)+".png", fn, 1024)
    if err != nil {
      log.Fatal(err)
    }
  }
  fmt.Println("Program ran in ", time.Since(timer))
}

func CreatePng(filename string, f ComplexFunc, n int) (err error) {
  file, err := os.Create(filename)
  if err != nil {
    return
  }
  defer file.Close()
  err = png.Encode(file, Julia(f, n))
  return
}

func Julia(f ComplexFunc, n int) image.Image {
  bounds := image.Rect(-n/2, -n/2, n/2, n/2)
  img := image.NewRGBA(bounds)
  s := float64(n/4)

  var wg sync.WaitGroup
  for iterX := 0; iterX < 8; iterX++ {
    for iterY := 0; iterY < 8; iterY++ {
      wg.Add(1)
      go func(x, y int) {
        for i := bounds.Min.X + (128 * x); i < bounds.Min.X + (128 * (x + 1)); i++ {
          for j := bounds.Min.Y + (128 * y); j < bounds.Min.Y + (128 * (y + 1)); j++ {
            n := Iterate(f, complex(float64(i)/s, float64(j)/s), 256)
            r := uint8(0)
            g := uint8(0)
            b := uint8(n % 32 * 8)
            img.Set(i, j, color.RGBA{r, g, b, 255})
          }
        }
        wg.Done()
      }(iterX, iterY)
    }
  }
  wg.Wait()
  return img
}

func Iterate(f ComplexFunc, z complex128, max int) (n int) {
  for ; n < max; n++ {
    if real(z)*real(z)+imag(z)*imag(z) > 4 {
      break
    }
    z = f(z)
  }
  return
}
