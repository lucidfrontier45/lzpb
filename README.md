# lzpb: Build AWS Lambda Zip Deploy Package

## Why

Recently, it has become very easy to cross-compile programs developed in languages like Python, Go, and Rust from Windows for AMD64 and Arm64 Linux targets. When combined with the AWS Lambda Web Adapter, we can seamlessly create ZIP packages for AWS Lambda deployment using only a Windows environment.

However, there are a couple of caveats. First, the Windows filesystem does not have an equivalent to Linux's execute permissions. This means that executable binaries and `run.sh` files created on Windows will not run on AWS Lambda out of the box. Additionally, for `run.sh` to be launched by bash, its line endings must be LF code, the format used by Linux.

This very small program solves both of these problems, creating a ZIP package ready for deployment to AWS Lambda.

## How

The principle is quite simple. The ZIP archive format allows you to specify permissions for each file. Similar to the `zip` command on Linux, this program packages all items in a specified directory. However, it also specifically forces the execute permission flag for two files: `run.sh` and `bootstrap`. Furthermore, it converts the line endings of `run.sh` to LF.

- `run.sh` used in Python or other script language settings.
- `bootstrap` used in Go, Rust or other native language settings. The executable is assumed to be compiled to this name.

## Usage

1. Get the latest preubilt binary of your platform from https://github.com/lucidfrontier45/lzpb/releases .
2. Place the binary to a directory that is included in your `PATH`
3. run `lzpb [--exec=files] <package_dir> <output_zip_file>`

### Options

- `--exec`: Comma-separated list of files to set executable permissions (default: "bootstrap,run.sh")
  - Use `--exec="file1,file2"` to specify custom files
  - Use `--exec=""` to disable exec permissions entirely

## Develop

The program is writen in Golang. It's a single file program, contains only around 100 lines, and can be easily compiled to any target platforms.

```sh
go build -o lzpb main.go
```
