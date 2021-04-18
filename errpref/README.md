# *errpref* - Adding error prefix, function execution chains and error context to Error Messages In Go

**The *errpref* software package is a collection of lightweight GoLang types designed to attach error prefix text, function execution chain lists and error context strings to error messages.**

The ***errpref*** software package was written in the [Go](https://golang.org/) programming language, a.k.a. ***Golang***.

***errpref*** supports [Go Modules]([Go Modules Reference - The Go Programming Language (golang.org)](https://golang.org/ref/mod)).

The current version of ***errpref*** is Version 1.6.0. Most notably, this version implements the [Left Margin Feature](#left-margin-feature) in error prefix string formatting.



## Table of Contents

 - [The Problem](#the-problem)
 - [The Solution](#the-solution)
 - [What *errpref* Does and Doesn't Do](#what-errpref-does-and-doesnt-do)
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
     - [Maximum Text Line Length](#maximum-text-line-length)
     - [Left Margin Feature](#left-margin-feature)
   - [ErrPref - Quick And Simple Solution](#errpref---quick-and-simple-solution)
 - [Exchanging Error Prefix Information with User Defined Types](#exchanging-error-prefix-information-with-user-defined-types)
 - [Usage Examples](#usage-examples)
   - [Test Code](#test-code)
   - [Example Application](#example-application)
 - [Package Configuration and Import](#package-configuration-and-import)
   - [Go Get Command](#go-get-command)
   - [Import Configuration](#import-configuration)
 - [External Dependencies](#external-dependencies)
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



## What *errpref* Does and Doesn't Do

The ***errpref*** package generates text strings which can be populated with error prefix and error context information. It does NOT create error messages. Error messages are still created by the user with the standard tool provided by the **Go Programming Language**.  

***errpref*** is used to generate the error prefix, function chain and error context information. After that, it is up to the user to add this information to the error message. 

```go
	if inComingFormatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'inComingFormatterCurrency' is"+
			" a 'nil' pointer!\n",
			errPrefixDto.String())
		return err
	}

```

In the example above, ***errPrefixDto*** is an instance of ***errpref*** type ***ErrPrefixDto*** previously configured with a series of method names documenting the execution path which led to this point in the code. As the example shows, the user is still responsible for configuring the error message in its entirety. 

While ***errpref*** was specifically designed to document function or method chains of execution, the example above shows that, as a practical matter, the user is free to position this text anywhere in the error message. Also, this text is not reserved exclusively for error messages. It could be used in informational messages or any other type of text display.

The typical error message decorated with error prefix and context information will typically look something like this:

```tex
main()-mainTest004()
TestFuncsDto.TestAlphaDto001()
 :  A->B
TestFuncDtoAlpha01.Tx1DoSomething()
testFuncDtoAlpha02.tx2DoSomethingElse()
testFuncDtoAlpha03.tx3DoAnything()
testFuncDtoAlpha04.tx4DoNothing()
 :  A/B==4
testFuncDtoAlpha05.tx5DoSomethingBig()
 :  A->B
testFuncDtoAlpha06.tx6TryGivingUp()
 :  A/B = C B==0
Example Error: An Error Ocurred! Something bad.
Like Divide by Zero
```

This example error message was taken from the ***errpref*** example application, [errorPrefixExamples](https://github.com/MikeAustin71/errorPrefixExamples).



## Definition of Terms

For our purposes, error prefix information consists of two elements: 

1. One or more function or method name listed in sequence of code execution. 
2. An optional error context string associated with a specific function or method name.

### Error Prefix

**Error Prefix** text is designed to be configured at the beginning of error messages and is most often used to document the thread of code execution by listing the calling sequence for a specific list of functions and/or methods.

### Error Context

**Error Context** strings are designed to provide additional information about the function or method identified by the
associated **Error Prefix** text. Typical context information might include variable names, variable values and additional details on function execution.

Error Context strings are optional and are always associated with an Error Prefix string.



## String Formatting Conventions

When using the methods provided by ***errpref*** types ***ErrPref*** and ***ErrPrefixDto***, the error prefix strings are parsed for function names, method names and error context information based on specific string delimiters. As of the current ***errpref*** release, these string delimiters are standardized, fixed and not subject to customization by the user. 

### Error Prefix Delimiters

Error Prefix Elements are delimited by one of two delimiters: 

- **New Line Delimiter** = "**\n**" Delimits function/method names on a single line.
- **In-Line Delimiter** = " **-** " (**[SPACE]-[SPACE]**) Delimits multiple function/method names on the same line.

### Error Context Delimiters

Associated Error Context Sub-Elements likewise have two delimiters: 

- **New Line Delimiter String** = "**\n[SPACE]:[SPACE]\[SPACE]**" Delimits error context information on a single line.
- **In-Line Delimiter String** = "**[SPACE]:[SPACE]**" Delimits error context information displayed with function/method name on same line.



## Getting Started

### Two Types To Choose From

Currently, there are two principal ***errpref*** types providing error prefix formatting services: ***ErrPrefixDto*** and ***ErrPref***. Type ***ErrPref*** offers basic error prefix formatting while type ***ErrPrefixDto*** covers a much wider range of formatting and data transmission capabilities.

### ErrPrefixDto - A Full Featured Solution

The Error Prefix Data Transfer Object, ***ErrPrefixDto***, offers the same services as type, ***ErrPref***, but is packaged in a different architecture. While ***ErrPref*** methods receive a string and instantly return a formatted string, ***ErrPrefixDto*** encapsulates error prefix information in an internal array of Error Prefix Information objects (***ErrorPrefixInfo***). Strings are only created when the type's ***String()*** method is called. Instances of ***ErrPrefixDto*** are designed to be passed as input parameters to subsidiary methods. The act of passing ***ErrPrefixDto*** objects to successive functions in the execution sequence effectively builds documentation of the execution function chain. 

The ***String()*** method has a value receiver.  Virtually all other ***ErrPrefixDto*** methods have pointer receivers.

#### Public Facing Methods
A public method may receive error prefix information in a variety of formats from sources both internal and external to the package. To date, the best use scenario for a Public Facing Function or Method follows this pattern:

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
  2. **Stringer** - The Stringer interface from the '**fmt**' package. This interface has only one method:

``` go
    	type Stringer interface {
    		String() string
    	}
```

  3. **string** - A string containing error prefix information. For maximum effectiveness, this should be delimited with the standard [String Delimiters](#string-formatting-conventions).
  5. **\[\]string** - A one-dimensional slice of strings containing error prefix information. Typically, the one-dimensional slice contains only function or method names.
  1. **\[\]\[2\]string** - A two-dimensional slice of strings containing error prefix and error context information. slice\[x\]\[0\] holds function or method names. slice\[x\]\[1\] holds associated error context information.
  2. ***strings.Builder*** - An instance of strings.Builder. Error prefix information extracted from ***strings.Builder*** will be imported into the new returned instance of ***\*ErrPrefixDto***.
  3. ***\*strings.Builder*** - A pointer to an instance of strings.Builder.  Extracted error prefix information will be imported into the new returned instance of ***\*ErrPrefixDto***.
  4. ***ErrPrefixDto*** - An instance of ***ErrPrefixDto***. The ***ErrorPrefixInfo*** collection extracted from this object will be copied to the new returned instance of ***\*ErrPrefixDto***.
  6.  **\*ErrPrefixDto** - A pointer to an instance of **ErrPrefixDto**. The ***ErrorPrefixInfo*** collection extracted from this object will be copied to the new returned instance of  ***\*ErrPrefixDto***.
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

Finally, notice how the currently executing method name (***NumStrBasic.GetCountryFormatters()***) is added to the error prefix method chain:

``` go
	ePrefix,
		err = ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrBasic.GetCountryFormatters()",
		"")
```

The final empty string parameter is optional and could be used to add error context information.



#### Internal or Private Methods

In the **Public Facing Methods** example, above, method ***GetCountryFormatters()*** makes a call to the  internal or private method, ***getCountryFormatters()***. Calls like this one to a supporting or subsidiary method usually pass a pointer to an instance of ***ErrPrefixDto***.  To date, the best use scenario for calls to subsidiary methods follows this pattern:

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



Notice that this pattern also allows for use of the **nil** value for parameter, ***ePrefix***. If no error prefix information is present or required, just pass a **nil** parameter value.

This pattern provides a separate function chain string for each method. This architecture allows for multiple calls from parent methods without adding unnecessary and irrelevant text to the function chain. If an error occurs, only the relevant error prefix and error context information will be returned.



#### Error Context Example

Recall that **Error Context** strings are designed to provide additional information about the function or method identified by the
associated **Error Prefix** text. Typical context information might include variable names, variable values and additional details on function execution.

In this example, a function chain is built by calls to multiple levels of the code hierarchy.  The final call to method **Tx3.DoSomething()** triggers an error thereby returning the names of all methods in the call chain plus error context information associated with those methods.

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



#### Maximum Text Line Length

Users have the option to set the Maximum Text Line Length for error prefix strings. The default Maximum Text Line Length is 40-characters. The following error prefix text display demonstrates the default 40-character maximum line length.

```tex
main()-mainTest004()
TestFuncsDto.TestAlphaDto001()
 :  A->B
TestFuncDtoAlpha01.Tx1DoSomething()
testFuncDtoAlpha02.tx2DoSomethingElse()
testFuncDtoAlpha03.tx3DoAnything()
testFuncDtoAlpha04.tx4DoNothing()
 :  A/B==4
testFuncDtoAlpha05.tx5DoSomethingBig()
 :  A->B
testFuncDtoAlpha06.tx6TryGivingUp()
 :  A/B = C B==0
Example Error: An Err
```



The next example shows the same error prefix information displayed with a custom Maximum Text Line Length of 70-characters.

```tex
main()-mainTest005() - TestFuncsDto.TestAlphaDto002()
TestFuncDtoAlpha01.Tx1DoSomething()
testFuncDtoAlpha02.tx2DoSomethingElse()
testFuncDtoAlpha03.tx3DoAnything()
testFuncDtoAlpha04.tx4DoNothing() : A/B==4
testFuncDtoAlpha05.tx5DoSomethingBig() : A->B
testFuncDtoAlpha06.tx6TryGivingUp() : A/B = C B==0
Example Error: An Error Ocurred! Something bad.
Like Divide by Zero!
```

The Maximum Text Line Length is monitored and controlled using methods:

-  ***ErrPrefixDto.SetMaxTextLineLen()***
-  ***ErrPrefixDto.GetMaxTextLineLen()***

Be advised that the minimum text line Length is 10-characters. Any attempt to set the Maximum Text Line Length to a value less than 10-characters will result in reinstatement of the default Maximum Text Line Length.

This example error messages shown above were taken from the ***errpref*** example application, [errorPrefixExamples](https://github.com/MikeAustin71/errorPrefixExamples).



#### Left Margin Feature

The default left margin length for text returned by the method ***ErrPrefixDto.String()*** is zero (0). In other words, by default, there is no left margin.

Example error prefix information text display with left margin length of zero:

```tex
Tx1.Something() - Tx2.SomethingElse()
Tx3.DoSomething() - Tx4() - Tx5()
Tx6.DoSomethingElse()
Tx7.TrySomethingNew()
 :  something->newSomething
Tx8.TryAnyCombination()
Tx9.TryAHammer() : x->y - Tx10.X()
Tx11.TryAnything() - Tx12.TryASalad()
Tx13.SomeFabulousAndComplexStuff()
Tx14.MoreAwesomeGoodness()
 :  A=7 B=8 C=9
```

The following ***ErrPrefixDto*** methods allow the user to control both the left margin length, and the character used to populate the left margin:

1. ***ErrPrefixDto.SetLeftMarginLength()*** 
2. ***ErrPrefixDto.SetLeftMarginChar()***
3. ***ErrPrefixDto.GetLeftMarginLength()***
4. ***ErrPrefixDto.GetLeftMarginChar()***

In the next example, the left margin length is set to three (3), and the left margin character is set to the empty space character (' ') or ASCII 0x20:

```tex
   Tx1.Something() - Tx2.SomethingElse()
   Tx3.DoSomething() - Tx4() - Tx5()
   Tx6.DoSomethingElse()
   Tx7.TrySomethingNew()
    :  something->newSomething
   Tx8.TryAnyCombination()
   Tx9.TryAHammer() : x->y - Tx10.X()
   Tx11.TryAnything() - Tx12.TryASalad()
   Tx13.SomeFabulousAndComplexStuff()
   Tx14.MoreAwesomeGoodness()
    :  A=7 B=8 C=9
```



In the final example, the left margin length is set to three (3), and the left margin character is set to the asterisk character ('*'): 

```tex
***Tx1.Something() - Tx2.SomethingElse()
***Tx3.DoSomething() - Tx4() - Tx5()
***Tx6.DoSomethingElse()
***Tx7.TrySomethingNew()
*** :  something->newSomething
***Tx8.TryAnyCombination()
***Tx9.TryAHammer() : x->y - Tx10.X()
***Tx11.TryAnything() - Tx12.TryASalad()
***Tx13.SomeFabulousAndComplexStuff()
***Tx14.MoreAwesomeGoodness()
*** :  A=7 B=8 C=9
```



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



## Exchanging Error Prefix Information with User Defined Types

Custom user defined types supporting the ***IBuilderErrorPrefix*** interface will be able to receive data from and insert data into instances of ***ErrPrefixDto***. The following ***ErrPrefixDto*** methods interoperate with the ***IBuilderErrorPrefix*** interface:

1. ***ErrPrefixDto.CopyInFromIBuilder()***
2. ***ErrPrefixDto.CopyOutToIBuilder()***
3. ***ErrPrefixDto.NewIEmpty()***
4. ***ErrPrefixDto.SetIBuilder()***

Taken together, these methods facilitate the import and export of error prefix and context information between ***ErrPrefixDto*** and user defined types implementing the ***IBuilderErrorPrefix*** interface.

The ***IBuilderErrorPrefix*** interface is defined as follows:

```go
type IBuilderErrorPrefix interface {
	GetEPrefStrings() [][2]string

	SetEPrefStrings(twoDStrArray [][2]string)

	String() string
}
```



## Usage Examples

### Test Code

Test Code is located in the **errpref** directory of the source code repository. All files beginning with the letters "**zzzt_**" and ending with "**_test.go**" contain test code. The **errpref** directory is located here: [Test Code](https://github.com/MikeAustin71/errpref)

Test Code provides many examples 

Currently, unit tests show code coverage at 83%.

To run the test code, first review the command syntax in [zzzzHowToRunTests](https://github.com/MikeAustin71/errpref/blob/main/zzzzHowToRunTests.md).

### Example Application

Additional code examples can be found in the Error Prefix Examples Application located at [https://github.com/MikeAustin71/errorPrefixExamples](https://github.com/MikeAustin71/errorPrefixExamples).  This application also contains a "Concurrency" example.



## Package Configuration and Import



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



## External Dependencies

None. This software package is not dependent on any external module or package.



## Source Code Documentation

[*errpref* Source Code Documentation](https://pkg.go.dev/github.com/MikeAustin71/errpref)



## Version

The latest version is Version 1.6.0.

As with all previous versions, this Version supports [Go Modules](https://golang.org/ref/mod).

For more details, see the [Release Notes](./releasenotes.md).



## License

Use of this source code is governed by the (open-source) MIT-style license which can be found in the LICENSE file
located in this directory.

[MIT License](./LICENSE)



## Comments And Questions

Send questions or comments to:

``` text
    mike.go@paladinacs.net
```
