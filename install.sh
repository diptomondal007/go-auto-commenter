#! /bin/bash

exec_curl(){
  echo "Found auto commenter Version: $VERSION"
  echo "Download may take few minutes depending on your internet speed"
  echo "Downloading go auto commenter  to $2"
  curl -L --silent --connect-timeout 30 --retry-delay 5 --retry 5 -o "$2" "$1"
}

OS=`uname`
ARCH=`uname -m`
VERSION=$1
URL=https://github.com/diptomondal007/go-auto-commenter
TARGET=/usr/local/bin/autocomment
INSTALL_START_MESSAGE="Installing go auto commenter ......"
INSTALL_END_MESSAGE="√ Installation finished ...."

if [ "$VERSION" == "" ]; then
  LATEST_RELEASE=$(curl -L -s -H 'Accept: application/json' $URL/releases/latest)
  VERSION=$(echo $LATEST_RELEASE | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')
fi

if [ "$OS" == "Darwin" ]; then
  exec_curl $URL/releases/download/"$VERSION"/mac_amd64 $TARGET
  echo "$INSTALL_START_MESSAGE"
  chmod +x $TARGET
  echo "$INSTALL_END_MESSAGE"
  autocomment

elif [ "$OS" == "Linux" ]; then
      if [ "$ARCH" == "x86_64" ]; then
        exec_curl $URL/releases/download/$VERSION/linux_amd64 $TARGET
        echo "$INSTALL_START_MESSAGE"
        chmod +x $TARGET
        echo "$INSTALL_END_MESSAGE"
        autocomment
      fi

      if [ "$ARCH" == "i386" ]; then
          exec_curl $URL/releases/download/$VERSION/linux_i386 $TARGET
          echo "$INSTALL_START_MESSAGE"
          chmod +x $TARGET
          echo "$INSTALL_END_MESSAGE"
          autocomment
      fi
fi