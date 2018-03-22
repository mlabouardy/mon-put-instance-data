## Metrics

### Memory

* Memory Utilization (%)
* Memory Used (Mb)
* Memory Available (Mb)

### Swap

* Swap Utilization (%)
* Swap Used (Mb)

### Disk

* Disk Space Utilization (%)
* Disk Space Used (Gb)
* Disk Space Available (Gb)

### Docker

* Memory Utilization per Container
* CPU User/System per Container

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
cloudwatch --memory --swap --disk --docker --duration 1
```

## TO DO

* Docker
* Network