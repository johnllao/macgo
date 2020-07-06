#### Create a self signed certificate

**Syntax**

    openssl req -newkey rsa:2048 -new -nodes -x509 -days <no of days> 
    -out <certificate key filename> 
    -keyout <private key filename> 
    -subj <subject>

**Example**

    openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 
    -out servercert.pem 
    -keyout serverpvtkey.pem 
    -subj "/C=SG/O=johnllao.com/OU=IT/CN=localhost"
    
#### Generate private key

**Example**

    openssl genrsa -des3 -out private.key 2048
    
#### Generate root certificate

**Example**

    openssl genrsa -des3 -out rootCAPrivate.key 2048
    openssl req -x509 -new -nodes -key rootCAPrivate.key -sha256 -days 3650 -out rootCA.pem

If you want a non password protected key just remove the -des3 option
    
#### Create a CSR

**Example**

    openssl genrsa -out dev.mergebot.com.key 2048
    openssl req -new -key dev.mergebot.com.key -out dev.mergebot.com.csr
    
#### Create a certificate and sign with root CA

**Example**

    openssl x509 -req 
    -in dev.mergebot.com.csr 
    -CA rootCA.pem -CAkey rootCAPrivate.key 
    -CAcreateserial 
    -out dev.mergebot.com.crt 
    -days 3650
    -sha256 -extfile dev.mergebot.com.ext
    
#### References

https://deliciousbrains.com/ssl-certificate-authority-for-local-https-development/

https://deliciousbrains.com/https-locally-without-browser-privacy-errors/

    
  