

# retrieve list of app versions to delete
VERSIONS := $(shell gcloud app versions list --service default --format="value(version.id)" --filter="traffic_split=0.0")

all:

devserver:
	python2.7 `which dev_appserver.py` $(CURDIR)/app.yaml

deploy:
	gcloud app deploy -q --promote --stop-previous-version

list:
	gcloud app versions list

delete:
	gcloud app versions delete $(VERSIONS)

versions:
	@echo $(VERSIONS)
