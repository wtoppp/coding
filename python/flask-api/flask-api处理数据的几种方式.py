��1��json��ʽ��������
from flask import Flask, request

app = Flask(import_name=__name__)


@app.route("/json", methods=["GET","Post"])
def echo():
    data = request.get_json()
    myname = data.get("name")
	age = data.get("age")

    response = "Hi, {}! You said you are {} years old.".format(myname, age)
    return response

app.run(host="0.0.0.0")


��2 form��ʽ��������,ʹ��postman��form-data����
from flask import Flask, request

app = Flask(import_name=__name__)


@app.route("/echo", methods=["POST"])
def echo():
    myname = request.form.get("name", "")
    age = request.form.get("age", "")

    response = "Hey there {}! You said you are {} years old.".format(myname, age)

    return response

app.run()

��3��args������ʽ
from flask import Flask, request

app = Flask(__name__)
@app.route("/echo")
def echo():

    to_echo = request.args.get("name", "")
    response = "{}".format(to_echo)

    return response

if __name__ == "__main__":
    app.run()
	
��4��pulls data from a form field name submitted along with the age field in the query string.
   ʹ��postman��Params����
    http://127.0.0.1:5000/echo?name=tim&age=89
from flask import Flask, request

app = Flask(__name__)
@app.route("/echo", methods=["POST"])
def echo():
    name = request.values.get("name", "")
    to_echo = request.values.get("age", "")

    response = "Hey there {}! You said {}".format(name, to_echo)

    return response

app.run()