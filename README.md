# What is this?

This is a blog application that uses [sqlc](https://sqlc.dev/) and [sqlite](https://www.sqlite.org/) to manage state.

# Prerequisites

### Install build tools

On Linux
```sh
sudo apt-get install build-essential
```

On Darwin (MacOS)

```sh
xcode-select --install
```

### Install  `frizzante` and `sqlc`

```sh
go install github.com/razshare/frizzante@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

>[!TIP]
>Remember to add Go binaries to your path.
>
> ```sh
> export GOPATH=$HOME/go
> export PATH=$PATH:$GOPATH/bin
> ```

# Get Started

Configure project

```sh
make configure
```

Generate sqlc code

```sh
make generate
```

Start development mode with

```sh
make dev
```

> [!NOTE]
> In development mode you can login with a default admin account
> with **email** `admin@admin.admin` and **password** `admin`.

# Build

Build for production with

```sh
make build
```

This will create a standalone `.gen/bin/app` binary file.