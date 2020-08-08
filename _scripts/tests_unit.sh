#!/usr/bin/env bash
set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
CUR_DIR=$(pwd)

. ./_scripts/env.sh

cd $BASE_DIR

if [[ "$1" == "total" ]]; then

  PASS="ok"

  testCmd="go test ./ -cover -tags=unit_test $@"

  count=0
  totalPercent=0

  while read line; do

    echo "${line}"
    COVERAGE=$(echo $line | sed 's/^.*coverage: \(.*\)\% of.*/\1/g')

    case "$COVERAGE" in
      *FAIL*)
        PASS=""
        echo "Aww :( dang."
        continue
        ;;
    esac

    totalPercent=$(echo $totalPercent + $COVERAGE | bc)

    count=$((count + 1))

  done  < <($testCmd)

  TOTAL_COVERAGE=$(bc <<< "scale=1; $totalPercent/$count")

  echo ""
  printf "  === Total Coverage of $TOTAL_COVERAGE for $count package"
  if [[ "$count" != "1" ]]; then
    printf "s"
  fi
  printf " ==="
  echo ""

  export UNIT_TEST_COVERAGE=$TOTAL_COVERAGE

  cd $CUR_DIR

  echo ""
  echo "-------------------------------------------"
  echo ""
  if [[ "$PASS" == "ok" ]]; then
    echo "  Tests passed with a coverage of $UNIT_TEST_COVERAGE%"
    echo ""
    echo "-------------------------------------------"
    echo ""
  else
    echo "  Tests failed!"
    echo ""
    echo "-------------------------------------------"
    echo ""
    exit 1;
  fi


else

  go test ./ -v -cover -tags=unit_test $@

fi
