package main

import (
	"fmt"

	"time"

	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

var (
	device       string = "lo0"
	snapshot_len int32  = 1024
	promiscuous  bool   = true
	err          error
	timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle
)

func handlePacket(packet gopacket.Packet) {
	if packet.TransportLayer() == nil {
		return
	}

	flow := packet.TransportLayer().TransportFlow()
	src, dst := flow.Endpoints()
	fmt.Println(src.String())
	fmt.Println(dst.String())
	//pk := gopacket.NewPacket(packet.Data(), layers.LayerTypeEthernet, gopacket.Lazy)
	//if el := pk.ErrorLayer(); el.Error() != nil {
	//	panic(el.Error().Error())
	//}
	//for _, layer := range pk.Layers() {
	//	fmt.Println(layer.LayerType())
	//}
}

func main() {

	file, _ := os.Create("./test.cap")
	w := pcapgo.NewWriter(file)
	w.WriteFileHeader(uint32(snapshot_len), layers.LinkTypeEthernet)
	defer file.Close()

	if handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout); err != nil {
		panic(err.Error())
	} else if err := handle.SetBPFFilter("tcp and port 8888"); err != nil {
		panic(err.Error())
	} else {
		ps := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range ps.Packets() {
			//w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
			//w.WritePacket(packet.Metadata().CaptureInfo, packet.Layer(layers.LayerTypeTCP).LayerContents())

			//show all layers
			//for _, lay := range packet.Layers() {
			//	fmt.Println(lay.LayerType().String())
			//}

			if lay := packet.Layer(gopacket.LayerTypePayload); lay != nil {
				l := lay.(*gopacket.Payload)
				fmt.Println(string(l.Payload()[:]))
			}

			//l := packet.Layer(layers.LayerTypeLoopback).(*layers.Loopback)
			//fmt.Println(string(l.Payload))
			//fmt.Println(string(packet.ApplicationLayer().LayerPayload()[:]))

			//fmt.Println(packet.Data())
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
