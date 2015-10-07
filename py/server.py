from thriftpy import load
from thriftpy.rpc import make_server
echo_thrift = load("../../echo.thrift", module_name="echo_thrift")


class echoHandler(object):

    def echo(self, msg):
        print ('PyServer: %s' % msg.text)
        return msg.text

print ('PyServer started on %s' % 'localhost:9999')
server = make_server(echo_thrift.Echo, echoHandler(), 'localhost', 9999)
server.serve()
