package papipes

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Source struct {
	Filename   string
	Name       string
	Properties map[string]interface{}
	Format     string
	Rate       int
	Channels   int
	// channelMap
	file        *os.File
	moduleIndex int
}

func (s *Source) Open() error {
	var err error

	if !filepath.IsAbs(s.Filename) {
		return errors.New("Filename is not absolue patch")
	}

	args := make([]string, 0)
	args = append(args, "load-module")
	args = append(args, "module-pipe-source")
	args = append(args, fmt.Sprintf("file=%s", s.Filename))
	if s.Name != "" {
		args = append(args, fmt.Sprintf("source_name=%s", s.Name))
	}
	if s.Format != "" {
		args = append(args, fmt.Sprintf("format=%s", s.Format))
	}
	if s.Rate > 0 {
		args = append(args, fmt.Sprintf("rate=%d", s.Rate))
	}
	if s.Channels > 0 {
		args = append(args, fmt.Sprintf("channels=%d", s.Channels))
	}

	var props string

	for k, v := range s.Properties {
		props = props + fmt.Sprintf("%s='%v'", k, v)
	}

	args = append(args, fmt.Sprintf("source_properties=\"%s\"", props))

	cmd := exec.Command("pactl", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	if _, err := fmt.Sscanf(string(out), "%d", &s.moduleIndex); err != nil {
		return err
	}

	if s.file, err = os.OpenFile(s.Filename, os.O_RDWR, 0755); err != nil {
		return err
	}

	return err
}

func (s *Source) Close() error {
	if err := s.file.Close(); err != nil {
		return err
	}

	args := make([]string, 0)
	args = append(args, "unload-module")
	args = append(args, fmt.Sprintf("%d", s.moduleIndex))

	cmd := exec.Command("pactl", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	return nil
}

func (s *Source) Write(p []byte) (n int, err error) {
	return s.file.Write(p)
}