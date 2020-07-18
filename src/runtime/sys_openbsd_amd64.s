// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// System calls and other sys.stuff for AMD64, OpenBSD
// /usr/src/sys/kern/syscalls.master for syscall numbers.
//

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"

// This is needed by asm_amd64.s
TEXT runtime·settls(SB),NOSPLIT,$8
	RET

// Call a library function with SysV calling conventions.
// The called function can take a maximum of 6 INTEGER class arguments,
// see
//   Michael Matz, Jan Hubicka, Andreas Jaeger, and Mark Mitchell
//   System V Application Binary Interface
//   AMD64 Architecture Processor Supplement
// section 3.2.3.
//
// Called by runtime·asmcgocall or runtime·cgocall.
// NOT USING GO CALLING CONVENTION.
TEXT runtime·asmsysvicall6(SB),NOSPLIT,$0
	// asmcgocall will put first argument into DI.
	PUSHQ	DI			// save for later
	MOVQ	libcall_fn(DI), AX
	MOVQ	libcall_args(DI), R11
	MOVQ	libcall_n(DI), R10

	get_tls(CX)
	MOVQ	g(CX), BX
	CMPQ	BX, $0
	JEQ	skiperrno1
	MOVQ	g_m(BX), BX
	MOVQ	(m_mOS+mOS_perrno)(BX), DX
	CMPQ	DX, $0
	JEQ	skiperrno1
	MOVL	$0, 0(DX)

skiperrno1:
	CMPQ	R11, $0
	JEQ	skipargs
	// Load 6 args into correspondent registers.
	MOVQ	0(R11), DI
	MOVQ	8(R11), SI
	MOVQ	16(R11), DX
	MOVQ	24(R11), CX
	MOVQ	32(R11), R8
	MOVQ	40(R11), R9
skipargs:

	// Call SysV function
	CALL	AX

	// Return result
	POPQ	DI
	MOVQ	AX, libcall_r1(DI)
	MOVQ	DX, libcall_r2(DI)

	get_tls(CX)
	MOVQ	g(CX), BX
	CMPQ	BX, $0
	JEQ	skiperrno2
	MOVQ	g_m(BX), BX
	MOVQ	(m_mOS+mOS_perrno)(BX), AX
	CMPQ	AX, $0
	JEQ	skiperrno2
	MOVL	0(AX), AX
	MOVQ	AX, libcall_err(DI)

skiperrno2:
	RET


TEXT runtime·sigfwd(SB),NOSPLIT,$0-32
	MOVQ	fn+0(FP),    AX
	MOVL	sig+8(FP),   DI
	MOVQ	info+16(FP), SI
	MOVQ	ctx+24(FP),  DX
	PUSHQ	BP
	MOVQ	SP, BP
	ANDQ	$~15, SP     // alignment for x86_64 ABI
	CALL	AX
	MOVQ	BP, SP
	POPQ	BP
	RET

TEXT runtime·sigtramp(SB),NOSPLIT,$72
	// Save callee-saved C registers, since the caller may be a C signal handler.
	MOVQ	BX,  bx-8(SP)
	MOVQ	BP,  bp-16(SP)  // save in case GOEXPERIMENT=noframepointer is set
	MOVQ	R12, r12-24(SP)
	MOVQ	R13, r13-32(SP)
	MOVQ	R14, r14-40(SP)
	MOVQ	R15, r15-48(SP)
	// We don't save mxcsr or the x87 control word because sigtrampgo doesn't
	// modify them.

	MOVQ	DX, ctx-56(SP)
	MOVQ	SI, info-64(SP)
	MOVQ	DI, signum-72(SP)
	CALL	runtime·sigtrampgo(SB)

	MOVQ	r15-48(SP), R15
	MOVQ	r14-40(SP), R14
	MOVQ	r13-32(SP), R13
	MOVQ	r12-24(SP), R12
	MOVQ	bp-16(SP),  BP
	MOVQ	bx-8(SP),   BX
	RET

// Runs on OS stack, called from runtime·osyield.
TEXT runtime·osyield1(SB),NOSPLIT,$0
	LEAQ	libc_sched_yield(SB), AX
	CALL	AX
	RET

// Called from runtime·usleep (Go). Can be called on Go stack, on OS stack,
// can also be called in cgo callback path without a g->m.
TEXT runtime·usleep1(SB),NOSPLIT,$0
	MOVL	usec+0(FP), DI
	MOVQ	$usleep2<>(SB), AX // to hide from 6l

	// Execute call on m->g0.
	get_tls(R15)
	CMPQ	R15, $0
	JE	noswitch

	MOVQ	g(R15), R13
	CMPQ	R13, $0
	JE	noswitch
	MOVQ	g_m(R13), R13
	CMPQ	R13, $0
	JE	noswitch
	// TODO(aram): do something about the cpu profiler here.

	MOVQ	m_g0(R13), R14
	CMPQ	g(R15), R14
	JNE	switch
	// executing on m->g0 already
	CALL	AX
	RET

switch:
	// Switch to m->g0 stack and back.
	MOVQ	(g_sched+gobuf_sp)(R14), R14
	MOVQ	SP, -8(R14)
	LEAQ	-8(R14), SP
	CALL	AX
	MOVQ	0(SP), SP
	RET

noswitch:
	// Not a Go-managed thread. Do not switch stack.
	CALL	AX
	RET

// Runs on OS stack. duration (in µs units) is in DI.
TEXT usleep2<>(SB),NOSPLIT,$0
	LEAQ	libc_usleep(SB), AX
	CALL	AX
	RET

