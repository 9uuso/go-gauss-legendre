package main

import (
	"fmt"
	"math"
)

func a(an, bn []float64) []float64 {
	a := (an[len(an)-1] + bn[len(bn)-1]) / float64(2)
	an = append(an, a)
	return an
}

func b(an, bn []float64) []float64 {
	b := math.Sqrt(an[len(an)-2] * bn[len(bn)-1])
	bn = append(bn, b)
	return bn
}

func t(an, bn, tn, pn []float64) []float64 {
	t := tn[len(tn)-1] - pn[len(pn)-1]*math.Pow((an[len(an)-2]-an[len(an)-1]), 2)
	tn = append(tn, t)
	return tn
}

func p(pn []float64) []float64 {
	p := 2 * pn[len(pn)-1]
	pn = append(pn, p)
	return pn
}

func pi(an, bn, tn, pin []float64) []float64 {
	pi := math.Pow((an[len(an)-1]+bn[len(bn)-1]), 2) / (4 * tn[len(tn)-1])
	pin = append(pin, pi)
	return pin
}

func main() {
	// initialize lists
	an := []float64{1}
	bn := []float64{}
	tn := []float64{}
	pn := []float64{1}
	bn = append(bn, float64(1)/math.Sqrt(2))
	tn = append(tn, float64(1)/float64(4))
	pin := []float64{}

	// run the algorithm ten times
	for i := 0; i < 10; i++ {
		an = a(an, bn)
		bn = b(an, bn)
		tn = t(an, bn, tn, pn)
		pn = p(pn)
		pin = pi(an, bn, tn, pin)
	}

	// print results
	for _, value := range pin {
		fmt.Printf("%.100f\n", value)
	}

}
