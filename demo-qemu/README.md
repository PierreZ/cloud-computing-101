# Cheatsheet

```shell
# Download an Linux image, like ubuntu

# Create a disk
qemu-img create -f qcow2 ubuntu.qcow2 16G

# Then start a vm
qemu-system-x86_64 \
    -enable-kvm \
    -m 8000 \
    -smp cpus=4 \
    -nic user,model=virtio \
    -drive file=ubuntu.qcow2,media=disk,if=virtio \
    -cdrom /home/pierrez/Downloads/
```