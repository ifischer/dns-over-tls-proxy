.DEFAULT_GOAL := all
CONTAINER  = dns_tls_proxy

all: test-tcp test-udp

clean:
	-docker-compose down --rmi all

test-tcp:
	docker-compose up --abort-on-container-exit

test-udp:
	KDIG_CMD='kdig -d @dns-tls-proxy:53 heise.de' PROXY_TRANSPORT=udp docker-compose up --abort-on-container-exit
