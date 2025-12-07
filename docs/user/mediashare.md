# MediaShare Service

MediaShare is an optional companion service for TeleIRC that enables sharing of media files (photos, videos, voice messages) from Telegram to IRC. Since IRC doesn't support inline media, MediaShare provides a self-hosted solution for storing and serving these files via web links.

## Contents

1. [Overview](#overview)
2. [Features](#features)
3. [Configuration](#configuration)
   1. [TeleIRC Configuration](#teleirc-configuration)
   2. [MediaShare Configuration](#mediashare-configuration)
4. [Deployment](#deployment)
   1. [Run with Docker](#run-with-docker)
   2. [Run as Binary](#run-as-binary)
   3. [Run with Systemd](#run-with-systemd)
5. [How It Works](#how-it-works)

## Overview

When a Telegram user shares a photo, video, or voice message, TeleIRC uploads it to the MediaShare service. MediaShare stores the file and returns a short URL that is posted to IRC. IRC users can click the link to view the media in their browser with a modern, responsive player.

**Key benefits:**
- Self-hosted - you control your data
- No external service dependencies (unlike Imgur)
- Supports video, audio, images, and other files
- Modern web player with dark/light mode
- Automatic cleanup of old files
- Bilingual support (English/Polish)

## Features

- **Media Player**: Modern responsive web player for video and audio
- **Image Viewer**: Lightbox-style image viewing with thumbnails
- **File List**: Browse recent uploads at the root URL
- **Auto-cleanup**: Configurable retention period for automatic file deletion
- **Internationalization**: English and Polish language support
- **Dark Mode**: Automatic dark/light mode based on system preference
- **API Key**: Optional authentication for uploads
- **SQLite Database**: Lightweight metadata storage

## Configuration

### TeleIRC Configuration

Add these settings to your TeleIRC environment file:

```bash
# Enable MediaShare integration
MEDIASHARE_ENABLED=true

# MediaShare server URL (where MediaShare is running)
MEDIASHARE_URL=http://localhost:8090

# API key for uploads (must match MediaShare's API key)
MEDIASHARE_API_KEY=your-secret-api-key

# Maximum file size in bytes (default: 50MB)
MEDIASHARE_MAX_SIZE=52428800

# Language for IRC messages (en or pl)
MEDIASHARE_LANG=en
```

### MediaShare Configuration

MediaShare uses these environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `MEDIASHARE_PORT` | `8090` | HTTP server port |
| `MEDIASHARE_API_KEY` | (empty) | API key for upload authentication. If empty, no authentication required |
| `MEDIASHARE_BASE_URL` | `http://localhost:8090` | Public URL for generating links |
| `MEDIASHARE_STORAGE_PATH` | `./uploads` | Directory for storing uploaded files |
| `MEDIASHARE_DB_PATH` | `./mediashare.db` | SQLite database file path |
| `MEDIASHARE_MAX_FILE_SIZE` | `52428800` | Maximum file size in bytes (50MB) |
| `MEDIASHARE_RETENTION_HOURS` | `168` | Hours to keep files (168 = 7 days). Set to 0 to disable cleanup |
| `MEDIASHARE_SERVICE_NAME` | `MediaShare` | Service name shown in web UI |
| `MEDIASHARE_LANG` | `en` | Interface language (`en` or `pl`) |

## Deployment

### Run with Docker

Build and run MediaShare with Docker:

```bash
# Build the image
docker build -t mediashare -f deployments/container/mediashare.Dockerfile .

# Run the container
docker run -d \
  --name mediashare \
  -p 8090:8090 \
  -v mediashare-data:/app/data \
  -e MEDIASHARE_API_KEY=your-secret-key \
  -e MEDIASHARE_BASE_URL=https://media.example.com \
  -e MEDIASHARE_STORAGE_PATH=/app/data/uploads \
  -e MEDIASHARE_DB_PATH=/app/data/mediashare.db \
  mediashare
```

### Run as Binary

Build MediaShare from source:

```bash
# Build
go build -o mediashare ./cmd/mediashare

# Run
MEDIASHARE_API_KEY=your-secret-key \
MEDIASHARE_BASE_URL=https://media.example.com \
./mediashare
```

### Run with Systemd

Create `/etc/systemd/system/mediashare.service`:

```ini
[Unit]
Description=MediaShare Service
After=network.target

[Service]
Type=simple
User=mediashare
WorkingDirectory=/opt/mediashare
ExecStart=/opt/mediashare/mediashare
Restart=always
RestartSec=5

# Environment
Environment=MEDIASHARE_PORT=8090
Environment=MEDIASHARE_API_KEY=your-secret-key
Environment=MEDIASHARE_BASE_URL=https://media.example.com
Environment=MEDIASHARE_STORAGE_PATH=/opt/mediashare/uploads
Environment=MEDIASHARE_DB_PATH=/opt/mediashare/data/mediashare.db
Environment=MEDIASHARE_RETENTION_HOURS=168
Environment=MEDIASHARE_LANG=en

[Install]
WantedBy=multi-user.target
```

Enable and start the service:

```bash
sudo systemctl daemon-reload
sudo systemctl enable mediashare
sudo systemctl start mediashare
```

## How It Works

### Upload Flow

1. Telegram user sends a photo/video/voice message
2. TeleIRC receives the message and downloads the file from Telegram
3. TeleIRC uploads the file to MediaShare via HTTP POST to `/upload`
4. MediaShare stores the file and returns a JSON response with the URL
5. TeleIRC posts the formatted message to IRC with the media URL

### URL Structure

MediaShare generates short, clean URLs:

- `/{id}` - View page with media player
- `/r/{id}` - Raw file (direct download/embed)
- `/` - File list (recent uploads)
- `/health` - Health check endpoint

Where `{id}` is a 5-character alphanumeric identifier (e.g., `Ab3xK`).

### File Storage

Files are stored in a date-based directory structure:

```
uploads/
├── 2024/
│   └── 12/
│       └── 07/
│           ├── Ab3xK.mp4
│           └── xY9pQ.jpg
```

### Cleanup

The cleanup worker runs every hour and removes files that:
- Haven't been opened within the retention period
- Or were uploaded but never opened within the retention period

This ensures disk space is managed automatically while keeping frequently accessed files available.

### Security

- **API Key**: Use `MEDIASHARE_API_KEY` to restrict uploads to authorized clients
- **CORS**: Enabled by default for cross-origin requests
- **Content-Type**: Properly set based on file type for safe browser handling
- **Path Traversal**: File IDs are randomly generated, preventing path traversal attacks
