deploy:
	gcloud functions deploy generate --entry-point Generate --runtime go113 --trigger-http --allow-unauthenticated --region asia-northeast1

deploy-with-docker:
	docker-compose run --rm gcloud make deploy
