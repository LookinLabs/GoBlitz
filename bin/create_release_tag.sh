#!/bin/sh

ENVIRONMENT=$1

if [[ -z $ENVIRONMENT || ($ENVIRONMENT != "test" && $ENVIRONMENT != "prod") ]]; then
    echo "Please specify the environment as either 'prod' or 'test'."
    echo "Example: ./bin/create_release_tag.sh test"
    exit 1
fi

# Perform a git pull
git pull origin master

if [[ $? -ne 0 ]]; then
    echo "Failed to pull the latest changes from the remote repository."
    exit 1
fi

# Get the latest tag
latest_tag=$(git describe --tags --abbrev=0)
echo "Latest tag: $latest_tag"

# Extract the version number from the latest tag
version=$(echo "$latest_tag" | sed -E 's/v([0-9]+\.[0-9]+\.[0-9]+).*/\1/')

# Split the version into major, minor, and patch
major=$(echo $version | cut -d. -f1)
minor=$(echo $version | cut -d. -f2)
patch=$(echo $version | cut -d. -f3)

# Increment the patch number
patch=$((patch + 1))

# Join the major, minor, and patch into the new version
new_version="$major.$minor.$patch"

# Create the new tag
if [[ $ENVIRONMENT == "prod" ]]; then
    new_tag="v$new_version"
else
    new_tag="v$new_version-alpha"
fi

echo "Creating release tag: $new_tag"
git tag $new_tag