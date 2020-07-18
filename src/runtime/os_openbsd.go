// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"runtime/internal/atomic"
	"runtime/internal/sys"
	"unsafe"
)

type mts struct {
	tv_sec  int64
	tv_nsec int64
}

type mOS struct {
	waitsemacount uint32
	perrno        *int32 // pointer to tls errno
}

type libcFunc uintptr

//go:linkname asmsysvicall6x runtime.asmsysvicall6
var asmsysvicall6x libcFunc // name to take addr of asmsysvicall6

func asmsysvicall6() // declared for vet; do NOT call

//go:nosplit
func sysvicall0(fn *libcFunc) uintptr {
	// Leave caller's PC/SP around for traceback.
	gp := getg()
	var mp *m
	if gp != nil {
		mp = gp.m
	}
	if mp != nil && mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = getcallerpc()
		// sp must be the last, because once async cpu profiler finds
		// all three values to be non-zero, it will use them
		mp.libcallsp = getcallersp()
	} else {
		mp = nil // See comment in sys_darwin.go:libcCall
	}

	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 0
	libcall.args = uintptr(unsafe.Pointer(fn)) // it's unused but must be non-nil, otherwise crashes
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	if mp != nil {
		mp.libcallsp = 0
	}
	return libcall.r1
}

//go:nosplit
func sysvicall1(fn *libcFunc, a1 uintptr) uintptr {
	r1, _ := sysvicall1Err(fn, a1)
	return r1
}

//go:nosplit

// sysvicall1Err returns both the system call result and the errno value.
// This is used by sysvicall1 and pipe.
func sysvicall1Err(fn *libcFunc, a1 uintptr) (r1, err uintptr) {
	// Leave caller's PC/SP around for traceback.
	gp := getg()
	var mp *m
	if gp != nil {
		mp = gp.m
	}
	if mp != nil && mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = getcallerpc()
		// sp must be the last, because once async cpu profiler finds
		// all three values to be non-zero, it will use them
		mp.libcallsp = getcallersp()
	} else {
		mp = nil
	}

	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 1
	// TODO(rsc): Why is noescape necessary here and below?
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	if mp != nil {
		mp.libcallsp = 0
	}
	return libcall.r1, libcall.err
}

//go:nosplit
func sysvicall2(fn *libcFunc, a1, a2 uintptr) uintptr {
	r1, _ := sysvicall2Err(fn, a1, a2)
	return r1
}

//go:nosplit
//go:cgo_unsafe_args

// sysvicall2Err returns both the system call result and the errno value.
// This is used by sysvicall2 and pipe2.
func sysvicall2Err(fn *libcFunc, a1, a2 uintptr) (uintptr, uintptr) {
	// Leave caller's PC/SP around for traceback.
	gp := getg()
	var mp *m
	if gp != nil {
		mp = gp.m
	}
	if mp != nil && mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = getcallerpc()
		// sp must be the last, because once async cpu profiler finds
		// all three values to be non-zero, it will use them
		mp.libcallsp = getcallersp()
	} else {
		mp = nil
	}

	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 2
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	if mp != nil {
		mp.libcallsp = 0
	}
	return libcall.r1, libcall.err
}

//go:nosplit
func sysvicall3(fn *libcFunc, a1, a2, a3 uintptr) uintptr {
	r1, _ := sysvicall3Err(fn, a1, a2, a3)
	return r1
}

//go:nosplit
//go:cgo_unsafe_args

// sysvicall3Err returns both the system call result and the errno value.
// This is used by sysicall3 and write1.
func sysvicall3Err(fn *libcFunc, a1, a2, a3 uintptr) (r1, err uintptr) {
	// Leave caller's PC/SP around for traceback.
	gp := getg()
	var mp *m
	if gp != nil {
		mp = gp.m
	}
	if mp != nil && mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = getcallerpc()
		// sp must be the last, because once async cpu profiler finds
		// all three values to be non-zero, it will use them
		mp.libcallsp = getcallersp()
	} else {
		mp = nil
	}

	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 3
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	if mp != nil {
		mp.libcallsp = 0
	}
	return libcall.r1, libcall.err
}

//go:nosplit
func sysvicall4(fn *libcFunc, a1, a2, a3, a4 uintptr) uintptr {
	// Leave caller's PC/SP around for traceback.
	gp := getg()
	var mp *m
	if gp != nil {
		mp = gp.m
	}
	if mp != nil && mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = getcallerpc()
		// sp must be the last, because once async cpu profiler finds
		// all three values to be non-zero, it will use them
		mp.libcallsp = getcallersp()
	} else {
		mp = nil
	}

	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 4
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	if mp != nil {
		mp.libcallsp = 0
	}
	return libcall.r1
}

//go:nosplit
func sysvicall5(fn *libcFunc, a1, a2, a3, a4, a5 uintptr) uintptr {
	// Leave caller's PC/SP around for traceback.
	gp := getg()
	var mp *m
	if gp != nil {
		mp = gp.m
	}
	if mp != nil && mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = getcallerpc()
		// sp must be the last, because once async cpu profiler finds
		// all three values to be non-zero, it will use them
		mp.libcallsp = getcallersp()
	} else {
		mp = nil
	}

	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 5
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	if mp != nil {
		mp.libcallsp = 0
	}
	return libcall.r1
}

//go:nosplit
func sysvicall6(fn *libcFunc, a1, a2, a3, a4, a5, a6 uintptr) uintptr {
	r1, _ := sysvicall6Err(fn, a1, a2, a3, a4, a5, a6)
	return r1
}

//go:nosplit
func sysvicall6Err(fn *libcFunc, a1, a2, a3, a4, a5, a6 uintptr) (r1, err uintptr) {
	// Leave caller's PC/SP around for traceback.
	gp := getg()
	var mp *m
	if gp != nil {
		mp = gp.m
	}
	if mp != nil && mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = getcallerpc()
		// sp must be the last, because once async cpu profiler finds
		// all three values to be non-zero, it will use them
		mp.libcallsp = getcallersp()
	} else {
		mp = nil
	}

	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 6
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	if mp != nil {
		mp.libcallsp = 0
	}
	return libcall.r1, libcall.err
}

//go:noescape
func setitimer(mode int32, new, old *itimerval)

//go:noescape
func sigaction(sig uint32, new, old *sigactiont)

//go:noescape
func sigaltstack(new, old *stackt)

//go:noescape
func obsdsigprocmask(how int32, new sigset) sigset

//go:nosplit
//go:nowritebarrierrec
func sigprocmask(how int32, new, old *sigset) {
	n := sigset(0)
	if new != nil {
		n = *new
	}
	r := obsdsigprocmask(how, n)
	if old != nil {
		*old = r
	}
}

func raiseproc(sig uint32)

func getthrid() int32

//go:noescape
func tfork(param *tforkt, psize uintptr, mm *m, gg *g, fn uintptr) int32

const (
	_ESRCH       = 3
	_EWOULDBLOCK = _EAGAIN
	_ENOTSUP     = 91

	// From OpenBSD's sys/time.h
	_CLOCK_REALTIME  = 0
	_CLOCK_VIRTUAL   = 1
	_CLOCK_PROF      = 2
	_CLOCK_MONOTONIC = 3
)

type sigset uint32

var sigset_all = ^sigset(0)

// From OpenBSD's <sys/sysctl.h>
const (
	_CTL_KERN   = 1
	_KERN_OSREV = 3

	_CTL_HW        = 6
	_HW_NCPU       = 3
	_HW_PAGESIZE   = 7
	_HW_NCPUONLINE = 25
)

func sysctlInt(mib []uint32) (int32, bool) {
	var out int32
	nout := unsafe.Sizeof(out)
	ret := sysctl(&mib[0], uint32(len(mib)), (*byte)(unsafe.Pointer(&out)), &nout, nil, 0)
	if ret < 0 {
		return 0, false
	}
	return out, true
}

func getncpu() int32 {
	// Try hw.ncpuonline first because hw.ncpu would report a number twice as
	// high as the actual CPUs running on OpenBSD 6.4 with hyperthreading
	// disabled (hw.smt=0). See https://golang.org/issue/30127
	if n, ok := sysctlInt([]uint32{_CTL_HW, _HW_NCPUONLINE}); ok {
		return int32(n)
	}
	if n, ok := sysctlInt([]uint32{_CTL_HW, _HW_NCPU}); ok {
		return int32(n)
	}
	return 1
}

func getPageSize() uintptr {
	if ps, ok := sysctlInt([]uint32{_CTL_HW, _HW_PAGESIZE}); ok {
		return uintptr(ps)
	}
	return 0
}

func getOSRev() int {
	if osrev, ok := sysctlInt([]uint32{_CTL_KERN, _KERN_OSREV}); ok {
		return int(osrev)
	}
	return 0
}

//go:nosplit
func semacreate(mp *m) {
}

//go:nosplit
func semasleep(ns int64) int32 {
	_g_ := getg()

	// Compute sleep deadline.
	var tsp *timespec
	if ns >= 0 {
		var ts timespec
		ts.setNsec(ns + nanotime())
		tsp = &ts
	}

	for {
		v := atomic.Load(&_g_.m.waitsemacount)
		if v > 0 {
			if atomic.Cas(&_g_.m.waitsemacount, v, v-1) {
				return 0 // semaphore acquired
			}
			continue
		}

		// Sleep until woken by semawakeup or timeout; or abort if waitsemacount != 0.
		//
		// From OpenBSD's __thrsleep(2) manual:
		// "The abort argument, if not NULL, points to an int that will
		// be examined [...] immediately before blocking. If that int
		// is non-zero then __thrsleep() will immediately return EINTR
		// without blocking."
		ret := thrsleep(uintptr(unsafe.Pointer(&_g_.m.waitsemacount)), _CLOCK_MONOTONIC, tsp, 0, &_g_.m.waitsemacount)
		if ret == _EWOULDBLOCK {
			return -1
		}
	}
}

//go:nosplit
func semawakeup(mp *m) {
	atomic.Xadd(&mp.waitsemacount, 1)
	ret := thrwakeup(uintptr(unsafe.Pointer(&mp.waitsemacount)), 1)
	if ret != 0 && ret != _ESRCH {
		// semawakeup can be called on signal stack.
		systemstack(func() {
			print("thrwakeup addr=", &mp.waitsemacount, " sem=", mp.waitsemacount, " ret=", ret, "\n")
		})
	}
}

// May run with m.p==nil, so write barriers are not allowed.
//go:nowritebarrier
func newosproc(mp *m) {
	stk := unsafe.Pointer(mp.g0.stack.hi)
	if false {
		print("newosproc stk=", stk, " m=", mp, " g=", mp.g0, " id=", mp.id, " ostk=", &mp, "\n")
	}

	// Stack pointer must point inside stack area (as marked with MAP_STACK),
	// rather than at the top of it.
	param := tforkt{
		tf_tcb:   unsafe.Pointer(&mp.tls[0]),
		tf_tid:   nil, // minit will record tid
		tf_stack: uintptr(stk) - sys.PtrSize,
	}

	var oset sigset
	sigprocmask(_SIG_SETMASK, &sigset_all, &oset)
	ret := tfork(&param, unsafe.Sizeof(param), mp, mp.g0, funcPC(mstart))
	sigprocmask(_SIG_SETMASK, &oset, nil)

	if ret < 0 {
		print("runtime: failed to create new OS thread (have ", mcount()-1, " already; errno=", -ret, ")\n")
		if ret == -_EAGAIN {
			println("runtime: may need to increase max user processes (ulimit -p)")
		}
		throw("runtime.newosproc")
	}
}

func osinit() {
	ncpu = getncpu()
	physPageSize = getPageSize()
	haveMapStack = getOSRev() >= 201805 // OpenBSD 6.3
}

var urandom_dev = []byte("/dev/urandom\x00")

//go:nosplit
func getRandomData(r []byte) {
	fd := open(&urandom_dev[0], 0 /* O_RDONLY */, 0)
	n := read(fd, unsafe.Pointer(&r[0]), int32(len(r)))
	closefd(fd)
	extendRandom(r, int(n))
}

func goenvs() {
	goenvs_unix()
}

// Called to initialize a new m (including the bootstrap m).
// Called on the parent thread (main thread in case of bootstrap), can allocate memory.
func mpreinit(mp *m) {
	mp.gsignal = malg(32 * 1024)
	mp.gsignal.m = mp
}

// Called to initialize a new m (including the bootstrap m).
// Called on the new thread, can not allocate memory.
func minit() {
	getg().m.procid = uint64(getthrid())
	minitSignals()
}

// Called from dropm to undo the effect of an minit.
//go:nosplit
func unminit() {
	unminitSignals()
}

func sigtramp()

type sigactiont struct {
	sa_sigaction uintptr
	sa_mask      uint32
	sa_flags     int32
}

//go:nosplit
//go:nowritebarrierrec
func setsig(i uint32, fn uintptr) {
	var sa sigactiont
	sa.sa_flags = _SA_SIGINFO | _SA_ONSTACK | _SA_RESTART
	sa.sa_mask = uint32(sigset_all)
	if fn == funcPC(sighandler) {
		fn = funcPC(sigtramp)
	}
	sa.sa_sigaction = fn
	sigaction(i, &sa, nil)
}

//go:nosplit
//go:nowritebarrierrec
func setsigstack(i uint32) {
	throw("setsigstack")
}

//go:nosplit
//go:nowritebarrierrec
func getsig(i uint32) uintptr {
	var sa sigactiont
	sigaction(i, nil, &sa)
	return sa.sa_sigaction
}

// setSignaltstackSP sets the ss_sp field of a stackt.
//go:nosplit
func setSignalstackSP(s *stackt, sp uintptr) {
	s.ss_sp = sp
}

//go:nosplit
//go:nowritebarrierrec
func sigaddset(mask *sigset, i int) {
	*mask |= 1 << (uint32(i) - 1)
}

func sigdelset(mask *sigset, i int) {
	*mask &^= 1 << (uint32(i) - 1)
}

//go:nosplit
func (c *sigctxt) fixsigcode(sig uint32) {
}

var haveMapStack = false

func osStackAlloc(s *mspan) {
	// OpenBSD 6.4+ requires that stacks be mapped with MAP_STACK.
	// It will check this on entry to system calls, traps, and
	// when switching to the alternate system stack.
	//
	// This function is called before s is used for any data, so
	// it's safe to simply re-map it.
	osStackRemap(s, _MAP_STACK)
}

func osStackFree(s *mspan) {
	// Undo MAP_STACK.
	osStackRemap(s, 0)
}

func osStackRemap(s *mspan, flags int32) {
	if !haveMapStack {
		// OpenBSD prior to 6.3 did not have MAP_STACK and so
		// the following mmap will fail. But it also didn't
		// require MAP_STACK (obviously), so there's no need
		// to do the mmap.
		return
	}
	a, err := mmap(unsafe.Pointer(s.base()), s.npages*pageSize, _PROT_READ|_PROT_WRITE, _MAP_PRIVATE|_MAP_ANON|_MAP_FIXED|flags, -1, 0)
	if err != 0 || uintptr(a) != s.base() {
		print("runtime: remapping stack memory ", hex(s.base()), " ", s.npages*pageSize, " a=", a, " err=", err, "\n")
		throw("remapping stack memory failed")
	}
}

//go:nosplit
func raise(sig uint32) {
	thrkill(getthrid(), int(sig))
}

func signalM(mp *m, sig int) {
	thrkill(int32(mp.procid), sig)
}
