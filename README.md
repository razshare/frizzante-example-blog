# What is this?

This is a blog application that uses [sqlc](https://sqlc.dev/) and [sqlite](https://www.sqlite.org/) to manage state.

# Prerequisites

Make sure you have `build-essential`, `curl` and `unzip` installed on your machine.

```sh
sudo apt install build-essential unzip curl
```

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