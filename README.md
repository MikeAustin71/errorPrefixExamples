# Error Prefix Examples

**errorPrefixExamples** - This application provides usage examples for the ***errpref*** software package. 

***errpref***  stands for **Error Prefix**.  The ***errpref*** package provides lightweight ***GoLang*** types designed to create and attach error prefix text, function chain lists and error context strings to error messages returned by ***Go*** functions.

The source code repository for the ***errpref*** package is located at  https://github.com/MikeAustin71/errpref. Source code documentation is maintained at [errpref · pkg.go.dev](https://pkg.go.dev/github.com/MikeAustin71/errpref).

The ***errpref*** package is written in the [The Go Programming Language (golang.org)](https://golang.org/) (a.k.a ***Golang***).

The source code for the **errorPrefixExamples** is located in the software repository at https://github.com/MikeAustin71/errorPrefixExamples .



## Purpose

The goal of this example code is to assist developers in quickly integrating the ***errpref*** package into their software projects. Feel free to clone this **errorPrefixExamples** repository and review the code examples. 



## Setup, Configuration And Import



### A Good Place To Start

If you have not done so, it is recommended that you first review the source code documentation for [errpref](https://pkg.go.dev/github.com/MikeAustin71/errpref).



### Setup and Configuration

The **errorPrefixExamples** application is designed to work with [***Go*** Modules](https://golang.org/ref/mod). This application therefore includes a **go.mod** file. A review of the go.mod file, located in the base **errorPrefixExamples** directory, will show the ***errpref*** package configured for the latest version.

```go

module github.com/MikeAustin71/errorPrefixExamples

go 1.16

require github.com/MikeAustin71/errpref v1.6.1

```



### Import

To use ***errpref***  in your code, you will need to import the software package ***github.com/MikeAustin71/errpref***.

```go

import (

 erPref "github.com/MikeAustin71/errpref"
)

```

The **errorPrefixExamples** application has already included in the necessary ***import*** statements.



### Troubleshooting *errpref* Package Setup

If problems are encountered configuring the ***errpref*** package, try the following:

1. From the command line run: 

   ```text
   go get github.com/MikeAustin71/errpref@v1.6.1 
   
   				or 
   
   go get github.com/MikeAustin71/errpref@latest
   ```

2. If problems continue, as a last resort, try clearing your cache:

   ```text
   go clean -modcache
   
   Re-Execute the 'go get' command in item No. 1 above.
   ```

   Be advised, the command will delete all modules in your cache. 

   

## Executing Example Code

### Background

All the test functions are located in ***testFunctions/testFunctionsDto.go*** and ***testFunctions/testFunctionsStrings.go***. They are designed to be called from source code file ***app/main.go***. 

The best usage examples are located at:

***examplesDto/examplesDtoAlpha.go***

***examplesDto/examplesDtoBravo.go***

***examplesDto/examplesDtoCharlie.go***

***examplesDto/examplesDtoDelta.go***

***examplesDto/examplesDtoEcho.go***

***examplesDto/examplesIBuilder.go***



### *main.go*

1. First configure ***app/main.go*** to call the desired test function from Type  ***MainTest*** also located in ***app/main.go***. Then configure function ***main()*** to call a specific test like ***MainTest.mainTest001()***, ***MainTest.mainTest002()***, **MainTest.mainTest003()** etc.
2. From the command line in directory ***errorPrefixExamples\app***, issue the command **go run main.go**. Test results will then be displayed in the command line interface. 



## Version

The latest version of the **errorPrefixExamples** application is Version 1.6.1.  This corresponds to the version number of the latest version of the ***errpref*** software package.



## License

Use of this source code is governed by the (open-source) MIT-style license which can be found in the LICENSE file
located in this directory.

[MIT License](./LICENSE)



## Comments And Questions

Send questions or comments to:

``` text
    mike.go@paladinacs.net
```