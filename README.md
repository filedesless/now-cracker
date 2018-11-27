# now-cracker

Little cracking challenge hosted on https://cracking-challenge.now.sh

Your goal is to crack the given md5 hash of a 7-digit PIN within 2 seconds.
You have to send it back to *URL*/*HASH* using POST request with the PIN as a body payload

## Requisites

A now secret named `flag`, for example:

	now secrets add flag FLAG{70F3B4E4962848859C917CE1D9C9A611}

## Run locally

It should listen on :8080

	make
	
## Deploy to now

It will output the url it has been deployed to

	now
