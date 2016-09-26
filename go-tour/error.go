package main

import "fmt"

// an error when Sqrt on negative float64
type ErrNegativeSqrt float64

func (val ErrNegativeSqrt) String() string {
	return fmt.Sprintf("Float64(%v)", float64(val))
}

func (val ErrNegativeSqrt) Error() string {
	// return a string rep of error for the float64
	// return fmt.Sprintf("cannot Sqrt negative value: %v", val) // this will infinitely loop, why?
	return fmt.Sprintf("cannot Sqrt negative value: %v", float64(val))
}

func mysqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 25; i++ {
		z -= ((z*z - x) / 2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(mysqrt(-2))
	fmt.Println(mysqrt(2))
	_, err := mysqrt(-3)
	if err != nil {
		fmt.Printf("%T(%q)\n", err, err)
	}
}
