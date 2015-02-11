FROM ubuntu:14.04

MAINTAINER Josh Braegger <rckclmbr@gmail.com>

RUN apt-get install -y python-flask

ADD app.py /app/app.py

EXPOSE 5000

CMD /usr/bin/env python /app/app.py
