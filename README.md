# Advent of Code

This is my repository to post my Advent of Code solutions.

To run my solutions:
- Make sure the language's requirements are installed.
- Execute the following command in the terminal:
```powershell
go run main.go
```

If you want to use my way of downloading the input:
- Copy the download_input.ps1 file to your own repository.
- Add a .env file to your repository and add the following variable (the session token from your Advent of Code, find using your browser's inspector):
```dotenv
SESSION_TOKEN=your_session_token_here
```
- Navigate to `_internal/input_downloader` in your terminal and run the following command:
```powershell
GOROOT=C:\Program Files\Go #gosetup
GOPATH=C:\Users\YOUR USERNAME HERE\go #gosetup
go build -o "path_to_output_exe_here" "REPOSITORY PATH + \_internal\input_downloader\main.go"
```
- Copy the created exe to your own repository. 
- Replace the powershell exe name with your exe name and replace the env path with your .env file