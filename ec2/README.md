# Help, I can't install docker

* You can use Ubuntu dual-boot
* You can use EC2 or other compute engine

# I want to use EC2, please guide me

* Create AWS Account
* Verify payment method (Jenius or any other CC with VISA logo should comply)
* Launch an EC2 instance (Using ubuntu is suggested, ensure you choose the free-tier plan)
* Create pair and download the key
* Look for your instance detail (IPv4 address/IPv4 DNS)
* Login to your EC2 by invoking: `ssh -i ssh -i personal-aws.pem ubuntu@ec2-18-142-95-76.ap-southeast-1.compute.amazonaws.com` (`personal-aws.pem` is your key, while `ubuntu@ec2-18-142-95-76.ap-southeast-1.compute.amazonaws.com` is your IPv4 DNS)
* Install docker:

```bash
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install docker.io
```

* Check if things work:

```bash
sudo docker ps
```

* To copy to/from your compute engine, you can use [SCP](https://linuxize.com/post/how-to-use-scp-command-to-securely-transfer-files/)

```bash
# copy README.md to ec2
scp -i ~/personal-aws.pem README.md ubuntu@ec2-18-142-95-76.ap-southeast-1.compute.amazonaws.com:/home/ubuntu

# copy something.txt from ec2 to local machine
scp -i ~/personal-aws.pem ubuntu@ec2-18-142-95-76.ap-southeast-1.compute.amazonaws.com:/home/ubuntu/something.txt ./
```