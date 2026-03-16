package proxy

import (
	"context"
	"io"

	"github.com/gliderlabs/ssh"

	"github.com/atherlock/koko/pkg/exchange"
)

type UserConnection interface {
	io.ReadWriteCloser
	ID() string
	WinCh() <-chan ssh.Window
	LoginFrom() string
	RemoteAddr() string
	Pty() ssh.Pty
	Context() context.Context
	HandleRoomEvent(event string, msg *exchange.RoomMessage)
}