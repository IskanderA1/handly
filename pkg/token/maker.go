package token

import "time"

type Maker interface {
	CreateToken(projectName string, duration time.Duration) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}
