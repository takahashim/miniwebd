#!/bin/sh

VERSION=`git describe --always`

ghr ${VERSION} ./pkg/dist/${VERSION}
