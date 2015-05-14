#!/bin/bash
#
# This is the InfluxDB-Viz CircleCI test script. Using this script allows total control
# the environment in which the build and test is run, and matches the official
# build process for InfluxDB-Viz.

BUILD_DIR=$HOME/influxdb-viz-build
GO_VERSION=go1.4.2
TIMEOUT="-timeout 300s"

# Executes the given statement, and exits if the command returns a non-zero code.
function exit_if_fail {
    command=$@
    echo "Executing '$command'"
    $command
    rc=$?
    if [ $rc -ne 0 ]; then
        echo "'$command' returned $rc."
        exit $rc
    fi
}

# 'go get' dependencies, will retry requested number of times.
function go_get {
    n=1
    retry_count=$1

    while [ $n -ne $retry_count ]; do
        go get -t -d -v ./...
        rc=$?
        echo "'go get -t -d -v' returned $rc on attempt #$n"
        if [ $rc -eq 0 ]; then
            break
        fi
        n=$((n + 1))
    done
    return $rc
}

source $HOME/.gvm/scripts/gvm
exit_if_fail gvm use $GO_VERSION

# Set up the build directory, and then GOPATH.
exit_if_fail mkdir $BUILD_DIR
export GOPATH=$BUILD_DIR
exit_if_fail mkdir -p $GOPATH/src/github.com/influxdb/influxdb-viz

exit 0
