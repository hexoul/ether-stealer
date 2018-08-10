build:
	xgo --deps=https://gmplib.org/download/gmp/gmp-6.0.0a.tar.bz2 \
			-targets=linux/arm,linux/amd64,windows/*,darwin/* -out bin/ether-stealer \
			./