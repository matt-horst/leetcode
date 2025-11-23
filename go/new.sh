#!/bin/bash

NAME=$1

mkdir -p $NAME
cd $NAME
echo "package ${NAME}" > "${NAME}.go"
echo "package ${NAME}" > "${NAME}_test.go"
cat ../test_template >> "${NAME}_test.go"
