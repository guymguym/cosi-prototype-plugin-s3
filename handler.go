package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/glog"
	"github.com/yard-turkey/cosi-prototype-interface/cosi"
)

// s3Handler is a singleton that implements the COSI interface to route grpc calls to s3 operations
type handler struct {
	s3 *s3.S3
}

var _ cosi.ProvisionerServer = &handler{}

func newHandler() *handler {
	sess, err := configureS3Endpoint()
	if err != nil {
		panic(err)
	}
	return &handler{s3: s3.New(sess)}
}

const pluginName = "s3-prototype-plugin"

func (s handler) GetPluginName(context.Context, *cosi.PluginNameRequest) (*cosi.PluginNameResponse, error) {
	return &cosi.PluginNameResponse{Name: pluginName}, nil
}

func (s handler) Provision(_ context.Context, req *cosi.ProvisionRequest) (*cosi.ProvisionResponse, error) {
	glog.Infof("provision request, bucket: %v", req.GetRequestBucketName())
	out, err := s.s3.CreateBucket(&s3.CreateBucketInput{
		Bucket: &req.RequestBucketName,
	})
	if err != nil {
		glog.Error("error creating bucket: " + err.Error())
	}
	val, err := s.s3.Client.Config.Credentials.Get()
	if err != nil {
		return nil, err
	}
	glog.Infof("create bucket success, bucket: %v", req.GetRequestBucketName())
	return s.provisionResponse(out, val), err
}

func (s handler) provisionResponse(out *s3.CreateBucketOutput, cred credentials.Value) *cosi.ProvisionResponse {
	return &cosi.ProvisionResponse{
		BucketName: *out.Location,
		Endpoint:   s.s3.Endpoint,
		Region:     *s.s3.Config.Region,
		EnvironmentCredentials: map[string]string{
			ENV_ACCESS_KEY_ID: cred.AccessKeyID,
			ENV_SECRET_KEY:    cred.SecretAccessKey,
		},
	}
}

func (s handler) Deprovision(_ context.Context, req *cosi.DeprovisionRequest) (*cosi.DeprovisionResponse, error) {
	glog.Infof("deprovision request, bucket: %v", req.GetBucketName())
	out, err := s.s3.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: &req.BucketName,
	})
	if err != nil {
		glog.Error("error deleting bucket: " + err.Error())
	}
	glog.Infof("delete bucket success, bucket %v:", req.GetBucketName())
	return s.deprovisionResponse(out), err
}

func (s handler) deprovisionResponse(out *s3.DeleteBucketOutput) *cosi.DeprovisionResponse {
	var msg string
	if out != nil {
		msg = out.String()
	} else {
		msg = ""
	}
	return &cosi.DeprovisionResponse{Message: msg}
}
