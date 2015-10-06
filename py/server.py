#!/usr/bin/env python
import sys
import glob
sys.path.append('../echo')

from Echo import Processor
from ttypes import Message

from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TCompactProtocol
from thrift.server import TServer


class EchoHandler:

    def echo(self, msg):
        print ('PyServer: %s' % msg.text)
        return msg.text


def runServer():
    # Set handler
    handler = EchoHandler()
    processor = Processor(handler)

    # Set transport. Should match what is in main.go.
    transport = TSocket.TServerSocket(port=9090)
    tfactory = TTransport.TBufferedTransportFactory()
    pfactory = TCompactProtocol.TCompactProtocolFactory()

    # Build and start server
    print ('PyServer started on port %s' % transport.port)
    server = TServer.TThreadedServer(processor, transport, tfactory, pfactory)
    server.serve()

# "Main"
runServer()
