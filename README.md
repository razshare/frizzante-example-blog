# What is this?

This is a blog application that uses [sqlc](https://sqlc.dev/) and [sqlite](https://www.sqlite.org/) to manage state.

# Prerequisites

> [!NOTE]
> #### Prerequisites
> Install `build-essential`, `sqlc` and `frizzante`.
>
> ```sh
> sudo apt install build-essential
> go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
> go install github.com/razshare/frizzante@latest
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