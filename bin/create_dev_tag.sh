#!/bin/sh

# Perform a git pull
git pull origin master

# Get the latest tag
latest_tag=$(git describe --tags --abbrev=0)

# Extract the version number from the latest tag
version=$(echo "$latest_tag" | sed -E 's/v([0-9]+\.[0-9]+\.[0-9]+).*/\1/')
echo $version

# Split the version into major, minor, and patch
major=$(echo $version | cut -d. -f1)
minor=$(echo $version | cut -d. -f2)
patch=$(echo $version | cut -d. -f3)

# Increment the patch number
patch=$((patch + 1))

# Join the major, minor, and patch into the new version
new_version="$major.$minor.$patch"
echo $new_version

# Create the new tag
new_tag="v$new_version-alpha"
git tag $new_tag