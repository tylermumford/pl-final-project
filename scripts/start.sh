cd /vagrant

scripts/kill.sh

bin/controller &
bin/caddy -pidfile='/vagrant/.caddypid' &
sleep 2
printf "Finished. Server is running in process "
cat /vagrant/.caddypid