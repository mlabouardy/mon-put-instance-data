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
mon-put-instance-data --memory --swap --disk --docker --duration 1
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