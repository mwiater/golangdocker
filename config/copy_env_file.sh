#! /bin/bash

# Need to include .env file in order to use it within this test directory.
# FS.embed does not allow including files in parent directories. Temporary workaround.
cp ../.env .env