

from wsgiref.simple_server import make_server
from hellowsgi import application

httpd = make_server('', 8989, application)
print "Serving HTTP on port 8989..."

httpd.serve_forever()
