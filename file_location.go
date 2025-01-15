package main

import (
	"github.com/go-faster/errors"
	"github.com/gotd/td/tg"
)

func getInputFileLocation(media tg.MessageMediaClass) (tg.InputFileLocationClass, error) {
	switch m := media.(type) {
	case *tg.MessageMediaPhoto:
		if photo, ok := m.Photo.(*tg.Photo); ok && len(photo.Sizes) > 0 {
			return &tg.InputPhotoFileLocation{
				ID:            photo.ID,
				AccessHash:    photo.AccessHash,
				FileReference: photo.FileReference,
				ThumbSize:     "w", // Example size; use "x" for the largest.
			}, nil
		}
	}
	return nil, errors.New("unsupported media type")
}
