FROM gliderlabs/alpine
MAINTAINER Josh Braegger <rckclmbr@gmail.com>

# Install Nginx
#RUN echo "deb http://ppa.launchpad.net/nginx/development/ubuntu trusty main" > /etc/apt/sources.list.d/nginx.list; \
#    apt-key adv --keyserver keyserver.ubuntu.com --recv-keys C300EE8C; \
#    apt-get update; \
#    apt-get install -y --force-yes nginx curl
RUN apk-install nginx curl bash

# Install confd
RUN curl -sL https://github.com/kelseyhightower/confd/releases/download/v0.6.3/confd-0.6.3-linux-amd64 > /usr/local/bin/confd; \
    chmod 755 /usr/local/bin/confd

# Create directories
RUN mkdir -p /etc/confd/conf.d
RUN mkdir -p /etc/confd/templates

# Remove default site
RUN rm -f /etc/nginx/sites-enabled/default

ADD ./nginx.conf.tmpl /etc/confd/templates/nginx.conf.tmpl
ADD ./nginx.toml /etc/confd/conf.d/nginx.toml
ADD ./boot.sh /opt/boot.sh

# Run the boot script
CMD /opt/boot.sh
