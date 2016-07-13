Golang :poop:
=======

Golang Hacks was created by, and is maintained by [Nyah Check](https://github.com/Ch3ck), and it's my mockup work and experiments on the Golang programming language.

Feel free to check out the [license](LICENSE), and make contributions by pull requests.


## Installing Golang on Linux | Mac | BSD

Get the [Tarball](https://golang.org/doc/install?download=go1.6.2.linux-amd64.tar.gz), and extract it into /usr/local, creating a Go tree in /usr/local/go.

```
$ tar -C /usr/local -xzf go1.6.2.linux-amd64.tar.gz

```

Run this command as the superuser. Add '/usr/local/go/bin' to the PATH environment variable. You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile

```
$ export PATH=$PATH:/usr/local/go/bin

```


## Hello Go!

In a file 'hello.go' write the hello world program.

package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}

Then compile it with the go tool:

```
$ go run /path/to/file/hello.go

hello, world

```
If you see the "hello, world" message then your Go installation is working.

## License

Golang-hacks is licensed under [The GNU Public License (GNU)](LICENSE).

