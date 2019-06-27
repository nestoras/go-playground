#!/bin/bash

protoc examplepb/example.proto --go_out=plugins=grpc:.
