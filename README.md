# wordlist-manager
Simple command line program to learn how do deal with files and apply some Clean Code concepts such as abstraction.

- [x] Receive file path as argument 
- [x] Show the word list
- [x] Add word on list
- [x] Update word on list
- [x] Remove word from list

## Run instructions

At the root of the project run the following shell code:
``` shell
go run main.go --path=C:\your\word\list\path.txt
```

Obs: In the case of the path not be defined, the program will try to search for a file called "wordlist.txt" on the root of the project.
