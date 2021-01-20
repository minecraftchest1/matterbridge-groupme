#!/bin/bash

echo "Calling ./getGroups with args -t $1"
./getGroups -t $1 | jq '.response[]' | grep -A 2 'group_id'
