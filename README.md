# envoy-proxy
Repository with testing services in different languages (Go, JavaScript) to test the Envoy proxy technology. Envoy configuration files are included.

## Services
These are services which will be running on the server and we will be trying to connect them all to one envoy daemon and configuration file. Below is the list of all services.

* **go-calc** – simple REST API calculator (operations +, -, *, /)
* **go-redir** – redirect API (you will be redirected to specified URL)
* **go-time** – get current time in different formats
* **js-greeting** – get random greeting from an array
* **js-redir** – JS version of the go-redir service
* **js-web** – simple website running on Node.js engine