
all:

devserver:
	python2.7 `which dev_appserver.py` $(CURDIR)/app.yaml

deploy:
	gcloud app deploy
