# Sysmo Backend

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)

## Configurations

### MongoDB Configurations

#### Enabling Replica Set

Open the `mongo.conf` file.

```bash
$ sudo nano /etc/mongo.conf
```

Uncomment the replication and add a replication set name.

```
replication:
  replSetName: "replicaset-01"
```

Restart the mongodb server and login and initiate the replication.

```bash
$ sudo systemctl restart mongodb
$ mongo
> rs.initiate()
```


### Environment File

Use the below example config file format to create the `.env` file.

```ini
HOST=0.0.0.0 # Server host
PORT=5000 # Server port

# Uncomment to enable production mode
# MODE=production

# Database --------------------------------------------------------------------------------

## Mongo
DB_USER=testUser
DB_PASS=testPass
DB_HOST=localhost
DB_PORT=27017
DB=sysmo
DB_SSL=false
DB_CONNECT_TIMEOUT_MS=5000
DB_AUTH_SOURCE=admin
DATABASE_URL=mongodb://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB}?ssl=${DB_SSL}&connectTimeoutMS=${DB_CONNECT_TIMEOUT_MS}&authSource=${DB_AUTH_SOURCE}

# Database --------------------------------------------------------------------------------


# Auth
JWT_ACCESS_SECRET=testJwtKey
JWT_REFRESH_SECRET=testJwtRefreshKey

# Firebase Config
FIREBASE_CONFIG_PATH=firebase_config.json
```

### Firebase Config File

Download the firebase config file in the below format.

```json
{
  "type": "service_account",
  "project_id": "xxxx",
  "private_key_id": "xxxx",
  "private_key": "-----BEGIN PRIVATE KEY-----\nxxxx\n-----END PRIVATE KEY-----\n",
  "client_email": "xxxx",
  "client_id": "xxxx",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "xxxxx"
}
```