# Go-astro

Golang plus Astro just feels right... so right I had to create a boilerplate to use every-time I need to spin up a new cool thing...

# Included

- A fiber-powered Golang JSON api + live reload in development
- Astro with Tailwind and React pre-installed

# Running (development)

## Running Golang API in Docker (development)

```bash
docker compose -f ./scripts/docker-compose.dev.yml up
```

The Astro project keeps complaining about ESBuild and refused to work, so you can also run that with:

```bash
make run-ui-dev
```

# Running (production)

You can easily build the binary with your UI embedded with:

```bash
make build
```

Or, if you have other services to run, or you just want to use Docker, it comes with a Dockerfile and a docker compose file you can use by running:

```bash
docker compose up
```

You'll find an example of the environment variables you can add in [.env.development](./.env.development), if the `.env` file is not found, it populates the config with defaults.
