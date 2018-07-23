package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func handlePacket(packet gopacket.Packet) {
	pk := gopacket.NewPacket(packet.Data(), layers.LayerTypeEthernet, gopacket.Lazy)
	for _, layer := range pk.Layers() {
		fmt.Println(layer.LayerType())
	}
}

func main() {
	if handle, err := pcap.OpenLive("any", 0, true, pcap.BlockForever); err != nil {
		panic(err.Error())
	} else if err := handle.SetBPFFilter("tcp and port 80"); err != nil {
		panic(err.Error())
	} else {
		ps := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range ps.Packets() {
			handlePacket(packet)
		}
	}
	//data := "..2..%,Z./\\?..Ed..y.@./...eY.i.+.-.P....2..~V.P...S...........;..y..,&._....tD,..B......k..o2...........|.A..xI.H.4`...wWqp./..v...........F..2..+.S.'O...2X.....j+X$............o..K.x....$.F.......6...F..b....Q.e..`HJ...x..7z..*....*x.....$o...r...e.......Q.*.Qp.n.O...:.G.m......;.....#.P....3z,...a..P.2..I'.<{.q......b.......`..d.\\.An......6.2.#..A.7.5.4g.M##.z..XT@.i?.X..&.,.s..A....[..L>...6pG....9];.%..J.....N....2.xG.2{...:u`.......W[.b.6.....xS.s.P..@@.,3.6..p..;.L..B.@..W.._.............(T....?...^(.7.\".x-..1T..;.t.t..a......ZYh..U.v+.o.."
	//
	//packet := gopacket.NewPacket([]byte(data), layers.LayerTypeEthernet, gopacket.Lazy)
	//
	//ls := packet.Layers()
	//for _, l := range ls {
	//	fmt.Println(l)
	//}

}
