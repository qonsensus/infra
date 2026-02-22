# Starting dev environment

```bash
cp dev/.env.example dev/.env
docker compose -f dev/docker-compose.yaml --env-file dev/.env up -d
```