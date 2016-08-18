#/usr/bin/bash
dst_dir="/opt/dynipd"
user="dynipd"

# Add dynipd user with no shell
useradd -r -s /bin/false $user

# Permissions
chown -R $user $dst_dir

# Configure service to start automatically
# NOTE: Debian assumed
cp extra/dynipd.service /lib/systemd/system
systemd enable dynipd.service