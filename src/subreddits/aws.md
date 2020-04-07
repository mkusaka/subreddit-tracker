# aws
## [1][How we run EKS with spot instances but fallback to on-demand](https://www.reddit.com/r/aws/comments/fw2pv1/how_we_run_eks_with_spot_instances_but_fallback/)
- url: https://blog.doit-intl.com/running-eks-workloads-on-spot-instances-with-on-demand-instances-fallback-14bef39ce689
---

## [2][How to limit data usage inside of a EC2 instance ?](https://www.reddit.com/r/aws/comments/fwil6o/how_to_limit_data_usage_inside_of_a_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/fwil6o/how_to_limit_data_usage_inside_of_a_ec2_instance/
---
Hello !  
We are currently deploying a VDI solution using Ubuntu instances + Nomachine (Amazon Workspaces is not available in our region)  


We would like to prevent our students from using more than 5 gb of data each day.   
I tried to use iptables with the "Quota" module, but for some reason, even if I explicitly allow SSH and Nomachine in the first rules, put the quota rules for the TCP protocol after that, and the DROP rules at the very end, the EC2 instance becomes inaccessible as soon as the quota is reached, even through SSH.  


Is there an Amazon service that could help ? Do you have idea I could try to get the result I want ?  
Have an excellent day :)
## [3][How do I gain access from aws-sdk to private bucket if IAM is not the creator of the bucket.](https://www.reddit.com/r/aws/comments/fwkg34/how_do_i_gain_access_from_awssdk_to_private/)
- url: https://www.reddit.com/r/aws/comments/fwkg34/how_do_i_gain_access_from_awssdk_to_private/
---
So, I am trying to retrieve files from private bucket on the frontend. 

I have created bucket from lets say accountA, and want to access it from accountB. my current configuration looks like this

bucket policy:
```json
{
    "Version": "2012-10-17",
    "Id": "http referer policy example",
    "Statement": [
        {
            "Sid": "Allow get requests originating from localhost",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::accountB:user/accountB"
            },
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::bucketname/*",
            "Condition": {
                "StringLike": {
                    "aws:Referer": [
                        "http://localhost:4200/*"
                    ]
                }
            }
        }
    ]
}
```

bucket cors config:
```xml
&lt;?xml version="1.0" encoding="UTF-8"?&gt;
&lt;CORSConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/"&gt;
&lt;CORSRule&gt;
    &lt;AllowedOrigin&gt;*&lt;/AllowedOrigin&gt;
    &lt;AllowedMethod&gt;GET&lt;/AllowedMethod&gt;
    &lt;AllowedMethod&gt;HEAD&lt;/AllowedMethod&gt;
    &lt;MaxAgeSeconds&gt;3000&lt;/MaxAgeSeconds&gt;
    &lt;AllowedHeader&gt;Authorization&lt;/AllowedHeader&gt;
&lt;/CORSRule&gt;
&lt;/CORSConfiguration&gt;
```
and i have attached this policy to accountB:
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AllowGroupToSeeBucketListAndAlsoAllowGetBucketLocationRequiredForListBucket",
            "Action": [
                "s3:ListAllMyBuckets",
                "s3:GetBucketLocation"
            ],
            "Effect": "Allow",
            "Resource": [
                "arn:aws:s3:::*"
            ]
        },
        {
            "Sid": "AllowRootLevelListingOfCompanyBucket",
            "Action": [
                "s3:ListBucket"
            ],
            "Effect": "Allow",
            "Resource": [
                "arn:aws:s3:::bucketname"
            ],
            "Condition": {
                "StringEquals": {
                    "s3:prefix": [
                        ""
                    ],
                    "s3:delimiter": [
                        "/"
                    ]
                }
            }
        },
        {
            "Sid": "AllowUserToReadWriteObjectData",
            "Action": [
                "s3:GetObject",
                "s3:PutObject"
            ],
            "Effect": "Allow",
            "Resource": [
                "arn:aws:s3:::bucketname/*"
            ]
        }
    ]
}
```


from my server side, where i am using accountA credentials to upload files to bucket, it works without any problems, however, when i use accountB credentials on my front end to do ```getObject``` on private files, i get forbidden 403 and ``` No'Access-Control-Allow-Origin'``` error at the same time. 

could anyone help?
## [4][Free tier ended up with a $17k/2400% increase on monthly bill due to UX bug in Redshift configurator, what are our odds of leniency? :O](https://www.reddit.com/r/aws/comments/fw367k/free_tier_ended_up_with_a_17k2400_increase_on/)
- url: https://www.reddit.com/r/aws/comments/fw367k/free_tier_ended_up_with_a_17k2400_increase_on/
---
I recently started working for a startup and had a plan to setup a Redshift DC2.large instance for a data warehouse project and to be eligible for the 2 month free tier as described [here](https://aws.amazon.com/redshift/free-trial/). After reading up on all the selling points of Amazon Redshift, logged into AWS console, clicked Redshift -&gt; Clusters -&gt; Create cluster. At this point I involved a colleague to have four eyes on the process. After a while I was able to get my colleague on the line to proceed creating the cluster. Selected DC2.large 1 node (estimated $266 p/m) and some other options and selected create, at which point I think my session had ran out and it asked me to reload and sign in again, which was somewhat frustrating at this point. Upon signing in is where I think a major UX/CX flaw (see below) caused me to proceed create a completely different cluster which also is extremely expensive. It accidentally proceeded with the RA3.16XL with 2 nodes since it reset the form and all prior config which ended up costing $17.331 for a service we did not intend to use or fully utilise. This is a 2400% increase on our monthly avg. and blows our entire annual budget in a month. What is your guys experiences? Is Amazon AWS understanding of such situations? Or should I prepare my vacancy over this embarrassing accident :/

https://reddit.com/link/fw367k/video/srbhgw1vh8r41/player
## [5][Single Page Application on Beanstalk](https://www.reddit.com/r/aws/comments/fwj52v/single_page_application_on_beanstalk/)
- url: https://www.reddit.com/r/aws/comments/fwj52v/single_page_application_on_beanstalk/
---
I'm using Flask as the web server, and Angular as my SPA in a /static directory. I implemented a catch-all end point on my Flask app to try solve the issue where users are given a 404 when they refresh a page, but it didnt solve the issue. Does anyone know how to fix this? 

The catch-all looks like:
application.route('/', defaults={'path': ''})
application.route('/static/&lt;path:path&gt;')
def main(path):
     return render_template('index.html')
## [6][AWS Serverless Applications with Code Commit](https://www.reddit.com/r/aws/comments/fwb0a7/aws_serverless_applications_with_code_commit/)
- url: https://www.reddit.com/r/aws/comments/fwb0a7/aws_serverless_applications_with_code_commit/
---
My question is....

**Can someone point me to a tutorial on how build a CI/CD pipeline for deploying code that is already hosed in AWS Commit that uses Lambda for compute?**  

Excuse my poor use of terminology, I am new to this.  We have all of the pieces to do it, just need to put it together.  We have CloudFormation Templates for all of the resources(API GW, LoadBalancer, Lambda, DynamoDB, SQS), Lambda Code checked into our AWS Commit repo, and S3 buckets to store artifacts.  I feel like there is some piece of technology that I am missing to glue these together into a deployed app.

I have been through the SAM tutorial, but I must have missed a key component during that session.
## [7][Apigee on aws](https://www.reddit.com/r/aws/comments/fwcaww/apigee_on_aws/)
- url: https://www.reddit.com/r/aws/comments/fwcaww/apigee_on_aws/
---
Did anyone use the apigee gateway in an AWS deployment as an api gateway.
## [8][[CloudFormation] What does the error "Transforms defined as maps require Name key." really mean?](https://www.reddit.com/r/aws/comments/fwfx83/cloudformation_what_does_the_error_transforms/)
- url: https://www.reddit.com/r/aws/comments/fwfx83/cloudformation_what_does_the_error_transforms/
---
I googled and searched everywhere and I cant find any reference to this error so far.

Does anyone know how to interpret this?

I get this error on the CLI as well as when I validate it in the designer. Unfortunately, the error does not give any indication of where the error occurs (like a line number), which would have been super helpful.

\[SOLVED\] Go figure, just shortly after I asked it, I solved it myself. Turns out you cant have an empty \`Transform\` clause in your JSON file, like such: `"Transform" : {},`
## [9][Reading workspaces tags within guest OS](https://www.reddit.com/r/aws/comments/fwfdbu/reading_workspaces_tags_within_guest_os/)
- url: https://www.reddit.com/r/aws/comments/fwfdbu/reading_workspaces_tags_within_guest_os/
---
I'm trying to find out if there is a way to read the tags set on our workspaces from within the guess OS itself?  We're trying to find a way to pull this data into our SCCM system so we can determine which software needs to be deployed and we're coming up short.  We've been going down the route of possibly using the powershell cmdlets to write registry keys and haven't gotten anything to work reliably.  Thank you in advance.
## [10][EC2 internet seems so slow](https://www.reddit.com/r/aws/comments/fw8hyg/ec2_internet_seems_so_slow/)
- url: https://www.reddit.com/r/aws/comments/fw8hyg/ec2_internet_seems_so_slow/
---
I have made a couple of instances, one with t3.medium and the other with m4.xlarge and the internet seems so lackluster, is there anything that I have to enable to get full speed or am I missing something?

btw I use wget to FTP files from a hosted directory into my EC2 instance

EDIT: I get a 5 megabyte/second internet, as opposed to the 10Gigabit/s that's promised
