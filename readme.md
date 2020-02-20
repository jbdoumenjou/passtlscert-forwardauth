# Goal

Test how we can pass the tls cert in header through the forward auth.

# Usage

## Launch the Stack

```bash
docker-compose up
```

## Stop the Stack

```bash
docker-compose down
```

## Check the Fake Auth Server
```bash
curl -kvv http://localhost:5858 -H "X-Forwarded-Tls-Client-Cert-Info: foo"
```

## Check Through Traefik

```bash
# check the testing app directly
curl -kvv https://whoami --cacert certs/ca.pem --cert certs/traefik.pem --key certs/traefik.key
```
