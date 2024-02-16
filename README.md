# Duplicate and Large File Finder

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
## Table of Contents
- [Description](#description)
- [Features](#features)
  - [Duplicate File Finder](#duplicate-file-finder)
  - [Large File Finder](#large-file-finder)
- [Installation & Prerequisites](#installation--prerequisites)
  - [Duplicate File Finder](#duplicate-file-finder-1)
  - [Large File Finder](#large-file-finder-1)
- [Usage](#usage)
  - [Duplicate File Finder](#duplicate-file-finder-2)
  - [Large File Finder](#large-file-finder-2)
- [Contributing](#contributing)
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

Download the pre-compiled executable (main.exe for Windows) the [Releases](https://github.com/lilsheepyy/Duplicate-File-Finder/releases/tag/Executable) section of this repository.
Place the executable in your desired location.
Usage
## Features
Download the pre-compiled executable (main.exe for Windows) the [Releases](https://github.com/lilsheepyy/Duplicate-File-Finder/releases/tag/Executable) section of this repository.
Place the executable in your desired location.
Usage
Basic Usage:

main.exe (Path to folder)

Replace (Path to folder) with the actual path to the directory you want to scan for duplicates.

Example:
```bash
main.exe C:\Users\ExampleUser\Documents
```
## Installation & Prerequisites
### Large File Finder
No additional prerequisites are required beyond those needed for the Duplicate File Finder.

Option 1: Running from Source
Clone this repository and navigate to the `largefilefinder` directory:
```bash
cd largefilefinder
```
Run the application:
```bash
go run main.go
```
Option 2: Using the Compiled Executable
Download the pre-compiled executable for the Large File Finder from the [Releases](https://github.com/lilsheepyy/Duplicate-File-Finder/releases/tag/Executable) section.
```

## Output
The utility will print a list of any discovered duplicate file groups, indicating the file paths within each group.

## Contributing
Contributions to improve Duplicate-File-Finder are welcome! Please follow these guidelines:

* Fork this repository.
* Create a branch for your changes.
* Submit a pull request with a detailed description of your contributions.
### Large File Finder
Basic Usage:
```bash
main.exe --dir (Path to directory) --size (Minimum file size)
```
Replace `(Path to directory)` with the actual path to the directory you want to scan for large files, and `(Minimum file size)` with the size threshold (e.g., `500MB`).

Example:
```bash
main.exe --dir C:\Users\ExampleUser\Documents --size 100MB
```
