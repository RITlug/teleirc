Using TeleIRC with Docker
=========================

The files included here are examples for you to use.
For more information on using them, [read the documentation](https://docs.teleirc.com/en/latest/deploy-teleirc/#docker).
Before using them, copy files you intend to use to the root directory of the repository.

## Included Dockerfiles

- `Dockerfile` - TeleIRC main bot
- `mediashare.Dockerfile` - MediaShare media hosting service (optional)

## Using MediaShare

MediaShare is an optional service for hosting Telegram media files. See [MediaShare documentation](../../docs/user/mediashare.md) for details.

To run both TeleIRC and MediaShare together, use the docker-compose example:

```bash
cp docker-compose.yml.example docker-compose.yml
# Edit .env to configure both services
docker-compose up -d
```
