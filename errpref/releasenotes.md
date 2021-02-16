# ErrPref Release Notes Version 1.4.1

This version of ***ErrPref*** was compiled and tested using Go 1.15.8.

This version supports ***Go*** modules.

## Version 1.4.0

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
