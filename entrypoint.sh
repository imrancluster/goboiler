#!/bin/sh

curl --request PUT --data-binary @config.yaml http://goboiler_consul:8500/v1/kv/goboiler

exec ./goboiler serve
