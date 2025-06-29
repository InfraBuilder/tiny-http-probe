# tiny-http-probe

[![Docker Image Size](https://img.shields.io/docker/image-size/infrabuilder/tiny-http-probe/latest)](https://hub.docker.com/r/infrabuilder/tiny-http-probe)
[![Docker Pulls](https://img.shields.io/docker/pulls/infrabuilder/tiny-http-probe)](https://hub.docker.com/r/infrabuilder/tiny-http-probe)
[![Docker Image Vulnerabilities](https://img.shields.io/badge/security-scanned-brightgreen)](https://hub.docker.com/r/infrabuilder/tiny-http-probe)
[![GitHub Repo](https://img.shields.io/badge/github-repo-blue?logo=github)](https://github.com/infrabuilder/tiny-http-probe)
[![CI](https://github.com/infrabuilder/tiny-http-probe/actions/workflows/docker-build-push.yml/badge.svg)](https://github.com/infrabuilder/tiny-http-probe/actions)
[![Docker Weekly Build](https://img.shields.io/badge/autobuild-weekly-brightgreen)](https://hub.docker.com/r/infrabuilder/tiny-http-probe)

---

## Overview

tiny-http-probe is a minimal, stateless HTTP server designed for health checks, cache testing, and network diagnostics. It provides a set of endpoints that return simple JSON payloads, allowing you to:
- Test HTTP connectivity and response headers
- Inspect client IP and proxy headers
- Validate cache and CDN behaviors

## Security Considerations
- **Rootless by default:** The container runs as a non-root user (UID/GID 1000) for improved security.
- **Minimal attack surface:** The final Docker image is built `FROM scratch` and contains only the statically compiled Go binary.
- **No persistent storage:** The app is stateless and does not write to disk.
- **No authentication:** Endpoints are public by design; do not expose to untrusted networks if sensitive.

## Build Instructions

### Build and Run with Docker

```sh
docker build -t tiny-http-probe .
docker run -p 8080:8080 tiny-http-probe
```

### Build Locally (requires Go 1.22+)

```sh
go mod tidy
go build -o tiny-http-probe
./tiny-http-probe
```

The server will listen on port 8080 by default. You can override the port with the `PORT` environment variable.

## Bug Reports

Please report any bugs or issues via [GitHub Issues](https://github.com/<your-github-username>/tiny-http-probe/issues).
