# Skynet

This is an experimental HFT cyrptocurrency trader. Currently there is no engine to trading, i.e., market-marker algorithm or any others. However, it pulls the GDAX OrderBook in real time reporting min/max on ask/bid spread. Provided that knowledge, you can write an algorithm around that to execute trades.

### Prerequisites
You need an account on GDAX. Edit the `Dockerfile` with your appropriate `key`, `secret` and `phrase` (yes, it is called `phrase` on GDAX) to run. 

You must export these as environment variables to run locally:

	export SKYNET__COINBASE_KEY={key}
	export SKYNET__COINBASE_SECRET={secret}
	export SKYNET__COINBASE_PHRASE={phrase}

Build and run via Docker, be sure to edit the `.env.local` to match the environment variables key, secret and phrase to Coinbase/GDAX.

	$> docker build -t amanelis/skynet:latest .
	$> docker run --env-file .env.local amanelis/skynet