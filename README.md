## TCP-Probe

External TCP probe implementation compatible with Google's active monitoring system [Cloudprober](https://github.com/google/cloudprober).

### Usage

```bash
$ go get
$ go run tcp_probe.go --address="google.com" --port=80
```

### Install

You can use [Docker image](https://hub.docker.com/r/splusminusx/cloudprober/) or simply install it with `go install`.

### Example Cloudprober config

Example once probe:
```
probe {
  name: "tcp_probe_once"
  type: EXTERNAL
  targets {
    host_names: "google.com"
  }
  external_probe {
    mode: ONCE
    command: "/probes/tcp_probe --address=@target@ --port=80"
  }
}
```

Example server probe:
```
probe {
  name: "tcp_probe_server"
  type: EXTERNAL
  targets {
    host_names: "google.com"
  }
  external_probe {
    mode: SERVER
    command: "/probes/tcp_probe --server"

    options {
      name: "address"
      value: "@target@"
    }

    options {
      name: "port"
      value: "80"
    }
  }
}
```