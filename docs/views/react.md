# Using React or any other framework instead of HTML Templates

In case of using React, Vue or any other framework, it's recommended to run [bin/prepare_other_frontend](../../bin/prepare_other_frontend.sh) script.

It automatically clean ups the views (welcome page only) and sets up basic index.html into public folder. It also creates `frontend` which should be used for the React, Vue or any other frontend framework.

When running `yarn build` or `npm run build` - you must move all the static files to the `public` folder.

## Example (Development)

In case of development you have to run go and frontend separately. You can also use docker-compose to do so.

But beware, that in case of docker-compose you'll not have live reload for the golang application.

### Docker Compose (Hot Reload Not Enabled in Go)

1. Configure .env

```
cp .env.example .env
```

2. Run docker compose

```
docker-compose up
```

3. Navigate to `http://localhost:8000` for the Go App

4. Navigate to `http://localhost:5173` for React App

### Native (Hot Reload enabled)

1. Run Go application

```
make
```

2. Setting up React + Typescript framework via Vite

```
➤ npm create vite@latest
✔ Project name: … frontend
✔ Select a framework: › React
✔ Select a variant: › TypeScript
```

3. Run frontend app

```
cd frontend
npm install
npm run dev
```

4. Navigate to `localhost:8000` for GoBlitz app

5. Navigate to `localhost:5173` for React App

## Example (Production)

1. Setting up React + Typescript framework via Vite

```
➤ npm create vite@latest
✔ Project name: … frontend
✔ Select a framework: › React
✔ Select a variant: › TypeScript
```

2. Making some changes in frontend app

Write some code in frontend folder

3. Building up application

```
cd frontend
npm install & npm run build
cp -rf dist/* ../public/
```

4. Run `make` to serve static files

```
make
```

5. Navigate to `localhost:8000` to check that everything's fine

6. Build your image from dockerfile and push to the remote