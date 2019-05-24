# changed-file-checker
Runs the command only if the file has changed.  

Initially created for gnatprep, but intended for general use.  
```bash
gnatprep [args with -] file1 file2
```
If __file1__ is modified then execute, otherwise ignore.  

Requirement: Do not modify the following command (Should be optionally plugged in).
