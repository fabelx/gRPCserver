package divider

import (
	"fmt"
	api "gRPCserver/pkg/api"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) Divide(ctx context.Context, req *api.DivideRequest) (*api.DivideResponse, error) {
	var errorMessage string

	result := func(x, y int32, em *string) int32 {
		defer func(em *string) {
			if r := recover(); r != nil {
				*em = fmt.Sprintf("%v", r)
				fmt.Println("Recovered in f", r)
			}
		}(em)
		r := x / y
		return r
	}(req.GetX(), req.GetY(), &errorMessage)

	return &api.DivideResponse{Result: result, Error: errorMessage}, nil
}
