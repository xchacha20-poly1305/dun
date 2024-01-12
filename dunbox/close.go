package dunbox

import (
	"github.com/sagernet/sing-box/common/taskmonitor"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/experimental/clashapi"
	E "github.com/sagernet/sing/common/exceptions"
)

func (s *Box) closeClashApi() error {
	if c, ok := s.router.ClashServer().(*clashapi.Server); ok {
		return c.Close()
	}
	return nil
}

func (s *Box) closeInboundListeners() error {
	var errors error
	for i, in := range s.inbounds {
		inType := in.Type()
		if inType == "tun" {
			continue
		}
		monitor := taskmonitor.New(s.logger, C.DefaultStopTimeout)
		monitor.Start("closeInboundListener inbound/", inType, "[", i, "]")
		errors = E.Append(errors, in.Close(), func(err error) error {
			return E.Cause(err, "closeInboundListener inbound/", inType, "[", i, "]")
		})
		monitor.Finish()
	}
	return errors
}
