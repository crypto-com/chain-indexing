#!/usr/bin/env bash

if [[ ! -f ./check-password-strength ]]; then
    make check-password-strength
fi
./check-password-strength
