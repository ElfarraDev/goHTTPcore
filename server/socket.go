/*
Example usage:
// web server socket
websocket := socket.Socket {
	IPAdress: net.ParseIP("203.0.113.1"),
	Port: 8080,
}
Client (192.168.1.100:49152) → Server (203.0.113.10:80)

+-------------------+        +------------------+
| Source IP/Port    |  TCP   | Destination     |
|------------------|  Packet |-----------------|
| 192.168.1.100    |   →    | 203.0.113.10    |
| Port: 49152      |        | Port: 80        |
+-------------------+        +------------------+
*/

package server

import "net"

type Socket struct {
	IPAdress net.IP
	Port     uint16
}
