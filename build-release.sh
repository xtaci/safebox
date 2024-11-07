#!bash

BUILD_DIR=$(dirname "$0")/build
mkdir -p $BUILD_DIR
cd $BUILD_DIR

sum="sha1sum"

COMPRESS="gzip"
if hash pigz 2>/dev/null; then
    COMPRESS="pigz"
fi

export GO111MODULE=on
VERSION=`git describe --tags --abbrev=0`
echo "BUILDING SAFEBOX $VERSION" 
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
    tar -cf safebox-${os}-amd64-$VERSION.tar safebox_${os}_amd64${suffix}
    ${COMPRESS} -f safebox-${os}-amd64-$VERSION.tar
    $sum safebox-${os}-amd64-$VERSION.tar.gz
done

#Apple M1 device
env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -mod=vendor -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o safebox_darwin_arm64$v github.com/xtaci/safebox
tar -cf safebox-darwin-arm64-$VERSION.tar safebox_darwin_arm64
${COMPRESS} -f safebox-darwin-arm64-$VERSION.tar
$sum safebox-darwin-arm64-$VERSION.tar.gz

# ARM64
env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -mod=vendor -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o safebox_linux_arm64  github.com/xtaci/safebox
if $UPX; then upx -9 safebox_linux_arm64 ;fi
tar -cf safebox-linux-arm64-$VERSION.tar safebox_linux_arm64
${COMPRESS} -f safebox-linux-arm64-$VERSION.tar
$sum safebox-linux-arm64-$VERSION.tar.gz
