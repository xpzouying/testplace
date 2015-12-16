#/usr/bin/env python
'''views.py

Just for views.
'''

from app import first_app
from flask import render_template


@first_app.route('/')
@first_app.route('/index')
def index():
    '''index

    Show index page
    '''

    return "Hello Zou, Ying - From blog"


@first_app.route('/about')
def about():
    '''about

    Show about page
    '''

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


@first_app.route('/hello')
def hello():
    '''hello

    Say hello to vister
    '''
    user = {'name': "Eden"}

    return render_template("hello.html",
                           title="Best_Title",
                           user=user)
