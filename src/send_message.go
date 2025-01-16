package src

import (
	"context"
	"fmt"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/message/html"
	"github.com/gotd/td/telegram/uploader"
)

func sendMessage(client *telegram.Client, ctx context.Context, photoPath string, chatID int64, caption string) {
	u := uploader.NewUploader(client.API())
	fmt.Println("Uploading photo:", photoPath)
	upload, err := u.FromPath(ctx, photoPath)
	if err != nil {
		fmt.Println("Failed to upload photo:", err)
		return
	}

	// Create the media with a caption
	document := message.UploadedPhoto(upload, html.String(nil, caption))

	// Use a message sender instance
	sender := message.NewSender(client.API()).WithUploader(u)

	target := sender.Resolve(fmt.Sprintf("@%s", username))

	// Send the photo with the caption
	fmt.Println("Sending photo to chat...")
	// if _, err := sender.To(inputPeer).Media(ctx, document); err != nil {
	// 	fmt.Println("Failed to send photo:", err)
	// 	return
	// }
	if _, err := target.Media(ctx, document); err != nil {
		fmt.Println("send: %w", err)
		return
	}

	fmt.Println("Photo sent successfully with caption!")
	EmptyDir(photoDIR)
}
