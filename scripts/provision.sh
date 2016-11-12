echo "Downloading package lists..."
apt-get update -qq

echo "Installing packages..."
apt-get install -y golang-go mono-runtime mono-devel mono-mcs git python-pip
export GOPATH=/vagrant/go
go get github.com/golang/crypto/scrypt
pip install pyjade

chmod +x /vagrant/scripts/*
mkdir /vagrant/data/users
mkdir /vagrant/data/storage

echo "Compiling source..."
/vagrant/scripts/build.sh