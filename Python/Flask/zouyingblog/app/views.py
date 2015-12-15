#/usr/bin/env python

from app import first_app
@first_app.route('/')
@first_app.route('/index')
def index():
    return "Hello Zou, Ying - From blog"


@first_app.route('/about')
def about():
    return '''
    <html>
        <head>
            <title>About Page</title>
        </head>

        <body>
            <h1>About Page</h1>
        </body>
    </html>
    '''
