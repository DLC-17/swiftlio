# swiftlio

A assignment tracking application for SMC students written in Go + HTMX + Templ and uses Canvas API for student data retrieval.

## How to run locally:

1. You must install four things:
   - [Golang](https://go.dev/doc/install)
   - [Taliwind CLI](https://tailwindcss.com/blog/standalone-cli)
   - [Air](https://github.com/air-verse/air)
   - [Templ](https://templ.guide/quick-start/installation)
   - [Docker](https://www.docker.com/products/docker-desktop/)
2. Fork the repo and clone it `git clone https://github.com/[your username]/swiftlio.git && cd swiftlio`
3. Make any local changes.
4. (Individual Make Commands) Run/Build/Test your changes using the Makefile in the root of the project.
   - `make run` This command calls the `build` command and then runs the binary which resides in the `./bin/` directory.
   - `make build` This command builds the whole project into a single binary in the `./bin/` folder.
   - `make tailwind` This command uses the Tailwind CLI to generate and watch for css changes.
   - `make templ` This command helps generate and parse the templ files and outputs the proxy URL that you can use to see changes
   - `make test` Tests all directories and files with `*_test.go`
5. Bootstrapping project with hot reloading for development
   - Create a `.env` file in the root. Copy and paste `.env_example` contents into `.env` and change env vars to your desired values.
   - Similarly, create a `docker-compose.yml` file in the root. You can simply copy and paste the contents from `docker-compose-sample.yml` into the `docker-compose.yml` file or change the ports if needed.
   - Then run `docker-compose up`. Running this (assuming you have docker and docker-compose) will successfully load the project. The default proxy URL is set to `127.0.0.1:7331` and any API calls come from `localhost:6969`. Change the ports if needed.

## Contributing

- Refer to [CONTRIBUTING.md](https://github.com/kelbwah/swiftlio/blob/master/CONTRIBUTING.md) for more.

## Bugs/Feedback

- Create an issue in this repo if you have any changes you want to make or found any bugs.
