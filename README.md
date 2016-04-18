# shell-wrapper
Shell wrapper that always succeed. It's useful to prevent repeated run by kubernetes job

Requires docker for build, but doesn't require Go

To build for linux amd64 jusn run
```
make
```

To build for another os/arch use GOOS and GOARCH
```
GOOS=darwin make
```

Thanks to *Nathan Leclaire* for his article [Shelled-out Commands In Golang](http://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/)
