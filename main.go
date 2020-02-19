package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/yard-turkey/cosi-prototype-interface/cosi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
)

const (
	ENV_CERT_PATH = "COSI_CERT_PATH"
	ENV_KEY_PATH  = "COSI_KEY_PATH"
)

var (
	certPath string
	keyPath  string
)

func init() {
	flag.Parse()
	var ok bool
	certPath, ok = os.LookupEnv(ENV_CERT_PATH)
	if ! ok {
		glog.Fatalf("env var %s not defined, required", ENV_CERT_PATH)
	}
	keyPath, ok = os.LookupEnv(ENV_KEY_PATH)
	if ! ok {
		glog.Fatal("env var %s not defined, required", ENV_KEY_PATH)
	}
}

func main() {
	glog.Info("starting s3 plugin services")

	glog.Infof("getting credentials from paths\n\tcert: %v\n\tkey: %v", certPath, keyPath)
	creds, err := credentials.NewServerTLSFromFile(certPath, keyPath)
	if err != nil {
		glog.Fatalf("couldn't get credentials: %v", err)
	}
	rpcServer := grpc.NewServer(grpc.Creds(creds))

	glog.Info("configuring provisioner server")
	cosi.RegisterProvisionerServer(rpcServer, newHandler())
	listener, err := configureHTTPListener()
	if err != nil {
		panic(err)
	}
	glog.Info("starting provisioner server, waiting for requests")
	if err = rpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
