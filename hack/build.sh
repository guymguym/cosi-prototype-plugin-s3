#! /usr/bin/env bash

set -e

ROOT=$(readlink -f $(dirname ${BASH_SOURCE[0]})"/..") 
OUTDIR="$ROOT/_output"
OUTBIN="$OUTDIR/plugin"
BUILDDIR="$OUTDIR/build"
[ ! -d "${OUTDIR}" ] && mkdir -p $OUTDIR

TAG="jcoperh/plugin"

(
	cd $ROOT
	[ ! -d "${BUILDDIR}" ] && mkdir -p $BUILDDIR 
	rm -rf "${BUILDDIR}/*"
	GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags="static"' -o $OUTBIN $ROOT/*.go
	echo "go build complete"	
	cp $OUTBIN $BUILDDIR
	cp $ROOT/Dockerfile $BUILDDIR  
	eval `minikube docker-env`
	docker build -t $TAG "${BUILDDIR}/"
)
