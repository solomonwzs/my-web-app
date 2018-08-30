#!/bin/sh
set -euo pipefail

SCRIPT=$(readlink -f "$0")
DIR=$(dirname "$SCRIPT")

FRONTEND_DIST="${DIR}/dist" \
    BACKEND_ADDR=":8089" \
    "${DIR}/main.goc"
