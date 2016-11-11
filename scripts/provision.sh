echo "Downloading package lists..."
apt-get update -qq

echo "Installing packages..."
apt-get install golang mono-runtime mono-mcs

echo "Compiling source..."
/vagrant/scripts/build.sh