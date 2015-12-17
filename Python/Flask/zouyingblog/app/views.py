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


@first_app.route('/hi_if')
def hi_if():
    '''hi_if

    Use if/else in template.
    '''

    user = {'name': 'UserName007'}

    return render_template("hi_if.html",
                           title="This is Title",
                           user=user)


@first_app.route('/hi_for')
def hi_for():
    '''hi_for

    Use for loop in html template.
    '''

    posts = [
                {
                    "author": {"name": "eden"},
                    "body": "Marten Eden"
                },
                {
                    "author": {"name": "bill"},
                    "body": "Bill Gates - Road in the way"
                }
            ]

    return render_template("hi_for.html",
            posts=posts)
