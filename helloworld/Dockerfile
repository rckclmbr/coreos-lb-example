FROM gliderlabs/alpine

MAINTAINER Josh Braegger <rckclmbr@gmail.com>

RUN apk-install py-flask

ADD app.py /app/app.py

EXPOSE 5000

CMD /usr/bin/env python /app/app.py
