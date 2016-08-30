from thriftpy import load
from thriftpy.rpc import make_server
from thriftpy.transport import TFramedTransportFactory
echo_thrift = load("/app/sh/echo.thrift", module_name="echo_thrift")


class echoHandler(object):

    def echo(self, msg):
        return msg.text

make_server(service=echo_thrift.Echo, handler=echoHandler(),
        host='127.0.0.1', port=9999,
        trans_factory=TFramedTransportFactory()).serve()
