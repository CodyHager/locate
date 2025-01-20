# locate
OS-agnostic CLI tool written in Go for finding files or directories in a filesystem.

## Building
To build an executable that you can add to your `PATH`, run

`go build -o locate`

Or on windows:

`go build -o locate.exe`

After adding the binary to your `PATH` you should be able to use the `locate` command from anywhere.

## Usage
Run `locate -h` for information about using `locate`.

If you want to search for a file in your working directory (including sub-directories), just run

`locate file fileName`

Flags:
- `-c`, `--case-sensitive`       If set to false, will ignore the case of file names. Default is false.
- `-e`, `--exlude`   Strings to exlude when searching through file names.
- `-h`, `--help`                 help for file
- `-p`, `--path`         Root directory to start searching at, defaults to your working directory
- `-s`, `--strict`               If set to false, will include file names that contain the name of the desired file. Default is false.

If you want to search for a folder in your working directory (including sub-directories), just run

`locate dir dirName`

Flags:
- `-c`, `--case-sensitive`       If set to false, will ignore the case of directory names. Default is false.
- `-e`, `--exlude`   Strings to exlude when searching through directory names.
- `-h`, `--help`                 help for dir
- `-p`, `--path`         Root directory to start searching at, defaults to your working directory
- `-s`, `--strict`               If set to false, will include directory names that contain the name of the desired directory. Default is false.

## Examples:
To search for a file called `main.go` (case sensitive and strict) starting at your working directory:

`locate file main.go -c -s`

To search for a directory called github (case sensitive) two directories up exluding directories that have a `.` in their name:

`locate dir github -p ../../ -c -e .`

> [!TIP]
> If you want to exlude more than one string, use `e` twice! Like this: `-e . -e a`
