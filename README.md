# What is this?

This is a blog application that uses [sqlc](https://sqlc.dev/) and [sqlite](https://www.sqlite.org/) to manage state.

# Prerequisites

> [!NOTE]
> #### Prerequisites
> Make sure you have `frizzante`, `air`, `sqlc`, `bun` and `build-essential` installed on your machine.
>
> ```sh
> sudo apt install build-essential
> go install github.com/razshare/frizzante@latest
> go install github.com/air-verse/air@latest
> go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
> which bun || curl -fsSL https://bun.sh/install | bash
> ```

# Get Started

Generate sqlc utilities with

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

This will create a standalone `bin/app` binary file.