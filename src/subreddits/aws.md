# aws
## [1][Week of Sept 21st - What are you building this week in AWS?](https://www.reddit.com/r/aws/comments/ix0dmq/week_of_sept_21st_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/ix0dmq/week_of_sept_21st_what_are_you_building_this_week/
---
Share what you are working on
## [2][Why were the pins removed from the navigation bar?](https://www.reddit.com/r/aws/comments/ixzfho/why_were_the_pins_removed_from_the_navigation_bar/)
- url: https://www.reddit.com/r/aws/comments/ixzfho/why_were_the_pins_removed_from_the_navigation_bar/
---
I had pins setup to 1 click navigate around the console. EC2, R53, Cloudwatch,cloudformation etc etc. Middle click the pin create a new tab to check some logs etc. 

Who in their right mind thought, you know what lets take away users ease of access to our services and make them click 2x as much. /rant
## [3][When using Step Functions, I find that the interaction between different paths is fiddly and hard to get exactly right. I made a web page where you can type in a payload, ResultPath, OutputPath etc and see the final and intermediate steps instantly](https://www.reddit.com/r/aws/comments/iy6wrq/when_using_step_functions_i_find_that_the/)
- url: https://www.mdavis.xyz/step-funcs/
---

## [4][My experience after 4 months of backing up Synology to Amazon Glacier](https://www.reddit.com/r/aws/comments/iy6yzd/my_experience_after_4_months_of_backing_up/)
- url: https://www.reddit.com/r/aws/comments/iy6yzd/my_experience_after_4_months_of_backing_up/
---
I created a small project to [setup AWS to host Synologies backups to Glacier](https://github.com/0x4447/0x4447_product_synology_backup). And after 4 months of using it I must say I'm happy with the results. As of now I stored 31.2 GB and the total costs of this 4 months is of $5.01 where the first month was the most expensive due to the initial backup. The backup is set to be daily at 1AM and I normally add one or two documents a day. Nothing major. And then the occasional few GB video file once a month. Bellow is the full brake down. 

### June 2020 - $3.69

```
    Amazon Glacier EarlyDelete-ByteHrs  
    $0.012 per GB - Early Delete 2.036 GB-Mo $0.01

    Amazon Glacier Requests-Tier1
    $0.050 per 1,000 Requests 72,578.000 Requests $3.63

    Amazon Glacier TimedStorage-ByteHrs$0.06
    $0.004 per GB / month - Storage 13.994 GB-Mo $0.06
```

### July 2020 - $0.61

```
    Amazon Glacier EarlyDelete-ByteHrs
    $0.012 per GB - Early Delete 5.920 GB-Mo $0.02

    Amazon Glacier Requests-Tier1
    $0.050 per 1,000 Requests 9,250.000 Requests $0.46

    Amazon Glacier TimedStorage-ByteHrs
    $0.004 per GB / month - Storage 29.967 GB-Mo $0.12
```

### August 2020 - $0.15

```
    Amazon Glacier EarlyDelete-ByteHrs$
    $0.012 per GB - Early Delete 5.968 GB-Mo $0.02

    Amazon Glacier Requests-Tier1
    $0.050 per 1,000 Requests 63.000 Requests $0.00

    Amazon Glacier TimedStorage-ByteHrs
    $0.004 per GB / month - Storage 31.091 GB-Mo $0.12
```

### September 2020 - $0.56

```
    Amazon Glacier EarlyDelete-ByteHrs
    $0.012 per GB - Early Delete 3.572 GB-Mo $0.01

    Amazon Glacier Requests-Tier1
    $0.050 per 1,000 Requests 31.000 Requests $0.00

    Amazon Glacier TimedStorage-ByteHrs
    $0.004 per GB / month - Storage 18.748 GB-Mo $0.07
```
## [5][I want to do a security audit on my company AWS account, not sure where to start](https://www.reddit.com/r/aws/comments/iy8jvr/i_want_to_do_a_security_audit_on_my_company_aws/)
- url: https://www.reddit.com/r/aws/comments/iy8jvr/i_want_to_do_a_security_audit_on_my_company_aws/
---
where do I start and check/test security vulnerability of our AWS account
## [6][Do you think AWS will open new regions in the US?](https://www.reddit.com/r/aws/comments/iy78oo/do_you_think_aws_will_open_new_regions_in_the_us/)
- url: https://www.reddit.com/r/aws/comments/iy78oo/do_you_think_aws_will_open_new_regions_in_the_us/
---
Hello, currently there are Oregon, N. California, Ohio and N. Virginia. (not counting GovCloud). However, I was wondering if it would be possible to add more regions, like Texas for example. Azure has regions in those places.
## [7][AWS Org testing accounts](https://www.reddit.com/r/aws/comments/iy9q9y/aws_org_testing_accounts/)
- url: https://www.reddit.com/r/aws/comments/iy9q9y/aws_org_testing_accounts/
---
Hi,

In our company we use an AWS org with Azure AD SSO for IAM. We have developers looking to test things like Control Tower, which means they need access to a master account of an AWS org. Obviously giving devs access to our prod master account isn't something we're going to do but creating new accounts solely for testing takes far longer than just spinning up a new sub account, as we still want some semblance of access management and billing access.

&amp;#x200B;

What are you guys doing for testing requirements like this? Do I just have to grin and bear it, or is there a solution I'm missing?
## [8][AWS Perspective - Visualize Workloads as Architecture Diagrams](https://www.reddit.com/r/aws/comments/ixocei/aws_perspective_visualize_workloads_as/)
- url: https://aws.amazon.com/solutions/implementations/aws-perspective/
---

## [9][Careful, ELBs are relatively expensive](https://www.reddit.com/r/aws/comments/iy8vck/careful_elbs_are_relatively_expensive/)
- url: https://www.reddit.com/r/aws/comments/iy8vck/careful_elbs_are_relatively_expensive/
---
Just a reminder to watch your billing stats as ELBs can get very expensive even with the most basic configuration.
## [10][Need help with Kinesis video streams for facial analysis](https://www.reddit.com/r/aws/comments/iy788q/need_help_with_kinesis_video_streams_for_facial/)
- url: https://www.reddit.com/r/aws/comments/iy788q/need_help_with_kinesis_video_streams_for_facial/
---
* I have to process the video stream to identify emotions. For which I need to use AWS Rekognition/custom model on SageMaker. With Kinesis WebRTC javascript sdk, currently video can be streamed only into the kinesis signalling channel. Signalling channel data is available for streaming only and not processing (ML). So, how can I get real time data for processing into Kinesis Streams from the frontend?
   * For streaming the video from frontend to backend into the Kinesis Video Streams for processing, I tested with Kinesis webRTC Javascript SDK and I am facing issues while implementing as mentioned above, so would Chime SDK serve as an alternative to this?
   * In Rekognition, "create-stream-processor" has a settings parameter. This currently only supports FaceSearch. I am looking to Detect and analyze faces. Is that possible with “create-stream-processor" in the Python SDK? Or do I have to use the Java SDK?
## [11][Security September: Cataclysms in the Cloud Formations – One Cloud Please](https://www.reddit.com/r/aws/comments/iy2w1o/security_september_cataclysms_in_the_cloud/)
- url: https://onecloudplease.com/blog/security-september-cataclysms-in-the-cloud-formations
---

