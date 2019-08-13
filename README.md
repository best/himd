# Himd
Himd is a http multithreading downloader develop by golang.

Project name is a recursive acronym for `Himd is a multithreading downloader`.



**Project Version**：0.0.1

> This project is under develop now.



## Usage

### Windows

On windows you can use the command below to show usage：

```powershell
himd.exe -h
 
Himd version: himd/0.0.1
Usage: himd [-u url] [-t thread]

Options:
  -h    command line help
  -t thread
        set downloader thread, default is number of logical CPUs (default 8)
  -u url
        set download url, must format as [https/http/ftp]://target.url
```



### Unix-like System

On Unix-like System, such as: CentOS, MacOS, you can use the comand below to show usage:

```bash
himd -h
 
Himd version: himd/0.0.1
Usage: himd [-u url] [-t thread]

Options:
  -h    command line help
  -t thread
        set downloader thread, default is number of logical CPUs (default 8)
  -u url
        set download url, must format as [https/http/ftp]://target.url
```