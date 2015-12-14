#!/usr/bin/env python
"""Test field keyword in python.

"""


"""
test01: show nothing
test02: use next(), see once until yield, but twice will throw error.
test03: use *send(msg)*, not next()

"""

def func01():
    print "In func01()"
    yield 11


def yield_test01():
    # This function will show nothing.
    func01()


def func02():
    print "In func02()"
    yield 12
    print "End func02()"


def yield_test02():
    c = func02()
    c.next()

    # Call twice will throw error.
    # c.next()


def func03():
    print "Begin func03()"
    msg = yield 13
    print msg
    
    msg2 = yield 113
    print msg2

    print "End func03()"


def yield_test03():
    c = func03()
    c.next() # == c.send(None)
    c.send("Fighting!") # (yield 13) will be "Fighting"
    # c.send("The second fighting!")


def func04():
    print 'func04 start'
    i = 1
    while True:
        print 'yielding...'
        m = yield i
        print 'yielded ', m
        i += 1
        if i > 10:
            break


def yield_test04():
    print "begin test04"
    c = func04()
    print "after c = func04()"

    ret = c.next() # must generate yield
    print "return of c.next() ==> ", ret
    ret2 = c.send("send_msg_from_out")
    print "return2 of send() ==> ", ret2
    ret3 = c.send("send_msg3333")
    print "return3 of send() ==> ", ret3

if __name__ == "__main__":
    # yield_test01()
    # yield_test02()
    # yield_test03()
    yield_test04()
