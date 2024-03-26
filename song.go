package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/dhowden/tag"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

type Song struct {
	path string
}

func loadSongs(settings *Settings) []*Song {
	songs := []*Song{}
	_ = filepath.WalkDir(settings.dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("[ERROR]: %v", err)
			return err
		}
		if !d.IsDir() {
			songs = append(songs, &Song{path: path})
		}
		return nil
	})
	return songs
}

func readSongMetadata(file *os.File) (tag.Metadata, error) {
	metadata, err := tag.ReadFrom(file)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

func getSongDuration(s *mp3.Stream) int {
	const sampleRate = 44100
	const sampleSize = 4
	samples := s.Length() / sampleSize
	duration := int(samples / int64(sampleRate))
	return duration
}
