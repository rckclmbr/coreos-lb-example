FROM gliderlabs/alpine
MAINTAINER Josh Braegger <rckclmbr@gmail.com>

# Install Nginx
RUN apk-install nginx curl bash

# Install confd
RUN curl -sL https://github.com/kelseyhightower/confd/releases/download/v0.8.0/confd-0.8.0-linux-amd64 > /usr/local/bin/confd; \
    chmod 755 /usr/local/bin/confd

# Create directories
RUN mkdir -p /etc/confd/conf.d; \
    mkdir -p /etc/confd/templates; \
    mkdir -p /tmp/nginx

ADD ./nginx.conf.tmpl /etc/confd/templates/nginx.conf.tmpl
ADD ./nginx.toml /etc/confd/conf.d/nginx.toml
ADD ./boot.sh /opt/boot.sh

# Run the boot script
CMD /opt/boot.sh
