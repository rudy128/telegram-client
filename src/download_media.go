package src

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/downloader"
	"github.com/gotd/td/tg"
)

func download_media(client *telegram.Client, ctx context.Context, msg *tg.Message) {
	d := downloader.NewDownloader()

	inputLocation, err := getInputFileLocation(msg.Media)
	if err != nil {
		fmt.Println("Failed to extract file location:", err)
		return
	}
	mediaPath := filepath.Join(photoDIR, fmt.Sprintf("media_%d.jpg", time.Now().Unix()))
	builder := d.Download(client.API(), inputLocation)

	bytesWritten, err := builder.ToPath(ctx, mediaPath)
	if err != nil {
		fmt.Println("Failed to download media:", err)
	} else {
		fmt.Printf("Media downloaded to: %s (%d bytes)\n", mediaPath, bytesWritten)
	}
}
