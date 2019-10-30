# Healthcheckers

Very simple client implementations for a variety of protocols (http, dns, rtsp)
meant to be used as basic health-checkers in Docker images providing services exposing protocols.

Implementation is in golang.

All clients consume a HEALTHCHECK_URL variable (typically `1.2.3.4:12345`) that should point to the service address to be tested.

The DNS client additionally support HEALTHCHECK_TYPE (`tcp` or `udp`) and HEALTHCHECK_QUESTION (a domain name to be resolved).

In case of any error, the clients will exit 1, without any detail.

In case of success, the body of the response may be returned for further processing.
