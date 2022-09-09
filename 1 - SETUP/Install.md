<!-- write a doc to install go on windows with chcoc  -->

# Install Go

## Install Go on Windows
 
### Install Chocolatey

Chocolatey is a package manager for Windows. It is a command line tool that allows you to install software on Windows. It is similar to Homebrew on Mac OS X or apt-get on Linux.

To install Chocolatey, open a command prompt as an administrator and run the following command:

    @powershell -NoProfile -ExecutionPolicy Bypass -Command "iex ((new-object net.webclient).DownloadString('https://chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"

### Install Go

To install Go, open a command prompt as an administrator and run the following command:

    choco install golang

### install extensions on VScode 

1. Open VS Code; from the extensions tab at the left, search for "go" and install it
> NOTE: the extension is called "Go" and not "golang" from the author "Go Team at Google"
2. Close VS Code completely and open it up again

3. Go to View menu; select **Command Palette**
    1. Or just press `ctrl+shift+p`
    2. Type: `go install`
    3. Select _"Go: Install/Update Tools"_
    4. Check all the checkboxes

