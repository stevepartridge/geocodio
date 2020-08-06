#!/usr/bin/env bash
set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

CUR_DIR=`pwd`

if [ -f "$BASE_DIR/_scripts/local.env" ]; then
  set -o allexport
  source "${BASE_DIR}/_scripts/local.env"
  set +o allexport
  # printenv
  echo "ENV"
else
  echo
  echo "!! Didn't find ${BASE_DIR}/_scripts/local.env !!"
  echo
  echo "  Create one by running the following:"
  echo
  echo "  cp ${BASE_DIR}/_scripts/template.env ${BASE_DIR}/_scripts/local.env"
  echo
  echo "  Update the values and run your command again."
  echo
  exit 1
fi


cd $CUR_DIR
