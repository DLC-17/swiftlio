# swiftlio
A assignment tracking application for SMC students written in Go + HTMX + Templ and uses Canvas API for student data retrieval.

## How to run locally: 
1. You must install three things:
    - [Golang](https://go.dev/doc/install)
    - [Taliwind CLI](https://tailwindcss.com/blog/standalone-cli)
    - [Air](https://github.com/air-verse/air)
2. Fork the repo and clone it `git clone https://github.com/[your username]/swiftlio.git && cd swiftlio`  
3. Make any local changes.
4. Run/Build/Test your changes using the Makefile in the root of the project.
    - `Make run` This command calls the `build` command and then runs the binary which resides in the `./bin/` directory.
    - `Make build` This command builds the whole project into a single binary in the `./bin/` folder.
    - `Make tailwind` This command uses the Tailwind CLI to generate and watch for css changes.
    - `Make templ` This command helps generate and parse the templ files and outputs the proxy URL that you can use to see changes
    - `Make test` Tests all directories and files with `*_test.go`
5. Note: Currently, you would need to open three different terminal windows/tabs and run these commands from the root of the project to get it working with hot reloading:
    - `air` Running this in one terminal will allow `air` to listen in on chosen files/directories as specified in the `air.toml`.
    - `Make tailwind` Watches for CSS changes 
    - `Make templ` Generates the templ/go files and outputs the proxy URL that you will use to see hot reloaded changes.

## Contributing
- Refer to [CONTRIBUTING.md](https://github.com/kelbwah/swiftlio/blob/master/CONTRIBUTING.md) for more. 

## Bugs/Feedback
- Create an issue in this repo if you have any changes you want to make or found any bugs.
