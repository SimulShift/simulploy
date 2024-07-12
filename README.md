# Simulploy

Simulploy is a command-line tool designed to manage several docker compose files.
In big projects that rely on complex docker environments, it is common to have 
multiple docker-compose files that define different services. Simulploy simplifies 
the process of managing these environments by providing a single interface to start, stop, and clean up Docker environments.

The main use case is being able to deploy docker containers anywhere in the terminal without having to navigate to the
directory containing the docker-compose files. 


## Installation

Clone the repository and navigate to the directory containing Simulploy:

```bash
git clone https://github.com/your-repo/simulploy.git
cd simulploy
```

Build and install the tool:

```bash
go build -o simulploy
```

Add simulploy in your PATH to use anywhere in the terminal.

## Configuration File

Simulploy uses a configuration file named `.simulploy.yaml` located in the user's home directory. This file contains essential settings that define how Simulploy interacts with Docker environments.

Here is  an example 
```yaml
filepath: C:\Users\penat\.simulploy.yaml
docker_dir: D:\Home\repos\docker-configs\docker-compose
project_root: D:\Home\repos\simulploy
docker_network: "NOT YET IMPLEMENTED"
metaservices:
  - postgres
  - chatbot
  - envoy
```

### Config File Setup

Create the `.simulploy.yaml` file in your home directory:

```plaintext
~/.simulploy.yaml
```

### Usage

```plaintext
simulploy [command]

Available Commands:
  clean       Delete Docker images for the profile
  completion  Generate completion script for Zsh
  db          Database operations
  down        Compose down the Docker environments
  help        Help about any command
  simulConfig Configure the CLI
  up          Bring up Docker environments

Flags:
  -D, --dev                  development build
  -h, --help                 help for simulploy
  -m, --metaservice string   choose a metaservice
  -P, --prod                 production build
  -p, --profile string       profile to use (default "development")

Use "simulploy [command] --help" for more information about a command.

```

### Available Commands

- `clean`: Deletes Docker images associated with the specified profile.
- `completion`: Generates a completion script for the Zsh shell.
- `db`: Performs database operations.
- `down`: Shuts down Docker environments specified in the profile.
- `help`: Provides help information about any command.
- `simulConfig`: Configures various aspects of the CLI tool.
- `up`: Starts up Docker environments according to the specified profile.

### Flags

- `-h, --help`: Displays help information for Simulploy.
- `-m, --metaservice string`: Specifies a metaservice to target.
- `-p, --profile string`: Sets the profile to use (default is "development").

### Examples

**Start Docker environments:**

```bash
simulploy up --metaservice postgres --dev
```


**Clean Docker images for a profile:**

```bash
simulploy clean -m chatbot --prod
```

## Configuration

Configure Simulploy by using the `simulConfig` command:

```bash
simulploy simulConfig --set key=value
```

## Support

For more information or to request help, join the discord and go to #simulploy
[![Discord](https://shields.io/badge/discord-join-blue)](https://discord.gg/KJwgaE34Cu)

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.
