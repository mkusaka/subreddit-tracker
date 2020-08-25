# aws
## [1][Week of Aug 24th - What are you building this week on AWS?](https://www.reddit.com/r/aws/comments/ifoebu/week_of_aug_24th_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/ifoebu/week_of_aug_24th_what_are_you_building_this_week/
---
Share what you are creating on AWS
## [2][Anyone else haivng large AWS incident / outtage causing issues. Appears mostly related to EU-WEST2a](https://www.reddit.com/r/aws/comments/ig90ou/anyone_else_haivng_large_aws_incident_outtage/)
- url: https://i.imgur.com/Ajikl0m.png
---

## [3][New EBS Volume Type (io2) – 100x Higher Durability and 10x More IOPS/GiB](https://www.reddit.com/r/aws/comments/ig0uib/new_ebs_volume_type_io2_100x_higher_durability/)
- url: https://aws.amazon.com/blogs/aws/new-ebs-volume-type-io2-more-iops-gib-higher-durability/?utm_source=feedburner&amp;utm_medium=feed&amp;utm_campaign=Feed%3A+AmazonWebServicesBlog+%28Amazon+Web+Services+Blog%29
---

## [4][EU-WEST-2 ECS/Fargate Outage?](https://www.reddit.com/r/aws/comments/ig91r8/euwest2_ecsfargate_outage/)
- url: https://www.reddit.com/r/aws/comments/ig91r8/euwest2_ecsfargate_outage/
---
Anyone else seeing problems with ECS/Fargate in EU-WEST-2?  All my tasks just went non-responsive simultaneously across 5 clusters.

I see a few "service xxxp failed to describe target health on target-group xxx with (error Internal failure)" messages in the events a few minutes ago, nothing else untoward.


Also seen a few tasks try to start then immediately stop with "Stopped reason ResourceInitializationError: unable to pull secrets or registry auth: pull command failed: : signal: killed"

Also seen a few tasks try to start then stop with "Stopped reason Timeout waiting for network interface provisioning to complete."
## [5][My experience with AWS GuardDuty IDS](https://www.reddit.com/r/aws/comments/ig9ie2/my_experience_with_aws_guardduty_ids/)
- url: https://www.reddit.com/r/aws/comments/ig9ie2/my_experience_with_aws_guardduty_ids/
---
# What is an Intrusion Detection System

An Intrusion Detection System is a device that monitors one or more  assets, in a host or/and a network level, in your environment, alerting  you for a potential malicious activity or a security policy violation.  An Intrusion Detection System can detect if an intrusion has occurred or  an intrusion is in progress and notifies you by generated findings  (alerts).

**Common types of Intrusion Detection Systems**

* **Host-Based IDS:** A type of IDS that gathers information from a machine and analyses this information for potential malicious activity.
* **Network IDS:** A type of IDS that captures network traffic and analyses this traffic to detect potential malicious activity.
* **Wireless IDS:** A type of IDS that captures your surrounding air traffic and analyses this traffic to detect potential malicious activity.

&amp;#x200B;

[Detection Techniques](https://preview.redd.it/ycaey5tsc4j51.png?width=476&amp;format=png&amp;auto=webp&amp;s=2aaf62b8b6d0287a47c407638c51c2869fab4d5e)

# Intrusion Detection Techniques

All of the above common types of IDS, it depends from the vendor,  technology, subject matter experts, problem to solve, etc., are using in  general the following detection techniques:

* **Signature Based Techniques** (detect known attacks, cannot detect unknown attacks)
* **Anomaly Based Techniques** (detect unknown attacks, can produce high number of false positives (over-fitting, imbalanced data, etc.))
* **Hybrid Approach** (combination of the above techniques)

# What is GuardDuty

An AWS threat detection service that continuously monitors for malicious  activity and unauthorized behavior to protect your AWS accounts and  workloads. GuardDuty is a type of IDS that captures various information,  such as API Calls, Network traffic etc., and analyses this traffic to  detect potential malicious activity. GuardDuty Intrusion Detection  System differs from the traditional common types that we described  above.

# How it works

GuardDuty consumes three types of [data sources](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_data-sources.html#guardduty_dns) in order to detect and notify malicious activities in the AWS environment.

* [**Cloudtrail Logs**](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-user-guide.html) are  generated on your AWS account when you create it. When an activity is  occurred then an event is recorded. You can view those events in the  Event History Console. Furthermore, you can create a new trail and  deliver the generated logs to a new created S3 bucket, where it can be  used, for instance, to do cloudtrail analytics through the Amazon Athena  service.
* [**VPC Flow Logs**](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records)  is a feature that enables you to capture information about the IP  traffic going to and from network interfaces in your VPC. Flow log data  can be published to Amazon CloudWatch Logs or Amazon S3.
* [**DNS Logs**](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/logging-monitoring.html)  monitoring (requests/responses) in a case that you use AWS DNS  resolvers for your EC2 instances (this is a the default setting).  GuardDuty can access and process your request and response DNS logs  through the internal AWS DNS resolvers. *If you are using a  3rd party DNS resolver, for example, OpenDNS or GoogleDNS, or if you  set up your own DNS resolvers, then GuardDuty cannot access and process  data from this data source.*

&amp;#x200B;

[How it works \(detection - Notification\)](https://preview.redd.it/j4ec2vbxc4j51.png?width=509&amp;format=png&amp;auto=webp&amp;s=e0414257a414b2e6696a0000e2f9a8e3d779d338)

GuardDuty  uses threat intelligence feeds, such as lists of malicious IP addresses and domains, and machine learning to identify unexpected and potentially unauthorized and malicious activity within your AWS  environment. In a case of a malicious activity, findings (alerts) are  generated in the console with the accompanied severity for  prioritization and further investigation.

It  seems that GuardDuty IDS uses the hybrid approach, including signature  based techniques and machine learning techniques respectively. By using  the hybrid approach GuardDuty IDS tries to detect known attacks  (signature based) and unknown attacks (machine learning).

# GuardDuty Finding Types

In a case of a detection the following finding types are generated:

&amp;#x200B;

https://preview.redd.it/tg83ket6c4j51.png?width=1400&amp;format=png&amp;auto=webp&amp;s=629b7f811316e50e1f6dc9474df444b9293eb63b

&amp;#x200B;

https://preview.redd.it/lqbidlpac4j51.png?width=1400&amp;format=png&amp;auto=webp&amp;s=509f520139c4a13523f7abfde186022075b14ff9

&amp;#x200B;

https://preview.redd.it/z2lrd9ngc4j51.png?width=1400&amp;format=png&amp;auto=webp&amp;s=806c00a738fb5d137815bad9246d1fd890b8eeb4

All the above findings accompanied with a severity as well.

&amp;#x200B;

https://preview.redd.it/6nsojxbkc4j51.png?width=638&amp;format=png&amp;auto=webp&amp;s=136c1f0f80aad3cb7fc59b17e50d6456baa15105

# Experiment (Generate Findings)

1. Place your public IP address to a threat list. This can be achieved by  placing your public IP or IPs to a text file, then create a S3 bucket  and upload your file in the S3 bucket, then go to Settings → Lists → Add  a Threat List (place a name for your list → Place the location of the  S3 bucket containing the threat list file → specify the format of your  list (plaintext, STIX, Fireeye, Proofpoint, OTX, Alienvault). Then you  can login in the AWS console and start playing around. GuardDuty will  start to generate findings mainly under the types **Recon** and **UnauthorizedAccess**.
2. You can use the script the [amazon-guardduty-tester](https://github.com/awslabs/amazon-guardduty-tester) from awlabs, you can change it according to your needs, to generate more finding types such as **CryptoCurrency**, **Trojan**, **Backdoor** and **Recon** as well.
3. Create a new trail from the cloudtrail service and then stop the created  trail. GuardDuty IDS will generate a finding for this malicious  activity under the type  **Stealth:IAMUser/CloudTrailLoggingDisabled**

&amp;#x200B;

[GuardDuty Findings during the experiment](https://preview.redd.it/c9az8j77d4j51.png?width=1377&amp;format=png&amp;auto=webp&amp;s=d5d6a1fad6706828c8833533ddabdeff3d3abc74)

&amp;#x200B;

&gt;!Finding Sample (Stealth:IAMUser/CloudTrailLoggingDisabled)!&lt;

&amp;#x200B;

[Finding Details](https://preview.redd.it/fh4hjbqcd4j51.png?width=814&amp;format=png&amp;auto=webp&amp;s=0ee7c5eb001ac658e24290bbae341fa1d9bb81b6)

&gt;!JSON output!&lt;

&amp;#x200B;

&gt;\[  
  {  
"schemaVersion": "2.0",  
"accountId": "XXXXXXXXXXX",  
"region": "eu-west-1",  
"partition": "aws",  
"id": "XXXXXXXXXXX",  
"arn": "arn:aws:guardduty:eu-west-1:XXXXXXXXXXXX:detector/f4b8d0c1b93c41d9da771dbad7c1f550/finding/72b8d0c4f192369329fa8c168cc71243",  
"type": "Stealth:IAMUser/CloudTrailLoggingDisabled",  
"resource": {  
"resourceType": "AccessKey",  
"accessKeyDetails": {  
"accessKeyId": "XXXXXXXXXXX",  
"principalId": "XXXXXXXXXXX",  
"userType": "IAMUser",  
"userName": "XXXXXXXXXXX"  
}  
},  
"service": {  
"serviceName": "guardduty",  
"detectorId": "f4b8d0c1b93c41d9da771dbad7c1f550",  
"action": {  
"actionType": "AWS\_API\_CALL",  
"awsApiCallAction": {  
"api": "StopLogging",  
"serviceName": "cloudtrail.amazonaws.com",  
"callerType": "Remote IP",  
"remoteIpDetails": {  
"ipAddressV4": "XXX.XXX.XXX.XXX",  
"organization": {  
"asn": "XXXXXXXXXXX",  
"asnOrg": "XXXXXXXXXXX",  
"isp": "XXXXXXXXXXX",  
"org": "XXXXXXXXXXX"  
},  
"country": {  
"countryName": "XXXXXXXXXXX"  
},  
"city": {  
"cityName": "XXXXXXXXXXX"  
},  
"geoLocation": {  
"lat": XXXXXXXXXXX,  
"lon": XXXXXXXXXXX  
}  
},  
"affectedResources": {  
"AWS::CloudTrail::Trail": "arn:aws:cloudtrail:eu-west-1:XXXXXXXXXXX:trail/cloudtrail-demo"  
}  
}  
},  
"resourceRole": "TARGET",  
"additionalInfo": {},  
"evidence": null,  
"eventFirstSeen": "2020-04-22T10:47:19Z",  
"eventLastSeen": "2020-04-22T10:52:53Z",  
"archived": false,  
"count": 3  
},  
"severity": 2,  
"createdAt": "2020-04-22T10:58:20.580Z",  
"updatedAt": "2020-04-22T11:04:06.476Z",  
"title": "AWS CloudTrail trail arn:aws:cloudtrail:eu-west-1:XXXXXXXXXXX:trail/cloudtrail-demo was disabled.",  
"description": "AWS CloudTrail trail arn:aws:cloudtrail:eu-west-1:148059315130:trail/cloudtrail-demo was disabled by XXXXXXXXXXX calling StopLogging under unusual circumstances. This can be attackers attempt to cover their tracks by eliminating any trace of activity performed while they accessed your account."  
  }  
\]

# Conclusion

It seems that GuardDuty helps incident responders, threat analysts and  security professionals in general to detect various threats in their AWS  environment with a few clicks. GuardDuty provides a simple but very  powerful user interface, as all the indicators about a threat are there  (either in the User Interface or in the JSON view). Users can upload and  activate their own or a third party threat list for strengthen their  detection strategy. Finding types are accompanied with a severity to  help analysts to prioritize their actions.
## [6][AWS S3 - All You Need To Know About This Service](https://www.reddit.com/r/aws/comments/ig9y85/aws_s3_all_you_need_to_know_about_this_service/)
- url: https://catalins.tech/aws-s3-all-you-need-to-know-about-this-service
---

## [7][New EBS Volume Type (io2) – 100x Higher Durability and 10x More IOPS/GiB](https://www.reddit.com/r/aws/comments/ig0q75/new_ebs_volume_type_io2_100x_higher_durability/)
- url: https://aws.amazon.com/blogs/aws/new-ebs-volume-type-io2-more-iops-gib-higher-durability/
---

## [8][OpenVPN to access RDS](https://www.reddit.com/r/aws/comments/ig9p87/openvpn_to_access_rds/)
- url: https://www.reddit.com/r/aws/comments/ig9p87/openvpn_to_access_rds/
---
Anyone have a good KB or rundown what needs to be setup so that the OpenVPN Access server (ec2) can route traffic to an RDS instance?

So far I have been able to setup a OpenVPN Access Server, connect to that server no problem. 

What I think is messed up is my security groups because I can not get traffic to access the RDS instance via the VPN. 

&amp;#x200B;

Any input would be helpful. Thanks.
## [9][Visualizing DynamoDB sensor data in a chart](https://www.reddit.com/r/aws/comments/ig9exq/visualizing_dynamodb_sensor_data_in_a_chart/)
- url: https://www.reddit.com/r/aws/comments/ig9exq/visualizing_dynamodb_sensor_data_in_a_chart/
---
I'm creating a web page with React. I'm gathering sensor data and storing it in DynamoDB. I would like to visualize the temperature, humidity etc. data on a line chart. I have "hourly", "daily", "weekly", "monthly", "yearly" etc. buttons to choose how much data is being visualized. Measurements are made every minute so we are gathering a lot of data. Right now I have a lambda function which queries the wanted amount of measurements between two timestamps. The timestamps in dynamo are strings atm. I'm visualizing the data on my chart but the query takes way too long as I'm using every data point. Should I somehow simplify the data if I want to show the weekly or monthly graph? What are the best ways to do this? Are there AWS tools to simplify this process and make it quicker and smoother? How would you visualize the data on a hourly, daily, monthly, weekly etc. basis? 

My query-function is below. Getting 2500 items takes over 5 seconds.

&amp;#x200B;

def lambda\_handler(event, context):

sensor\_id = event\["params"\]\["path"\]\["lambda\_time"\]

timestamp\_from = event\["params"\]\["querystring"\]\["timestamp\_from"\]

timestamp\_to = event\["params"\]\["querystring"\]\["timestamp\_to"\]

response = table.query(

KeyConditionExpression=Key('sensorID').eq(sensor\_id) &amp; Key('timeStamp').between(timestamp\_from, timestamp\_to),

)

return response\['Items'\]
## [10][Can you export the Organizations accounts information from the Management Console?](https://www.reddit.com/r/aws/comments/ifw8o4/can_you_export_the_organizations_accounts/)
- url: https://www.reddit.com/r/aws/comments/ifw8o4/can_you_export_the_organizations_accounts/
---
Hello Community,

I desperately need a way to export the information (specifically the account name/alias) for each account under my Organization. I can't store the exported data on my personal laptop since this data includes my company's customers. 

I have two solutions set up to do this programmatically where the export never leaves our approved security boundaries, but my solutions are caught up in my "corporate approval red tape". 

I would be able to leverage AWS CLI or AWS Tools for PowerShell from my company workstation, but I use a virtual desktop and I need admin rights to install the CLI and/or appropriate PowerShell Modules. Again, I'm waiting on an approval from our tech shop to install these for me.

  
In the meantime, is there a way to effectively do this through the Management Console so that I can simply access the Management Console from my company virtual desktop and dump the file to a company share?
## [11][Send all Cloudwatch Data to the root account in an organization](https://www.reddit.com/r/aws/comments/ig92jk/send_all_cloudwatch_data_to_the_root_account_in/)
- url: https://www.reddit.com/r/aws/comments/ig92jk/send_all_cloudwatch_data_to_the_root_account_in/
---
Hello,

I'm new to the AWS Services in General.We are using 1 Main Account managing an organization with 30 Sub IAM User, each IAM User has 1-2 EC2 Instances with ELB and RDS or a Webserver on Nginx running.I want to Monitor every Cloudwatch activity with my Main account to simplify exporting data to a Monitoring Solution, or to simplify Cloudwatch Activity in General.

I've read through following article  [https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/CloudWatchEvents-CrossAccountEventDelivery.html](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/CloudWatchEvents-CrossAccountEventDelivery.html)

and manually created a Event rule for each acc which includes the cloudwatch service and the main acc-id as the target. In the Main-acc i createt an Eventbus receiving form everybody in the organization. But I didn't get any data in the cloudwatch screen on the main acc, besides the servers running on the root acc.  Is there a way to automatically send every CW data from any subIAM-ACC to the Main-Acc ?

Using the AWS-CLI getting following error: "The config profile(sender) could not be found" if I want to creat a new role
