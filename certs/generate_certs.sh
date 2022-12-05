#!/bin/bash
#
# From root of project, run: `bash certs/generate_certs.sh`

clear

openssl req -x509 -newkey rsa:4096 -sha256 -days 365 -nodes \
  -keyout ./certs/192.168.0.99.key -out ./certs/192.168.0.99.crt -subj "/CN=192.168.0.99" \
  -addext "subjectAltName=DNS:192.168.0.99,IP:192.168.0.99"