#!/usr/bin/env bash
set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
CUR_DIR=$(pwd)

. ./_scripts/env.sh

cd $BASE_DIR
   

if [ ! -z ${1+x} ]; then
  if [ -d "$BASE_DIR/$1" ]; then
    cd $BASE_DIR/$1
  fi
fi

go test -cover -coverprofile cover.out -tags=unit_test && go tool cover -html=cover.out

cd $CUR_DIR