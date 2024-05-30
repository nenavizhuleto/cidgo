package main

import (
	"flag"
	"log"

	"github.com/nenavizhuleto/cidgo"
)

var (
	addr  = flag.String("addr", "localhost:9123", "CID TCP Server") // may use netcat: `nc -lvp 9123` - starts tcp listener on port 9123
	proto = flag.String("proto", "surguard", "Protocol")

	dev_id = flag.String("devid", "1234", "Device ID")
	sector = flag.String("sector", "00", "Sector")
	zone   = flag.String("zone", "000", "Zone")

	recv_id = flag.String("recvid", "00", "Receiver ID")
	line    = flag.String("line", "1", "Line")
)

func main() {
	flag.Parse()
	client := cidgo.NewClient(*addr, cidgo.Protocol(*proto))

	dev := cidgo.NewDevice(*dev_id, *sector, *zone)
	recv := cidgo.NewReceiver(*recv_id, *line)

	log.Println(dev)
	log.Println(recv)

	if err := client.SendCommand(dev, recv, cidgo.PanicCommand); err != nil {
		log.Fatalln(err)
	}

	if err := client.SendCommand(dev, recv, cidgo.OCByUser); err != nil {
		log.Fatalln(err)
	}

	if err := client.SendCommand(dev, recv, cidgo.PeriodicTestCommand); err != nil {
		log.Fatalln(err)
	}

}
