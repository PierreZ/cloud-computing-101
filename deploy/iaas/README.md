# Deploy my-app by hand

```shell
curl -L --output /bin/my-app https://github.com/PierreZ/cloud-computing-101/releases/download/latest/my-app
chmod +x /bin/my-app
curl -L --output /etc/systemd/system/demo.service https://raw.githubusercontent.com/PierreZ/cloud-computing-101/refs/heads/main/deploy/iaas/my-app.service
systemctl daemon-reload
systemctl enable --now demo
```
