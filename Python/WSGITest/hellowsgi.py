

def application(environ, start_response):
    start_response('200 OK', [('Content-Type', 'text/html')])
    return '<h1>Hello WSGI! %s.</h1>' % (environ['PATH_INFO'][1:] or 'Web')
