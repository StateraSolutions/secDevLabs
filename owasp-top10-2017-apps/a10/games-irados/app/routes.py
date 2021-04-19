#!/usr/bin/env python
# -*- coding: utf-8 -*-
from functools import wraps
import uuid
import datetime
from flask import (
    Flask,
    render_template,
    request,
    redirect,
    flash,
    make_response,
    session
)
from util.init_db import init_db
from flask.logging import default_handler
from flask_bootstrap import Bootstrap
from model.password import Password
from model.db import DataBase
import logging
import os
import socket
from flask_cors import CORS, cross_origin
from model.db import DataBase

app = Flask(__name__)
bootstrap = Bootstrap(app)

app.config.from_pyfile('config.py')

FORMAT = '%(asctime)-15s %(method)s %(error)s %(message)s'
logging.basicConfig(format=FORMAT)

def loggingOut(method, route, user, status):
    IP = request.remote_addr
    msg = '- STATUS: '+ str(status) + ' - ROUTE: ' + route + ' - IP: '+ str(IP) + ' - USER: ' + str(user)
    info = {'method': "METHOD: "+method, 
            'error': msg
    }
    logging.warning('STATUS LOG', extra=info)


def generate_csrf_token():
    '''
        Generate csrf token and store it in session
    '''
    if '_csrf_token' not in session:
        session['_csrf_token'] = str(uuid.uuid4())
    return session.get('_csrf_token')


app.jinja_env.globals['csrf_token'] = generate_csrf_token

@app.before_request
def csrf_protect():
    '''
        CSRF PROTECION
    '''
    if request.method == "POST":
        token_csrf = session.get('_csrf_token')
        form_token = request.form.get('_csrf_token')
        if not token_csrf or str(token_csrf) != str(form_token):
            return "ERROR: Wrong value for csrf_token"

def login_required(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        if 'username' not in session:
            flash('oops, session expired', "danger")
            return redirect('/login')
        return f(*args, **kwargs)
    return decorated_function

@app.route('/', methods=['GET'])
def root():
    return redirect('/login')

@app.route('/logout', methods=['GET'])
@login_required
def logout():
    session.clear()
    return redirect('/login')

@app.route('/login', methods=['GET', 'POST'])
def login():
    if request.method == 'POST':
        username = request.form.get('username').encode('utf-8')
        psw = Password(request.form.get('password').encode('utf-8'))
        user_password, success = database.get_user_password(username)
        if not success or user_password == None or not psw.validate_password(str(user_password[0])):
            flash("Usuario ou senha incorretos", "danger")
            status = str(401)+" - Usuario ou senha incorretos"
            loggingOut(request.method,'/login', username, status)
            return render_template('login.html')
        session['username'] = username
        status = str(200)+" - Login realizado"
        loggingOut(request.method,'/login', username, status)
        return redirect('/home')
    else:
        return render_template('login.html')

@app.route('/register', methods=['GET', 'POST'])
def newuser():
    if request.method == 'POST':
        username = request.form.get('username').encode('utf-8')
        psw1 = request.form.get('password1').encode('utf-8')
        psw2 = request.form.get('password2').encode('utf-8')

        if psw1 == psw2:
            psw = Password(psw1)
            hashed_psw = psw.get_hashed_password()
            message, success = database.insert_user(username, hashed_psw)
            if success == 1:
                flash("Novo usuario adicionado!", "primary")
                status = str(201)+" - Novo usuario adicionado"
                loggingOut(request.method,'/register', username, status)
                return redirect('/login')
            else:
                flash(message, "danger")
                status = str(500)+" - Erro ao inserir no DB"
                loggingOut(request.method,'/register', username, status)
                return redirect('/register')

        flash("Passwords must be the same!", "danger")
        status = "Erro: passwords nao sao iguais"
        loggingOut(request.method,'/register', username, status)
        return redirect('/register')
    else:
        return render_template('register.html')

@app.route('/home', methods=['GET'])
@login_required
def home():
    return render_template('index.html')

@app.route('/coupon', methods=['GET', 'POST'])
@login_required
def cupom():
    if request.method == 'POST':
        coupon = request.form.get('coupon')
        user = session.get('username')
        rows, success = database.get_game_coupon(coupon, user)
        if not success or rows == None or rows == 0:
            flash("Cupom invalido", "danger")
            status = str(404)+" - Cupom invalido"
            loggingOut(request.method,'/coupon', user, status)
            return render_template('coupon.html')
        game, success = database.get_game(coupon, user)
        if not success or game == None:
            flash("Cupom invalido", "danger")
            status = str(404)+" - Cupom invalido"
            loggingOut(request.method,'/coupon', user, status)
            return render_template('coupon.html')
        flash("Voce ganhou {}".format(game[0]), "primary")
        status = str(200)+" - Cupom valido"
        loggingOut(request.method,'/coupon', user, status)
        return render_template('coupon.html')
    else:
        return render_template('coupon.html')

if __name__ == '__main__':
    dbEndpoint = os.environ.get('MYSQL_ENDPOINT')
    dbUser = os.environ.get('MYSQL_USER')
    dbPassword = os.environ.get('MYSQL_PASSWORD')
    dbName = os.environ.get('MYSQL_DB')
    database = DataBase(dbEndpoint, dbUser, dbPassword, dbName)
    init_db(database)
    app.run(host='0.0.0.0',port=10010, debug=True)
