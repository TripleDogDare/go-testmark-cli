go-testmark-cli
---

A simple CLI for [go-testmark](https://github.com/warpfork/go-testmark).

This will allow you to list and extract testmark hunks from the command line.

This is experimental and the UI will likely change greatly if this is developed further.


Example list:

[testmark]:# (example-from-go-testmark-repository)
```
~/repos/misc/go-testmark (master) $ testmark list $(find -name '*.md')
./testexec/selfexercise.md:whee/fs/a:5
./testexec/selfexercise.md:whee/script:12
./testexec/selfexercise.md:whee/output:20
./testexec/selfexercise.md:whee/then-more-files/fs/b:30
./testexec/selfexercise.md:whee/then-more-files/script:35
./testexec/selfexercise.md:whee/then-more-files/output:40
./testexec/selfexercise.md:whee/then-touching-files/script:50
./testexec/selfexercise.md:whee/then-touching-files/output:59
./testexec/selfexercise.md:whee/then-touching-files/then-subtesting-again/script:69
./testexec/selfexercise.md:whee/then-touching-files/then-subtesting-again/output:75
./testexec/selfexercise.md:using-stdin/input:86
./testexec/selfexercise.md:using-stdin/script:92
./testexec/selfexercise.md:using-stdin/output:97
./testdata/exampleWithDirs.md:one/two:10
./testdata/exampleWithDirs.md:one/three:15
./testdata/exampleWithDirs.md:one:24
./testdata/exampleWithDirs.md:really/deep/dirs/wow:31
./testdata/exampleWithDirs.md:one/four/bang:38
./testdata/example.md:this-is-the-data-name:12
./testdata/example.md:more-data:35
./testdata/example.md:cannot-describe-no-linebreak:69
./README.md:this-is-testmark-btw:70
```

Example extract:

[testmark]:# (example-extract)
```
~/repos/misc/go-testmark (master) $ testmark extract this-is-testmark-btw $(find -name '*.md')
./README.md:this-is-testmark-btw:70
{"these things": "you know?",
 "syntax highlighed, typically": "etc, etc"}
```

