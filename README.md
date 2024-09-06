# Swiftlio

A assignment tracking application for SMC students written in Go + HTMX + Templ and uses Canvas API for student data retrieval.

## How to run locally: 

### Initial Steps:

1. Ensure you have the latest stable version of [Golang](https://go.dev/doc/install) installed
2. Fork the repo and clone it `git clone https://github.com/[your username]/swiftlio.git && cd swiftlio`
3. Create a `.env` file in the root. Copy and paste `.env_example` contents into `.env` and change env vars to your desired values.
4. <b><i>(If you want to use Docker)</i></b> Create a `docker-compose.yml` file in the root. You can simply copy and paste the contents from `docker-compose-sample.yml` into the `docker-compose.yml` file or change the ports if needed.
     
### Using Docker Compose:

1. Ensure you have Docker Desktop or Docker installed on your local machine (ensure that you can run the `docker-compose` command)
   - [Install Docker](https://www.docker.com/products/docker-desktop/)
2. Running on your machine
   - Run `docker-compose up`. Running this (assuming you have docker and docker-compose) will successfully load the project. The default proxy URL is set to `127.0.0.1:7331` and the API endpoints start come from `localhost:6969`. Change the ports if needed.    

### Without Docker Compose:

1. You must install three things to enable live hot-reloading while developing:
   - [Taliwind CLI](https://tailwindcss.com/blog/standalone-cli)
   - [Air](https://github.com/air-verse/air)
   - [Templ](https://templ.guide/quick-start/installation)
2. (Individual Make Commands) Run/Build/Test your changes using the Makefile in the root of the project.
   - `make build` This command builds the whole project into a single binary in the `./bin/` folder.
   - `make air` This command starts the `air` hot reloading
   - `make tailwind` This command uses the Tailwind CLI to generate and watch for css changes.
   - `make templ` This command helps generate and parse the templ files and outputs the proxy URL that you can use to see changes
   - `make dev` This command runs `make tailwind`, `make templ`, and `air` all in paralllel
   - `make test` Tests all directories and files with `*_test.go`
3. However, all you will need to run is `make dev` to get all the main commands running in parallel for live development. The default proxy URL is set to `127.0.0.1:7331` and the API endpoints start come from `localhost:6969`. Change the ports if needed

## Contributing

- Refer to [CONTRIBUTING.md](https://github.com/kelbwah/swiftlio/blob/master/CONTRIBUTING.md) for more.

## Bugs/Feedback

- Create an issue in this repo if you have any changes you want to make or found any bugs.
