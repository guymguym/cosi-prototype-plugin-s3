#! /usr/bin/env bash

set -e

ROOT=$(readlink -f $(dirname ${BASH_SOURCE[0]})"/..") 
OUTDIR="$ROOT/_output"
OUTBIN="$OUTDIR/plugin"
BUILDDIR="$OUTDIR/build"

TAG="jcoperh/plugin"

(
	cd $ROOT
	[ ! -d "${OUTDIR}" ] && mkdir -p $OUTDIR
	[ ! -d "${BUILDDIR}" ] && mkdir -p $BUILDDIR 
	rm -rf "${BUILDDIR}/*"
	GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags="static"' -o $OUTBIN $ROOT/*.go
	echo "go build complete"	
	cp $OUTBIN $BUILDDIR
	cp $ROOT/Dockerfile $BUILDDIR
	docker build -t $TAG "${BUILDDIR}/"
)
