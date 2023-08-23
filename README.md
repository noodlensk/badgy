# badgy

Badgy is a simple tool to check for new notifications in multiple accounts like gmail, slack, etc.

Name `Badgy` came from the idea of having a `badge` on the desktop with the number of unread notifications in different
accounts.

## Installation

```bash
make dep
make install
```

Ensure that you have set the `$GOPATH` environment variable.

## Configuration

Export the following environment variables:

```bash
# gmail provider
BADGY_GMAIL_TOKEN="<your token>"
BADGY_GMAIL_CREDENTIALS="<your gapp credentials>"

# slack provider
BADGY_SLACK_TOKEN="<your token>"
BADGY_SLACK_COOKIE="<your 'd' value cookie>"
```
