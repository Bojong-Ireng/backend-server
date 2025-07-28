# Backend server for Bojongireng's tracking app
---
This backend server is intended to serve all frontend apps, like mobile app (and web app if present).  
Using Go v1.24 and PostgreSQL v1.17, and is containerized using docker.

## Database relational schema
![RS Bojongireng](https://github.com/user-attachments/assets/afab6925-1b8c-42eb-8591-b010fbfd3e37)

## Prerequisite
- Install Docker: https://www.docker.com/products/docker-desktop/  
You don't need to install Go or PostgreSQL.

## How to build and run the server
1. Clone the repository  
<pre>
git clone https://github.com/Bojong-Ireng/backend-server.git
</pre>

2. Create .env file  
This file is for the backend-server.
Copy below code and fill blank values with yours
> or you can modify values in compose.yaml, the important thing is it should match with this .env file
<pre>
DB_HOST=db
DB_USER=postgres
DB_PASSWD=
DB_NAME=bojong_ireng
DB_PORT=5432
DB_SSL_MODE=disable
DB_TIMEZONE=
</pre>

3. Create password.txt for database
Create a password.txt file inside /docker/db and put your password directly there.  
Below is an example of using bash + vim.
<pre>
  cd ./docker/db
  touch password.txt
  # Put your password directly inside the file
  vim password.txt
</pre>

4. Build database image  
You may need to start docker dekstop / docker service first.  

<pre>
cd ./docker
docker build -f db.Dockerfile -t preloaded-db .
# Go back to root
cd ../
</pre>

5. Compose the server  
This step depends on whether you're running it on production or development (see steps below).

### Build mode / Production
<pre>
# Make sure you are on root of the project
docker compose build
</pre>
After building it, you can freely run / stop the container. Your builds will not be lost if you shutdown your computer. 
> Below steps are for running / stopping the container from CLI, if you are using GUI tools you can skip these steps.  
1. Running the server  
Below is a way to run it in detached mode (-d), you may modify the flag to your needs.

<pre>
# Make sure you are on root of the project
docker compose up -d
</pre>

2. Stopping the server  
<pre>
# Make sure you are on root of the project
docker compose down
</pre>

### Backend development
<pre>
# Make sure you are on root of the project
docker compose -f docker/dev-compose.yaml build
</pre>
After building it, you can freely run / stop the container. Your builds will not be lost if you shutdown your computer.  
> Below steps are for running / stopping the container from CLI, if you are using GUI tools you can skip these steps.

1. Running the server  
Below is a way to run it in detached mode (-d), you may modify the flag to your needs.  
<pre>
# Make sure you are on root of the project
docker compose -f docker/dev-compose.yaml up -d
</pre>
2. Stopping the server  
<pre>
# Make sure you are on root of the project
docker compose -f docker/dev-compose.yaml down
</pre>
---
Thats it, you've successfully build and run the backend-server with its database using docker container.  
If you are using docker dekstop GUI, you can check see the containers and builds now.  
