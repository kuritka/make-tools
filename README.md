# make-tools 

```Makefile
SHELL := /bin/bash

init=@./make-tools --init -m "starting"
save=@./make-tools --message "saving value" --save
load=@./make-tools --debug --load
exists=@./make-tools -m "environment variable doesn't exist" --env-exists
# Options

export ENV_SECRET="secret"
t:
	go build -o make-tools main.go
	$(init)
	$(exists) ENV_SECRET
	$(save) encoded1=`echo -n yyy | base64`
	$(save) encoded2=`echo bubu | base64`
	$(load) encoded1
	@echo
	$(save) hello=SGVsbG8gZnJvbSBFTkNPREVE
	@echo "the encoded message: `./make-tools -l hello | base64 -d`"
	make y
	
```

y:
	echo 123

```shell
❯ make t
go build -o make-tools main.go
starting
saving value
saving value
12:21PM DBG Debug mode enabled
12:21PM DBG loaded 'encoded1':'eXl5'
eXl5
saving value
the encoded message: Hello from ENCODED


```