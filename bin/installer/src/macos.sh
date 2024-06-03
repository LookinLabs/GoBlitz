#!/bin/sh

function install_macos_utilities() {

    if ! command -v brew &> /dev/null; then
        echo "Homebrew is not installed. Installing Homebrew..."
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    fi

    for common_utils in curl git go; do
        if ! command -v $common_utils &> /dev/null; then
            echo "$common_utils is not installed. Installing $common_utils..."
            brew install $common_utils
        fi
    done

    for linter_utils in golangci-lint gosec; do
        if ! command -v $linter_utils &> /dev/null; then
            echo "$linter_utils is not installed. Installing $linter_utils..."
            brew install $linter_utils
        fi
    done

    for migration_utils in goose; do
        if ! command -v $migration_utils &> /dev/null; then
            echo "$migration_utils is not installed. Installing $migration_utils..."
            brew install $migration_utils
        fi
    done

    if ! command -v air &> /dev/null; then
        echo "air is not installed. Installing air..."
        curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
    fi


    if ! command -v docker &> /dev/null; then
        echo "Docker is not installed. Please install Docker Desktop."
        # Open the Docker Desktop download page
        open https://hub.docker.com/editions/community/docker-ce-desktop-mac/
    fi
}