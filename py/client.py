#!/usr/bin/env python
import sys, uuid, threading
sys.path.append('../echo')

from Echo import Client
from ttypes import Message

from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TCompactProtocol

# Number of concurrent clients to run
clientCount = 10


class clientThread(threading.Thread):

    def __init__(self, num):
        threading.Thread.__init__(self)
        self.num = num

    def run(self):
        # Open independent client connection
        transport = TSocket.TSocket(port=9090)
        transport = TTransport.TBufferedTransport(transport)
        protocol = TCompactProtocol.TCompactProtocol(transport)
        client = Client(protocol)
        transport.open()

        while self.num > 0:
            # Make thrift call and output result
            msg = Message(text=str(uuid.uuid4()))
            ret = client.echo(msg)
            if msg.text == ret:
                print ('PyClient: %s' % ret)
            else:
                print ('PyClient: ERROR for %s', msg.text)
            self.num -= 1

        transport.close()


def runClient():
    # Parse command line arguments
    # num = Number of requests each client will make
    try:
        num = int(sys.argv[1])
    except:
        num = 1

    # Spawn client connections
    threads = []
    global clientCount
    while clientCount > 0:
        threads.append(clientThread(num))
        clientCount -= 1
    [t.start() for t in threads]
    [t.join() for t in threads]

# "Main"
runClient()
