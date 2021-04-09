# ErrPref Release Notes Version 1.5.2

This version of ***ErrPref*** was compiled and tested using Go 1.16.3.

This version supports ***Go*** modules.

## Version 1.5.2
1. Fixed link to source code documentation.

2. Added documentation providing more usage examples.

3. Added new method ErrPrefixDto.NewEPrefCtx()

## Version 1.5.1
1. Fixed Go Mod File to assign correct version.

## Version 1.5.0
1. Modified interface IErrorPrefix. Added Methods:
   - GetEPrefStrings() \[\]\[2\]string
   - SetEPrefStrings(twoDStrArray \[\]\[2\]string)
   
2. Added interface IBuilderErrorPrefix
   
3. Added interface IBasicErrorPrefix

4. Added methods on Type ErrPrefixDto
   - AddEPrefStrings()
   - GetEPrefStrings()
   - NewIBasicErrorPrefix()
   - NewIEmpty()
   - SetEPrefStrings()


## Version 1.4.0

This version of ***ErrPref*** was compiled and tested using Go 1.15.8.

This version supports ***Go*** modules.

1. Modified the interface IErrorPrefix. Removed references to ErrorPrefixInfo.
   
2. Added 'Z' methods which return an ErrPrefixDto by value.
   - ZCtx()
   - ZCtxEmpty()
   - ZEPref()
   - ZEPrefCtx()
   - ZEPrefOld()

3. Added tests


## Version 1.3.0

1. Added IErrorPrefix interface. 

2. Added ErrPrefixDto method XCtxEmpty.

3. Added ErrPrefixDto method XSetEmpty.

4. Changed behavior of ErrPrefixDto method SetCtx. Empty input strings will now delete last error context.

5. Changed behavior of ErrPrefixDto method XCtx. Empty input strings will now delete last error context.

6. Added ErrPrefixDto method NewFromIErrorPrefix.

## Version 1.2.0

1. Removed Maximum Line Length operation from ErrPrefixDto.String()

2. Added Ptr() method to ErrPrefixDto.

3. Added tests for ErrPrefixDto methods CopyIn() and CopyOut()

4. Added new method ErrPrefixDto.StrMaxLineLen().

5. Changed method name and signature for ErrPrefixDto.SetMaxTextLineLen()

6. Changed method name for ErrPrefixDto.SetMaxTextLineLenToDefault()

7. Added new method ErrPrefixDto.Copy()

## Version 1.1.0

This version of ***ErrPref*** was compiled and tested using Go 1.15.8.

This version supports ***Go*** modules.

### ErrPrefixDto

1. Converted method ErrPrefixDto.String() from a pointer receiver to a value receiver.

2. Added 'X' Methods:
   - XCtx() - Returns pointer to current instance.

   - XEPref() - Returns pointer to current instance.

   - XEPrefCtx() - Returns pointer to current instance.

   - XEPrefOld() - Returns pointer to current instance.

     



## Version 1.0.0 

This version of ***ErrPref*** was compiled and tested using Go 1.15.8.

This version supports ***Go*** modules.

Initial release of ***ErrPref***.
