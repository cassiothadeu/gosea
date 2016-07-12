all: install run

key:
	openssl genrsa -out server.key 2048
	openssl ecparam -genkey -name secp384r1 -out server.key

cert:
	openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650

install: key cert
	go install

run:
	gosea
