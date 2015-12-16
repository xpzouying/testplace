#/usr/bin/env python


from app import app
from flask import render_template


@app.route('/')
@app.route('/index')
def index():
    return "Hello Flask!"


@app.route('/hi_for')
def hi_for():
    posts = [{"name": "zouying", "info": "I'm zouying"},
             {"name": "eden", "info": "Martin eden info."}]

    return render_template("hi_for.html",
                          posts=posts)
