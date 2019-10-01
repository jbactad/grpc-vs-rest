#! /usr/bin/env bash

/bin/bash -c "cd restserver && go build && ./restserver";
/bin/bash -c "cd restservertls && go build && ./restservertls";
/bin/bash -c "cd grpcserver && go build && ./grpcserver";
/bin/bash -c "cd grpcservertls && go build && ./grpcservertls";
