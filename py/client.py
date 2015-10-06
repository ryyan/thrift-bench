#!/usr/bin/env python
import sys
import glob
sys.path.append('../echo')

from Echo import Client
from ttypes import Message

from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TCompactProtocol

# Number of concurrent clients to run
clientCount = 10

def handleClient(num):
    # Open independent client connection
    transport = TSocket.TSocket(port=9090)
    transport = TTransport.TBufferedTransport(transport)
    protocol = TCompactProtocol.TCompactProtocol(transport)
    client = Client(protocol)
    transport.open()

    # Make thrift call and output result
    msg = Message(text=str(num))
    ret = client.echo(msg)
    transport.close()

    if msg.text == ret:
        print ('PyClient: %s' % ret)
    else:
        print ('PyClient: ERROR for %s', msg.text)

def runClient():
    # Parse command line arguments
    try:
        num = int(sys.argv[1])
    except:
        num = 1

    while num > 0:
        # TODO: Make this multi-threaded to mirror goroutines
        handleClient(num)
        num -= 1

# "Main"
runClient()
