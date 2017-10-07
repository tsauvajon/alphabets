#!/bin/bash

# Sanity checks
relog=false
# Check for docker group
if ! $( id -Gn | grep -wq docker ); then
  sudo usermod -aG docker linux1
  echo "ID linux1 was not a member of the docker group. This has been corrected."
  relog=true
fi
# Check PATH for /data/npm/bin
if ! $( echo $PATH | grep -q /data/npm/bin ); then
  echo "export PATH=/data/npm/bin:$PATH" >> $HOME/.profile
  echo "PATH was missing '/data/npm/bin'. This has been corrected."
  relog=true
fi
# Relog needed?
if [[ "$relog" = true ]]; then
  echo "Some changes have been made that require you to log out and log back in."
  echo "Please do this now and then re-run this script."
  exit 1
fi
# Ensure /data exists
if [[ ! -d "/data" ]]; then
  echo "/data disk is missing. It could take up to 10 minutes to format and mount the /data disk. Issue 'df -h' to verify the /data disk is available before running this script again. When /data is available, please run this script again."
  exit 2
fi
# END Sanity checks

printf "

IBM Master the Mainframe

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
:::::::::::''  ''::'      '::::::  ::::::::::::::'.:::::::::::::::
:::::::::' :. :  :         :::: :  :::::::::::.:::':::::::::::::::
::::::::::  :   :::.       ::: M :::::::..::::'     :::: : :::::::
::::::::    :':  '::'     '' M   M :::::: :'           '' ':::::::
:'        : '   :  ::    . M       M   '                        .:
:               :  .:: . M           M                         :::
:. .,.        :::  ':: M M M       M M M                 .:...::::
:::::::.      '      M   M   M   M   M   M               :: :::::.
::::::::           M     M     M     M     M   '    '   .:::::::::
::::::::.        ::: M   M           M   M :         ''' :::::::::
::::::::::      :::::: M M           M M             :::::::::::::
: .::::::::.   .:''::::: M           M   ::   :   '::.::::::::::::
:::::::::::::::. '  '::::: M       M   :::::.:.:.:.:.:::::::::::::
:::::::::::::::: :     ':::: M   M  ' ,:::::::::: : :.:'::::::::::
::::::::::::::::: '     :::::: M    . :'::::::::::::::' ':::::::::
::::::::::::::::::''   :::::::: : :' : ,:::vem:::::'      ':::::::
:::::::::::::::::'   .::::::::::::  ::::::::::::::::       :::::::
:::::::::::::::::. .::::::::::::::::::::::::::::::::::::.'::::::::

IBM Master the Mainframe

"


#Install NodeJS
echo -e “*** install_nodejs ***”
cd /tmp
wget -q https://nodejs.org/dist/v6.9.5/node-v6.9.5-linux-s390x.tar.gz
cd /usr/local && sudo tar --strip-components=1 -xzf /tmp/node-v6.9.5-linux-s390x.tar.gz
echo -e “*** Done withe NodeJS ***\n”


#Setup and install docker-compose
echo -e “*** Installing docker-compose. ***\n”
sudo zypper install -y python-pyOpenSSL python-setuptools
sudo easy_install pip
sudo pip install docker-compose==1.13.0
echo -e “*** Done with docker-compose. ***\n”

#Install Hyperledger Composer Components
echo -e “*** Installing Hyperledger Composer command line tools. ***\n”
mkdir /data/linux1/ 
npm config set prefix '/data/npm'
npm config set cache /data/linux1/.npm
export PATH=/data/npm/bin:$PATH
cd /data/linux1/
npm install -g composer-cli@0.13.1

echo -e “*** Installing Hyperledger Composer rest server. ***\n”
npm install -g composer-rest-server@0.13.1

echo -e “*** Installing Hyperledger Composer playground. ***\n”
npm install -g composer-playground@0.13.1

echo -e "*** Clone and install the Coposer Tools repository.***\n"
git clone https://github.com/hyperledger/composer-tools
cd composer-tools/
npm install
cd packages/fabric-dev-servers/
npm install
cd /data/linux1/composer-tools/packages/fabric-dev-servers/
./downloadFabric.sh
./startFabric.sh
./createComposerProfile.sh
mkdir /data/playground/
nohup composer-playground >/data/playground/playground.stdout 2>/data/playground/playground.stderr & disown
sudo iptables -I INPUT 1 -p tcp --dport 8080 -j ACCEPT
sudo iptables -I INPUT 1 -p tcp --dport 3000 -j ACCEPT
sudo iptables -I INPUT 1 -p tcp --dport 1880 -j ACCEPT
sudo bash -c "iptables-save > /etc/linuxone/iptables.save"

#Install NodeRed
echo -e "*** Installing NodeRed. ***\n"
npm install -g node-red
nohup node-red >/data/playground/nodered.stdout 2>/data/playground/nodered.stderr & disown

# Persist PATH setting
echo "export PATH=/data/npm/bin:$PATH" >> $HOME/.profile

# Persist docker group addition
sudo usermod -aG docker linux1

echo "Please log out of this system and log back in to pick up the group and PATH changes."

