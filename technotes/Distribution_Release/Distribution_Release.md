[TOC]



# Distribution Release

Check the [Release Notes](../../errpref/releasenotes.md) and [errpref/README.md](../../errpref/README.md) first.



## ***errPref*** Source Code Repository

Distribution releases are made from the source code repository for ***errPref*** is **github.com/MikeAustin71/errpref/** . This repository is used to distribute release versions.

The development repository is ***errprefops*** located at ***github.com/MikeAustin71/errprefops***

## 1. Do NOT merge ***errprefops*** branch ***dev***  to branch ***main*** before You Transfer Files 

You might make changes to **github.com/MikeAustin71/errpref/** that need be transferred back to ***github.com/MikeAustin71/errprefops***.



## 2. Make Certain **github.com/MikeAustin71/errpref** is on branch ***dev***



## 3. Transfer Files From ***errprefops***  to ***errPref***

Copy all the files in the ***github.com/MikeAustin71/errprefops/errPref*** directory to the ***github.com/MikeAustin71/errPref***.

Be sure to overwrite files in the ***github.com/MikeAustin71/errPref***. Do not delete the files as ***github.com/MikeAustin71/errPref*** contains some additional files that should NOT be deleted. 



## 3. Verify Tests on  ***errPref*** 

Run the tests on ***errPref*** and store results in ***errPref/zzzzz_tests.txt***. Verify that all tests are passing. 



## 4. Push local ***errPref*** branch ***dev*** to Remote ***errPref*** branch ***dev*** 



## 5. Switch from local ***errPref*** branch ***dev*** to local ***errPref*** branch ***main*** 



## 6. Merge local ***errPref*** branch ***dev*** to local ***errPref*** branch ***main***



## 7. Push local ***errPref*** branch ***main*** to Remote ***errPref*** branch ***main***



## 8. Tag and Version this Release

Example:

```
# List All Tags
git tag --list

# Create New Tag
git tag v2.0.1

# Push New Tag to Remote Repo
git push origin v2.0.1

# List All Tags
git tag --list
```



## 9. Important: Switch from local ***errPref*** branch ***main*** to local ***errPref*** branch ***dev***

Don't leave the repo on branch ***main***.



## JOB DONE!

