# gen-d
Simple CLI to generate txt file based on day

## WHY?
Lately I started writing daily journal. So every single day I have to create the file manually
and if the month has changed I need to create the folder with format `yyyy_m` again and again.
So that's why I created this simple program to automatically generate the file, I just need run this program
and start writing.

## Installation
```
git clone git@github.com:poncorobbin/gen-d.git
cd gen-d
go build ./cmd/gend.go
```

## Usage
```
./gend
```
File will be generated in root folder along with the folder with format `yyyy_m/d.txt`
