# Building My Own Shell

## Introduction

This project is a challenge to build a small shell written in Go. It is a learning project focused on understanding how a shell reads input, recognizes built-in commands, finds executables, and launches external programs.

## My Current Progress

The shell currently supports:

- An interactive read-evaluate-print loop (REPL) with a `$` prompt
- Parsing user input into a command and its arguments
- Reporting commands that cannot be found
- Looking up executables using the directories in `PATH`
- Running external programs with their arguments
- Forwarding output and errors from external programs to the terminal

### Built-in Commands

| Command | Description |
| --- | --- |
| `exit` | Exits the shell |
| `echo <text>` | Prints text to standard output |
| `type <command>` | Reports whether a command is a shell built-in or an external executable |

## Requirements
- Golang version 1.26
- A .env at project root

From the project root run:

```sh
touch .env
go run ./app
```

Example session:

```text
$ echo hello world
hello world
$ type echo
echo is a shell builtin
$ type ls
ls is /usr/bin/ls
$ ls
...
$ exit
Exiting shell!
```

## Current Limitations

Command parsing is intentionally basic at this stage. The shell does not yet support:

- Quoted or escaped arguments
- Pipes and input/output redirection
- Environment-variable expansion
- Command history or autocompletion
- Job control or background processes
- Built-ins such as `cd` and `pwd`

PATH discovery also prints temporary debugging output during startup.

## Planned Work

Future improvements include more robust command parsing, additional built-ins, piping, redirection, command history, and autocompletion.
