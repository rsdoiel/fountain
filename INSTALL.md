
Installation
------------

This project is experimental. Get the latest release from [GitHub](https://github.com/rsdoiel/fountain/releases/). 

Quick install with curl or irm
------------------------------

If you are using macOS or Linux you maybe able to install fountain using the following curl command.

~~~shell
curl https://rsdoiel.github.io/fountain/installer.sh | sh
~~~

On Windows you would use Powershell and the following.

~~~
irm https://rsdoiel.github.io/fountain/installer.ps1 | iex
~~~

Install from source
-------------------

## Requirements

- Golang >= 1.20
- Pandoc >= 3
- GNU Make
- Git

## Steps

1. Clone the Git repository for the project
2. change directory into the cloned project
3. Run `make`, `make test` and `make install`

Here's what that looks like for me.

~~~
git clone https://github.com/rsdoiel/fountain src/github.com/rsdoiel/fountain
cd src/github.com/rsdoiel/fountain
make
make test
make install
~~~

By default it will install the programs in `$HOME/bin`. `$HOME/bin` needs
to be included in your `PATH`. E.g.

~~~
export PATH="$HOME/bin:$PATH"
~~~

Can be added to your `.profile`, `.bashrc` or `.zshrc` file depending on your system's shell.


