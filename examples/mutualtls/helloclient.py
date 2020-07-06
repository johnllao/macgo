import http.client
import ssl

ctx = ssl.SSLContext(ssl.PROTOCOL_TLSv1_2)
ctx.load_cert_chain(
    "/Users/johnllao/src/macgo/examples/mutualtls/client1cert.pem",
    "/Users/johnllao/src/macgo/examples/mutualtls/client1pvtkey.pem",
    None
)
ctx.load_verify_locations(
    "/Users/johnllao/src/macgo/examples/mutualtls/servercert.pem"
)
ctx.verify_mode = ssl.CERT_REQUIRED
ctx.check_hostname = True

conn = http.client.HTTPSConnection(host="localhost", port=8443, context=ctx)
conn.request("GET", "/")
res = conn.getresponse()
data = res.read()
print(data)