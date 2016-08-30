from sys import argv
from uuid import uuid4
from multiprocessing import Process, Value
from thriftpy import load
from thriftpy.rpc import make_client
from thriftpy.transport import TFramedTransportFactory
echo_thrift = load("/app/sh/echo.thrift", module_name="echo_thrift")


def handleClient(num, actual):
    # Open independent client connection
    client = make_client(service=echo_thrift.Echo, host='127.0.0.1',
                         port=9999, trans_factory=TFramedTransportFactory())

    # UUID
    uid = str(uuid4())

    for i in range(num):
        # Make thrift call and increment atomic count
        txt = uid + str(i)
        ret = client.echo(echo_thrift.Message(text=txt))
        if (txt == ret):
            with actual.get_lock():
                actual.value += 1


# Parse command line arguments
# Number of requests each client will make
num = int(argv[1])

# Number of concurrent clients to run
clientCount = 10

# Number of request completed; Using a Value to ensure atomicity
actual = Value('I', 0)

# Spawn client connections
processes = []
[processes.append(Process(target=handleClient, args=(num, actual))) for _ in range(clientCount)]
[p.start() for p in processes]
[p.join() for p in processes]

if actual.value != (clientCount * num):
    print("ERROR: Actual and expected completed requests mismatch")
    print("Expected: " + str(clientCount * num))
    print("Actual: " + str(actual.value))
