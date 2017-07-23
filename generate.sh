#!/bin/bash
protoc -I service/ service/service.proto --go_out=plugins=grpc:service
grpc_tools_ruby_protoc -I service/ --ruby_out=ruby/lib --grpc_out=ruby/lib service/service.proto