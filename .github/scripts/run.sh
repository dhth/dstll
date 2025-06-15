#!/usr/bin/env bash

# shellcheck disable=SC2046
dstll $(git ls-files '**/*.go' | head -n 10) -p
dstll write $(git ls-files '**/*.go' | head -n 10) -o /var/tmp/dstll
