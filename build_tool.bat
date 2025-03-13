@echo off
REM ----------------------------
REM Docker Compose Rebuild Tool
REM ----------------------------

REM Set the default compose file
set COMPOSE_FILE=docker-compose.dev.yml

REM If a different compose file is provided as an argument, use it
if not "%1"=="" set COMPOSE_FILE=%1

echo [INFO] Using compose file: %COMPOSE_FILE%

REM -------------------------------------------------
REM Stop and remove containers, images, and volumes
REM -------------------------------------------------
docker-compose -f %COMPOSE_FILE% down --rmi all --volumes --remove-orphans

REM -------------------
REM Build retry logic
REM -------------------
set MAX_RETRIES=3
set RETRY_COUNT=0

:BUILD_RETRY
echo [INFO] Build attempt #%RETRY_COUNT%

docker-compose -f %COMPOSE_FILE% build --no-cache
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

REM ----------------------------------
REM Recreate and start all containers
REM ----------------------------------
echo [INFO] Starting services...
docker-compose -f %COMPOSE_FILE% up --force-recreate

REM --------------------------------------
REM Script ends without further questions
REM --------------------------------------
echo [INFO] All done. Containers are running.
exit /b 0