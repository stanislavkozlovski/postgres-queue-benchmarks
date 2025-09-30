# Postgres Queue Benchmarks

A simple Go benchmark that stress-tests PostgreSQL as a queue (insert + claim rows with `FOR UPDATE SKIP LOCKED`).

## Design

A simple design would use a `BOOLEAN is_read` field to denote whether a field is read or not.
In my testing, I found this lowers performance. The EXPLAIN query shows that each read query starts paying a lot of latency to filter through many pages of already `is_read=true` rows.
Vacuum apparently takes time to prune them away, and until it does - reads are slowed.

A more performant approach is to delete rows once read.
This keeps the table small and performant. Since we don't necessarily want to omit the data (some companies have a policy of never deleting, at least before a backup) - we will move the data to another table.
In my simple tests, I find this increases read throughput by **50%** at larger scale.
---

## 1. Prepare Server Environment
Create a Server EC2. Pick the instance carefully
```bash
export PRIVATE_IP=172.31.20.198
### Mount fast volume for PostgreSQL data
DEV=/dev/nvme1n1    # adjust if lsblk shows a different device
MNT=/pgdata

sudo mkfs.xfs -f $DEV
sudo mkdir -p $MNT
sudo mount $DEV $MNT

sudo mkdir -p $MNT/pg17
sudo chown -R postgres:postgres $MNT/pg17
sudo chmod 700 $MNT/pg17
sudo chown root:postgres $MNT
sudo chmod 750 $MNT
# persist mount
UUID=$(sudo blkid -s UUID -o value $DEV)
echo "UUID=$UUID $MNT xfs defaults,nofail 0 2" | sudo tee -a /etc/fstab


### Install PostgreSQL 17 (PGDG repo)
sudo apt-get update -y
sudo apt-get install -y curl ca-certificates gnupg lsb-release
curl https://www.postgresql.org/media/keys/ACCC4CF8.asc | \
  gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/pgdg.gpg >/dev/null
echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" | \
  sudo tee /etc/apt/sources.list.d/pgdg.list

sudo apt-get update -y
sudo apt-get install -y postgresql-17 postgresql-client-17 postgresql-common


### Create cluster on mounted volume
sudo pg_dropcluster --stop 17 main
sudo pg_createcluster 17 main --datadir=$MNT/pg17 --port=5432 -- -c listen_addresses='$PRIVATE_IP,localhost'

pg_lsclusters

### allow all traffic
echo "host all all 0.0.0.0/0 md5" | sudo tee -a /etc/postgresql/17/main/pg_hba.conf
sudo systemctl reload postgresql@17-main

### Configure database
sudo -u postgres createdb benchmark
sudo -u postgres psql -c "ALTER USER postgres WITH PASSWORD 'postgres';"
```

## 2. Prepare Client Environment

Create a Client EC2, in the same AZ as the server.
Open the security group's postgresql port (5432)
Install Go
```bash
cd /tmp
wget https://go.dev/dl/go1.24.7.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.24.7.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
```

Build the Benchmark
```bash
git clone https://github.com/stanislavkozlovski/postgres-queue-benchmarks.git
cd postgres-queue-benchmarks

go build -o ./pg_queue_bench .
chmod +x ./pg_queue_bench
```
Run the Benchmark

export HOST="172.31.20.198"  # adjust to your server's private IP
```bash
./pg_queue_bench \
  --host=$HOST \
  --port=5432 \
  --db=benchmark \
  --user=postgres \
  --password=postgres \
  --writers=50 \
  --readers=50 \
  --duration=120s \
  --payload=1024 \
  --report=5s
```

**Flags**

- `--host` – PostgreSQL host (default `localhost`)
- `--port` – PostgreSQL port (default `5432`)
- `--db` – database name (default `benchmark`)
- `--user` – database user (default `postgres`)
- `--password` – password for user
- `--writers` – number of concurrent writer goroutines
- `--readers` – number of concurrent reader goroutines
- `--duration` – test duration (e.g. `120s`)
- `--payload` – payload size in bytes (default `1024`)
- `--report` – report interval (default `5s`)