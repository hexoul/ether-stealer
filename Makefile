build:
	xgo --deps=https://gmplib.org/download/gmp/gmp-6.0.0a.tar.bz2 \
			-targets=linux/arm,windows/*,darwin/* -out bin/ether-stealer \
			./