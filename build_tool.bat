@echo off
REM ----------------------------
REM docker compose rebuild tools
REM ----------------------------

REM Stop and remove all containers, networks, volumes, and remove orphans
docker-compose -f docker-compose.dev.yml down --rmi all --volumes --remove-orphans

REM Build all services
docker-compose -f docker-compose.dev.yml build --no-cache

REM Recreate and start all containers
docker-compose -f docker-compose.dev.yml up --force-recreate

REM Access the backend container
REM start cmd /k "docker exec -it backend bash"

REM Access the frontend container
REM start cmd /k "docker exec -it frontend bash"

REM Access the database container (using psql)
REM start cmd /k "docker exec -it db psql -U postgres"

REM View the logs for all services
start cmd /k "docker-compose -f docker-compose.dev.yml logs -f"