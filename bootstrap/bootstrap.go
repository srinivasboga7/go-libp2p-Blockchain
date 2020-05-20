package bootstrap

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	multiaddr "github.com/multiformats/go-multiaddr"
	"golang.org/x/net/context"
)

// NewBootstrapPeer spins up a DHT node which will be used by other nodes in peer discovery
func NewBootstrapPeer(hostAddr multiaddr.Multiaddr, peers []multiaddr.Multiaddr) error {
	ctx := context.Background()

	host, err := libp2p.New(ctx,
		libp2p.ListenAddrs([]multiaddr.Multiaddr{hostAddr}...),
	)
	// connect with other bootstrap peers
	peerAddrs, err := peer.AddrInfosFromP2pAddrs(peers...)

	for _, peerAddr := range peerAddrs {
		host.Connect(ctx, peerAddr)
	}

	if err != nil {
		return err
	}

	kademliaDHT, err := dht.New(ctx, host)
	if err != nil {
		return err
	}

	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		return err
	}

	return nil
}
