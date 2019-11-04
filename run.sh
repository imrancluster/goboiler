# booting up consul, redis, rds containers
docker-compose up -d consul db redis

# building app
go build -v .

# setting Key=>Value, dependency of app
curl --request PUT --data-binary @config.local.yaml http://localhost:8500/v1/kv/goboiler

# start app
export GOBOILER_CONSUL_URL="127.0.0.1:8500"
export GOBOILER_CONSUL_PATH="goboiler"
export DEBUG="true"

# change your actual binary name
# make sure this name is same as for your go build name
./goboiler serve
