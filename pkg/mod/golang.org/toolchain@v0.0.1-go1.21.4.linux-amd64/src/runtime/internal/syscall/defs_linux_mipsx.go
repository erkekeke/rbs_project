// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (mips || mipsle)

package syscall

const (
	SYS_FCNTL         = 4055
	SYS_EPOLL_CTL     = 4249
	SYS_EPOLL_PWAIT   = 4313
	SYS_EPOLL_CREATE1 = 4326
	SYS_EPOLL_PWAIT2  = 4441

	EPOLLIN       = 0x1
	EPOLLOUT      = 0x4
	EPOLLERR      = 0x8
	EPOLLHUP      = 0x10
	EPOLLRDHUP    = 0x2000
	EPOLLET       = 0x80000000
	EPOLL_CLOEXEC = 0x80000
	EPOLL_CTL_ADD = 0x1
	EPOLL_CTL_DEL = 0x2
	EPOLL_CTL_MOD = 0x3
)

type EpollEvent struct {
	Events    uint32
	pad_cgo_0 [4]byte
	Data      uint64
}
