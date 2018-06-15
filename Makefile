

# retrieve list of app versions to delete
VERSIONS = $(shell gcloud app versions list --service default --format="value(version.id)" --filter="traffic_split=0.0")

CUR_VENDOR := $(CURDIR)/vendor
TMP_VENDOR := /tmp/vendor

all:

devserver: $(TMP_VENDOR)/.touch
	GOPATH=$(TMP_VENDOR) python2.7 `which dev_appserver.py` $(CURDIR)/app.yaml

deploy: $(TMP_VENDOR)/.touch
	GOPATH=$(TMP_VENDOR) gcloud app deploy -q --promote --stop-previous-version

$(TMP_VENDOR)/.touch:
	dep ensure
	mkdir -p $(CUR_VENDOR)/src
	mv $(CUR_VENDOR)/* $(CUR_VENDOR)/src/. 2> /dev/null || true
	mv $(CUR_VENDOR) $(TMP_VENDOR)
	touch $(TMP_VENDOR)/.touch

list:
	gcloud app versions list

delete:
	gcloud app versions delete $(VERSIONS)

versions:
	@echo $(VERSIONS)
