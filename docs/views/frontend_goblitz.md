# FrontEnd development via GoBlitz

## Assets

Assets (images, gifs, css, js files etc) are served from public folder. The relative path is `/assets`.

You can look an example of importing assets in the [Welcome Page Code](../../views/mainPage/welcome.html).

## Using HTML Templates to setup pages

In GoBlitz Views are similar to react, where you can define template components in views/components folder and then use them on the page.
You can take a look for the [Welcome Page](../../views/mainPage/welcome.html)

The pages are separated into folders like views/mainPage/ views/statusPage, you can add there views/AboutPage etc.

HTML Template values are generated via [views/templates folder](../../views/templates/).

## Error pages

Error pages are handled by the [error controller](../../controller/error/http_errors.go). 

Controller look whether the path is available or not and returns the view into middleware, which returns the response to user.