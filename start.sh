#!/usr/bin/env sh
export MICRO_REGISTRY=etcd
export MICRO_REGISTRY_ADDRESS=localhost:2379
export MICRO_CLIENT=grpc
export MICRO_API_NAMESPACE=codetube.cn.api
micro api