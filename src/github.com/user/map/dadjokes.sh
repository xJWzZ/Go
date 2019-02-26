#!/bin/bash

go build map.go

i=$(./map $@)
echo $i
exit 1
