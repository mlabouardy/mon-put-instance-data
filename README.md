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

## How to use

```
cloudwatch --memory --swap --disk --duration 1
```

* IAM Policy

```
{
    "Effect": "Allow",
    "Statement": "cloudwatch:PutMetricData"
    "Resource": "*"
}
```

## TO DO

* Docker
* Network