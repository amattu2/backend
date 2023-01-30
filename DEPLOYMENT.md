# Introduction

This guide will briefly walk you through deploying the server as a Systemd service on Linux. The repository should be cloned on your server already.

<details>
  <summary>Systemd Service Template</summary>

  ```ini
  [Unit]
  Description=Backend API Service systemd manager
  After=network.target

  [Service]
  Type=simple
  ExecStart=/var/www/backend/bin/linux
  WorkingDirectory=/var/www

  Restart=always
  RestartSec=3
  KillSignal=SIGINT

  StandardOutput=syslog
  StandardError=syslog

  SyslogIdentifier=api-service
  PrivateTmp=true

  Environment=SSLCERT=/etc/Cloudflare/cert.pem
  Environment=SSLKEY=/etc/Cloudflare/key.pem
  Environment=PORT=443
  Environment=ADDR=0.0.0.0
  Environment=REQUESTMAX=10

  [Install]
  WantedBy=multi-user.target
  ```
</details>

## Compile The App

Run one of the following commands

- `make build_linux`
- `make build`

## Setup The Service

Copy the example file to the Systemd folder

```bash
sudo cp backend/setup/API.service /etc/systemd/system/API.service
```

**Note:** See the default service file and make any changes necessary. It's setup to use SSL.

## Enable The Service

```bash
sudo systemctl start API
sudo systemctl enable API
```
