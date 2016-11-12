echo "Downloading package lists..."
apt-get update -qq

echo "Installing packages..."
apt-get install -y mono-runtime mono-devel mono-mcs git python-pip
go get github.com/golang/crypto/scrypt
pip install pyjade

echo "Installing Go..."
curl -# https://storage.googleapis.com/golang/go1.7.3.linux-386.tar.gz > /tmp/go.tar.gz
tar -C /usr/local/ -xzf /tmp/go.tar.gz
rm /tmp/go.tar.gz

echo 'export GOPATH=/vagrant/go
export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
source /etc/profile
chmod +x /vagrant/scripts/*
mkdir /vagrant/data/users
mkdir /vagrant/data/storage

echo "Compiling source..."
/vagrant/scripts/build.sh