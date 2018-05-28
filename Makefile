
all:

devserver:
	python2.7 `which dev_appserver.py` app.yaml

deploy:
	gcloud app deploy
