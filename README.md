# learn-cicd-starter (Notely)

![code coverage badge](https://github.com/castoldie/learn-cicd-starter/.github/workflows/ci.yml/badge.svg)


This repo contains the starter code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

## Local Development

Make sure you're on Go version 1.20+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8000"
```

Run the server:

```bash
go build -o notely && ./notely
```

*This starts the server in non-database mode.* It will serve a simple webpage at `http://localhost:8000`.

You do *not* need to set up a database or any interactivity on the webpage yet. Instructions for that will come later in the course!

# PROJECT FINISHED!

## What i did in this course:

set up a continuous integration pipeline with GitHub Actions that ensures new PRs pass certain checks before they are merged to main:
  
- Unit tests pass
- Formatting checks pass
- Linting checks pass
- Security checks pass
  
configured a cloud-based MySQL database hosted on PlanetScale
set up a continuous deployment pipeline with GitHub Actions that does the following whenever changes are merged into main:

- Builds a new server binary
- Builds a new Docker image for the server
- Pushes the Docker image to the Google Artifact Registry
- Deploys a new Cloud Run revision with the new image and serves the app to the public interne
