@echo off
REM ----------------------------
REM Docker Compose Rebuild Tool
REM ----------------------------

REM Set the default compose file
set COMPOSE_FILE=docker-compose.dev.yml

REM Check if a different compose file is provided as an argument
if not "%1"=="" set COMPOSE_FILE=%1

echo Using compose file: %COMPOSE_FILE%

REM Confirm before proceeding
set /p CONTINUE=This will stop and remove all containers, networks, volumes, and images. Continue? (y/n): 
if /i not "%CONTINUE%"=="y" exit /b

REM Confirm removing Swagger docs
echo Generating Swagger documentation...
IF EXIST ".\backend\docs" (
    echo Removing existing docs directory...
    rmdir /s /q ".\backend\docs"
)

REM Swagger docs generation
echo Initializing Swagger documentation...
cd backend
swag init --parseDependency --dir ./server --generalInfo main.go
cd ..

REM Stop and remove all containers, networks, volumes, and remove orphans
docker-compose -f %COMPOSE_FILE% down --rmi all --volumes --remove-orphans

REM Set max retries for build process
set MAX_RETRIES=3
set RETRY_COUNT=0

:BUILD_RETRY
echo [INFO] Build attempt #%RETRY_COUNT%

REM Build all services without cache
docker-compose -f %COMPOSE_FILE% build --no-cache

REM Check if build was successful
if %ERRORLEVEL% NEQ 0 (
    set /a RETRY_COUNT+=1
    echo [ERROR] Build failed! Attempt #%RETRY_COUNT% of %MAX_RETRIES%

    if %RETRY_COUNT% LSS %MAX_RETRIES% (
        echo [INFO] Retrying the build...
        goto BUILD_RETRY
    ) else (
        echo [FATAL] Build failed after %MAX_RETRIES% attempts. Exiting.
        exit /b 1
    )
) else (
    echo [SUCCESS] Build completed successfully.
)

REM Recreate and start all containers
docker-compose -f %COMPOSE_FILE% up --force-recreate

REM Provide options to access containers
echo.
echo Access Containers:
echo 1. Backend
echo 2. Frontend
echo 3. Database
echo 4. View Logs
echo 5. Exit
set /p CHOICE=Choose an option (1-5): 

if "%CHOICE%"=="1" start cmd /k "docker exec -it backend bash"
if "%CHOICE%"=="2" start cmd /k "docker exec -it frontend bash"
if "%CHOICE%"=="3" start cmd /k "docker exec -it db psql -U postgres"
if "%CHOICE%"=="4" start cmd /k "docker-compose -f %COMPOSE_FILE% logs -f"
if "%CHOICE%"=="5" exit /b