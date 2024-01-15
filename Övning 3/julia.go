// Stefan Nilsson 2013-02-27

// This program creates pictures of Julia sets (en.wikipedia.org/wiki/Julia_set).
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"strconv"
	"fmt"
	"time"
	"sync"
	"runtime"

)

type ComplexFunc func(complex128) complex128

var Funcs []ComplexFunc = []ComplexFunc{
	func(z complex128) complex128 { return z*z - 0.61803398875 },
	func(z complex128) complex128 { return z*z + complex(0, 1) },
	func(z complex128) complex128 { return z*z + complex(-0.835, -0.2321) },
	func(z complex128) complex128 { return z*z + complex(0.45, 0.1428) },
	func(z complex128) complex128 { return z*z*z + 0.400 },
	func(z complex128) complex128 { return cmplx.Exp(z*z*z) - 0.621 },
	func(z complex128) complex128 { return (z*z+z)/cmplx.Log(z) + complex(0.268, 0.060) },
	func(z complex128) complex128 { return cmplx.Sqrt(cmplx.Sinh(z*z)) + complex(0.065, 0.122) },
}

func main() {
    //Added startTime to time function
    startTime := time.Now()
	for n, fn := range Funcs {
		err := CreatePng("picture-"+strconv.Itoa(n)+".png", fn, 1024)
		if err != nil {
			log.Fatal(err)
		}
	}
	time := time.Since(startTime)
	fmt.Printf("Time: ")
	fmt.Println(time)
}

// CreatePng creates a PNG picture file with a Julia image of size n x n.
func CreatePng(filename string, f ComplexFunc, n int) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()
	err = png.Encode(file, Julia(f, n))
	return
}

// Julia returns an image of size n x n of the Julia set for f.
func Julia(f ComplexFunc, n int) image.Image {
	bounds := image.Rect(-n/2, -n/2, n/2, n/2)
	img := image.NewRGBA(bounds)
	s := float64(n / 4)

	wg := new(sync.WaitGroup) //Adding waitgroup for use in go routine

	for i := bounds.Min.X; i < bounds.Max.X; i++ {
	    wg.Add(1)
	    go func(i int){
            for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
                n := Iterate(f, complex(float64(i)/s, float64(j)/s), 256)
                r := uint8(0)
                g := uint8(0)
                b := uint8(n % 32 * 8)
                img.Set(i, j, color.RGBA{r, g, b, 255})
            }
		wg.Done() //continue when all go routines are done
	    }(i)
	}
	wg.Wait() // return img when all go routines are finished
	return img
}

// Iterate sets z_0 = z, and repeatedly computes z_n = f(z_{n-1}), n â‰¥ 1,
// until |z_n| > 2  or n = max and returns this n.
func Iterate(f ComplexFunc, z complex128, max int) (n int) {
	for ; n < max; n++ {
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			break
		}
		z = f(z)
	}
	return
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Try to use all available CPUs.
	fmt.Println("Number of cpu's: " + strconv.Itoa(runtime.NumCPU()))
}


/*
How many CPUs does you program use?
    - 8 CPUs
*/

/*
How much faster is your parallel version?
    - Initial version time: 12.454643903s
    - Modified version time: 2.707423849s
*/