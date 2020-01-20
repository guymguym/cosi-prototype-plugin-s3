package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/yard-turkey/cosi-prototype-interface/cosi"
)

// s3Handler is a singleton that implements the COSI interface to route grpc calls to s3 operations
type handler struct {
	s3 *s3.S3
}

func newHandler(sess *session.Session) *handler {
	return &handler{s3: s3.New(sess)}
}

func (s handler) Provision(context.Context, *cosi.ProvisionRequest) (*cosi.ProvisionResponse, error) {
	return &cosi.ProvisionResponse{
		Credentials:  nil,
		Metadata:     nil,
		Endpoint:     "",
		Failed:       false,
		ErrorMessage: "",
	}, nil
}

func (s handler) Deprovision(context.Context, *cosi.DeprovisionRequest) (*cosi.DeprovisionResponse, error) {
	panic("implement me")
}
