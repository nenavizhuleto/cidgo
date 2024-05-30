package cidgo

import (
	"fmt"
	"regexp"
)

type Protocol string

const (
	Ademco685     = Protocol("ademco685")
	SurGuard      = Protocol("surguard")
	SurGuardTimed = Protocol("surguard-timed")
)

var is_numeric = regexp.MustCompile(`^[0-9]+$`)

// NOTE: there are more elegant solutions to this problem, but it is what it is.
func must_be_numeric(value string) string {
	if is_numeric.MatchString(value) {
		return value
	} else {
		panic(fmt.Sprintf("expected numeric value, but got: %s", value))
	}
}

type Device struct {
	id     string
	sector string
	zone   string
}

func NewDevice(id, sector, zone string) Device {
	return Device{
		id:     must_be_numeric(id),
		sector: must_be_numeric(sector),
		zone:   must_be_numeric(zone),
	}
}

type Receiver struct {
	id   string
	line string
}

func NewReceiver(id, line string) Receiver {
	return Receiver{
		id:   must_be_numeric(id),
		line: must_be_numeric(line),
	}
}

type CommandType string

const (
	CreateCommand = CommandType("create")
	UpdateCommand = CommandType("update")
	ReadCommand   = CommandType("read")
)

type Command struct {
	code string
	t    CommandType
}

func NewCommand(code string, t CommandType) Command {
	return Command{
		code: must_be_numeric(code),
		t:    t,
	}
}

// Commands
var (
	PanicCommand        = NewCommand("120", CreateCommand) // Panic/Alarm
	OCByUser            = NewCommand("401", CreateCommand) // Open/Close By User
	PeriodicTestCommand = NewCommand("602", CreateCommand) // Periodic Test Report
)
