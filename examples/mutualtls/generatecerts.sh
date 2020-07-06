#!/bin/bash

openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -out servercert.pem -keyout serverpvtkey.pem -subj "/C=SG/O=johnllao.com/OU=IT/CN=localhost"

openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -out client1cert.pem -keyout client1pvtkey.pem -subj "/C=SG/O=johnllao.com/OU=IT/CN=localhost"

openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -out client2cert.pem -keyout client2pvtkey.pem -subj "/C=SG/O=johnllao.com/OU=IT/CN=localhost"

openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -out client3cert.pem -keyout client3pvtkey.pem -subj "/C=SG/O=johnllao.com/OU=IT/CN=localhost"