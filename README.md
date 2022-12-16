# ParrotOS Packages Filter

This software is a project made for ParrotOS (parrotsec.org). It should allow you to download the Packages file for all the architectures available for Parrot (amd64, arm64, armhf, i386), filter some informations about each package and display them on the browser in JSON format.

## How does it work?

In [this link](https://download.parrot.sh/parrot/dists/parrot/) it is possible to see the ParrotOS repository. 

<img width="632" alt="image" src="https://user-images.githubusercontent.com/45731605/208175369-cc078e4b-4142-4cfc-b7ca-bb5886f480c9.png">

Inside it are the `contrib`, `main` and `non-free` branches. Each of them contains specific software, a large part essential for the operation of the OS, others simply optional (such as security tools). 

The informations about these software (name, version, maintainer, depends, description, and much more) are contained inside the Packages file. Looking further into the repository folders, you can see that there is a Packages file for each hardware architecture. 

As previously described, this package filter makes it possible to use this information in the convenient JSON format.

## How can I start it?

Download the repository with `git clone` and then decide whether to build the source `go build main.go` or just run it with `go run main.go`

This is a preview. Again, it is a WIP.
