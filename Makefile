build:
	xgo --deps=https://gmplib.org/download/gmp/gmp-6.0.0a.tar.bz2 \
			-targets=linux/arm,linux/amd64,windows/*,darwin/* -out bin/ether-stealer \
			./

update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
	go get github.com/hexoul/ether-stealer@v2.0.0
