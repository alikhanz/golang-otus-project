#!/bin/bash

#@ToDo: use APP_ROOT env variable, after app dockerize.

protoc -I/usr/local/include \
		-I../ \
		-I../api/protobuf/ \
		-I${GOPATH}/src \
		--go_out=plugins=grpc,paths=source_relative:../pkg/pb rotator.proto