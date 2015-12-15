from flask import Flask
from middleware import HelloMiddleware

app = Flask(__name__)
print "Begin: app in main: ", app, ", app.wsgi_app: ", app.wsgi_app
app.wsgi_app = HelloMiddleware(app.wsgi_app)
print "After: app in main: ", app, ", app.wsgi_app: ", app.wsgi_app

@app.route('/')
def hello_app():
    return "Hello App!"


if __name__ == "__main__":
    app.run()
