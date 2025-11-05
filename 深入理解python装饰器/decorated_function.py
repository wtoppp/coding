#-*-coding:UTF-8-*-
'''
python3 decorated_function.py

【装饰器初始化】outer 返回 xyz: <function outer.<locals>.xyz at 0x000001C9AD8B3A60>
【装饰器初始化】outer 返回 xyz: <function outer.<locals>.xyz at 0x000001C9AD8B3BA0>
【1】装饰器开始执行，原函数为send_wechat
【2】xyz函数ID: 1965711637088
【3】原函数func对象: <function send_wechat at 0x000001C9AD8B3920>
【4】原函数func的ID: 1965711636768
【5】接收到的参数: args=('hello weixin',), kwargs={}
*************** 以上是原函数执行前 ***************

【6】准备执行原函数 send_wechat...
     this is msg: hello weixin
【7】原函数执行完成，返回值: weinxin200
【8】装饰器执行结束，返回结果
weinxin200
'''
def outer(func):
    def xyz(*args,**kwargs):  #
        print(f"【1】装饰器开始执行，原函数为{func.__name__}")
        print(f"【2】xyz函数ID: {id(xyz)}")
        print(f"【3】原函数func对象: {func}")
        print(f"【4】原函数func的ID: {id(func)}")
        print(f"【5】接收到的参数: args={args}, kwargs={kwargs}")
        print("*" * 15 + " 以上是原函数执行前 " + "*" * 15 +"\n")
        print(f"【6】准备执行原函数 {func.__name__}...")

        result = func(*args, **kwargs)  # 这才是真正执行原函数的地方,func代表原函数,func()这样加个括号代表执行原函数
        print(f"【7】原函数执行完成，返回值: {result}") #如果在原函数中没有return值就返回None

        print(f"【8】装饰器执行结束，返回结果")
        return result

    print(f"【装饰器初始化】outer 返回 xyz: {xyz}")
    return xyz


@outer #1.第一次装饰器初始化,遇到@outer立即执行 outer(send_wechat) ,2.打印 【装饰器初始化】outer 返回 xyz...3.返回 xyz wrapper 函数替换原来的 send_wechat即send_wechat=outer(send_wechat)
def send_wechat(msg):
    print("       this is msg:",msg)
    return "weinxin200"

@outer #第二次装饰器初始化,即send_email=outer(send_email)
def send_email(to,msg):
    print(to,msg)

if __name__ == '__main__':
    res_wechat=send_wechat("hello weixin")
    #res_email = send_email("myaile@163.com","emailserver")
    print(res_wechat)
    #print(res_email)
