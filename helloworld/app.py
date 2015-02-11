from flask import Flask, request
import socket

app = Flask(__name__)
app.config.from_object(__name__)


@app.route('/', methods=['GET','POST'])
def index():
    return "Hello world from host: %s" % socket.getfqdn()


if __name__ == "__main__":
    app.run(host='0.0.0.0')
