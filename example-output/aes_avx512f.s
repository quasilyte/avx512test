// Code generated by avx512test. DO NOT EDIT.

#include "../../../../../../runtime/textflag.h"

TEXT asmtest_aes_avx512f(SB), NOSPLIT, $0
	//TODO: VAESDEC Y22, Y9, Y8                                // 62323528dec6 or 6232b528dec6
	//TODO: VAESDEC Z27, Z3, Z11                               // 62126548dedb or 6212e548dedb
	//TODO: VAESDEC Z15, Z3, Z11                               // 62526548dedf or 6252e548dedf
	//TODO: VAESDEC 99(R15)(R15*1), Z3, Z11                    // 62126548de9c3f63000000 or 6212e548de9c3f63000000
	//TODO: VAESDEC (DX), Z3, Z11                              // 62726548de1a or 6272e548de1a
	//TODO: VAESDEC Z27, Z12, Z11                              // 62121d48dedb or 62129d48dedb
	//TODO: VAESDEC Z15, Z12, Z11                              // 62521d48dedf or 62529d48dedf
	//TODO: VAESDEC 99(R15)(R15*1), Z12, Z11                   // 62121d48de9c3f63000000 or 62129d48de9c3f63000000
	//TODO: VAESDEC (DX), Z12, Z11                             // 62721d48de1a or 62729d48de1a
	//TODO: VAESDEC Z27, Z3, Z25                               // 62026548decb or 6202e548decb
	//TODO: VAESDEC Z15, Z3, Z25                               // 62426548decf or 6242e548decf
	//TODO: VAESDEC 99(R15)(R15*1), Z3, Z25                    // 62026548de8c3f63000000 or 6202e548de8c3f63000000
	//TODO: VAESDEC (DX), Z3, Z25                              // 62626548de0a or 6262e548de0a
	//TODO: VAESDEC Z27, Z12, Z25                              // 62021d48decb or 62029d48decb
	//TODO: VAESDEC Z15, Z12, Z25                              // 62421d48decf or 62429d48decf
	//TODO: VAESDEC 99(R15)(R15*1), Z12, Z25                   // 62021d48de8c3f63000000 or 62029d48de8c3f63000000
	//TODO: VAESDEC (DX), Z12, Z25                             // 62621d48de0a or 62629d48de0a
	//TODO: VAESDECLAST Z8, Z23, Z23                           // 62c24540dff8 or 62c2c540dff8
	//TODO: VAESDECLAST Z28, Z23, Z23                          // 62824540dffc or 6282c540dffc
	//TODO: VAESDECLAST -17(BP)(SI*8), Z23, Z23                // 62e24540dfbcf5efffffff or 62e2c540dfbcf5efffffff
	//TODO: VAESDECLAST (R15), Z23, Z23                        // 62c24540df3f or 62c2c540df3f
	//TODO: VAESDECLAST Z8, Z6, Z23                            // 62c24d48dff8 or 62c2cd48dff8
	//TODO: VAESDECLAST Z28, Z6, Z23                           // 62824d48dffc or 6282cd48dffc
	//TODO: VAESDECLAST -17(BP)(SI*8), Z6, Z23                 // 62e24d48dfbcf5efffffff or 62e2cd48dfbcf5efffffff
	//TODO: VAESDECLAST (R15), Z6, Z23                         // 62c24d48df3f or 62c2cd48df3f
	//TODO: VAESDECLAST Z8, Z23, Z5                            // 62d24540dfe8 or 62d2c540dfe8
	//TODO: VAESDECLAST Z28, Z23, Z5                           // 62924540dfec or 6292c540dfec
	//TODO: VAESDECLAST -17(BP)(SI*8), Z23, Z5                 // 62f24540dfacf5efffffff or 62f2c540dfacf5efffffff
	//TODO: VAESDECLAST (R15), Z23, Z5                         // 62d24540df2f or 62d2c540df2f
	//TODO: VAESDECLAST Z8, Z6, Z5                             // 62d24d48dfe8 or 62d2cd48dfe8
	//TODO: VAESDECLAST Z28, Z6, Z5                            // 62924d48dfec or 6292cd48dfec
	//TODO: VAESDECLAST -17(BP)(SI*8), Z6, Z5                  // 62f24d48dfacf5efffffff or 62f2cd48dfacf5efffffff
	//TODO: VAESDECLAST (R15), Z6, Z5                          // 62d24d48df2f or 62d2cd48df2f
	//TODO: VAESENC X0, X21, X16                               // 62e25500dcc0 or 62e2d500dcc0
	//TODO: VAESENC 99(R15)(R15*8), X21, X16                   // 62825500dc84ff63000000 or 6282d500dc84ff63000000
	//TODO: VAESENC 7(AX)(CX*8), X21, X16                      // 62e25500dc84c807000000 or 62e2d500dc84c807000000
	//TODO: VAESENC Z12, Z16, Z21                              // 62c27d40dcec or 62c2fd40dcec
	//TODO: VAESENC Z27, Z16, Z21                              // 62827d40dceb or 6282fd40dceb
	//TODO: VAESENC 7(SI)(DI*8), Z16, Z21                      // 62e27d40dcacfe07000000 or 62e2fd40dcacfe07000000
	//TODO: VAESENC -15(R14), Z16, Z21                         // 62c27d40dcaef1ffffff or 62c2fd40dcaef1ffffff
	//TODO: VAESENC Z12, Z13, Z21                              // 62c21548dcec or 62c29548dcec
	//TODO: VAESENC Z27, Z13, Z21                              // 62821548dceb or 62829548dceb
	//TODO: VAESENC 7(SI)(DI*8), Z13, Z21                      // 62e21548dcacfe07000000 or 62e29548dcacfe07000000
	//TODO: VAESENC -15(R14), Z13, Z21                         // 62c21548dcaef1ffffff or 62c29548dcaef1ffffff
	//TODO: VAESENC Z12, Z16, Z5                               // 62d27d40dcec or 62d2fd40dcec
	//TODO: VAESENC Z27, Z16, Z5                               // 62927d40dceb or 6292fd40dceb
	//TODO: VAESENC 7(SI)(DI*8), Z16, Z5                       // 62f27d40dcacfe07000000 or 62f2fd40dcacfe07000000
	//TODO: VAESENC -15(R14), Z16, Z5                          // 62d27d40dcaef1ffffff or 62d2fd40dcaef1ffffff
	//TODO: VAESENC Z12, Z13, Z5                               // 62d21548dcec or 62d29548dcec
	//TODO: VAESENC Z27, Z13, Z5                               // 62921548dceb or 62929548dceb
	//TODO: VAESENC 7(SI)(DI*8), Z13, Z5                       // 62f21548dcacfe07000000 or 62f29548dcacfe07000000
	//TODO: VAESENC -15(R14), Z13, Z5                          // 62d21548dcaef1ffffff or 62d29548dcaef1ffffff
	//TODO: VAESENCLAST X7, X22, X28                           // 62624d00dde7 or 6262cd00dde7
	//TODO: VAESENCLAST (AX), X22, X28                         // 62624d00dd20 or 6262cd00dd20
	//TODO: VAESENCLAST 7(SI), X22, X28                        // 62624d00dda607000000 or 6262cd00dda607000000
	//TODO: VAESENCLAST Y31, Y19, Y0                           // 62926520ddc7 or 6292e520ddc7
	//TODO: VAESENCLAST 7(SI)(DI*1), Y19, Y0                   // 62f26520dd843e07000000 or 62f2e520dd843e07000000
	//TODO: VAESENCLAST 15(DX)(BX*8), Y19, Y0                  // 62f26520dd84da0f000000 or 62f2e520dd84da0f000000
	//TODO: VAESENCLAST Z25, Z6, Z22                           // 62824d48ddf1 or 6282cd48ddf1
	//TODO: VAESENCLAST Z12, Z6, Z22                           // 62c24d48ddf4 or 62c2cd48ddf4
	//TODO: VAESENCLAST 7(SI)(DI*1), Z6, Z22                   // 62e24d48ddb43e07000000 or 62e2cd48ddb43e07000000
	//TODO: VAESENCLAST 15(DX)(BX*8), Z6, Z22                  // 62e24d48ddb4da0f000000 or 62e2cd48ddb4da0f000000
	//TODO: VAESENCLAST Z25, Z8, Z22                           // 62823d48ddf1 or 6282bd48ddf1
	//TODO: VAESENCLAST Z12, Z8, Z22                           // 62c23d48ddf4 or 62c2bd48ddf4
	//TODO: VAESENCLAST 7(SI)(DI*1), Z8, Z22                   // 62e23d48ddb43e07000000 or 62e2bd48ddb43e07000000
	//TODO: VAESENCLAST 15(DX)(BX*8), Z8, Z22                  // 62e23d48ddb4da0f000000 or 62e2bd48ddb4da0f000000
	//TODO: VAESENCLAST Z25, Z6, Z11                           // 62124d48ddd9 or 6212cd48ddd9
	//TODO: VAESENCLAST Z12, Z6, Z11                           // 62524d48dddc or 6252cd48dddc
	//TODO: VAESENCLAST 7(SI)(DI*1), Z6, Z11                   // 62724d48dd9c3e07000000 or 6272cd48dd9c3e07000000
	//TODO: VAESENCLAST 15(DX)(BX*8), Z6, Z11                  // 62724d48dd9cda0f000000 or 6272cd48dd9cda0f000000
	//TODO: VAESENCLAST Z25, Z8, Z11                           // 62123d48ddd9 or 6212bd48ddd9
	//TODO: VAESENCLAST Z12, Z8, Z11                           // 62523d48dddc or 6252bd48dddc
	//TODO: VAESENCLAST 7(SI)(DI*1), Z8, Z11                   // 62723d48dd9c3e07000000 or 6272bd48dd9c3e07000000
	//TODO: VAESENCLAST 15(DX)(BX*8), Z8, Z11                  // 62723d48dd9cda0f000000 or 6272bd48dd9cda0f000000
	RET
