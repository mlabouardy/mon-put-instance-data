[![Docker Stars](https://img.shields.io/docker/pulls/mlabouardy/mon-put-instance-data.svg)](https://hub.docker.com/r/mlabouardy/mon-put-instance-data/) 
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE) [![CircleCI](https://circleci.com/gh/mlabouardy/mon-put-instance-data/tree/master.svg?style=svg&circle-token=84815df17fc7fed1212bedec59274ddcd9146d57)](https://circleci.com/gh/mlabouardy/mon-put-instance-data/tree/master) [![Go Report Card](https://goreportcard.com/badge/github.com/mlabouardy/mon-put-instance-data)](https://goreportcard.com/report/github.com/mlabouardy/mon-put-instance-data) [![Docker Stars](https://img.shields.io/github/issues/mlabouardy/mon-put-instance-data.svg)](https://github.com/mlabouardy/mon-put-instance-data/issues)  

## Download

Below are the available downloads for the latest version of CLI (1.0.0). Please download the proper package for your operating system and architecture.

### Linux:

```
wget https://s3.us-east-1.amazonaws.com/mon-put-instance-data/1.0.0/linux/mon-put-instance-data
```

### Windows:

```
wget https://s3.us-east-1.amazonaws.com/mon-put-instance-data/1.0.0/windows/mon-put-instance-data
```

## How to use

* Setup an IAM Policy:

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "1",
            "Effect": "Allow",
            "Action": "cloudwatch:PutMetricData",
            "Resource": "*"
        }
    ]
}
```

* Start metrics collector:

```
mon-put-instance-data --memory --swap --disk --network --docker --duration 1
```

## Metrics

* Memory
    * Memory Utilization (%)
    * Memory Used (Mb)
    * Memory Available (Mb)
* Swap
    * Swap Utilization (%)
    * Swap Used (Mb)
* Disk
    * Disk Space Utilization (%)
    * Disk Space Used (Gb)
    * Disk Space Available (Gb)
* Network
    * Bytes In/Out
    * Packets In/Out
    * Errors In/Out
* Docker
    * Memory Utilization per Container
    * CPU User/System per Container

## Supported AMI

* Amazon Linux
* Amazon Linux 2
* Ubuntu 16.04
* Microsoft Windows Server