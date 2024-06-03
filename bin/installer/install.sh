#!/bin/bash

source src/*.sh

OS=$(uname -s)
VERSION=$1

if [[ -z $VERSION ]]; then
    echo "No version specified. Please specify the version to install."
    echo "Available versions can be tag or branch name. Example: v1.0.0 or master"
    exit 1
fi

if [[ $OS == "Darwin" ]]; then
    install_on_macos
    install_goblitz
fi