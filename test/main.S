"".main STEXT size=154 args=0x0 locals=0x68
	0x0000 00000 (main.go:7)	TEXT	"".main(SB), ABIInternal, $104-0
	0x0000 00000 (main.go:7)	MOVQ	TLS, CX
	0x0009 00009 (main.go:7)	PCDATA	$0, $-2
	0x0009 00009 (main.go:7)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (main.go:7)	PCDATA	$0, $-1
	0x0010 00016 (main.go:7)	CMPQ	SP, 16(CX)
	0x0014 00020 (main.go:7)	PCDATA	$0, $-2
	0x0014 00020 (main.go:7)	JLS	144
	0x0016 00022 (main.go:7)	PCDATA	$0, $-1
	0x0016 00022 (main.go:7)	SUBQ	$104, SP
	0x001a 00026 (main.go:7)	MOVQ	BP, 96(SP)
	0x001f 00031 (main.go:7)	LEAQ	96(SP), BP
	0x0024 00036 (main.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0024 00036 (main.go:7)	FUNCDATA	$1, gclocals·26c19b003b4032a46d3e8db29831f3fe(SB)
	0x0024 00036 (main.go:7)	FUNCDATA	$3, "".main.stkobj(SB)
	0x0024 00036 (main.go:20)	XORPS	X0, X0
	0x0027 00039 (main.go:20)	MOVUPS	X0, ""..autotmp_0+56(SP)
	0x002c 00044 (main.go:20)	LEAQ	""..autotmp_0+56(SP), AX
	0x0031 00049 (main.go:20)	MOVQ	AX, ""..autotmp_2+48(SP)
	0x0036 00054 (main.go:20)	TESTB	AL, (AX)
	0x0038 00056 (main.go:20)	LEAQ	type.string(SB), CX
	0x003f 00063 (main.go:20)	MOVQ	CX, ""..autotmp_0+56(SP)
	0x0044 00068 (main.go:20)	LEAQ	""..stmp_0(SB), CX
	0x004b 00075 (main.go:20)	MOVQ	CX, ""..autotmp_0+64(SP)
	0x0050 00080 (main.go:20)	TESTB	AL, (AX)
	0x0052 00082 (main.go:20)	JMP	84
	0x0054 00084 (main.go:20)	MOVQ	AX, ""..autotmp_1+72(SP)
	0x0059 00089 (main.go:20)	MOVQ	$1, ""..autotmp_1+80(SP)
	0x0062 00098 (main.go:20)	MOVQ	$1, ""..autotmp_1+88(SP)
	0x006b 00107 (main.go:20)	MOVQ	AX, (SP)
	0x006f 00111 (main.go:20)	MOVQ	$1, 8(SP)
	0x0078 00120 (main.go:20)	MOVQ	$1, 16(SP)
	0x0081 00129 (main.go:20)	PCDATA	$1, $0
	0x0081 00129 (main.go:20)	CALL	fmt.Println(SB)
	0x0086 00134 (main.go:21)	MOVQ	96(SP), BP
	0x008b 00139 (main.go:21)	ADDQ	$104, SP
	0x008f 00143 (main.go:21)	RET
	0x0090 00144 (main.go:21)	NOP
	0x0090 00144 (main.go:7)	PCDATA	$1, $-1
	0x0090 00144 (main.go:7)	PCDATA	$0, $-2
	0x0090 00144 (main.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x0095 00149 (main.go:7)	PCDATA	$0, $-1
	0x0095 00149 (main.go:7)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 76 7a 48 83 ec 68 48 89 6c 24 60 48  H;a.vzH..hH.l$`H
	0x0020 8d 6c 24 60 0f 57 c0 0f 11 44 24 38 48 8d 44 24  .l$`.W...D$8H.D$
	0x0030 38 48 89 44 24 30 84 00 48 8d 0d 00 00 00 00 48  8H.D$0..H......H
	0x0040 89 4c 24 38 48 8d 0d 00 00 00 00 48 89 4c 24 40  .L$8H......H.L$@
	0x0050 84 00 eb 00 48 89 44 24 48 48 c7 44 24 50 01 00  ....H.D$HH.D$P..
	0x0060 00 00 48 c7 44 24 58 01 00 00 00 48 89 04 24 48  ..H.D$X....H..$H
	0x0070 c7 44 24 08 01 00 00 00 48 c7 44 24 10 01 00 00  .D$.....H.D$....
	0x0080 00 e8 00 00 00 00 48 8b 6c 24 60 48 83 c4 68 c3  ......H.l$`H..h.
	0x0090 e8 00 00 00 00 e9 66 ff ff ff                    ......f...
	rel 12+4 t=17 TLS+0
	rel 59+4 t=16 type.string+0
	rel 71+4 t=16 ""..stmp_0+0
	rel 130+4 t=8 fmt.Println+0
	rel 145+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.string."Hello world" SRODATA dupok size=11
	0x0000 48 65 6c 6c 6f 20 77 6f 72 6c 64                 Hello world
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 08 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
""..stmp_0 SRODATA size=16
	0x0000 00 00 00 00 00 00 00 00 0b 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."Hello world"+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·26c19b003b4032a46d3e8db29831f3fe SRODATA dupok size=9
	0x0000 01 00 00 00 06 00 00 00 00                       .........
"".main.stkobj SRODATA size=24
	0x0000 01 00 00 00 00 00 00 00 d8 ff ff ff ff ff ff ff  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
	rel 16+8 t=1 type.[1]interface {}+0
