@echo off
echo Building MaejicCode application...
echo.

echo Step 1: Installing frontend dependencies...
cd ui
call npm install
if %errorlevel% neq 0 (
    echo Failed to install frontend dependencies
    exit /b %errorlevel%
)

echo.
echo Step 2: Building frontend...
call npm run build
if %errorlevel% neq 0 (
    echo Failed to build frontend
    exit /b %errorlevel%
)

cd ..

echo.
echo Step 3: Building Go executable...
go build -o maejiccode.exe ./cmd/web
if %errorlevel% neq 0 (
    echo Failed to build Go executable
    exit /b %errorlevel%
)

echo.
echo ========================================
echo Build completed successfully!
echo Executable: maejiccode.exe
echo ========================================
echo.
echo To run the application:
echo   maejiccode.exe
echo.
echo The application will be available at:
echo   http://localhost:8080
echo ========================================
