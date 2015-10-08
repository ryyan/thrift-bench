from sys import argv
from multiprocessing import Process
from thriftpy import load
from thriftpy.rpc import make_client
echo_thrift = load("../../echo.thrift", module_name="echo_thrift")


def handleClient(num):
    # Open independent client connection
    client = make_client(echo_thrift.Echo, 'localhost', 9999)
    
    for i in range(num):
        # Make thrift call and output result
        msg = echo_thrift.Message(text=str(i))
        ret = client.echo(msg)
        print ('PyClient: %s' % ret)

# Parse command line arguments
# Number of requests each client will make
num = int(argv[1])

# Number of concurrent clients to run
clientCount = 10

# Spawn client connections
processes = []
[processes.append(Process(target=handleClient, args=(num,))) for _ in range(clientCount)]
[p.start() for p in processes]
[p.join() for p in processes]
