# make-tools 

```Makefile
# Set the shell to bash always
SHELL := /bin/bash

init= @ $(MK) --init
exists= @ $(MK) -m "define environment variable" --env-exists
save= @ $(MK) -d --save
load= @ $(MK) --load
HELM_CHART_PATH := ./mychart/Chart.yaml
appver=`$(MK) --get-helm-appversion=$(HELM_CHART_PATH)`
helmver=`$(MK) --get-helm-version=$(HELM_CHART_PATH)`
killif=@ $(MK) -d --fail-if

export ENV_SECRET="secret"

my: maketools
	$(init)
	$(exists) ENV_SECRET
	$(save) encoded=`echo "hello from make-tools " | base64`
	$(load) encoded
	@echo
	@echo "encoded message: `$(MK) -l encoded | base64 -d`"
	@echo $(helmver)
	$(killif) v$(appver)=v1.16.0
	@echo v$(appver) $(helmver)

# Tooling
MK=$(shell which make-tools)

maketools:
	go install github.com/kuritka/make-tools@v0.0.4
```

```shell
❯ make my
go install github.com/kuritka/make-tools@v0.0.4
4:25PM DBG Debug mode enabled
4:25PM DBG saving data [encoded aGVsbG8gZnJvbSBtYWtlLXRvb2xzIAo ]
aGVsbG8gZnJvbSBtYWtlLXRvb2xzIAo
encoded message: hello from make-tools
0.1.0
4:25PM DBG Debug mode enabled
4:25PM DBG Fail if: v1.16.0=v1.16.0
make: *** [my] Error 1
```