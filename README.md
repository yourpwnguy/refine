<div align="center">

![GoStyle LOGO](https://i.imgur.com/wveX8z8.png)

</div>
<h4 align="center">Simple, Ultra-fast file handling utility for text deduplication.</h4>
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/yourpwnguy/refine">
<!-- <a href="https://github.com/iaakanshff/crtfinder/releases"><img src="https://img.shields.io/github/downloads/iaakanshff/crtfinder/total"> -->
<a href="https://github.com/yourpwnguy/refine/graphs/contributors"><img src="https://img.shields.io/github/contributors-anon/yourpwnguy/refine">
<!-- <a href="https://github.com/iaakanshff/crtfinder/releases/"><img src="https://img.shields.io/github/release/iaakanshff/crtfinder"> -->
<a href="https://github.com/yourpwnguy/refine/issues"><img src="https://img.shields.io/github/issues-raw/yourpwnguy/refine">
<a href="https://github.com/yourpwnguy/refine/stars"><img src="https://img.shields.io/github/stars/yourpwnguy/refine">
<!-- <a href="https://github.com/iaakanshff/crtfinder/discussions"><img src="https://img.shields.io/github/discussions/iaakanshff/crtfinder"> -->
</p>

---

Meet `refine`, a powerful and user-friendly tool for the process of removing duplicate lines from files. `refine` is designed with efficiency in mind, making file manipulation a seamless experience for users. Below is a comprehensive guide on installing and using Seek to enhance your file processing tasks.

## Features 💡

- Efficient file deduplication and management.
- Includes inbuilt sorting.
- Support for diverse input methods.
- Advanced wildcard sorting with exception handling.

## Installation 🛠️
To install the refine tool, you can simply use the following command.
````bash
go install -v "github.com/yourpwnguy/refine/cmd/refine@latest"
cp ~/go/bin/refine /usr/local/bin/
````

## Usage 📘
```yaml
Usage: refine [options]

Options: [flag] [argument] [Description]

DIRECT:
  refine file.txt                       (Read and write to the same file)
  refine file1.txt file2.txt            (Read from file1 and write to the file2)

STDIN:
  cat file.txt | refine                 (Read from stdin and display to stdout)
  cat file.txt | refine newfile.txt     (Read from stdin and write to a specific file)

FEATURES: (ONLY DIRECT MODE)
  refine -w, --wildcard                 (Sort all files in the directory)
  refine -we, --wildcard-exception      (Specify files to be skipped while using wildcard)

DEBUG:
  refine -v, --version                  (Check current version)
```
### DIRECT MODE:

1) Using `refine` to read and write the deduplicated ouptut to the same file:
```js
$ cat file.txt
https://google.com
https://google.com
https://microsoft.com
https://twitter.com
https://github.com
https://github.com

$ refine file.txt
[INF] File: test.txt, Original: 6 lines, Unique: 4 lines, Processed in 108.2µs

$ cat file.txt
https://github.com
https://google.com
https://microsoft.com
https://twitter.com
```

2) Using `refine` to read from file1 and write to deduplicated output to file2:
```js
$ cat file1.txt
https://google.com
https://google.com
https://microsoft.com
https://twitter.com
https://github.com
https://github.com

$ refine file1.txt file2.txt
[INF] File: file2.txt, Original: 6 lines, Unique: 4 lines, Processed in 101.1µs

$ cat file2.txt
https://github.com
https://google.com
https://microsoft.com
https://twitter.com
```

3) Using `refine` for wildcard sorting (-w), which sorts all files in a directory. This feature is limited to direct mode, as during the tool's development, no use case for the pipeline mode was found.
```js
$ refine -w .\test\
[INF] File: 1.txt, Original: 2203598 lines, Unique: 308 lines, Processed in 355.4838ms
[INF] File: 2.txt, Original: 2193548 lines, Unique: 308 lines, Processed in 357.8736ms
[INF] File: 3.txt, Original: 2176797 lines, Unique: 308 lines, Processed in 360.693ms
[INF] File: 4.txt, Original: 2229058 lines, Unique: 308 lines, Processed in 353.194ms
```

4) Using `refine` for wildcard sorting (-w), which sorts all files in a directory except for the specified exceptions. The exceptions, meaning the files to be skipped, can be provided through the -we (wildcard exception) flag with filenames comma-separated.
```js
$ refine -w .\test\ -we 1.txt,2.txt
[INF] File: 3.txt, Original: 2176797 lines, Unique: 306 lines, Processed in 265.4093ms
[INF] File: 4.txt, Original: 2229058 lines, Unique: 308 lines, Processed in 376.9528ms
```
---

### STDIN MODE:

1) Using `refine` for sorting the lines from the standard input (stdin). The deduplicated output is displayed on the terminal without modifying the original file. This method is ideal for viewing results without altering the source file.

```js
$ cat file.txt | refine
https://github.com
https://google.com
https://microsoft.com
https://twitter.com
[INF] Original: 6 lines, Unique: 4 lines, Processed in 0s
```


2) Using the `refine` for sorting the lines from the standard input (stdin), and writes the deduplicated output to a new file specified as an argument. This allows users to create a new file with cleaned data while preserving the original content. `Note:` If the specified file already exists and contains data, it will also be sorted.

```js
$ cat file1.txt | refine file2.txt
[INF] File: file2.txt, Original: 10 lines, Unique: 4 lines, Processed in 150.3µs
```
          
## But Why Use Our Tool❓ 

Well, I understand that there are already popular solutions out there like `tomnomnom's` [anew](https://github.com/tomnomnom/anew), but here's why I think `refine` stands out: I've taken inspiration from those tools and built `refine` to be even more flexible and feature-rich. It's all about making features like wildcard sorting as straightforward and powerful as possible. With `refine`, I've aimed to make it easy for you to manage and sort all your files exactly how you need them.

## Contributing 🤝

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, feel free to open an issue or submit a pull request.
