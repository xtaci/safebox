#!/bin/bash

BUILD_DIR=$(dirname "$0")/build
mkdir -p $BUILD_DIR
cd $BUILD_DIR

sum="sha1sum"

export GO111MODULE=on
echo "BUILDING SAFEBOX" 
echo "Setting GO111MODULE to" $GO111MODULE

if ! hash sha1sum 2>/dev/null; then
	if ! hash shasum 2>/dev/null; then
		echo "I can't see 'sha1sum' or 'shasum'"
		echo "Please install one of them!"
		exit
	fi
	sum="shasum"
fi

UPX=false
if hash upx 2>/dev/null; then
	UPX=true
fi

VERSION="1.0.4"
LDFLAGS="-X main.VERSION=$VERSION -s -w"
GCFLAGS=""

# AMD64 
OSES=(linux darwin windows freebsd openbsd)
for os in ${OSES[@]}; do
	suffix=""
	if [ "$os" == "windows" ]
	then
		suffix=".exe"
	fi
	env CGO_ENABLED=0 GOOS=$os GOARCH=amd64 go build -mod=vendor -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o safebox_${os}_amd64${suffix} github.com/xtaci/safebox
	if $UPX; then upx -9 safebox_${os}_amd64${suffix};fi
	tar -zcf safebox-${os}-amd64-$VERSION.tar.gz safebox_${os}_amd64${suffix}
	$sum safebox-${os}-amd64-$VERSION.tar.gz
done

#Apple M1 device
env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -mod=vendor -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o safebox_darwin_arm64$v github.com/xtaci/safebox
tar -zcf safebox-darwin-arm64-$VERSION.tar.gz safebox_darwin_arm64
$sum safebox-darwin-arm64-$VERSION.tar.gz

# ARM64
env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -mod=vendor -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o safebox_linux_arm64  github.com/xtaci/safebox
if $UPX; then upx -9 safebox_linux_arm64 ;fi
tar -zcf safebox-linux-arm64-$VERSION.tar.gz safebox_linux_arm64
$sum safebox-linux-arm64-$VERSION.tar.gz
