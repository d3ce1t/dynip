#/usr/bin/bash
dst_dir="/opt/dynipd"
mkdir -p $dst_dir
cp dynipd $dst_dir
cp extra/config.example.yaml $dst_dir
./extra/post-install.sh
echo 'DynIP Daemon installed'
echo 'Please run systemctl start dynipd.service to star the daemon right now'