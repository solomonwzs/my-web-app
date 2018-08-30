#!/bin/sh
set -euo pipefail

SCRIPT=$(readlink -f "$0")
DIR=$(dirname "$SCRIPT")

BUILD_DIR="$DIR/build"
FRONTEND_DIR="$DIR/frontend"
BACKEND_DIR="$DIR/backend"

[ -d "$BUILD_DIR" ] && rm -rf "$BUILD_DIR"
mkdir -p "$BUILD_DIR"

cd "$FRONTEND_DIR"
npm run build

cd "$BACKEND_DIR"
make

cd "$DIR"
cp -r "${FRONTEND_DIR}/dist" "$BUILD_DIR"
cp -r "${BACKEND_DIR}/bin/main.goc" "$BUILD_DIR"
cp "${DIR}/start.sh" "$BUILD_DIR"
