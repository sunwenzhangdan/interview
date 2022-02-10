<!--
 * @Date: 2022-01-29 12:42:15
 * @LastEditors: seven sun 
 * @LastEditTime: 2022-01-29 14:13:30
 * @FilePath: /interview/k8s/k8s.md
-->

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

You can now join any number of control-plane nodes by copying certificate authorities
and service account keys on each node and then running the following as root:

  kubeadm join k8s-master:6443 --token p68lm6.wl6qcdq30prymlpg \
    --discovery-token-ca-cert-hash sha256:ae6a7ab46165a48d4c82dc65e8ddcb3f8fe56b2a400028aa1336816eadfb2323 \
    --control-plane

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join k8s-master:6443 --token p68lm6.wl6qcdq30prymlpg \
    --discovery-token-ca-cert-hash sha256:ae6a7ab46165a48d4c82dc65e8ddcb3f8fe56b2a400028aa1336816eadfb2323


所有的机器安装docker

sudo yum remove docker*
sudo yum install -y yum-utils
 
#配置docker的yum地址
sudo yum-config-manager \
--add-repo \
http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
 
 
#安装指定版本
sudo yum install -y docker-ce-20.10.7 docker-ce-cli-20.10.7 containerd.io-1.4.6
 
#   启动&开机启动docker
systemctl enable docker --now
 
# docker加速配置
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://82m9ar63.mirror.aliyuncs.com"],
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker



基础设置（三台机器全部都要设置）

1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
# 将 SELinux 设置为 permissive 模式（相当于将其禁用）
sudo setenforce 0
sudo sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config
 
#关闭swap关闭分区
swapoff -a 
sed -ri 's/.*swap.*/#&/' /etc/fstab
 
#允许 iptables 检查桥接流量
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF
 
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sudo sysctl --system

#关闭防火墙
#适用于当前pod只有一个容器
kubectl exec -it superset-3fpnq -- /bin/bash      
