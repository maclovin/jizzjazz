#!/bin/bash

BASE_PATH="github.com/maclovin/jizzjiazz/pkg"
PKG_DIR="../pkg"

for dir in "$PKG_DIR"/*/; do
  dir=${dir%/}
  package_name=$(basename "$dir")
  
  (
    cd "$dir" || exit
    go mod init "$BASE_PATH/$package_name"
  )
done