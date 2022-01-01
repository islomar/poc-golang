#!/bin/bash
#---
# Excerpted from "Powerful Command-Line Applications in Go",
# published by The Pragmatic Bookshelf.
# Copyrights apply to this code. It may not be used to create training material,
# courses, books, articles, and the like. Contact us if you are in doubt.
# We make no guarantees that this code is fit for any purpose.
# Visit https://pragprog.com/titles/rggo for more book information.
#---

OSLIST="linux windows darwin"
ARCHLIST="amd64 arm arm64"

for os in ${OSLIST}; do
  for arch in ${ARCHLIST}; do
    if [[ "$os/$arch" =~ ^(windows/arm64|darwin/arm)$ ]]; then continue; fi

    echo Building binary for $os $arch
    mkdir -p releases/${os}/${arch}
    CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -tags=inmemory \
      -o releases/${os}/${arch}/
  done
done
