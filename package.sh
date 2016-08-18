#/usr/bin/bash
env GOOS=linux GOARCH=amd64 ./build.sh
mkdir -p dynip-dist/extra
cp dynipd/dynipd dynip-dist/
cp dynipd/extra/config.example.yaml dynip-dist/extra
cp dynipd/extra/dynipd.service dynip-dist/extra
cp dynipd/extra/post-install.sh dynip-dist/extra
cp dynipd/extra/install.sh dynip-dist
tar cvzf dynip-dist.tgz dynip-dist
rm -r dynip-dist