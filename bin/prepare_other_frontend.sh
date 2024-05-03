#!/bin/bash

rm -rf views/welcome_page
echo "<!DOCTYPE html>
<html>
<head>
    <title>My Website</title>
</head>
<body>
    <h1>Welcome to my website!</h1>
</body>
</html>" > public/index.html
mkdir frontend