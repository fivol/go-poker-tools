# Poker combinations by board and list of hands


### Build
```
cd cmd/go-poker-combinations
go build
chmod +x go-poker-combinations
mv go-poker-combinations go-poker-combinations-macos
GOOS=linux GOARCH=386 go build
chmod +x go-poker-combinations
mv go-poker-combinations go-poker-combinations-linux
```
[Build for different Architectures](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04) 

### Examples
```
./go-poker-combinations --hands 2h2c Ks7s2s
./go-poker-combinations --hands 2h2c --combos set Ks7s2s
echo 2h2c | ./go-poker-combinations Ks7s2s
```

