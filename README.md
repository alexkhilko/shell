# shell

Interactive shell written in Go, that allows to run commands on operating system

This is one of the John Cricket's Coding Challenges solutions https://codingchallenges.fyi/challenges/challenge-shell/ 

# Usage

To build an executable do `make build`

It should create executable under `/bin/shell` executable in current directory

Example of usage
```
❯ ./bin/shell
sh> ls
LICENSE         Makefile        README.md       bin             go.mod          go.sum          main            main.go         word.txt
sh> pwd
/Users/alexkhilko/coding_challenges/shell
sh> cd ..
sh> ls
shell           test.txt        wc
sh> echo 'hey there' | wc
       1       2      12
sh> history
ls
pwd
cd ..
ls
echo 'hey there' | wc
history
sh> exit
```