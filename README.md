# Conservation mode switcher tool for Lenovo Ideapad laptops.
It's tested on Fedora Linux, but it should work on most Linux distributions.
To install it, run "go install", this will install it in "$HOME/go/bin" by default.
To run it from anywhere, change the bash profile by adding : "export $GOPATH=$HOME/bin" and add $GOPATH/bin to PATH
To run it with sudo, add "/home/user/go/bin" to the secure path from /etc/sudoers
To run it, type "conservation" and add one of the two flags: "-on" or "-off".
