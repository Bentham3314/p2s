# p2s
Post a message using webhook to slack 

## Env:
linux x86_64

```
$ GOOS=linux GOARCH=amd64 go build
```

## Usage

- GET

```
$ go get -u github.com/Bentham3314/p2s
```

- SET environment variable

```
$ export SLACK_POST_USERNAME="webhook-test"
$ export SLACK_POST_ICON=":robot_face:"
$ export SLACK_POST_CHANNEL="#bots"
$ export SLACK_POST_WEBHOOK_URL=""
```

- Post to SLACK

```
$ echo "TEST MESSAGE" | p2s 

$ cat /tmp/message.txt | p2s 

$ p2s /tmp/message.txt
```

### a

- a.sh

```
#!/bin/sh

. ~/.p2s-env

TODAY=`date +%Y%m%d`
mysqldump --single-transaction --all-databases > mysqldump.$TODAY

if [ $? -ne 0 ]; then
  cat ~/mysqldump-faild.mess | p2s
fi
```
