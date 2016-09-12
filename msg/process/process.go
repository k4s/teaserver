package process

import (
	"log"

	"github.com/k4s/tea/network/protocol"
	"github.com/k4s/teaserver/conf"
)

var Processor protocol.Processor

func init() {
	switch conf.Protocol {
	case "json":
		Processor = protocol.NewJson()
	case "protobuf":
		Processor = protocol.NewProto()
	default:
		log.Fatal("unknown Protocol: %v", conf.Protocol)
	}

}
