package packet_parser

type PacketParser interface {
	parse(data []byte) []byte
}

type HeaderFilter interface {
	filter(data []byte) []byte
}

type packetParser func(data []byte) []byte
type headerFilter func(data []byte) []byte

func (p packetParser) parse(data []byte) []byte {
	return p(data)
}

func (h headerFilter) filter(data []byte) []byte {
	return h(data)
}
