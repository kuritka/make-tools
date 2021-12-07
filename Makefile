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
