package src

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
)

func moneyChannel(msg *tg.Message, client *telegram.Client, ctx context.Context) {
	chatID, _ := extractUserID(msg.PeerID)
	senderID := msg.FromID
	text := msg.Message
	receptionist_ID, err := strconv.ParseInt(os.Getenv("RECEPTIONIST_ID"), 10, 64)
	chat_ID, err := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	if chatID == chat_ID {
		if senderID == nil {
			if strings.Contains(text, "Task") && strings.Contains(text, "youtube") {
				// EmptyDir(photoDIR)
				messageCount = 0
				collecting = true
			} else {
			}
		} else {
			if messageCount <= 4 && collecting {
				if msg.Media != nil {
					wg.Add(1)
					go func(msg *tg.Message) {
						defer wg.Done()
						download_media(client, ctx, msg)
					}(msg)
				}
				if strings.Contains(strings.ToLower(text), "t") {
					caption = text
				}

			} else if collecting {
				collecting = false
				// Send Message
				photoPath, err := getRandomFileFromDir(photoDIR)
				if err != nil {
					fmt.Println("Error getting random file:", err)
					return
				}
				sendMessage(client, ctx, photoPath, receptionist_ID, caption)
				messageCount = 0

			}
		}
		messageCount += 1
	}
}
