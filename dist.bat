go build -o dist/hs.exe -ldflags="-w -s"
upx -9 dist/hs.exe