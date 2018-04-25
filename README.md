# OCSP Forwarder

A workaround for web servers without HTTP proxy support for OCSP requests to use an HTTP proxy nonetheless.

## Setup

```sh
# add system user for OCSP Forwarder
useradd --system -s /bin/false -M ocsp-forwarder

# clone and build git source
git clone https://gitlab.stusta.de/julienschmidt/ocsp-forwarder.git
cd ocsp-forwarder
go build -o /usr/local/bin/ocsp-forwarder

# install and start systemd service
cp systemd/ocsp-forwarder.service /etc/systemd/system/
systemctl enable ocsp-forwarder.service
systemctl start ocsp-forwarder.service
```

In the nginx config (server block):

```
ssl_stapling_responder http://127.0.0.1:8234;
```

[`ssl_trusted_certificate`](https://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_trusted_certificate) must also be set!
