#!/bin/bash

curl -v -H "Content-Type: application/yaml" --data-binary "@$(dirname "$0")/scrape.yml" localhost:10888/api/prometheus
echo
