package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"net"
	"os"
)

const (
	ENV_S3_ENDPOINT = "S3_ENDPOINT"
	ENV_LISTEN      = "COSI_GRPC_LISTEN"
)
const defaultListen = "localhost:8080"

func configureHTTPListener() (listener net.Listener, err error) {
	logf.Info("initializing grpc handler")
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

func configureS3Session() (sess *session.Session, err error) {
	ep, _ := os.LookupEnv(ENV_S3_ENDPOINT)
	if ep == "" {
		sess, err = session.NewSession(aws.NewConfig().WithMaxRetries(retries))
		logf.Info(fmt.Sprintf("configuring session for default endpoint (%s)", *sess.Config.Endpoint))
	} else {
		logf.Info(fmt.Sprintf("configuring session for custom endpoint (%s)", *sess.Config.Endpoint))
		sess, err = session.NewSession(aws.NewConfig().WithMaxRetries(retries).WithEndpoint(ep))
	}
	return sess, err
}
