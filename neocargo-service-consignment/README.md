# neoCargo Consignment Microservice

## Prerequisite

We are using private GitHub repository for our Go Modules.

To make container works with private Go Modules, we had to find a secure
solution for fetching Go external packages from private GitHub repository within
Docker images. We will use GitHub personal access token for that.

The `GITHUB_TOKEN` environment variable is injected to Docker image dynamically.

So, you need to [create a GitHub personal access token](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token) before you proceed with the next steps.

## Usage

**Installation**
- Copy the files `config.env` and `deploy.env` to your repo
- Replace the variables (e.g.: `GITHUB_TOKEN`) in `config.env`

Build the container:

```sh
$ make build
```

Run the container:

```sh
$ make run
```
