#/usr/bin/env python

from flask import Flask

first_app = Flask(__name__)

from app import views
