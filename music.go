package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/dhowden/tag"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

type Music struct {
	path  string
	isHot bool
}

func loadMusics(settings *Settings) []Music {
	musics := make([]Music, 0)
	err := filepath.WalkDir(settings.dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("[ERROR]: %v", err)
			return err
		}
		if !d.IsDir() {
			musics = append(musics, Music{path: path, isHot: false})
		}
		return nil
	})
	if err != nil {
		fmt.Printf("[ERROR]: Could not load the musics from %v\n%v\n", settings.dir, err)
		os.Exit(0)
	}
	if len(musics) == 0 {
		fmt.Println("[ERROR]: The specified directory has no any music to play.")
		os.Exit(0)
	}
	return musics
}

func decode(f *os.File) (*mp3.Stream, error) {
	stream, err := mp3.DecodeWithSampleRate(44100, f)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

func readMusicMetadata(file *os.File) (tag.Metadata, error) {
	metadata, err := tag.ReadFrom(file)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

func getMusicDuration(s *mp3.Stream) time.Duration {
	const sampleRate = 44100
	const sampleSize = 4
	samples := s.Length() / sampleSize
	duration := int(samples / int64(sampleRate))
	return time.Duration(duration * int(time.Second))
}

func checkMusicFinished(p *Player) bool {
	if p.duration-p.player.Position() > 0 {
		return false
	} else {
		return true
	}
}
