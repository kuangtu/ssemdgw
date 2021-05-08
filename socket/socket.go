packet mdgwsocket

import (
    "io"
    "net"
)


type MdgwSock struct {
    laddr, raddr *net.TCPAddr
    lconn, rconn io.ReadWriteCloser
}



