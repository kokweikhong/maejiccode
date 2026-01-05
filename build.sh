#!/bin/bash

echo "Building MaejicCode application..."
echo ""

echo "Step 1: Installing frontend dependencies..."
cd frontend
npm install
if [ $? -ne 0 ]; then
    echo "Failed to install frontend dependencies"
    exit 1
fi

echo ""
echo "Step 2: Building frontend..."
npm run build
if [ $? -ne 0 ]; then
    echo "Failed to build frontend"
    exit 1
fi

cd ..

echo ""
echo "Step 3: Building Go executable..."
go build -o maejiccode ./cmd/web
if [ $? -ne 0 ]; then
    echo "Failed to build Go executable"
    exit 1
fi

echo ""
echo "========================================"
echo "Build completed successfully!"
echo "Executable: maejiccode"
echo "========================================"
echo ""
echo "To run the application:"
echo "  ./maejiccode"
echo ""
echo "The application will be available at:"
echo "  http://localhost:8080"
echo "========================================"
