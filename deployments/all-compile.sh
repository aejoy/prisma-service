#!/bin/bash

platforms=(
  "linux/amd64"
  "linux/386"
  "linux/arm"
  "linux/arm64"
  "darwin/amd64"
  "darwin/arm64"
  "windows/amd64"
  "windows/386"
  "freebsd/amd64"
  "freebsd/386"
)

for platform in "${platforms[@]}"
do
  GOOS=$(echo "$platform" | cut -d/ -f1)
  GOARCH=$(echo "$platform" | cut -d/ -f2)

  out=prisma-"$GOOS-$GOARCH"
  if [ "$GOOS" == "windows" ]; then
    out="$out.exe"
  fi

  echo "Building for $GOOS/$GOARCH"

  go build -ldflags "-s -w" -o "./release/$out" ../cmd/api/main.go
  sha256sum "./release/$out" > "release/$out.sha256"

  echo "Build completed for $GOOS/$GOARCH"
done

echo "All builds are done!"