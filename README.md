# SAVvie

## To run the website:

First-time setup requires several gigabytes of disk space and 5-30 minutes, depending on your Internet connection and computer speed.

1. Install [Vagrant](https://www.vagrantup.com/).
1. Install [VirtualBox](https://www.virtualbox.org/).
2. Download the source code (with Git, or the "Clone or download" button on [GitHub](https://github.com/tylermumford/pl-final-project))
3. Unzip the source code if necessary. This is the "working directory" (the one with the README.md file inside).
4. Open a terminal/command prompt and use the `cd` command to navigate to the working directory.
5. Run `vagrant box add ubuntu/trusty32`. This may take a while.
6. Run `vagrant up`. This may also take a while.
7. Run `vagrant ssh`. This shouldn't take too long. If you're on Windows, this command may require Cygwin (or some other implementation of the `ssh` command) to be installed.
8. You should now have a terminal for an Ubuntu virtual machine.
9. In that terminal, run `/vagrant/scripts/restart.sh`. If you see errors, try running `dos2unix`.
10. You should see a message about a server running on `localhost:8080`.
11. You can now open a browser (such as Chrome) on your computer and navigate to `http://localhost:8080`. You should see the SAVvie homepage.

## Dev notes:

- Vagrant passes port 8080 from host to guest. 
- Caddy runs on port 8080. It proxies some requests to the controller.
- The controller runs on port 8000.