from sys import argv
from multiprocessing import Process
from uuid import uuid4
from thriftpy import load
from thriftpy.rpc import make_client
echo_thrift = load("../../echo.thrift", module_name="echo_thrift")

# Number of concurrent clients to run
clientCount = 10

def handleClient(num):
    client = make_client(echo_thrift.Echo, 'localhost', 9999)
    while num > 0:
        # Make thrift call and output result
        msg = echo_thrift.Message(text=str(uuid4()))
        ret = client.echo(msg)
        if msg.text == ret:
            print ('PyClient: %s' % ret)
        else:
            print ('PyClient: ERROR for %s', msg.text)
        num -= 1

def runClient():
    # Parse command line arguments
    # num = Number of requests each client will make
    try:
        num = int(argv[1])
    except:
        num = 1

    # Spawn client connections
    processes = []
    global clientCount
    while clientCount > 0:
        p = Process(target=handleClient, args=(num,))
        p.start()
        p.join()
        clientCount -= 1

# "Main"
runClient()
