package main

import (
	"github.com/yard-turkey/cosi-prototype-interface/cosi"
	"google.golang.org/grpc"
	"k8s.io/klog/klogr"
)

var logf = klogr.New().WithName("cosi-s3-plugin")

func main() {
	sess, err := configureS3Session()
	if err != nil {
		panic(err)
	}

	rpcServer := grpc.NewServer()
	cosi.RegisterProvisionerServer(rpcServer, newHandler(sess))

	listener, err := configureHTTPListener()
	if err != nil {
		panic(err)
	}
	if err = rpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
