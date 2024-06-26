#!/bin/sh

function install_goblitz() {
    if ! [ -d "$(pwd)/GoBlitz" ]; then
        echo "GoBlitz directory does not exist. Pulling GoBlitz..."
        git config --global advice.detachedHead false
        git clone git@github.com:LookinLabs/GoBlitz.git --branch $VERSION
        rm -rf GoBlitz/.git
    fi
}

function update_goblitz() {
    if `pwd | grep -q "GoBlitz"` && [ "$VERSION" == "master" ]; then
        echo "GoBlitz directory found. Pulling updates..."
        git merge origin/master
    fi
    
    if `pwd | grep -q "GoBlitz"` && [ "$VERSION" != "master" ]; then
        echo "GoBlitz directory found. Checking out $VERSION..."
        git fetch
        git merge $VERSION
    fi
}