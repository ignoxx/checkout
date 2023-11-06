package utils

import (
	"errors"
	"runtime/debug"
	"time"
)

func GetBuildInfo() (string, time.Time, string, error) {
	var revision, modified string
	var revisionTime time.Time

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", time.Time{}, "", errors.New("failed to read build info")
	}
	for _, kv := range info.Settings {
		if kv.Value == "" {
			continue
		}
		switch kv.Key {
		case "vcs.revision":
			revision = kv.Value
		case "vcs.time":
			revisionTime, _ = time.Parse(time.RFC3339, kv.Value)
		case "vcs.modified":
			modified = kv.Value
		}
	}

	return revision, revisionTime, modified, nil
}
