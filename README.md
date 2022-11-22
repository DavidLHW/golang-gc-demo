<div id="top" />

# GO Garbage Collection 1.19 Demo

This project explores the practical use of `GOMEMLIMIT` and visualizes the changes to heap memory in realtime.

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#setup">Setup</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#configuring">Configuring</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#glossary">Glossary</a></li>
    <li><a href="#technical-essay">Technical Essay</a></li>
    <li><a href="#references">References</a></li>
  </ol>
</details>

## Setup

### Prerequisites

#### General

- Go >= v1.19
- Docker

#### Windows

- WSL / cygwin / git bash (to run shell scripts)

### Configuring

The estimated memory required to run the http server is around `10 MiB`, hence the ideal initial memory limit should be set to be more than `10 MiB` to prevent GC from constantly running, leaving no space for new heap allocation.

Default configuration

```env
GOMEMLIMIT=30MiB   # soft memory limit for the runtime in human readable bytes
GOGC=100           # initial garbage collection target percentage, 100 => 100% => 2x
GODEBUG=gctrace=1  # enables debug mode with garbage collection tracing
```

<p align="right">(<a href="#top">back to top</a>)</p>

## Usage

Run the docker container.

```bash
# builds image and run container
$ ./scripts/run.sh
```

Run tests.

```bash
$ ./scripts/test.sh
```

### API Endpoints

By default the gin app will listen on port 8080 on localhost. The `base_url` mentioned from here should be http://localhost:8080 unless explicitly modified.
Refer to [API Docs](docs/api-endpoints.md) for more detailed documentation.

<p align="right">(<a href="#top">back to top</a>)</p>

## Glossary

<p align="right">(<a href="#top">back to top</a>)</p>

## Technical Essay

<p align="right">(<a href="#top">back to top</a>)</p>

## References

- [GO GC Guide](https://go.dev/doc/gc-guide#Memory_limit)
- [GO Runtime Soft Memory Limit Proposal](https://github.com/golang/go/issues/48409)
- [GO Test GC Memory Limit - Source Implementation](https://github.com/golang/go/blob/4585bf96b4025f18682122bbd66d4f2a010b3ac9/src/runtime/testdata/testprog/gc.go#L325)
- [Statsviz - GO Runtime Metrics Visualizer](https://github.com/arl/statsviz)

<p align="right">(<a href="#top">back to top</a>)</p>
