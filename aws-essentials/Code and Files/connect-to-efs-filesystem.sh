sudo yum -y install amazon-efs-utils
sudo mkdir /mnt/efs
sudo mount -t efs fs-af29dd97:/ /mnt/efs