package packet_parser

var (
	TcpHeader = headerFilter(tcpHeader)
)

func tcpHeader(data []byte) []byte {
	//gopacket.RegisterLayerType(1111,)
	return nil
}
