# ErrPref
A lightweight golang type for attaching custom function chains and error prefix
strings to error messages.

## The Problem
As my Go programs, types and methods have grown in sophistication and complexity,
the need for improved error tracking, detailed error messages, and a record of code
execution has become more important. The term 'Function Chains' as used here describes
a list of the method names called prior to encountering a particular error. Adding
function chain documentation to the beginning of returned error messages seem to make
error tracking and management a much easier proposition.

## The Solution
The type, 'ErrPref' is designed to affix custom text to the beginnings of error messages. 
'ErrPref' is lightweight, easy to use and seems to work in a variety of situations.

The idea is not to go overboard. I tried my hand at an error handler for Go Programs earlier
and quickly realized that complexity was growing exponentially. At least at the beginning,
the idea is to keep 'ErrPref' simple as a fast and efficient mechanism for adding error 
prefix text to the beginning of error messages. That said, 'ErrPref' is a work in progress. 
The current code is pre-version 1.0. We shall see where this leads.

For examples and documentation, refer to the source code file, 'errpref.go'.

