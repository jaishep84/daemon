

import (
	"golang.org/x/net/context"
	"net"
)

// A KnownHostsIPWalker shall retrieve all known hosts IPs
// and send them down to the returned channel
type KnownHostsIPWalker func(ctx context.Context) <-chan net.IP
