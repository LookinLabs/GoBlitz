#!/bin/sh

ENVIRONMENT=$1
INCREMENT=$2

if [[ -z $ENVIRONMENT || ($ENVIRONMENT != "test" && $ENVIRONMENT != "prod") ]]; then
    echo "Please specify the environment as either 'prod' or 'test'."
    echo "Example: ./bin/create_release_tag.sh test patch"
    exit 1
fi

if [[ -z $INCREMENT || ($INCREMENT != "major" && $INCREMENT != "minor" && $INCREMENT != "patch") ]]; then
    echo "Please specify what to increment as either 'major', 'minor', or 'patch'."
    echo "Example: ./bin/create_release_tag.sh test patch"
    exit 1
fi

# Perform a git pull
git pull origin master

if [[ $? -ne 0 ]]; then
    echo "Failed to pull the latest changes from the remote repository."
    exit 1
fi

# Get the latest tag
if [[ $ENVIRONMENT == "prod" ]]; then
    latest_tag=$(git tag -l 'v[0-9]*.[0-9]*.[0-9]*' | grep -v -- '-alpha' | sort -V | tail -n 1)
else
    latest_tag=$(git tag -l 'v[0-9]*.[0-9]*.[0-9]*-alpha' | sort -V | tail -n 1)
fi
echo "Latest tag: $latest_tag"

# Check if latest_tag is empty
if [[ -z $latest_tag ]]; then
    major=0
    minor=0
    patch=0
else
    # Extract the version number from the latest tag
    version=$(echo "$latest_tag" | sed -E 's/v([0-9]+\.[0-9]+\.[0-9]+).*/\1/')

    # Split the version into major, minor, and patch
    major=$(echo $version | cut -d. -f1)
    minor=$(echo $version | cut -d. -f2)
    patch=$(echo $version | cut -d. -f3)
fi

# Increment the specified version part
if [[ $INCREMENT == "major" ]]; then
    major=$((major + 1))
    minor=0
    patch=0
elif [[ $INCREMENT == "minor" ]]; then
    minor=$((minor + 1))
    patch=0
elif [[ $INCREMENT == "patch" ]]; then
    patch=$((patch + 1))
fi

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