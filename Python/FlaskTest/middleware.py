from flask import request

class HelloMiddleware(object):

    def __init__(self, app):
        self.app = app

    def __call__(self, environ, start_response):
        print "I'm in middleware"
        return self.app(environ, start_response)
