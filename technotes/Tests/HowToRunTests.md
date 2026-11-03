[TOC]

# Running Tests in Project ***errprefops***

Open a command prompt in this directory (***MikeAustin71/errprefops/errpref***) and run the
following commands.

## Run Basic Tests

Open a command prompt in this directory:
 ***MikeAustin71/errpref***

and run the following command:

***go test -v > zzzzz_tests.txt***

This will generate a '.txt' file in the current directory which 
contains all the test results. 

The **-v** option signals a request for verbose output.

## Test Execution WITHOUT Code Coverage
Run these test commands in the ***MikeAustin71/errprefops/errpref*** directory:

### Windows Command Without Coverage
```go
go test -v > zzzzz_tests.txt
```



### Linux Command Without Coverage
```go
go test -v | tee zzzzz_tests.txt
```

This will generate test results in the ***MikeAustin71/errprefops/errpref***
directory. The test results will be stored in the text file, ***MikeAustin71/errprefops/errpref/zzzzz_tests.txt***. 



## Running Tests with code coverage

First, pull down and install the `cover` package.

```go
go get golang.org/x/tools/cmd/cover
```

Next, follow the test execution protocol.  

Run these test commands in the ***MikeAustin71/errprefops/errpref*** directory:



## Test Execution WITH Code Coverage

Run this in the ***MikeAustin71/errprefops/errpref*** directory:

### Windows Command With Coverage
```go
go test -cover -v > zzzzz_tests.txt

// Tests for race conditions
go test -cover -v -race > zzzzz_tests.txt
```

### Linux Command With Coverage
```go
go test -cover -v | tee zzzzz_tests.txt
```




## Cover Profile

Generate the code coverage detail. Run this command
in the ***MikeAustin71/errprefops/errpref*** directory:

```go
go test -coverprofile=zzzzz_coverage.out
```

This will output results to file: **`zzzzz_coverage.out`**

The following provides for code coverage display in your
browser. Run this on the terminal command line and run it
in the ***MikeAustin71/errprefops/errpref*** directory:

```go
go tool cover -html=zzzzz_coverage.out
```

This will output results to file: **`zzzzz_coverage.out`**