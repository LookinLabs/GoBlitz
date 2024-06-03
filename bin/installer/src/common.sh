#!/bin/sh

function install_goblitz() {
    if ! [ -d "$(pwd)/GoBlitz" ]; then
        echo "GoBlitz directory does not exist. Pulling GoBlitz..."
        git clone git@github.com:LookinLabs/GoBlitz.git --depth 1 --branch $VERSION
    fi
}

function update_goblitz() {
    if `pwd | grep -q "GoBlitz"` && [ "$VERSION" == "master" ]; then
        echo "GoBlitz directory found. Pulling updates..."
        git pull origin master
    fi
    
    if `pwd | grep -q "GoBlitz"` && [ "$VERSION" != "master" ]; then
        echo "GoBlitz directory found. Checking out $VERSION..."
        git fetch
        git checkout $VERSION
        git pull
    fi
}