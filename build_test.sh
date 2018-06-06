#!/bin/sh

export PATH="$PATH:$GOPATH/bin"

OLD_GOPATH=$GOPATH

scriptPath=$(cd `dirname $0`; pwd)
cd $scriptPath/../../

NEW_GOPATH=$('pwd')
export GOPATH=$OLD_GOPATH:$NEW_GOPATH

cd src

if [ -f ".server_admin.pid" ]; then
  kill -9 `cat .server_admin.pid`
  rm -rf ".server_admin.pid"
fi



echo "==========run=========="
if [ -f "nohup.out" ]; then
  rm -rf "nohup.out"
fi


nohup revel run class-wechat test &


export GOPATH=$OLD_GOPATH
