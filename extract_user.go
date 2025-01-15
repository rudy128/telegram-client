package main

import (
	"fmt"

	"github.com/gotd/td/tg"
)

func extractUserID(peer tg.PeerClass) (int64, error) {
	// Define the regular expression pattern to capture the numerical value after 'UserID:'
	switch p := peer.(type) {
	case *tg.PeerChannel:
		return p.ChannelID, nil
	case *tg.PeerChat:
		return p.ChatID, nil
	case *tg.PeerUser:
		return p.UserID, nil
	default:
		return 0, fmt.Errorf("unsupported peer type: %T", peer)
	}
}
