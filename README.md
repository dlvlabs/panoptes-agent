# Panoptes Agent

A lightweight monitoring agent designed to collect validator node status data and send it to a central monitoring system.

## Current Status

Currently, the agent only provides local logging functionality for collected monitoring data. Integration with the central monitoring system is on hold.

## Features

Currently supported monitoring items:

- Disk Monitoring
  - Disk usage monitoring (space and inodes)
  - Support for monitoring multiple paths
- Block Height Monitoring
  - Track current block height
  - Detect block creation delays
- Validator Status Monitoring
  - Track missed blocks
  - Monitor validator status changes
  - Track moniker information

Currently, all monitoring data is stored in local logs. Once integration with the central monitoring system is implemented, the collected data will be sent to the central system at configured intervals.

## Requirements

- Go 1.22.7 or higher
- `$GOPATH` environment variable set
- Access to node RPC endpoint

## Installation

1. Clone the repository:

```bash
git clone https://github.com/dlvlabs/panoptes-agent.git
cd panoptes-agent
```

2. Configure the agent:

```bash
cp config.toml.example config.toml
```

Configure required settings in `config.toml`:

- `name`: Optional - Set agent identifier
- `data_send_interval`: Required - Set data collection interval in minutes
- `rpc_url`: Required - Set node RPC endpoint to request data
- `[feature]` section: Required - Enable/disable monitoring features
  - When a feature is enabled (true), its corresponding configuration section is required
  - `disk_space`: When true, requires `paths` setting in `[disk]` section
  - `validator_massage`: When true, requires `acc_address` setting in `[validator]` section

Configuration example:

```toml
[agent]
# Optional: Agent identifier
name = ""

# Required: Data collection interval (minutes)
data_send_interval = 10

# Required: Node RPC endpoint to request data
rpc_url = "http://localhost:26657"

[feature]
# Required: Enable/disable monitoring features
block_height = true
disk_space = true
validator_massage = true

# Required when disk_space is true
[disk]
paths = ["/home", "/var/lib/cosmovisor"]

# Required when validator_massage is true
[validator]
acc_address = "cosmos1..."
```

3. Install the agent:

```bash
make install
```

This command will:

- Install binary to `$GOPATH/bin`
- Create config directory at `~/.config/panoptes/`
- Copy configuration if not exists

## Project Structure

```
.
├── app/            # Application core
├── cmd/            # Command line interface
├── config/         # Configuration management
├── infrastructure/ # External service clients
│   └── client/
│       └── rpc/    # RPC client implementation
├── internal/       # Internal packages
│   ├── agent/      # Agent core logic
│   ├── block/      # Block height monitoring
│   ├── disk/       # Disk usage monitoring
│   └── validator/  # Validator status monitoring
└── utils/          # Utility packages
    └── scheduler/  # Scheduling utilities
```

## Running

Direct execution:

```bash
go run main.go
```

Run installed binary:

```bash
panoptes
```

## Development

### Running Tests

```bash
make test
```

### Cleanup

Remove binary:

```bash
make clean
```

Remove binary and configuration:

```bash
make clean-all
```

## License

This project is licensed under the MIT License.
