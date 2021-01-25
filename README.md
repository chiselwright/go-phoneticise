# phoneticise

A simple library to convert ASCII text to its equivalent NATO phonetic
alphabet representation.

For example:

> Hello World!

becomes:

> Hotel Echo Lima Lima Oscar (space) Whiskey Oscar Romeo Lima Delta

## cli-client

You can experiment with the library by running:

```sh
go run ./cmd/cli-client/main.go
```

You'll see something like this:

<!-- markdownlint-disable MD013 -->
```txt
❯ go run ./cmd/cli-client/main.go
{empty input or ^C to quit}
Enter text: Hello World!
» Hotel Echo Lima Lima Oscar (space) Whiskey Oscar Romeo Lima Delta
{empty input or ^C to quit}
Enter text: Hello World!
» Hotel Echo Lima Lima Oscar (space) Whiskey Oscar Romeo Lima Delta
{empty input or ^C to quit}
Enter text: Not My cIRcus
» November Oscar Tango (space) Mike Yankee (space) Charlie India Romeo Charlie Uniform Sierra
{empty input or ^C to quit}
Enter text:
```
<!-- markdownlint-enable MD013 -->
