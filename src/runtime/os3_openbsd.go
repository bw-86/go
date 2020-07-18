// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"unsafe"
)

//go:cgo_import_dynamic libc_clock_gettime clock_gettime "libc.so"
//go:cgo_import_dynamic libc_exit exit "libc.so"
//go:cgo_import_dynamic libc_kevent kevent "libc.so"
//go:cgo_import_dynamic libc_kqueue kqueue "libc.so"
//go:cgo_import_dynamic libc_madvise madvise "libc.so"
//go:cgo_import_dynamic libc_mmap mmap "libc.so"
//go:cgo_import_dynamic libc_munmap munmap "libc.so"
//go:cgo_import_dynamic libc_open open "libc.so"
//go:cgo_import_dynamic libc_read read "libc.so"
//go:cgo_import_dynamic libc_sched_yield sched_yield "libc.so"
//go:cgo_import_dynamic libc_usleep usleep "libc.so"
//go:cgo_import_dynamic libc_write write "libc.so"
//go:cgo_import_dynamic libc_pipe pipe "libc.so"
//go:cgo_import_dynamic libc_pipe2 pipe2 "libc.so"

//go:linkname libc_clock_gettime libc_clock_gettime
//go:linkname libc_exit libc_exit
//go:linkname libc_kevent libc_kevent
//go:linkname libc_kqueue libc_kqueue
//go:linkname libc_madvise libc_madvise
//go:linkname libc_mmap libc_mmap
//go:linkname libc_munmap libc_munmap
//go:linkname libc_open libc_open
//go:linkname libc_read libc_read
//go:linkname libc_sched_yield libc_sched_yield
//go:linkname libc_usleep libc_usleep
//go:linkname libc_write libc_write
//go:linkname libc_pipe libc_pipe
//go:linkname libc_pipe2 libc_pipe2

var (
	libc_clock_gettime,
	libc_exit,
	libc_kevent,
	libc_kqueue,
	libc_madvise,
	libc_mmap,
	libc_munmap,
	libc_open,
	libc_read,
	libc_sched_yield,
	libc_usleep,
	libc_write,
	libc_pipe,
	libc_pipe2 libcFunc
)

func osyield1()

//go:nosplit
func osyield() {
	_g_ := getg()

	// Check the validity of m because we might be called in cgo callback
	// path early enough where there isn't a m available yet.
	if _g_ != nil && _g_.m != nil {
		sysvicall0(&libc_sched_yield)
		return
	}
	osyield1()
}

//go:nosplit
func nanotime1() int64 {
	var ts mts
	sysvicall2(&libc_clock_gettime, _CLOCK_MONOTONIC, uintptr(unsafe.Pointer(&ts)))
	return ts.tv_sec*1e9 + ts.tv_nsec
}

//go:nosplit
func mmap(addr unsafe.Pointer, n uintptr, prot, flags, fd int32, off uint32) (unsafe.Pointer, int) {
	p, err := doMmap(uintptr(addr), n, uintptr(prot), uintptr(flags), uintptr(fd), uintptr(off))
	if p == ^uintptr(0) {
		return nil, int(err)
	}
	return unsafe.Pointer(p), 0
}

//go:nosplit
func doMmap(addr, n, prot, flags, fd, off uintptr) (uintptr, uintptr) {
	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(&libc_mmap))
	libcall.n = 6
	libcall.args = uintptr(noescape(unsafe.Pointer(&addr)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	return libcall.r1, libcall.err
}

//go:nosplit
func munmap(addr unsafe.Pointer, n uintptr) {
	sysvicall2(&libc_munmap, uintptr(addr), uintptr(n))
}

//go:nosplit
func madvise(addr unsafe.Pointer, n uintptr, flags int32) {
	sysvicall3(&libc_madvise, uintptr(addr), uintptr(n), uintptr(flags))
}

//go:nosplit
func exit(r int32) {
	sysvicall1(&libc_exit, uintptr(r))
}

//go:nosplit
func pipe() (r, w int32, errno int32) {
	var p [2]int32
	_, e := sysvicall1Err(&libc_pipe, uintptr(noescape(unsafe.Pointer(&p))))
	return p[0], p[1], int32(e)
}

//go:nosplit
func pipe2(flags int32) (r, w int32, errno int32) {
	var p [2]int32
	_, e := sysvicall2Err(&libc_pipe2, uintptr(noescape(unsafe.Pointer(&p))), uintptr(flags))
	return p[0], p[1], int32(e)
}

//go:nosplit
func closeonexec(fd int32) {
	fcntl(fd, _F_SETFD, _FD_CLOEXEC)
}

//go:nosplit
func setNonblock(fd int32) {
	flags := fcntl(fd, _F_GETFL, 0)
	fcntl(fd, _F_SETFL, flags|_O_NONBLOCK)
}

//go:nosplit
func kqueue() int32 {
	return int32(sysvicall0(&libc_kqueue))
}

//go:nosplit
func kevent(kq int32, ch *keventt, nch int32, ev *keventt, nev int32, ts *timespec) int32 {
	return int32(sysvicall6(&libc_kevent, uintptr(kq), uintptr(unsafe.Pointer(ch)), uintptr(nch), uintptr(unsafe.Pointer(ev)), uintptr(nev), uintptr(unsafe.Pointer(ts))))
}

func fcntl(fd, cmd, arg int32) int32 {
	return int32(sysvicall3(&libc_fcntl, uintptr(fd), uintptr(cmd), uintptr(arg)))
}

func exitThread(wait *uint32) {
	// We should never reach exitThread on OpenBSD because we let
	// libc clean up threads.
	throw("exitThread")
}

//go:nosplit
func open(path *byte, mode, perm int32) int32 {
	return int32(sysvicall3(&libc_open, uintptr(unsafe.Pointer(path)), uintptr(mode), uintptr(perm)))
}

//go:nosplit
func read(fd int32, buf unsafe.Pointer, nbyte int32) int32 {
	r1, err := sysvicall3Err(&libc_read, uintptr(fd), uintptr(buf), uintptr(nbyte))
	if c := int32(r1); c >= 0 {
		return c
	}
	return -int32(err)
}

//go:nosplit
func closefd(fd int32) int32 {
	return int32(sysvicall1(&libc_close, uintptr(fd)))
}

func usleep1(usec uint32)

//go:nosplit
func usleep(µs uint32) {
	usleep1(µs)
}

//go:nosplit
func write1(fd uintptr, buf unsafe.Pointer, nbyte int32) int32 {
	r1, err := sysvicall3Err(&libc_write, fd, uintptr(buf), uintptr(nbyte))
	if c := int32(r1); c >= 0 {
		return c
	}
	return -int32(err)
}
