# Preface
`check_consul_node` is a Nagios check for reporting the health state of a node in Consul cluster.

# Build requirements
This tool is implemented in Go so, obviously, a Go compiler is required.

# Command line parameters
| *Option* | *Description* | *Note* |
|:---------|:--------------|:-------|
| `--ca-file=<ca>` | CA file if the CA is not present in the CA store of the system | |
| `--consul-url=<uri>` | Consul URL | Default is `http://localhost:8500` |
| `--datacenter=<dc>` | Data center | |
| `--help` | Show help text | - |
| `--insecure` | Skip SSL verification | |
| `--namespace=<ns>` | Consul name space to use | Only available for Consul Enterprise! |
| `--node=<name>` | Node name | Default: hostname |
| `--token=<token>` | Consul ACL token | Required if ACLs are configured |
| `--version` | Show version information |

# Licenses
## check_consul_node

Copyright (C) 2021 by Andreas Maus

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

