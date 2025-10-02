# Equal Experts Operability Take-Home Exercise

## Completed By: Chris Clarkson (<chris@caxa-technologies.co.uk>)

## Date: 2nd October 2025

## Overview

The Go code, Dockerfile, and README.md have been created to meet the requirements of the Technical Test Brief provided [here](/TASK_README.md).

### Testing the Code (No Dockerfile)

This code has been built with Go 1.24.3 and can be tested using the following commands;

```bash
$> go mod download
$> go test ./... -v

?       ee/tech-test-cjc        [no test files]
?       ee/tech-test-cjc/cmd    [no test files]
=== RUN   TestUnmarshalModel
--- PASS: TestUnmarshalModel (0.00s)
=== RUN   TestGistClient
--- PASS: TestGistClient (0.00s)
=== RUN   TestGithubGistServe
--- PASS: TestGithubGistServe (0.00s)
=== RUN   TestGithubGistServeErrorFromGithub
--- PASS: TestGithubGistServeErrorFromGithub (0.00s)
=== RUN   TestGithubGistServeErrorNoUser
--- PASS: TestGithubGistServeErrorNoUser (0.00s)
PASS
ok      ee/tech-test-cjc/internal/github        (cached)
```

#### Congratulations - you have successfully tested the code

### Building the code (No Dockerfile)

This code has been built using Go 1.24.4 and can be built and run using the following commands:
from the root of the repository enter the following commands

```bash
$> go mod download
$> go build -o tech-test

```

Note: You should see no errors during this build.

To test that the build is successful, run the following command and should should the following help output:

```bash
$> ./tech-test

A command line application that runs a simple web server to
return a list of public GISTS for a given GitHub username.

Usage:
  tech-test [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  serve       Start the webserver on the default port

Flags:
  -h, --help     help for tech-test
  -t, --toggle   Help message for toggle

Use "tech-test [command] --help" for more information about a command.
```

#### Congratulations - you have successfully built the code

## Build and test using the Dockerfile

The build process for this code has been completed using a multi-stage Dockerfile. The Dockerfile uses the official Golang image to build the code and then copies the binary to a smaller Alpine image for running. the Tests are executed during the build to validate that no errors are present and that the code is functioning as expected.

NOTE: the Builder and Output container bases are pinned to the speicific hash of the base image to ensure that the build is repeatable.

To build the Docker image, run the following command from the root of the repository:

```bash
$> docker buildx build -t tech-test .
```

## Running the Docker container

To run the Docker container, use the following command:

```bash
$> docker run -d -p 8080:8080 tech-test
```

This will start the container in detached mode and map port 8080 on the host to port 8080 in the container.

To call the api hosted in the container, use the following command:

```bash
$> curl http://localhost:8080/{github-username}

// For Example
$> curl http://localhost:8080/clarksoncj
[
    {
        "id": "36031daf4f217238782c3c89d9ed47d4",
        "description": "",
        "public": true,
        "url": "https://api.github.com/gists/36031daf4f217238782c3c89d9ed47d4",
        "forks_url": "https://api.github.com/gists/36031daf4f217238782c3c89d9ed47d4/forks",
        "commits_url": "https://api.github.com/gists/36031daf4f217238782c3c89d9ed47d4/commits",
        "node_id": "MDQ6R2lzdDM2MDMxZGFmNGYyMTcyMzg3ODJjM2M4OWQ5ZWQ0N2Q0",
        "git_pull_url": "https://gist.github.com/36031daf4f217238782c3c89d9ed47d4.git",
        "git_push_url": "https://gist.github.com/36031daf4f217238782c3c89d9ed47d4.git",
        "html_url": "https://gist.github.com/ClarksonCJ/36031daf4f217238782c3c89d9ed47d4",
        "files": {
            "keybase.md": {
                "filename": "keybase.md",
                "type": "text/markdown",
                "language": "Markdown",
                "raw_url": "https://gist.githubusercontent.com/ClarksonCJ/36031daf4f217238782c3c89d9ed47d4/raw/f806886a56f4fb45ef3b9ddfd02b8586fc6238a1/keybase.md",
                "size": 3122
            }
        },
        "created_at": "2018-03-28T07:48:28Z",
        "updated_at": "2018-03-28T07:48:28Z",
        "owner": {
            "login": "ClarksonCJ",
            "id": 3652990,
            "node_id": "MDQ6VXNlcjM2NTI5OTA=",
            "avatar_url": "https://avatars.githubusercontent.com/u/3652990?v=4",
            "gravatar_id": "",
            "url": "https://api.github.com/users/ClarksonCJ",
            "html_url": "https://github.com/ClarksonCJ",
            "followers_url": "https://api.github.com/users/ClarksonCJ/followers",
            "following_url": "https://api.github.com/users/ClarksonCJ/following{/other_user}",
            "gists_url": "https://api.github.com/users/ClarksonCJ/gists{/gist_id}",
            "starred_url": "https://api.github.com/users/ClarksonCJ/starred{/owner}{/repo}",
            "subscriptions_url": "https://api.github.com/users/ClarksonCJ/subscriptions",
            "organizations_url": "https://api.github.com/users/ClarksonCJ/orgs",
            "repos_url": "https://api.github.com/users/ClarksonCJ/repos",
            "events_url": "https://api.github.com/users/ClarksonCJ/events{/privacy}",
            "received_events_url": "https://api.github.com/users/ClarksonCJ/received_events",
            "type": "User",
            "site_admin": false
        },
        "comments": 0,
        "comments_url": "https://api.github.com/gists/36031daf4f217238782c3c89d9ed47d4/comments",
        "truncated": false
    },
... [<TRUNCATED FOR BREVITY>]
]
```
