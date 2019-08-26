package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/miekg/dns"
)

type config struct {
	ProxyHost string `env:"PROXY_HOST" envDefault:"localhost"`
	ProxyPort string `env:"PROXY_PORT" envDefault:"8053"`
	ProxyTransport string `env:"PROXY_TRANSPORT" envDefault:"tcp"`
	DNSServer string `env:"DNS_SERVER" envDefault:"1.1.1.1:853"`
}

func main() {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("Failed to parse environment: %s\n ", err.Error())
	}

	server := &dns.Server{Addr: cfg.ProxyHost+":"+cfg.ProxyPort, Net: cfg.ProxyTransport}
	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		dnsClient := new(dns.Client)
		dnsClient.Net = "tcp-tls"
		response, _, err := dnsClient.Exchange(r, cfg.DNSServer)
		if err != nil {
			fmt.Printf("Failed getting DNS response: %s\n", err.Error())
		}
		w.WriteMsg(response)
	})
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed to start server: %s\n ", err.Error())
	}
}
