//go:build !noasm && amd64
// AUTO-GENERATED BY GOAT -- DO NOT EDIT

TEXT ·dot_256(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ res+16(FP), DX
	MOVQ len+24(FP), CX
	BYTE $0x55               // pushq	%rbp
	WORD $0x8948; BYTE $0xe5 // movq	%rsp, %rbp
	LONG $0xf8e48348         // andq	$-8, %rsp
	WORD $0x8b4c; BYTE $0x09 // movq	(%rcx), %r9
	LONG $0x07f98341         // cmpl	$7, %r9d
	JG   LBB0_9
	LONG $0xff418d45         // leal	-1(%r9), %r8d
	LONG $0x03c1f641         // testb	$3, %r9b
	JE   LBB0_2
	WORD $0x8944; BYTE $0xc8 // movl	%r9d, %eax
	WORD $0xe083; BYTE $0x03 // andl	$3, %eax
	LONG $0xc057f8c5         // vxorps	%xmm0, %xmm0, %xmm0
	WORD $0xc931             // xorl	%ecx, %ecx

LBB0_4:
	LONG $0x0f10fac5             // vmovss	(%rdi), %xmm1                   # xmm1 = mem[0],zero,zero,zero
	LONG $0xb971e2c4; BYTE $0x06 // vfmadd231ss	(%rsi), %xmm1, %xmm0    # xmm0 = (xmm1 * mem) + xmm0
	LONG $0x04c78348             // addq	$4, %rdi
	LONG $0x04c68348             // addq	$4, %rsi
	LONG $0x01c18348             // addq	$1, %rcx
	WORD $0xc839                 // cmpl	%ecx, %eax
	JNE  LBB0_4
	WORD $0x2941; BYTE $0xc9     // subl	%ecx, %r9d
	LONG $0x03f88341             // cmpl	$3, %r8d
	JAE  LBB0_7

LBB0_31:
	LONG $0x0211fac5         // vmovss	%xmm0, (%rdx)
	WORD $0x8948; BYTE $0xec // movq	%rbp, %rsp
	BYTE $0x5d               // popq	%rbp
	BYTE $0xc3               // retq

LBB0_9:
	LONG $0xc057f8c5               // vxorps	%xmm0, %xmm0, %xmm0
	LONG $0x20f98341               // cmpl	$32, %r9d
	JB   LBB0_10
	LONG $0xe0498d41               // leal	-32(%r9), %ecx
	WORD $0xc1f6; BYTE $0x20       // testb	$32, %cl
	JNE  LBB0_12
	LONG $0x1f10fcc5               // vmovups	(%rdi), %ymm3
	LONG $0x5710fcc5; BYTE $0x20   // vmovups	32(%rdi), %ymm2
	LONG $0x4f10fcc5; BYTE $0x40   // vmovups	64(%rdi), %ymm1
	LONG $0x4710fcc5; BYTE $0x60   // vmovups	96(%rdi), %ymm0
	LONG $0xe457d8c5               // vxorps	%xmm4, %xmm4, %xmm4
	LONG $0x985de2c4; BYTE $0x1e   // vfmadd132ps	(%rsi), %ymm4, %ymm3    # ymm3 = (ymm3 * mem) + ymm4
	LONG $0x985de2c4; WORD $0x2056 // vfmadd132ps	32(%rsi), %ymm4, %ymm2  # ymm2 = (ymm2 * mem) + ymm4
	LONG $0x985de2c4; WORD $0x404e // vfmadd132ps	64(%rsi), %ymm4, %ymm1  # ymm1 = (ymm1 * mem) + ymm4
	LONG $0x985de2c4; WORD $0x6046 // vfmadd132ps	96(%rsi), %ymm4, %ymm0  # ymm0 = (ymm0 * mem) + ymm4
	LONG $0x80ef8348               // subq	$-128, %rdi
	LONG $0x80ee8348               // subq	$-128, %rsi
	WORD $0x8941; BYTE $0xc9       // movl	%ecx, %r9d
	WORD $0xf983; BYTE $0x20       // cmpl	$32, %ecx
	JAE  LBB0_20
	JMP  LBB0_15

LBB0_10:
	LONG $0xc957f0c5 // vxorps	%xmm1, %xmm1, %xmm1
	LONG $0xd257e8c5 // vxorps	%xmm2, %xmm2, %xmm2
	LONG $0xdb57e0c5 // vxorps	%xmm3, %xmm3, %xmm3
	JMP  LBB0_16

LBB0_2:
	LONG $0xc057f8c5 // vxorps	%xmm0, %xmm0, %xmm0
	LONG $0x03f88341 // cmpl	$3, %r8d
	JB   LBB0_31

LBB0_7:
	WORD $0x8944; BYTE $0xc8 // movl	%r9d, %eax
	WORD $0xc931             // xorl	%ecx, %ecx

LBB0_8:
	LONG $0x0c10fac5; BYTE $0x8f               // vmovss	(%rdi,%rcx,4), %xmm1            # xmm1 = mem[0],zero,zero,zero
	LONG $0x5410fac5; WORD $0x048f             // vmovss	4(%rdi,%rcx,4), %xmm2           # xmm2 = mem[0],zero,zero,zero
	LONG $0x9979e2c4; WORD $0x8e0c             // vfmadd132ss	(%rsi,%rcx,4), %xmm0, %xmm1 # xmm1 = (xmm1 * mem) + xmm0
	LONG $0xb969e2c4; WORD $0x8e4c; BYTE $0x04 // vfmadd231ss	4(%rsi,%rcx,4), %xmm2, %xmm1 # xmm1 = (xmm2 * mem) + xmm1
	LONG $0x5410fac5; WORD $0x088f             // vmovss	8(%rdi,%rcx,4), %xmm2           # xmm2 = mem[0],zero,zero,zero
	LONG $0x9971e2c4; WORD $0x8e54; BYTE $0x08 // vfmadd132ss	8(%rsi,%rcx,4), %xmm1, %xmm2 # xmm2 = (xmm2 * mem) + xmm1
	LONG $0x4410fac5; WORD $0x0c8f             // vmovss	12(%rdi,%rcx,4), %xmm0          # xmm0 = mem[0],zero,zero,zero
	LONG $0x9969e2c4; WORD $0x8e44; BYTE $0x0c // vfmadd132ss	12(%rsi,%rcx,4), %xmm2, %xmm0 # xmm0 = (xmm0 * mem) + xmm2
	LONG $0x04c18348                           // addq	$4, %rcx
	WORD $0xc839                               // cmpl	%ecx, %eax
	JNE  LBB0_8
	JMP  LBB0_31

LBB0_12:
	LONG $0xc057f8c5         // vxorps	%xmm0, %xmm0, %xmm0
	LONG $0xc957f0c5         // vxorps	%xmm1, %xmm1, %xmm1
	LONG $0xd257e8c5         // vxorps	%xmm2, %xmm2, %xmm2
	LONG $0xdb57e0c5         // vxorps	%xmm3, %xmm3, %xmm3
	WORD $0xf983; BYTE $0x20 // cmpl	$32, %ecx
	JB   LBB0_15

LBB0_20:
	LONG $0x2710fcc5                           // vmovups	(%rdi), %ymm4
	LONG $0x6f10fcc5; BYTE $0x20               // vmovups	32(%rdi), %ymm5
	LONG $0x7710fcc5; BYTE $0x40               // vmovups	64(%rdi), %ymm6
	LONG $0x7f10fcc5; BYTE $0x60               // vmovups	96(%rdi), %ymm7
	LONG $0x9865e2c4; BYTE $0x26               // vfmadd132ps	(%rsi), %ymm3, %ymm4    # ymm4 = (ymm4 * mem) + ymm3
	LONG $0x986de2c4; WORD $0x206e             // vfmadd132ps	32(%rsi), %ymm2, %ymm5  # ymm5 = (ymm5 * mem) + ymm2
	LONG $0x9875e2c4; WORD $0x4076             // vfmadd132ps	64(%rsi), %ymm1, %ymm6  # ymm6 = (ymm6 * mem) + ymm1
	LONG $0x987de2c4; WORD $0x607e             // vfmadd132ps	96(%rsi), %ymm0, %ymm7  # ymm7 = (ymm7 * mem) + ymm0
	QUAD $0x000000809f10fcc5                   // vmovups	128(%rdi), %ymm3
	QUAD $0x000000a09710fcc5                   // vmovups	160(%rdi), %ymm2
	QUAD $0x000000c08f10fcc5                   // vmovups	192(%rdi), %ymm1
	QUAD $0x000000e08710fcc5                   // vmovups	224(%rdi), %ymm0
	QUAD $0x0000809e985de2c4; BYTE $0x00       // vfmadd132ps	128(%rsi), %ymm4, %ymm3 # ymm3 = (ymm3 * mem) + ymm4
	QUAD $0x0000a0969855e2c4; BYTE $0x00       // vfmadd132ps	160(%rsi), %ymm5, %ymm2 # ymm2 = (ymm2 * mem) + ymm5
	QUAD $0x0000c08e984de2c4; BYTE $0x00       // vfmadd132ps	192(%rsi), %ymm6, %ymm1 # ymm1 = (ymm1 * mem) + ymm6
	QUAD $0x0000e0869845e2c4; BYTE $0x00       // vfmadd132ps	224(%rsi), %ymm7, %ymm0 # ymm0 = (ymm0 * mem) + ymm7
	LONG $0xc0c18341                           // addl	$-64, %r9d
	LONG $0x00c78148; WORD $0x0001; BYTE $0x00 // addq	$256, %rdi                      # imm = 0x100
	LONG $0x00c68148; WORD $0x0001; BYTE $0x00 // addq	$256, %rsi                      # imm = 0x100
	LONG $0x1ff98341                           // cmpl	$31, %r9d
	JA   LBB0_20
	WORD $0x8944; BYTE $0xc9                   // movl	%r9d, %ecx

LBB0_15:
	WORD $0x8941; BYTE $0xc9 // movl	%ecx, %r9d
	WORD $0xf983; BYTE $0x08 // cmpl	$8, %ecx
	JB   LBB0_18

LBB0_16:
	WORD $0x8944; BYTE $0xc9 // movl	%r9d, %ecx

LBB0_17:
	LONG $0x2710fcc5             // vmovups	(%rdi), %ymm4
	LONG $0xb85de2c4; BYTE $0x1e // vfmadd231ps	(%rsi), %ymm4, %ymm3    # ymm3 = (ymm4 * mem) + ymm3
	WORD $0xc183; BYTE $0xf8     // addl	$-8, %ecx
	LONG $0x20c78348             // addq	$32, %rdi
	LONG $0x20c68348             // addq	$32, %rsi
	WORD $0xf983; BYTE $0x07     // cmpl	$7, %ecx
	JA   LBB0_17

LBB0_18:
	WORD $0xc985             // testl	%ecx, %ecx
	JE   LBB0_19
	LONG $0xff418d44         // leal	-1(%rcx), %r8d
	WORD $0xc1f6; BYTE $0x03 // testb	$3, %cl
	JE   LBB0_23
	WORD $0x8941; BYTE $0xc9 // movl	%ecx, %r9d
	LONG $0x03e18341         // andl	$3, %r9d
	LONG $0xe457d8c5         // vxorps	%xmm4, %xmm4, %xmm4
	WORD $0xc031             // xorl	%eax, %eax

LBB0_25:
	LONG $0x2f10fac5             // vmovss	(%rdi), %xmm5                   # xmm5 = mem[0],zero,zero,zero
	LONG $0xb951e2c4; BYTE $0x26 // vfmadd231ss	(%rsi), %xmm5, %xmm4    # xmm4 = (xmm5 * mem) + xmm4
	LONG $0x04c78348             // addq	$4, %rdi
	LONG $0x04c68348             // addq	$4, %rsi
	LONG $0x01c08348             // addq	$1, %rax
	WORD $0x3941; BYTE $0xc1     // cmpl	%eax, %r9d
	JNE  LBB0_25
	WORD $0xc129                 // subl	%eax, %ecx
	LONG $0x03f88341             // cmpl	$3, %r8d
	JAE  LBB0_28
	JMP  LBB0_30

LBB0_19:
	LONG $0xe457d8c5 // vxorps	%xmm4, %xmm4, %xmm4
	JMP  LBB0_30

LBB0_23:
	LONG $0xe457d8c5 // vxorps	%xmm4, %xmm4, %xmm4
	LONG $0x03f88341 // cmpl	$3, %r8d
	JB   LBB0_30

LBB0_28:
	WORD $0xc889 // movl	%ecx, %eax
	WORD $0xc931 // xorl	%ecx, %ecx

LBB0_29:
	LONG $0x2c10fac5; BYTE $0x8f               // vmovss	(%rdi,%rcx,4), %xmm5            # xmm5 = mem[0],zero,zero,zero
	LONG $0x7410fac5; WORD $0x048f             // vmovss	4(%rdi,%rcx,4), %xmm6           # xmm6 = mem[0],zero,zero,zero
	LONG $0x9959e2c4; WORD $0x8e2c             // vfmadd132ss	(%rsi,%rcx,4), %xmm4, %xmm5 # xmm5 = (xmm5 * mem) + xmm4
	LONG $0xb949e2c4; WORD $0x8e6c; BYTE $0x04 // vfmadd231ss	4(%rsi,%rcx,4), %xmm6, %xmm5 # xmm5 = (xmm6 * mem) + xmm5
	LONG $0x7410fac5; WORD $0x088f             // vmovss	8(%rdi,%rcx,4), %xmm6           # xmm6 = mem[0],zero,zero,zero
	LONG $0x9951e2c4; WORD $0x8e74; BYTE $0x08 // vfmadd132ss	8(%rsi,%rcx,4), %xmm5, %xmm6 # xmm6 = (xmm6 * mem) + xmm5
	LONG $0x6410fac5; WORD $0x0c8f             // vmovss	12(%rdi,%rcx,4), %xmm4          # xmm4 = mem[0],zero,zero,zero
	LONG $0x9949e2c4; WORD $0x8e64; BYTE $0x0c // vfmadd132ss	12(%rsi,%rcx,4), %xmm6, %xmm4 # xmm4 = (xmm4 * mem) + xmm6
	LONG $0x04c18348                           // addq	$4, %rcx
	WORD $0xc839                               // cmpl	%ecx, %eax
	JNE  LBB0_29

LBB0_30:
	LONG $0xd358ecc5               // vaddps	%ymm3, %ymm2, %ymm2
	LONG $0xc058f4c5               // vaddps	%ymm0, %ymm1, %ymm0
	LONG $0xc258fcc5               // vaddps	%ymm2, %ymm0, %ymm0
	LONG $0xc07cffc5               // vhaddps	%ymm0, %ymm0, %ymm0
	LONG $0xc07cffc5               // vhaddps	%ymm0, %ymm0, %ymm0
	LONG $0x197de3c4; WORD $0x01c1 // vextractf128	$1, %ymm0, %xmm1
	LONG $0xc158fac5               // vaddss	%xmm1, %xmm0, %xmm0
	LONG $0xc058dac5               // vaddss	%xmm0, %xmm4, %xmm0
	LONG $0x0211fac5               // vmovss	%xmm0, (%rdx)
	WORD $0x8948; BYTE $0xec       // movq	%rbp, %rsp
	BYTE $0x5d                     // popq	%rbp
	WORD $0xf8c5; BYTE $0x77       // vzeroupper
	BYTE $0xc3                     // retq
