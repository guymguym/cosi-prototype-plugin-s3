package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/golang/glog"
	"net"
	"os"
)

const (
	ENV_S3_ENDPOINT = "S3_ENDPOINT"
	ENV_LISTEN      = "COSI_GRPC_LISTEN"
	ENV_REGION = "AWS_REGION"
	ENV_ACCESS_KEY_ID = "AWS_ACCESS_KEY_ID"
	ENV_SECRET_KEY = "AWS_SECRET_ACCESS_KEY"
)
const defaultListen = "localhost:8080"
const defaultRegion = "us-east-1"

func configureHTTPListener() (listener net.Listener, err error) {
	glog.Info("initializing grpc handler")
	listen := defaultListen
	if l, ok := os.LookupEnv(ENV_LISTEN); ok {
		listen = l
	}

	listener, err = net.Listen("tcp", listen)
	if err != nil {
		panic(err)
	}
	return listener, err
}

const retries = 3

func configureS3Endpoint() (sess *session.Session, err error) {
	ep, _ := os.LookupEnv(ENV_S3_ENDPOINT)
	if ep == "" {
		// If endpoint undefined, assume AWS endpoint
		sess, err = session.NewSession(
			aws.NewConfig().
				WithMaxRetries(retries).
				WithRegion(defaultRegion))
		if err != nil {
			return nil, err
		}
		glog.Info(fmt.Sprintf("configuring session for default AWS endpoint "))
	} else {
		sess, err = session.NewSession(aws.NewConfig().WithMaxRetries(retries).WithEndpoint(ep))
		if err != nil {
			return nil, err
		}
		glog.Info(fmt.Sprintf("configuring session for custom endpoint (%v)", *sess.Config.Endpoint))
	}
	return sess, err
}
