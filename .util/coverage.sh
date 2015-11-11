#!/usr/bin/env bash
# Hack to get coveage report on all or only tests matching a shell glob
# Usage: Either call with no args or with something like .util/coverage.sh auth*

# Run tests with verbose output and optionally pass a fileglob to match
go test -v -coverprofile=c.out "$@"

# Hack to get the output file to work when using a fileglob
sed -i.bak 's/command-line-arguments/github.com\/spohnan\/ci-bot-01/' c.out

# Open in a browser
go tool cover -html=c.out

# Clean up temp files
rm -f c.out*
