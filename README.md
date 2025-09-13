# vmopt-cli

VMware Optimizer CLI

## Description

`vmopt-cli` is a Go-based command-line tool designed to help optimize resource usage when running VMware. I built this tool because my computer has **limited resources** (low RAM and CPU). Running multiple VMs with `vmopt-cli`, I can monitor and control resource usage to keep my system responsive while working.

## Features

- Real-time monitoring of CPU and RAM usage.
- Limit CPU usage of VMs (e.g., max 80%).
- Limit RAM usage of VMs (e.g., max 90%).
- Logging system usage for later review.
- Lightweight and simple CLI written in Go.

## Installation

1. Make sure you have **Go 1.21+** installed.
2. Clone the repository:

   ```bash
   git clone https://github.com/nepile/vmopt-cli.git
   ```

3. Navigate into the project directory:

   ```bash
   cd vmopt-cli
   ```

4. Build the executable:

   ```bash
   go build -o vmopt
   ```

5. Run the CLI:

   ```bash
   ./vmopt --help
   ```

## Usage

### General format:

```bash
./vmopt [flags]
```

### Available Args

| Args         | Description                       | Example            |
|--------------|-----------------------------------|--------------------|
| `monitor`    | Monitoring CPU and RAM usage      | `./vmopt monitor`  |
| `optimize`   | Optimizing resources usage of VMs | `./vmopt optimize` |

### Available Flags

| Flag         | Description                  | Example         |
|--------------|------------------------------|-----------------|
| `--max-cpu`  | Maximum CPU usage (in percentage) | `--max-cpu 80`  |
| `--max-ram`  | Maximum RAM usage (in percentage) | `--max-ram 90`  |
| `--log`      | Log file location            | `--log ./usage.log` |
| `--interval` | Monitoring interval in seconds | `--interval 5`  |
| `--help`     | Show help and available options | `--help`        |

### Example Commands

- Run monitoring with a 2-second interval:

  ```bash
  ./vmopt monitor --interval 2
  ```

- Limit CPU usage to 80% and RAM usage to 90% and set process to lower-priority:

  ```bash
  ./vmopt optimize --max-cpu=80 --max-ram=90 --action=lower-priority
  ```

- Limit CPU usage to 80% and RAM usage to 90% and kill the process (VM):

  ```bash
  ./vmopt optimize --max-cpu=80 --max-ram=90 --action=kill
  ```

## Why I Built This

- My computer has limited resources, and opening multiple VMs often caused it to freeze.
- I wanted a simple tool to monitor and limit resource usage so my host system stays responsive.
- This project also serves as a portfolio piece to showcase my ability to build CLI applications with Go.

## Example Scenarios

- Running 3 VMs on a 4GB RAM laptop: Use `--max-ram 85` to prevent VMware from consuming all memory.
- High CPU load during VM boot: Set `--max-cpu 70` to avoid 100% CPU spikes.
- Debugging system freezes: Run with `--log ./usage.log` and review logs after freezes.

## Contributing

Contributions are welcome!

1. Fork this repository
2. Create a new branch (`git checkout -b feature-xyz`)
3. Commit your changes (`git commit -m "feature: add xyz"`)
4. Push to your branch (`git push origin feature-xyz`)
5. Open a Pull Request

## License

This project is now licensed under the MIT License. See the LICENSE file for details.

## Notes

- This tool directly interacts with your system host resources — use with caution.
- Not recommended for production servers without proper testing.
- Still under development, features may change over time.