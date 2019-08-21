package msa

import (
	"context"
	"log"

	somefunction "github.com/seriwb/front-bff-msa/msa/gen/some_function"
)

// some-function service example implementation.
// The example methods log the requests and return zero values.
type someFunctionsrvc struct {
	logger *log.Logger
}

// NewSomeFunction returns the some-function service implementation.
func NewSomeFunction(logger *log.Logger) somefunction.Service {
	return &someFunctionsrvc{logger}
}

// Add implements add.
func (s *someFunctionsrvc) Add(ctx context.Context, p *somefunction.AddPayload) (res int, err error) {
	s.logger.Print("someFunction.add")
	return p.A + p.B, nil
}
