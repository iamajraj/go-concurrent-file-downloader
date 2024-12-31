# Concurrent File Downloader
A simple Go application to download multiple files concurrently with progress tracking.

## Features
- Download files from multiple URLs concurrently.
- Displays real-time download progress for each file.
- Save downloaded files to a specified output folder.

## Requirements
- Go 1.18 or later.

## Installation
1. Clone this repository.
   ```bash
   git clone https://github.com/iamajraj/go-concurrent-file-downloader
   ```
2. Navigate to the project directory.
   ```bash
   cd go-concurrent-file-downloader
   ```
3. Build the application.
   ```bash
   go build -o downloader
   ```

## Usage
Run the application with a list of URLs and an optional output directory:
```bash
./downloader --output=<output_folder> <url1> <url2> ...
```

Example:
```bash
./downloader --output=downloads https://example.com/file1 https://example.com/file2
```

## Project Structure
- `main.go`: Entry point of the application.
- `downloader/downloader.go`: Handles downloading logic.
- `utils/worker.go`: Contains reusable worker-related functionality.

## License
This project is licensed under the MIT License.
