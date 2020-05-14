# aws
## [1][I wrote a guide on building your own resume site using AWS S3, CloudFront, and Route53. Its almost free to host and looks great on a resume.](https://www.reddit.com/r/aws/comments/gjl5wi/i_wrote_a_guide_on_building_your_own_resume_site/)
- url: https://seanjziegler.com/how-to-build-a-free-static-resume-site-with-aws-s3-cloudfront-and-route-53/
---

## [2][I made an app that lets user stores images (I´m using S3 buckets for that), how can I prevent users to upload too many GB of photos so I don´t get charged a lot in my account?](https://www.reddit.com/r/aws/comments/gj904w/i_made_an_app_that_lets_user_stores_images_im/)
- url: https://www.reddit.com/r/aws/comments/gj904w/i_made_an_app_that_lets_user_stores_images_im/
---
Is there any option to prevent this? My intent when doing the app was as a showcase for my portfolio (trying to get my first job as a dev) so the link will be "out there" in public for everyone to see and I would like to have some sort of control regarding how much my users write data in the AWS services I´m using (my particular concern is S3, since photos nowdays are so heavy).
## [3][RRAS begins AWS ELB](https://www.reddit.com/r/aws/comments/gjmbkd/rras_begins_aws_elb/)
- url: https://www.reddit.com/r/aws/comments/gjmbkd/rras_begins_aws_elb/
---
Hi all, 

Currently rolling out RRAS for an always on VPN in AWS. I currently have one server behind a network load balancer and that works, but I will need to load balancer a couple of RRAS servers, but I know typical round robin can cause issues. 

using normal load balancers requires Port Following, or what might be called sticky ports where a client connection will only subsequently connect to the same server, and this needs to be able to forward UDP 500 and UDP 4500 to same box to complete the tunnel. 

Does the elastic load balancer/Network load balancer in AWS allow this?
## [4][How to deploy an AWS EKS cluster with Terraform?](https://www.reddit.com/r/aws/comments/gjm6pa/how_to_deploy_an_aws_eks_cluster_with_terraform/)
- url: https://www.reddit.com/r/aws/comments/gjm6pa/how_to_deploy_an_aws_eks_cluster_with_terraform/
---
If you are looking to learn about how to deploy an [EKS cluster on AWS, this tutorial](https://www.padok.fr/en/blog/aws-eks-cluster-terraform) will help you. It explains a simple way to create an AWS EKS cluster with Terraform.
## [5][Question about managing Windows WorkSpaces](https://www.reddit.com/r/aws/comments/gjm6o7/question_about_managing_windows_workspaces/)
- url: https://www.reddit.com/r/aws/comments/gjm6o7/question_about_managing_windows_workspaces/
---
I am deploying some Windows 10 WorkSpaces.   With physical PCs, after imaging I need to log on as admin to install some software that cannot be included in an image, like our enterprise AV.  

Is there a way to do the same for a WorkSpace?   Ours are integrated with our domain, and as far as I can tell I can’t even turn them on without logging in as the domain user they are assigned to, which obviously I cannot do.   I would like to be able to log-in as a local or domain admin before handing it over to the user.
## [6][Kinesis to S3 vs direct writes to S3](https://www.reddit.com/r/aws/comments/gjm2uf/kinesis_to_s3_vs_direct_writes_to_s3/)
- url: https://www.reddit.com/r/aws/comments/gjm2uf/kinesis_to_s3_vs_direct_writes_to_s3/
---
We are working with an application which is responding to user clicks on the front end. Based on every user click, an average of 3 events are generated in the backend. We need to store all of these events into a S3 as a JSON file.  
As a team our idea was to directly create the JSON object on the application itself and then write it to S3. However we spoke with someone from AWS and he said we should put the data into a Kinesis and then stream it into S3. We are not really sure about the advantage of this process. Has anyone here faced a similar challenge? What would you suggest?
## [7][Web developer says they're developing on m3 (which has been phased out?)](https://www.reddit.com/r/aws/comments/gja4w5/web_developer_says_theyre_developing_on_m3_which/)
- url: https://www.reddit.com/r/aws/comments/gja4w5/web_developer_says_theyre_developing_on_m3_which/
---
Hi, my company is working with a web developer contractor to create an ecommerce site. They said they would use a m3.medium EC2 instance to for hosting. But I'm searching this up and it looks like these instances can't be started up anymore. What could be going on?
## [8][Cloudfront not updating files through origin domain name](https://www.reddit.com/r/aws/comments/gjlvzs/cloudfront_not_updating_files_through_origin/)
- url: https://www.reddit.com/r/aws/comments/gjlvzs/cloudfront_not_updating_files_through_origin/
---
Can anyone help me with explaining how can I see my updated files on cloudfront when using origin domain?

Thank you!
## [9][Create Security group with Self ref - Not working](https://www.reddit.com/r/aws/comments/gjlm1l/create_security_group_with_self_ref_not_working/)
- url: https://www.reddit.com/r/aws/comments/gjlm1l/create_security_group_with_self_ref_not_working/
---
I am trying to create a security group with cloudformation and this is my code

&amp;#x200B;

Scenario1:

      XXXX:
        Type: AWS::EC2::SecurityGroup
        Properties:
          GroupName: XX-XX
          GroupDescription: Allow ssh traffic
          SecurityGroupIngress:
          - IpProtocol: tcp
            FromPort: 22
            ToPort: 22
            SourceSecurityGroupName: !Ref ZZZZ
      SGAPIGWIngress:
        Type: AWS::EC2::SecurityGroupIngress
        Properties:
          Description: XXXX inbound rule
          GroupId: !Ref XXXX
          IpProtocol: tcp
          FromPort: -1
          ToPort: -1
          SourceSecurityGroupId: !Ref XXXX

According to the documentation, it should work. But I get a reply with the Group name in the error

    Invalid id: "XX-XX" (expecting "sg-...") (Service: AmazonEC2; Status Code: 400; Error Code: InvalidGroupId.Malformed; Request ID: 6e2f50fe-1fbf-484c-8d7c-5dc13f4b12ca)

In the resources tab in cloudformation, i see the Group name instead of the group id.

When i try to get the GroupId with !GetAtt XXXX.GroupId, i still get the group name.

Scenario 2:

The same code, but the security group also has a VPCId property specified. Now, it simply gets stuck when it tries to create the security group.

In the resources tab, the security groups with VPCId specified get their Physical ID as sg-... and the Security groups without their VPCId specified get their Group name as Physical ID.

Am i missing something? or has aws made recent updates that is not in the documentation?

Update: In Scenario2, the CFN fails to find the security groups. Fails with sg..does not exist in VPC. I only have one VPC and all the sgs are created in this vpc. 
## [10][Policy help](https://www.reddit.com/r/aws/comments/gjlfd9/policy_help/)
- url: https://www.reddit.com/r/aws/comments/gjlfd9/policy_help/
---
Im writing a Service Control Policy that needs to apply to everyone EXCEPT those using a specific admin role. How would one go about writing such a condition?  


I have everything else done, which currently means im denying everyone some given actions, however i want to allow those with this role to perform the action.
