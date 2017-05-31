package remote

import (
	"fmt"
	"math"
)

// TCPArgs is structured around the client's provided parameters
// The struct's fields need to be exported too!
type PiCalcParams struct {
	NumIterations int
}

// Compose is our RPC functions return type
type Pi float64

// Details is our exposed RPC function
func (p *Pi) CalcPi(args PiCalcParams, reply *float64) error {
	fmt.Println("Recieved pi computation request")
	var f float64 = 0.0
	for k := 0; k <= args.NumIterations; k++ {
		f += 4.0 * math.Pow(-1, float64(k)) / float64(2*k+1)
	}
	*reply = f
	return nil
}
