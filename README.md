# Backend Server for Bojongireng's app
---
This backend server is intended to serve all frontend apps, like mobile app (and web app if present).
Written in Go v1.24.5.

## Requirements
- Docker Installed: https://www.docker.com/products/docker-desktop/
You don't need to install Go or PostgreSQL.

## How to run server
1. Clone the repository
<pre> ```sh
git clone https://github.com/Bojong-Ireng/backend-server.git
```</pre>

2. Build database image 
You may need to start docker dekstop / docker service first. 

<pre> ```sh
cd ./docker/db
docker build -f db.Dockerfile -t preloaded-db .
# Go back to root
cd ../../
```</pre>

3. Compose the server
This step depends on whether you're running it on production or development.

### Build mode / Production

### Backend development
<pre> ```sh
# Make sure you're on root directory
docker compose -f docker/dev/dev-compose.yaml build
# You may remove -d flag or use other flags to your needs.
docker compose -f docker/dev/dev-compose.yaml up -d
```</pre>

If you use docker dekstop, you can check it now.