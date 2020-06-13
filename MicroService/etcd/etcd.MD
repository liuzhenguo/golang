centos:
安装
yum install -y etcd

开机自启动：
sudo systemctl enable etcd
启动
sudo systemctl start etcd
查看是否启动
sudo systemctl status etcd

etcd的systemd配置文件
/usr/lib/systemd/system/etcd.service
etcd的配置文件
/etc/etcd/etcd.conf
修改ip
ETCD_LISTEN_CLIENT_URLS="http://localhost:2379"
为(自己本机的ip地址)
ETCD_LISTEN_CLIENT_URLS="http://localhost:2379,http://192.168.60.38：2379"
重启etcd
systemctl restart etcd