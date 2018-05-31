
all:

devserver:
	python2.7 `which dev_appserver.py` $(CURDIR)/app.yaml

deploy:
	gcloud app deploy -q --promote --stop-previous-version

list:
	gcloud app versions list

delete:
	gcloud app versions delete $(VERSIONS)
