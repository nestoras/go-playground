#!/bin/bash

protoc newsfeedpb/newsfeed.proto --go_out=plugins=grpc:.
