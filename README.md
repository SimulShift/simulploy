# Simulploy

Simulploy is a command-line tool designed to streamline the management of Docker environments. It provides a set of commands to efficiently handle Docker operations tailored for various development profiles.

## Installation

Clone the repository and navigate to the directory containing Simulploy:

```bash
git clone https://github.com/your-repo/simulploy.git
cd simulploy
```

Build and install the tool:

```bash
make install
```

## Usage

```plaintext
simulploy [command]
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
simulploy up --profile production
```

**Generate Zsh completion script:**

```bash
simulploy completion
```

**Clean Docker images for a profile:**

```bash
simulploy clean --profile staging
```

## Configuration

Configure Simulploy by using the `simulConfig` command:

```bash
simulploy simulConfig --set key=value
```

## Support

For more information or to request help, use:

```bash
simulploy help
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

Simulploy is released under the MIT License. See the LICENSE file for more details.
