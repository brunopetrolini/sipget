# SIPGET

## Introduction

SIPGET is a command line interface (CLI) application built with Go. It is designed to return the IP address and public server name of a given web address. This tool was built with the aim of studying the fundamentals of the Go language.

## Installation

To install SIPGET, you need to have [Go](https://go.dev/doc/install) installed on your machine. Once Go is installed, you can use the `git clone` command to download SIPGET to your machine:

```bash
git clone https://github.com/brunopetrolini/sipget.git
```

## Usage

To use SIPGET, open the terminal in the project folder and type the following commands:

```bash
./sipget ip --url example.com
```

The above command will result in the public IP addresses of the web address provided.

For server names, use the command:

```bash
./sipget server --url example.com
```

_Note: Replace `example.com` with the web address you want to get the IP or public server name for._
