# Duplicate-File-Finder

A command-line utility for identifying and managing duplicate files.

## Description
Duplicate-File-Finder is a utility designed to help you declutter your storage by efficiently finding and organizing duplicate files. It utilizes robust MD5 hashing for accurate comparisons, ensuring that duplicates are reliably detected.

## Features
* Efficient MD5 Hashing: Employs the well-established MD5 algorithm to create unique fingerprints of files, guaranteeing accurate duplicate detection.
* Flexible Usage: Can be run directly with Go or used as a standalone executable, providing options for different workflows.
* Clear Output: Presents discovered duplicates in an organized and informative manner, enabling easy review and decision-making.
* Silly anotations and output.

## Installation & Prerequisites

Golang (if you plan to run the source code directly)
Option 1: Running from Source

Clone this repository: 
```bash
git clone https://github.com/lilsheepyy/Duplicate-File-Finder.git
```

Navigate to the project directory: 
```bash
cd Duplicate-File-Finder
```
Run the application: 
```bash
go run main.go
```
Option 2: Using the Compiled Executable

Download the pre-compiled executable (main.exe for Windows) from the "Releases" section of this repository.
Place the executable in your desired location.
Usage
Basic Usage:

main.exe (Path to folder)

Replace (Path to folder) with the actual path to the directory you want to scan for duplicates.

Example:
```bash
main.exe C:\Users\ExampleUser\Documents
```

## Output
The utility will print a list of any discovered duplicate file groups, indicating the file paths within each group.

## Contributing
Contributions to improve Duplicate-File-Finder are welcome! Please follow these guidelines:

* Fork this repository.
* Create a branch for your changes.
* Submit a pull request with a detailed description of your contributions.
