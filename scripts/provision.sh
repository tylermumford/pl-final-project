echo "Downloading package lists..."
apt-get update -qq

echo "Installing packages..."
apt-get install -y golang-go mono-runtime mono-devel mono-mcs python-pip
pip install pyjade

echo "Compiling source..."
/vagrant/scripts/build.sh

echo "Final steps..."
chmod +x /vagrant/scripts/*