# avx512test

Utility that was used to generate initial [Go AVX-512 encoder test suite](https://github.com/golang/go/tree/master/src/cmd/asm/internal/asm/testdata/avx512enc).

Note that it should probably not be used for new tests generation.

Published as a response to request made in <https://github.com/golang/go/issues/25724>.

## Requirements / Preparations

* [Intel XED](https://github.com/intelxed/xed) of 088c48a2efa447872945168272bcd7005a7ddd91 revision.
  Not guaranteed to work with newer XED.
  Recommended to be in sync with [xeddata](https://github.com/golang/arch/blob/master/x86/xeddata/doc.go#L48).

* Latest `x86.csv` file. Ideally, either x86csv generator will evolve to use XED or we will
  avoid using `x86.csv` in future. Right now, we only have the old version and one that was generated
  by [uncommited generator](https://go-review.googlesource.com/c/arch/+/104496/) ([x86.csv](/x86.csv)).

To install XED, do something like this:

```sh
# Go to some folder where you would like to have XED cloned.

git clone https://github.com/intelxed/xed.git xed
git clone https://github.com/intelxed/mbuild.git mbuild
cd xed
./mfile.py --shared install

# Now install the generated `libxed.so`.
# It's located under "./kits/xed-install-$blah/lib/libxed.so".
# (You may want to run ldconfig.)
```

You don't need XED includes as I vendored the relevant headers.

If you're running Linux on AMD64 machine, you may even take a look at `libxed.so` in the root
of this repository. If it works, you may skip the compilation step, but you still need to
make sure that relevant shared library can be loaded.

## Running the generator

If everything is OK, single command is enough to generate everything:

```sh
export GOPATH=$(go env GOPATH)

cd $GOPATH/github.com/quasilyte/avx512test

go install ./cmd/avx512test/main.go

$GOPATH/bin/avx512test
```

By default, it expects to find `x86.csv` inside current directory.
Use `-x86csv` parameter to set other location.

Produced output is written to `./output` directory.
Normally, its file list may look like:

```sh
$ cd output
$ ls
aes_avx512f.s    avx512cd.s     avx512pf.s          gfni_avx512f.s
avx512_4fmaps.s  avx512dq.s     avx512_vbmi2.s      vpclmulqdq_avx512f.s
avx512_4vnniw.s  avx512er.s     avx512_vbmi.s
avx512_bitalg.s  avx512f.s      avx512_vnni.s
avx512bw.s       avx512_ifma.s  avx512_vpopcntdq.s
```

## How does it work

TODO: describe how does avx512test works.
