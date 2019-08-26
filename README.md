DNS over TLS proxy
==================

A minimal DNS proxy that proxies DNS requests to a remote DNS Server (e.g. Cloudflare) via DNS over TLS.

Components:
* services/dns-tls-proxy: Minimal Golang based DNS server
* services/dns-client: Minimal Ubuntu image containing to perform kdig requests against server

Run example
-----------

The example can be run within Docker (recommended) or on bare metal.

**Run example via Docker**:

```
# TCP transport
$ make test-tcp

# UDP transport
$ make test-udp
```
This will spin up a Docker compose stack, containing DNS-proxy-server and the minimal DNS client.
The client will wait until the DNS server is up and running and then performs one sample DNS query against the server.

**Run example on bare metal**

Running on bare metal requires Go1.11 to be installed, since it uses Go modules.

```
cd services/dns-tls-proxy

# Run server (TCP), perform sample request via TCP 
PROXY_TRANSPORT=tcp go run .
kdig +tcp -d @localhost:8053 heise.de

# Run server (UDP), perform sample request via UDP
PROXY_TRANSPORT=udp go run .
kdig -d @localhost:8053 heise.de
```
