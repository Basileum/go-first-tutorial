# Get-started go tutorial (go-first-tutorial)
Following this get-started tutorial : https://go.dev/doc/tutorial/getting-started

Done in UNIX - Ubuntu

GO Installation : https://go.dev/doc/install
Installed with this command : 
```
sudo  rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
```

GO version : **go version go1.20.5 linux/amd64**

Once we wrote our code, we can run it with :
```
go run
```

If we use some external package, we can build like this and then run it again
```
go mod tidy
go run
```
