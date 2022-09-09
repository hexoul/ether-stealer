build:
	xgo --deps=https://gmplib.org/download/gmp/gmp-6.0.0a.tar.bz2 \
			-targets=linux/arm,linux/amd64,windows/*,darwin/* -out bin/ether-stealer \
			./

update-pkg-cache:
	GOPROXY=proxy.golang.org go list -m github.com/hexoul/ether-stealer@v2.0.0
