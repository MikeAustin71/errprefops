# ErrPref Release Notes Version 1.2.0

This version of ***ErrPref*** was compiled and tested using Go 1.15.8.

This version supports ***Go*** modules.

## Version 1.1.1

1. Removed Maximum Line Length operation from ErrPrefixDto.String()

2. Added Ptr() method to ErrPrefixDto.

3. Added tests for ErrPrefixDto methods CopyIn() and CopyOut()


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
