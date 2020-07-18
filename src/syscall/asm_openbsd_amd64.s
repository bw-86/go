// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

//
// System calls for openbsd/amd64 are implemented in ../runtime/syscall_openbsd.go
//

TEXT ·sysvicall6(SB),NOSPLIT,$0
	JMP	runtime·syscall_sysvicall6(SB)

TEXT ·rawSysvicall6(SB),NOSPLIT,$0
	JMP	runtime·syscall_rawsysvicall6(SB)
