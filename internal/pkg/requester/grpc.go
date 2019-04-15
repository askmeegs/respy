package requester

import (
	"fmt"
)

type GrpcRequester struct {
	url string
}

func NewGrpcRequester(u string, proto string) *GrpcRequester {
	// import local proto file

	// run protoc command to generate go client --> put in the same folder as the source .proto
	// protoc -I routeguide/ routeguide/route_guide.proto --go_out=plugins=grpc:routeguide

	return &GrpcRequester{
		url: u,
	}
}

func (g *GrpcRequester) OneRequest() (string, error) {
	return "", nil

	// take JSON input

	// call endpoint of the format <ip>/<port>/<name>.<subname>/<function>
	// eg 127.0.0.1:3550/example.SearchProducts
	/*
		with request body

		{
			"query": "camera"
		}

	*/

}

// TODO
// remove the generated go code
// func (g *GrpcRequester) Cleanup() error {
// 	return nil
// }
