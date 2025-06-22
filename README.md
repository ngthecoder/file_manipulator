# File Manipulator

## Description
This program manipulates files based on the input user provides.

## System Requirements
Developed with go1.24.4 linux/amd64

## Usage
./program \<mode\> [arguments]

## Modes
### reverse
Read the contents of input file and save the reversed contents to output file<br>
./program reverse <input file (string)> <output file (string)>
### copy
Read the contents of input file and copy the contents to output file<br>
./program copy <input file (string)> <output file (string)>
### duplicate-contents
Read the contents of input file and save the duplicated contents to input file<br>
./program duplicate-contents <input file (string)> <how many times to duplicate (int)>
### replace-string
Read the contents of input file and replace old string with new string then save the updated contents to input file<br>
./program replace-string <input file (string)> <old string (string)> <new string (string)>
 
## Examples
./program reverse input.txt output.txt<br>
./program copy input.txt output.txt<br>
./program duplicate-contents input.txt 5<br>
./program replace-string input.txt cat dog
