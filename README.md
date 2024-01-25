## Introduction: Unduplication Tool (Seek) 🚀
Meet Seek, a powerful and user-friendly tool created by Aakansh to simplify the process of removing duplicate lines from files. Seek is designed with efficiency in mind, making file manipulation a seamless experience for users. Below is a comprehensive guide on installing and using Seek to enhance your file processing tasks. A tool for processing files and writing to them.

## Installation 🛠️
### Download the files into your system
````
go install github.com/iaakanshff/seek@v1.0.2
````
### Copy the binary to the path folder
````
sudo cp ~/go/bin/seek /bin/
````

## Usage 📘
### 1) Directly on files:
````
─$ cat file.txt                 |      ─$ seek file.txt 
https://google.com              |      ─$ cat file.txt 
https://google.com              |      https://google.com 
https://google.com              |      https://microsoft.com 
https://microsoft.com           |      https://twitter.com 
https://twitter.com             |      https://github.com 
https://twitter.com             |
https://github.com              |
https://github.com              |
````
Seek directly operates on files, removing duplicates and updating the original file with clean output. It's a quick and efficient way to sanitize your data.

### 2) Using Piping (|):
````
─$ cat file.txt                 |      ─$ cat file.txt | seek
https://google.com              |      https://google.com 
https://google.com              |      https://microsoft.com     
https://google.com              |      https://twitter.com        
https://microsoft.com           |      https://github.com     
https://twitter.com             |       
https://twitter.com             |
https://github.com              |
https://github.com              |
````
Here, Seek processes input from the standard output (stdout) using piping. The unduplicated output is displayed on the terminal without modifying the original file. This method is ideal for viewing results without altering the source file.

### 3) Piping with Output to a New File:
````
─$ cat file.txt                 |      ─$ cat file.txt | seek newfile.txt 
https://google.com              |      ─$ cat newfile.txt 
https://google.com              |      https://google.com 
https://google.com              |      https://microsoft.com 
https://microsoft.com           |      https://twitter.com 
https://twitter.com             |      https://github.com 
https://twitter.com             |
https://github.com              |
https://github.com              |
````
In this scenario, Seek processes input through piping and writes the unduplicated output to a new file specified as an argument. This approach allows users to create a new file with cleaned data while preserving the original content

## Unistallation 🗑️
````
sudo rm ~/go/bin/seek
sudo rm /bin/seek
````
You can uninstall the tool anytime. Just follow the above steps and remember to provide feedback :)

## Conclusion 🌟
I Hope that you'll like this tool, any feedback will be appreciated. I find this tool very useful when I have to process a large chunk of files to remove duplicate lines.
