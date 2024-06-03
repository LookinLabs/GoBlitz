#!/bin/bash

for i in $(pwd)/bin/installer/src/*.sh; do
    source $i
done

OS=$(uname -s)
OPTION=$1
VERSION=$2

if [[ -z $VERSION ]]; then
    echo "No version specified. Please specify the version to install or update."
    echo "Available versions can be tag or branch name. Example: v1.0.0 or master"
    exit 1
fi

case $OPTION in
    --install)
        if [[ $OS == "Darwin" ]]; then
            install_macos_utilities $VERSION
            install_goblitz $VERSION
        fi
        ;;
    --update)
        if [[ $OS == "Darwin" ]]; then
            update_goblitz $VERSION
        fi
        ;;
    *)
        echo "Invalid option. Please use --install or --update."
        exit 1
        ;;
esac