package entity

import "github.com/o1egl/paseto"

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}
