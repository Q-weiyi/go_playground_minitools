# go_playground_minitools

This project demonstrates how to create a simple GUI application using Fyne in Golang. The application is for checking the hashed signature of the body in an api call. The shared signkey is needed to perform the hashing.

## Prerequisites

1. **Install Golang**: Follow the [official installation guide](https://golang.org/doc/install) to set up Golang on your system.
2. **Install Fyne**: Run the following command to install the Fyne package:

```
go get -u fyne.io/fyne/v2/cmd/fyne
```

3. **Add the `GOPATH` binaries to your system's `PATH`**:

- Locate the appropriate configuration file for your shell (e.g., `.bashrc`, `.bash_profile`, `.zshrc`).

  For example, if you are using Zsh, you can find or create the `.zshrc` file in your home directory:

  ```
  touch ~/.zshrc
  ```

- Open the configuration file in a text editor and add the following line:

  ```
  export PATH=$PATH:$(go env GOPATH)/bin
  ```

- Save the file and restart your terminal or run `source ~/.zshrc` (or the appropriate configuration file for your shell) to apply the changes.

4. **Install xgo**:

```
go install github.com/karalabe/xgo@latest
```

5. **Install Docker**: Follow the [official installation guide](https://docs.docker.com/get-docker/) to set up Docker on your system.

## Build the Project for Windows 64-bit

1. Navigate to your project directory:

```
cd path/to/your/project/go_playground_minitools
```

Replace `path/to/your/project` with the actual path to your project directory.

2. Build the Windows 64-bit executable using xgo:

```
GOPATH=$GOPATH xgo --targets=windows/amd64 -out SignKeyMiniTool github.com/Q-weiyi/go_playground_minitools
```

This command will generate a Windows 64-bit executable named `SignKeyMiniTool-windows-4.0-amd64.exe.exe` in the current directory.

## Build the Project for Mac

1. In your terminal, build the project:

```
go build -o SignKeyMiniTool
```
