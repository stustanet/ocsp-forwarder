# OCSP Forwarder

A workaround for web servers without HTTP proxy support for OCSP requests to use an HTTP proxy nonetheless.

## Setup

```sh
# add system user for OCSP Forwarder
useradd --system -s /bin/false -M ocsp-forwarder

# Make sure that Go and git are installed. For Debian Stretch:
apt install git
apt install -t stretch-backports golang

# install go package
GOPATH=/usr/local/src/go GOBIN=/usr/local/bin go get github.com/stustanet/ocsp-forwarder

# install and start systemd service
cp /usr/local/src/go/src/github.com/stustanet/ocsp-forwarder/systemd/ocsp-forwarder.service /etc/systemd/system/
# edit /etc/systemd/system/ocsp-forwarder.service
systemctl enable --now ocsp-forwarder.service
```

Adjust the parameters in the `/etc/systemd/system/ocsp-forwarder.service` as nedeed. For Let's Encrypt X3 certificates the `responder_url` is `http://r3.o.lencr.org/`.


In the nginx config (server block):

```
ssl_stapling_responder http://127.0.0.1:8234;
```

[`ssl_trusted_certificate`](https://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_trusted_certificate) (containing the chain of intermediate certificates) must also be set!

Verify that OCSP stapling works:
```
openssl s_client -connect example.com:443 -tls1_2  -tlsextdebug  -status | grep -i "OCSP Response"
```
