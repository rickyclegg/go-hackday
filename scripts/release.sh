#!/usr/bin/env bash

set -e

echo "Releasing go-hackday"
docker tag go-hackday:latest 911610643236.dkr.ecr.eu-west-1.amazonaws.com/gohackday:latest
docker push 911610643236.dkr.ecr.eu-west-1.amazonaws.com/gohackday:latest
