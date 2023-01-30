# Introduction

To setup the backend API as a SystemD service (Linux), copy the following three commands. This expects that the build is in `/backend/bin` already.

```bash
sudo cp backend/setup/API.service /etc/systemd/system/API.service
sudo systemctl start API
sudo systemctl enable API
```
