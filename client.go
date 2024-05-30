package cidgo

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/nenavizhuleto/cidgo/protocol/surguard"
)

var (
	ErrNotAllDataWritten      = errors.New("not all data written")
	ErrUnsupportedProtocol    = errors.New("protocol not supported")
	ErrUnsupportedCommandType = errors.New("command not supported")
)

type Client struct {
	addr  string
	proto Protocol

	timeout time.Duration
}

func NewClient(addr string, proto Protocol) *Client {
	return &Client{
		addr:  addr,
		proto: proto,

		timeout: 10 * time.Second,
	}
}

func (c *Client) SetTimeout(d time.Duration) {
	c.timeout = d
}

func (c *Client) SendCommand(device Device, receiver Receiver, command Command) error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}

	if err := conn.SetWriteDeadline(time.Now().Add(c.timeout)); err != nil {
		return err
	}

	switch c.proto {
	case SurGuard:
		return c.send_proto_surguard(conn, device, receiver, command)
	default:
		return fmt.Errorf("%s: %s", ErrUnsupportedProtocol, c.proto)
	}
}

func (c *Client) send_proto_surguard(conn net.Conn, device Device, receiver Receiver, command Command) error {

	var event_type surguard.EventType
	switch command.t {
	case CreateCommand:
		event_type = surguard.NewEvent
	case UpdateCommand:
		event_type = surguard.RestoreEvent
	case ReadCommand:
		event_type = surguard.StatusEvent
	default:
		return fmt.Errorf("%s: %s", ErrUnsupportedCommandType, command.t)
	}

	pkt := surguard.NewPacket(
		receiver.id,
		receiver.line,
		device.id,
		string(event_type),
		command.code,
		device.sector,
		device.zone,
	)

	log.Println("new: ", string(pkt[:]))

	n, err := conn.Write(pkt[:])
	if err != nil {
		return err
	}

	if len(pkt) < n {
		return fmt.Errorf("%s: %d of %d", ErrNotAllDataWritten, n, len(pkt))
	}

	return nil
}
