# aws
## [1][A public data lake for analysis of COVID-19 data | Amazon Web Services](https://www.reddit.com/r/aws/comments/fxmklc/a_public_data_lake_for_analysis_of_covid19_data/)
- url: https://aws.amazon.com/blogs/big-data/a-public-data-lake-for-analysis-of-covid-19-data/
---

## [2][Amazon Elastic Container Service now supports Amazon EFS file systems](https://www.reddit.com/r/aws/comments/fxdpcu/amazon_elastic_container_service_now_supports/)
- url: https://aws.amazon.com/blogs/aws/amazon-ecs-supports-efs/
---

## [3][The DynamoDB Book: Not mine (written by a friend), but genuinely one of the BEST resources for DynamoDB I have read yet.](https://www.reddit.com/r/aws/comments/fx7kws/the_dynamodb_book_not_mine_written_by_a_friend/)
- url: https://dynamodbbook.com/
---

## [4][Amazon RDS Proxy with PostgreSQL Compatibility (Preview)](https://www.reddit.com/r/aws/comments/fxi0iz/amazon_rds_proxy_with_postgresql_compatibility/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/04/amazon-rds-proxy-with-postgresql-compatibility-preview/
---

## [5][AWS Directory Service w/ MFA](https://www.reddit.com/r/aws/comments/fxs87m/aws_directory_service_w_mfa/)
- url: https://www.reddit.com/r/aws/comments/fxs87m/aws_directory_service_w_mfa/
---
We currently utilize a Microsoft based AWS AD for all of our user management including AWS Client Endpoint VPN. I'd like to add MFA to the AD/VPN. Any recommendations on RADIUS setup for this?
## [6][Creating Serverless React Apps on AWS](https://www.reddit.com/r/aws/comments/fxs7ub/creating_serverless_react_apps_on_aws/)
- url: https://medium.com/@c.andrewlong/creating-serverless-react-apps-on-aws-bd038ce76a81
---

## [7][MS SQL License Key](https://www.reddit.com/r/aws/comments/fxmxge/ms_sql_license_key/)
- url: https://www.reddit.com/r/aws/comments/fxmxge/ms_sql_license_key/
---
We have a few instances running Windows 2016 with Microsoft Enterprise SQL Server. I am trying to install PowerBI Report Server on them. During the installation process, PBI  says it requires software assurance and the license key for the MS SQL. 

Am I missing something here? I am not sure how to proceed since I didn't use any BYOL.
## [8][How to reference existing manually created resources in SAM cloud formation template.yaml?](https://www.reddit.com/r/aws/comments/fxquqz/how_to_reference_existing_manually_created/)
- url: https://www.reddit.com/r/aws/comments/fxquqz/how_to_reference_existing_manually_created/
---
I've a dynamoDB table setup manually and has some data in it. And I want to reference this table in in the new SAM application without creating another table. 

All the articles I've seen on this relate to importing these resources in Cloud Formation stack. If I do, can I use this template in SAM application?

or is there a way to reference the existing resources in new SAM application?
## [9][Can't run mysql with an EBS volume in EKS kubernetes](https://www.reddit.com/r/aws/comments/fxqtdj/cant_run_mysql_with_an_ebs_volume_in_eks/)
- url: https://www.reddit.com/r/aws/comments/fxqtdj/cant_run_mysql_with_an_ebs_volume_in_eks/
---
I am trying to run mysql in EKS (kubernetes v 1.15). It runs well with a local volume (`hostPath`).

    apiVersion: v1
    kind: PersistentVolume
    metadata:
      name: mysql
      labels:
        type: local
    spec:
      storageClassName: manual
      capacity:
        storage: 500Mi
      accessModes:
        - ReadWriteOnce
      hostPath:
        path: "/data/mysql"

But when I try to move to an EBS volume it ceases to work. First I tried using a pvc to the default `StorageClass`, the I created my own `StorageClass`

    apiVersion: v1
    kind: PersistentVolumeClaim
    metadata:
      name: mysql
    spec:
      storageClassName: manual
      accessModes:
        - ReadWriteOnce
      resources:
        requests:
          storage: 500Mi

then I tried to attach it to the pod directly using `awsBlockStore`

    volumes:
        - name: mysql-data
          awsElasticBlockStore:
            volumeID: "vol-09df5a8d4c0f2da62"
            fsType: ext4

And nothing seamed to work. the output of `kubectl logs mysql-pod` is he following:

    2020-04-09 11:23:26+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.7.29-1debian10 started.
    2020-04-09 11:23:26+00:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
    2020-04-09 11:23:26+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.7.29-1debian10 started.
    2020-04-09 11:23:26+00:00 [Note] [Entrypoint]: Initializing database files
    2020-04-09T11:23:26.919072Z 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
    2020-04-09T11:23:26.920523Z 0 [ERROR] --initialize specified but the data directory has files in it. Aborting.
    2020-04-09T11:23:26.920548Z 0 [ERROR] Aborting

What am I doing wrong ?
## [10][AWS started running Folding@Home on Spot Instances](https://www.reddit.com/r/aws/comments/fx1aw0/aws_started_running_foldinghome_on_spot_instances/)
- url: https://folding.extremeoverclocking.com/team_summary.php?s=&amp;t=238068
---

