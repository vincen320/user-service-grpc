#Server Environment, eg:api.example.com
SERVER_CN=localhost

#Step 1: Generate Certificaat Authority (CA) + Trust Certificate (ca.crt)
openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
openssl req -passin pass:1111 -new -x509 -days 365 -key ca.key -out ca.crt -subj "/CN=${SERVER_CN}"
#or for Windows use double slash
# openssl req -passin pass:1111 -new -x509 -days 365 -key ca.key -out ca.crt -subj "//CN=${SERVER_CN}"


#Step 2: Generate Server Private Key(server.key)
openssl genrsa -passout pass:1111 -des3 -out server.key 4096

#Step 3: Get certificate signing request, to req signing to CA (w/ server.csr -> this generated file)
openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "/CN=${SERVER_CN}"
#or for Windows use double slash
# openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "//CN=${SERVER_CN}"

#Step 4: Sign the csr to CA that have been created. - use as certFile
openssl x509 -req -passin pass:1111 -days 265 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt

#Step 5:
#Convert Server certificate to .pem - use as keyFile by gRPC
openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem