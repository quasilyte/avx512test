// Code generated by avx512test. DO NOT EDIT.

#include "../../../../../../runtime/textflag.h"

TEXT asmtest_avx512_vbmi(SB), NOSPLIT, $0
	//TODO: VPERMB X13, X7, K1, X15                            // 625245098dfd
	//TODO: VPERMB 7(SI)(DI*4), X7, K1, X15                    // 627245098dbcbe07000000
	//TODO: VPERMB -7(DI)(R8*2), X7, K1, X15                   // 623245098dbc47f9ffffff
	//TODO: VPERMB Y18, Y13, K7, Y30                           // 6222152f8df2
	//TODO: VPERMB 17(SP)(BP*1), Y13, K7, Y30                  // 6262152f8db42c11000000
	//TODO: VPERMB -7(CX)(DX*8), Y13, K7, Y30                  // 6262152f8db4d1f9ffffff
	//TODO: VPERMB Z3, Z8, K1, Z3                              // 62f23d498ddb
	//TODO: VPERMB Z27, Z8, K1, Z3                             // 62923d498ddb
	//TODO: VPERMB 7(AX), Z8, K1, Z3                           // 62f23d498d9807000000
	//TODO: VPERMB (DI), Z8, K1, Z3                            // 62f23d498d1f
	//TODO: VPERMB Z3, Z2, K1, Z3                              // 62f26d498ddb
	//TODO: VPERMB Z27, Z2, K1, Z3                             // 62926d498ddb
	//TODO: VPERMB 7(AX), Z2, K1, Z3                           // 62f26d498d9807000000
	//TODO: VPERMB (DI), Z2, K1, Z3                            // 62f26d498d1f
	//TODO: VPERMB Z3, Z8, K1, Z21                             // 62e23d498deb
	//TODO: VPERMB Z27, Z8, K1, Z21                            // 62823d498deb
	//TODO: VPERMB 7(AX), Z8, K1, Z21                          // 62e23d498da807000000
	//TODO: VPERMB (DI), Z8, K1, Z21                           // 62e23d498d2f
	//TODO: VPERMB Z3, Z2, K1, Z21                             // 62e26d498deb
	//TODO: VPERMB Z27, Z2, K1, Z21                            // 62826d498deb
	//TODO: VPERMB 7(AX), Z2, K1, Z21                          // 62e26d498da807000000
	//TODO: VPERMB (DI), Z2, K1, Z21                           // 62e26d498d2f
	//TODO: VPERMI2B X7, X24, K7, X8                           // 62723d0775c7
	//TODO: VPERMI2B 17(SP), X24, K7, X8                       // 62723d0775842411000000
	//TODO: VPERMI2B -17(BP)(SI*4), X24, K7, X8                // 62723d077584b5efffffff
	//TODO: VPERMI2B Y24, Y5, K2, Y24                          // 6202552a75c0
	//TODO: VPERMI2B 15(R8)(R14*1), Y5, K2, Y24                // 6202552a7584300f000000
	//TODO: VPERMI2B 15(R8)(R14*2), Y5, K2, Y24                // 6202552a7584700f000000
	//TODO: VPERMI2B Z12, Z9, K4, Z3                           // 62d2354c75dc
	//TODO: VPERMI2B Z22, Z9, K4, Z3                           // 62b2354c75de
	//TODO: VPERMI2B -17(BP)(SI*8), Z9, K4, Z3                 // 62f2354c759cf5efffffff
	//TODO: VPERMI2B (R15), Z9, K4, Z3                         // 62d2354c751f
	//TODO: VPERMI2B Z12, Z19, K4, Z3                          // 62d2654475dc
	//TODO: VPERMI2B Z22, Z19, K4, Z3                          // 62b2654475de
	//TODO: VPERMI2B -17(BP)(SI*8), Z19, K4, Z3                // 62f26544759cf5efffffff
	//TODO: VPERMI2B (R15), Z19, K4, Z3                        // 62d26544751f
	//TODO: VPERMI2B Z12, Z9, K4, Z30                          // 6242354c75f4
	//TODO: VPERMI2B Z22, Z9, K4, Z30                          // 6222354c75f6
	//TODO: VPERMI2B -17(BP)(SI*8), Z9, K4, Z30                // 6262354c75b4f5efffffff
	//TODO: VPERMI2B (R15), Z9, K4, Z30                        // 6242354c7537
	//TODO: VPERMI2B Z12, Z19, K4, Z30                         // 6242654475f4
	//TODO: VPERMI2B Z22, Z19, K4, Z30                         // 6222654475f6
	//TODO: VPERMI2B -17(BP)(SI*8), Z19, K4, Z30               // 6262654475b4f5efffffff
	//TODO: VPERMI2B (R15), Z19, K4, Z30                       // 624265447537
	//TODO: VPERMT2B X6, X16, K7, X11                          // 62727d077dde
	//TODO: VPERMT2B (BX), X16, K7, X11                        // 62727d077d1b
	//TODO: VPERMT2B -17(BP)(SI*1), X16, K7, X11               // 62727d077d9c35efffffff
	//TODO: VPERMT2B Y16, Y17, K6, Y27                         // 622275267dd8
	//TODO: VPERMT2B 7(SI)(DI*4), Y17, K6, Y27                 // 626275267d9cbe07000000
	//TODO: VPERMT2B -7(DI)(R8*2), Y17, K6, Y27                // 622275267d9c47f9ffffff
	//TODO: VPERMT2B Z20, Z1, K3, Z6                           // 62b2754b7df4
	//TODO: VPERMT2B Z9, Z1, K3, Z6                            // 62d2754b7df1
	//TODO: VPERMT2B (CX), Z1, K3, Z6                          // 62f2754b7d31
	//TODO: VPERMT2B 99(R15), Z1, K3, Z6                       // 62d2754b7db763000000
	//TODO: VPERMT2B Z20, Z9, K3, Z6                           // 62b2354b7df4
	//TODO: VPERMT2B Z9, Z9, K3, Z6                            // 62d2354b7df1
	//TODO: VPERMT2B (CX), Z9, K3, Z6                          // 62f2354b7d31
	//TODO: VPERMT2B 99(R15), Z9, K3, Z6                       // 62d2354b7db763000000
	//TODO: VPERMT2B Z20, Z1, K3, Z9                           // 6232754b7dcc
	//TODO: VPERMT2B Z9, Z1, K3, Z9                            // 6252754b7dc9
	//TODO: VPERMT2B (CX), Z1, K3, Z9                          // 6272754b7d09
	//TODO: VPERMT2B 99(R15), Z1, K3, Z9                       // 6252754b7d8f63000000
	//TODO: VPERMT2B Z20, Z9, K3, Z9                           // 6232354b7dcc
	//TODO: VPERMT2B Z9, Z9, K3, Z9                            // 6252354b7dc9
	//TODO: VPERMT2B (CX), Z9, K3, Z9                          // 6272354b7d09
	//TODO: VPERMT2B 99(R15), Z9, K3, Z9                       // 6252354b7d8f63000000
	//TODO: VPMULTISHIFTQB X0, X1, K5, X8                      // 6272f50d83c0
	//TODO: VPMULTISHIFTQB 17(SP)(BP*1), X1, K5, X8            // 6272f50d83842c11000000
	//TODO: VPMULTISHIFTQB -7(CX)(DX*8), X1, K5, X8            // 6272f50d8384d1f9ffffff
	//TODO: VPMULTISHIFTQB Y2, Y24, K7, Y3                     // 62f2bd2783da
	//TODO: VPMULTISHIFTQB 17(SP)(BP*2), Y24, K7, Y3           // 62f2bd27839c6c11000000
	//TODO: VPMULTISHIFTQB -7(DI)(R8*4), Y24, K7, Y3           // 62b2bd27839c87f9ffffff
	//TODO: VPMULTISHIFTQB Z7, Z2, K7, Z18                     // 62e2ed4f83d7
	//TODO: VPMULTISHIFTQB Z13, Z2, K7, Z18                    // 62c2ed4f83d5
	//TODO: VPMULTISHIFTQB 7(AX)(CX*4), Z2, K7, Z18            // 62e2ed4f83948807000000
	//TODO: VPMULTISHIFTQB 7(AX)(CX*1), Z2, K7, Z18            // 62e2ed4f83940807000000
	//TODO: VPMULTISHIFTQB Z7, Z21, K7, Z18                    // 62e2d54783d7
	//TODO: VPMULTISHIFTQB Z13, Z21, K7, Z18                   // 62c2d54783d5
	//TODO: VPMULTISHIFTQB 7(AX)(CX*4), Z21, K7, Z18           // 62e2d54783948807000000
	//TODO: VPMULTISHIFTQB 7(AX)(CX*1), Z21, K7, Z18           // 62e2d54783940807000000
	//TODO: VPMULTISHIFTQB Z7, Z2, K7, Z24                     // 6262ed4f83c7
	//TODO: VPMULTISHIFTQB Z13, Z2, K7, Z24                    // 6242ed4f83c5
	//TODO: VPMULTISHIFTQB 7(AX)(CX*4), Z2, K7, Z24            // 6262ed4f83848807000000
	//TODO: VPMULTISHIFTQB 7(AX)(CX*1), Z2, K7, Z24            // 6262ed4f83840807000000
	//TODO: VPMULTISHIFTQB Z7, Z21, K7, Z24                    // 6262d54783c7
	//TODO: VPMULTISHIFTQB Z13, Z21, K7, Z24                   // 6242d54783c5
	//TODO: VPMULTISHIFTQB 7(AX)(CX*4), Z21, K7, Z24           // 6262d54783848807000000
	//TODO: VPMULTISHIFTQB 7(AX)(CX*1), Z21, K7, Z24           // 6262d54783840807000000
	RET
