echo "Creating Linux version..."
GOOS=linux go build
zip unpaydisco_linux.zip unpaydisco

# echo "Creating Windows version..."
# GOOS=windows GOARCH=386 go build -o solrdora.exe
# zip -r unpaydisco_win.zip unpaydisco.exe settings.json public/* views/* data/*

# echo "Creating Mac OS X version..."
# go build
# zip -r unpaydisco_mac.zip unpaydisco settings.json public/* views/* data/*
