#/usr/bin/bash
dynip_pkg="peeple/dynip/dynip"
dynipd="peeple/dynip/dynipd"

function build_and_install {
  echo "Build $1"
  go build $1
}

build_and_install $dynip_pkg
cd dynipd
build_and_install $dynipd
cd ..