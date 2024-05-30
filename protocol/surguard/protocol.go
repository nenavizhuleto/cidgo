package surguard

const (
	packet_size        = 24
	recv_size          = 2
	line_size          = 1
	abonent_size       = 4
	event_type_size    = 1
	event_code_size    = 3
	sector_number_size = 2
	zone_number_size   = 3

	dc4 = 0x14
)

// 5 RR L s 18 AAAA Q XYZ GG CCC DC4
type Packet [packet_size]byte

func NewPacket(
	receiver,
	line,
	abonent,
	event_type,
	event_code,
	sector_number,
	zone_number string,
) (pkt Packet) {
	i := 0
	pkt[i] = '5'
	i += 1
	copy(pkt[i:], receiver[:recv_size])
	i += recv_size
	copy(pkt[i:], line[:line_size])
	i += line_size
	pkt[i] = ' '
	i += 1
	copy(pkt[i:], "18")
	i += 2
	copy(pkt[i:], abonent[:abonent_size])
	i += abonent_size
	copy(pkt[i:], event_type[:event_type_size])
	i += event_type_size
	copy(pkt[i:], event_code[:event_code_size])
	i += event_code_size
	copy(pkt[i:], sector_number[:sector_number_size])
	i += sector_number_size
	copy(pkt[i:], zone_number[:zone_number_size])
	i += zone_number_size
	pkt[i] = dc4
	return
}

// type Packet struct {
// 	_5    byte    // packet type, always set to '5'
// 	_RR   [2]byte // receiver number
// 	_L    byte    // line number
// 	_s    byte    // space byte, always set to ' '
// 	_18   [2]byte // format identificator, always set to '18'
// 	_AAAA [4]byte // abonent number
// 	_Q    byte    // event type, 'E' - new event, 'R' - restore event, 'P' status
// 	_XYZ  [3]byte // event code
// 	_GG   [2]byte // sector number
// 	_CCC  [3]byte // zone number
// 	_DC4  uint16  // always set to 0x14
// }

type EventType byte

const (
	NewEvent     = EventType('E')
	RestoreEvent = EventType('R')
	StatusEvent  = EventType('P')
)

// func NewPacket(
// 	receiver [2]byte,
// 	line byte,
// 	abonent [4]byte,
// 	event_type EventType,
// 	event_code [3]byte,
// 	sector [2]byte,
// 	zone [3]byte,
// ) Packet {
// 	return Packet{
// 		_5:    '5',
// 		_RR:   receiver,
// 		_L:    line,
// 		_s:    ' ',
// 		_18:   [2]byte{'1', '8'},
// 		_AAAA: abonent,
// 		_Q:    byte(event_type),
// 		_XYZ:  event_code,
// 		_GG:   sector,
// 		_CCC:  zone,
// 		_DC4:  0x14,
// 	}
// }
