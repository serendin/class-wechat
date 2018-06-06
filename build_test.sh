#!/bin/sh

export PATH="$PATH:$GOPATH/bin"

OLD_GOPATH=$GOPATH

scriptPath=$(cd `dirname $0`; pwd)
cd $scriptPath/../../

NEW_GOPATH=$('pwd')
export GOPATH=$OLD_GOPATH:$NEW_GOPATH

cd src

if [ -f ".server_cw.pid" ]; then
  kill -9 `cat .server_cw.pid`
  rm -rf ".server_cw.pid"
fi



echo "==========run=========="
if [ -f "nohup.out" ]; then
  rm -rf "nohup.out"
fi


nohup revel run class-wechat test &


export GOPATH=$OLD_GOPATH
