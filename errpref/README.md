# Error Prefix (errpref) - Adding error prefix, function chains and error context to Error Messages 

**The *errpref* software package is a collection of lightweight GoLang types designed to attach error prefix text, function chain lists and error context strings to error messages.**

The ***errpref*** software package was written in the [Go](https://golang.org/) programming language, a.k.a. ***Golang***.

***errpref*** supports [Go Modules]([Go Modules Reference - The Go Programming Language (golang.org)](https://golang.org/ref/mod)).

The current version of ***errpref*** is Version 1.6.0



## Table of Contents

 - [The Problem](#the-problem)
 - [The Solution](#the-solution)
 - [Definition of Terms](#definition-of-terms)
   - [Error Prefix](#error-prefix)
   - [Error Context](#error-context)
 - [String Formatting Conventions](#string-formatting-conventions)
   - [Error Prefix Delimiters](#error-prefix-delimiters)
   - [Error Context Delimiters](#error-context-delimiters)
 - [Getting Started](#getting-started)
   - [Two Types To Choose From](#two-types-to-choose-from) 
   - [ErrPrefixDto - A Full Featured Solution](#errprefixdto---a-full-featured-solution)
     - [Public Facing Methods](#public-facing-methods)
     - [Internal or Private Methods](#internal-or-private-methods)
     - [Error Context Example](#error-context-example)
     - [Code Examples](#code-examples)
       - [Test Code](#test-code)
       - [Example Application](#Example-application)
   - [ErrPref - Quick And Simple Solution](#errpref---quick-and-simple-solution)
     - [Example Usage Summary](#example-usage-summary)
 - [More Examples](#more-examples)
 - [ErrPref Package Configuration and Import](#errpref-package-configuration-and-import)
   - [Go Get Command](#go-get-command)
   - [Import Configuration](#import-configuration)
 - [Source Code Documentation](#source-code-documentation)
 - [Version](#version)
 - [License](#license)
 - [Comments And Questions](#comments-and-questions)



## The Problem

As my **Go** programs, types and methods have grown in sophistication and complexity, the need for improved error tracking, detailed error messages, and a record of code execution has become more important. The terms 'Function Chains' or 'Method Chains' as used here describes a list of the functions called prior to encountering a particular error. Adding function chain documentation to returned error messages makes error tracking and management a much easier proposition.

Basically, when I see an error message, especially during the development phase, the first question that comes to mind is, **How the Hell did I get here?** Answering that question quickly, and with certainty, usually requires a list of functions ordered by their execution sequence.  Adding that information to error messages is the focus of this project.



## The Solution

The ***errpref*** or Error Prefix project is intended to provide better function or method documentation in error messages returned by Go functions. Two types have been defined for this purpose: ***ErrPref*** and ***ErrPrefixDto***. Both of these types are designed to receive function chain information, format it and return the formatted strings for inclusion in error messages.

The idea is not to go overboard. I tried my hand at an error handler for Go Programs earlier and quickly realized that the complexity was growing exponentially. At least at the outset of this project, the idea is to keep the ***errpref*** package simple, as a fast and efficient mechanism for adding error prefix text and function chain lists to error messages.

### Wait a Minute - Let's Think About This

So, if you only take a quick look at the ***errpref*** package you might be tempted to say, **"Whoopee - all you've done is concatenate a series of strings!"**

That's what I thought, at first. However, there is the problem of delimiting error prefixes and error context information plus the problem of determining how much information to show on each line. 

For types ***ErrPref*** and  ***ErrPrefixDto***, the user may configure the line length which will determine whether a new line character (**\n**) or in-line delimiter string (**[SPACE]-[SPACE]**) is used to separate error prefix information.

To summarize: If you are serious about error management and tracking, ***ErrPref*** might be worth a second look.



## Definition of Terms

For our purposes, error prefix information consists of two elements: a function name, and an optional error context string.

### Error Prefix

**Error Prefix** text is designed to be configured at the beginning of error messages and is most often used to document the thread of code execution by listing the calling sequence for a specific list of functions and/or methods.

### Error Context

**Error Context** strings are designed to provide additional information about the function or method identified by the
associated **Error Prefix** text. Typical context information might include variable names, variable values and additional details on function execution.

Error Context strings are optional and are always associated with an Error Prefix string.



## String Formatting Conventions

When using the methods provided by types ***ErrPref*** and ***ErrPrefixDto*** it's best to use the following formatting
conventions:

### Error Prefix Delimiters

Error Prefix Elements are delimited by one of two delimiters: new line character (**\n**) or in-line delimiter string (**[SPACE]-[SPACE]**).

### Error Context Delimiters

Associated Error Context Sub-Elements likewise have two delimiters: new line (**\n[SPACE]:[SPACE]\[SPACE]**) or in-line delimiter **([SPACE]:[SPACE]**)



## Getting Started

### Two Types To Choose From

Currently, there are two principal types providing error prefix formatting services: ***ErrPrefixDto*** and ***ErrPref***. Type ***ErrPref*** offers basic error prefix formatting while type ***ErrPrefixDto*** covers a much wider range of formatting and data transmission capabilities.

### ErrPrefixDto - A Full Featured Solution

The Error Prefix Data Transfer Object, ***ErrPrefixDto***, offers the same services as type, ***ErrPref***, but is packaged in a different architecture. While ***ErrPref*** methods receive a string and instantly return a formatted string, ***ErrPrefixDto*** encapsulates error prefix information in an internal array of Error Prefix Information objects (***ErrorPrefixInfo***). Strings are only created when the type's ***String()*** method is called. Instances of ***ErrPrefixDto*** are designed to be passed as input parameters to subsidiary methods.

The ***String()*** method has a value receiver.  Virtually all other ***ErrPrefixDto*** methods have pointer receivers.

#### Public Facing Methods
A public method may receive error prefix information in a variety of formats from sources both internal and external to the package. To date, the best use scenario follows this pattern:

``` go
func (numStrBasic *NumStrBasic) GetCountryFormatters(
  fmtCollection *FormatterCollection,
  countryCulture CountryCultureId,
  errorPrefix interface{}) error {

  var ePrefix *ErrPrefixDto
  var err error

  ePrefix,
    err = ErrPrefixDto{}.NewIEmpty(
    errorPrefix,
    "NumStrBasic.GetCountryFormatters()",
    "")

  if err != nil {
    return err
  }

  return numStrBasicQuark{}.ptr().
    getCountryFormatters(
      fmtCollection,
      countryCulture,
      ePrefix)
}
```

By calling ***ErrPrefixDto{}.NewIEmpty()*** with an empty interface, functions can pass error prefix information in any of 10-different formats:

  1. **nil** - A nil value is valid and generates an empty collection of error prefix and error context information.
  2. **Stringer** - The Stringer interface from the 'fmt' package. This interface has only one method:

``` go
    	type Stringer interface {
    		String() string
    	}
```

  3. **string** - A string containing error prefix information.
  5. **\[\]string** - A one-dimensional slice of strings containing error prefix information.
  1. **\[\]\[2\]string** - A two-dimensional slice of strings containing error prefix and error context information.
  2. **strings.Builder** - An instance of strings.Builder. Error prefix information will be imported into the new returned instance of ErrPrefixDto.
  3. **\*strings.Builder** - A pointer to an instance of strings.Builder.  Error prefix information will be imported into the new returned instance of ErrPrefixDto.
  4. **ErrPrefixDto** - An instance of **ErrPrefixDto**. The **ErrorPrefixInfo** from this object will be copied to the new returned instance of **ErrPrefixDto*.
  6.  **\*ErrPrefixDto** - A pointer to an instance of **ErrPrefixDto**. ErrorPrefixInfo from this object will be copied to the new returned instance of **ErrPrefixDto*.
  7.   **IBasicErrorPrefix** - An interface to a method generating a two-dimensional slice of strings containing error prefix and error context information.

``` go
             type IBasicErrorPrefix interface {
               GetEPrefStrings() [][2]string
             }

```

The Method Signature for ***ErrPrefixDto{}.NewIEmpty()*** is shown below:

``` go
func (ePrefDto ErrPrefixDto) NewIEmpty(
	iEPref interface{},
	newErrPrefix string,
	newErrContext string) (
	*ErrPrefixDto,
	error)
```

This pattern allows for use of the 'nil' value. This means that parameter ***iEPref*** will accept a 'nil' value. If no error prefix information is present or required, just pass a 'nil' value for the ***iEPref*** parameter.

Finally, notice how the currently executing method name (***NumStrBasic.GetCountryFormatters()***) is added to the error prefix chain:

``` go
	ePrefix,
		err = ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrBasic.GetCountryFormatters()",
		"")
```

The final empty string parameter is optional and could be used to add error context information.



#### Internal or Private Methods

In the **Public Facing Methods** example, above, method ***GetCountryFormatters()*** makes a call to the  internal or private method, ***getCountryFormatters()***. Calls like this one to a supporting or subsidiary method usually pass a pointer to an instance of ***ErrPrefixDto***.  To date, the best use scenario for subsidiary methods follows this pattern:

``` go

func (nStrBasicQuark numStrBasicQuark) getCountryFormatters(
	fmtCollection *FormatterCollection,
	countryCulture CountryCultureId,
	ePrefix *ErrPrefixDto) (
	err error) {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"numStrBasicQuark." +
			"getCountryFormatters()")

	if fmtCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCollection' is "+
			"a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

    ...
```

Notice how the function name is added to the error prefix chain:

``` go
ePrefix.SetEPref(
		"numStrBasicQuark." +
			"getCountryFormatters()")
```



Notice that this pattern also allows for use of the 'nil' value for parameter, ***ePrefix***. If no error prefix information is present or required, just pass a 'nil' parameter value.

This pattern provides a separate function chain string for each method. This architecture allows for multiple calls from parent methods without adding unnecessary and irrelevant text to the function chain. If an error occurs, only the relevant error prefix and error context information will be returned.



#### Error Context Example

In this example, a function chain is built by calls to multiple levels of the code hierarchy.  The final call to method **Tx3.DoSomething()** triggers an error thereby returning the names of all methods in the call chains plus error context information.

``` go
func(tx1 *Tx1) Something() {

  ePrefDto := ErrPrefixDto{}.New()

  ePrefDto.SetEPref("Tx1.Something()")

  tx2 := Tx2{}

  err := Tx2.SomethingElse(&ePrefDto)

  if err !=nil {
    fmt.Printf("%v\n",
    err.Error())
    return
  }

}

func(tx2 *Tx2) SomethingElse(ePrefDto *ErrPrefixDto) error {

	if ePrefDto == nil {
		ePrefDto = ErrPrefixDto{}.Ptr()
	} else {
		ePrefDto = ePrefDto.CopyPtr()
	}

  	ePrefDto.SetEPref("Tx2.SomethingElse()")

  	tx3 := Tx3{}

  	err := Tx3.DoSomething(ePrefDto)

  	if err !=nil {
    	return err
  	}
}

func(tx3 *Tx3) DoSomething(ePrefDto *ErrPrefixDto) error {

	if ePrefDto == nil {
		ePrefDto = ErrPrefixDto{}.Ptr()
	} else {
		ePrefDto = ePrefDto.CopyPtr()
	}

    // Add error prefix and error context
    // information.
    ePrefDto.SetEPrefCtx(
       "Tx3.DoSomething()",
       "A=B/C C== 0.0")

    .... bad code ....

    if isDivideByZero {
      return fmt.Errorf("%v\n" +
      "I divided by zero and got this error.\n",
      ePrefDto.String())
    }
}


```

When this error is returned up the function chain and finally printed out, the text will look like this:

``` text
 Tx1.Something() - Tx2.SomethingElse()
 Tx3.DoSomething() : A=B/C C== 0.0
 I divided by zero and got this error.

```
#### Code Examples

##### Test Code
Test Code is located in the **errpref** directory of the source code repository. All files beginning with the letters "**zzzt_**" contain test code. The **errpref** directory is located here: [Test Code](https://github.com/MikeAustin71/errpref)

To run the test code, first review the command syntax in [zzzzHowToRunTests](https://github.com/MikeAustin71/errpref/blob/main/zzzzHowToRunTests.md).

##### Example Application
Additional code examples can be found in the Error Prefix Examples Application located in the [errorPrefixExamples](https://github.com/MikeAustin71/errorPrefixExamples) source code repository.

### ErrPref - Quick And Simple Solution

Type ***ErrPref*** is designed to format error prefix text strings for use in error messages. ***ErrPref*** is simple, lightweight, easy to use and seems to work in a variety of situations. The concept is straight forward, "**Put raw text strings in - Get formatted error prefix information out**".

The ***ErrPref*** methods do not use pointer receivers. All receivers are value receivers.

The ***ErrPref*** type does not store error prefix information like the ***ErrPrefixDto*** type. Again, the concept is, "Put raw strings in - Get formatted error prefix strings out".

#### Example Usage Summary

***ErrPref*** passes error prefix information by string. An example calling sequence is shown as follows:

``` go

func(tx1 *Tx1) Something() {

 tx2 := Tx2{}

 err := Tx2.SomethingElse("Tx1.Something()")

 if err !=nil {
   fmt.Printf("%v\n",
   err.Error())
   return
 }

}

func(tx2 *Tx2) SomethingElse(errPrefix string) error {

  // errorPrefix now equals "Tx1.Something()"
  errorPrefix := ErrPref{}.EPrefOld(errPrefix)

  // errorPrefix now equals
  // "Tx1.Something() - Tx2.SomethingElse()"
  errorPrefix = ErrPref{}.EPref(
    errorPrefix,
    "Tx2.SomethingElse()")

  tx3 := Tx3{}

  err := Tx3.DoSomething(errorPrefix)

  if err !=nil {
   return err
  }
}

func(tx3 *Tx3) DoSomething(errorPrefix string) error {

    // Adding Error Context
    errorPrefix = ErrPref{}.EPrefCtx(
                   "Tx3.DoSomething()",
                   "A=B/C C== 0.0")

    .... bad code ....

    if isDivideByZero {
      return fmt.Errorf("%v\n" +
      "I divided by zero and got this error.\n",
      errorPrefix)
    }
}
```

When this error is returned up the function chain and finally printed out, the text will look like this:

``` go
 Tx1.Something() - Tx2.SomethingElse()
 Tx3.DoSomething() : A=B/C C== 0.0
 I divided by zero and got this error.

```



## More Examples

For additional examples, clone and review the example project located in source code repository [errorPrefixExamples](https://github.com/MikeAustin71/errorPrefixExamples)



## ErrPref Package Configuration and Import



### Go Get Command

```go
go get github.com/MikeAustin71/errpref/@v1.6.0
```

​	-- or --

```go
go get github.com/MikeAustin71/errpref/@latest
```



### Import Configuration

```go

import (

	erPref "github.com/MikeAustin71/errpref"
)

```

The import example above shows an alias of 'erPref' which is optional.





## Source Code Documentation

[*errpref* Source Code Documentation](https://pkg.go.dev/github.com/MikeAustin71/errpref)



## Version

The latest version is Version 1.6.0.

As with all previous versions, this Version supports [Go Modules](https://golang.org/ref/mod).

[Version 1.6.0 Release Notes](./releasenotes.md)



## License

Use of this source code is governed by the (open-source) MIT-style license which can be found in the LICENSE file
located in this directory.

[MIT License](./LICENSE)



## Comments And Questions

Send questions or comments to:

``` text
    mike.go@paladinacs.net
```
