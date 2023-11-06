package types

import (
	"fmt"
	"time"
)

type Environment struct {
	WebServerPort string
    RpcURL        string
    Mnemonic      string
	Env           string
	CommitHash    string
	CommitTime    time.Time
	Modified      bool
}

func (e Environment) BuildVersion() string {
	if e.CommitHash == "" {
		e.CommitHash = "0000000"
	}
	if e.CommitTime.IsZero() {
		e.CommitTime = time.Now()
	}

	return fmt.Sprintf("%s.%s.%s.%t", e.ShortCommitHash(), e.Env, e.CommitTime.Format("20060102150405"), e.Modified)
}

func (e Environment) ShortCommitHash() string {
	if len(e.CommitHash) < 7 {
		return e.CommitHash
	}

	return e.CommitHash[:7]
}
