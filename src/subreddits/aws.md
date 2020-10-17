# aws
## [1][AWS Wish List 2020](https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/)
- url: https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/
---
&amp;#x200B;

AWS always releases a bunch of features, sometimes everyday or atleast once a week. Here is my wish list of the features I want to see as a part of AWS infrastructure

1: AWS Managed Proxy Server(Rather than spinning own squid server)

2: EBS replication across different availability zones(Possible? Legal constraints?)

3: Multi-region VPC(Possible? Legal constraints?)

4: UI to debug boot issues(Better then EC2 Get Instance Screenshot and Instance logs)

5: Support tagging for every individual service(It's improving)

6: VPC endpoints support for every service (EKS?) 

7: EC2 instance live migration

8: Display AWS Cli  while resource creation(Similar to GCP)

9: Cost calculation while resource creation(AWS start supporting(for example, RDS) this feature but not for every service

10: More features in App Mesh(Circuit breaker, Rate Limiting)

P.S: Not sure if some features are already available, but if something is missing, please feel free to add
## [2][How many servers does AWS own now?](https://www.reddit.com/r/aws/comments/jclk7n/how_many_servers_does_aws_own_now/)
- url: https://www.reddit.com/r/aws/comments/jclk7n/how_many_servers_does_aws_own_now/
---
According to wikipedia, they have 1.4M servers in 2014. Does anyone know the latest figure?
## [3][[Noob Question] Network Routing through the internet](https://www.reddit.com/r/aws/comments/jctgv4/noob_question_network_routing_through_the_internet/)
- url: https://www.reddit.com/r/aws/comments/jctgv4/noob_question_network_routing_through_the_internet/
---
Hi, Noob here, just have a question how the network is routed through the internet when the client hits the url.

I want to know how the traffic is routed using with/without WAF, Route 53(hostedzone) and cloudfront

This is my understanding so far

with cloudfront, route 53, Waf

CLient(http request) ----&gt; Domain registrar DNS servers ----&gt; Route53 DNS ----&gt; cloudfront ----&gt; WAF ----&gt; EC2 live server

THIS MIGHT VERY WELL BE A BLUNDER. I want to know the correct traffic travel path. Thank You. Experts please help!!
## [4][What proportion of your bill goes on different services?](https://www.reddit.com/r/aws/comments/jcuu45/what_proportion_of_your_bill_goes_on_different/)
- url: https://www.reddit.com/r/aws/comments/jcuu45/what_proportion_of_your_bill_goes_on_different/
---
I'm yet to work on any projects that create significant bills each month. Purely out of curiosity, I'd love to get some anecdotal data what on what makes up bills, i.e. roughly what proportion goes to VMs, databases, storage, etc, though I know this will vary wildly depending on use case. If you don't mind also sharing the type and size of the business, that would be great!
## [5][Did you notice AWS Golang SDK v0.26.0 is broken???](https://www.reddit.com/r/aws/comments/jctl5s/did_you_notice_aws_golang_sdk_v0260_is_broken/)
- url: https://www.reddit.com/r/aws/comments/jctl5s/did_you_notice_aws_golang_sdk_v0260_is_broken/
---
For a long time we have been using V1 SDK (Golang) , and recently stated migrating to V2 (v0.23.0). We regularly update SDK versions, however our move to V0.26.0 was a disaster.

It seems v0.26.0 is completely broken. All pagination related functionality is completely gone and there is no proper communication as well. I see issue opened for the same, but no proper feedback or action yet.

Wondering if AWS SDK V2 team tests code or not? And how reliable would be going forward on V2.
## [6][AWS Account Password Changed By Someone Else](https://www.reddit.com/r/aws/comments/jcd8jg/aws_account_password_changed_by_someone_else/)
- url: https://www.reddit.com/r/aws/comments/jcd8jg/aws_account_password_changed_by_someone_else/
---
I got an email today stating "We received a request to reset the password for the AWS account". Shortly followed by another email "Your Amazon Web Services Password Has Been Updated". Although this is not an account I use much, naturally I was very troubled by this email. I had to use the "Forgot password" link after not being able to log in, and was able to reset my password. Is there an audit trail where I can see who might have done what in my account? How can I be sure there are no browser sessions still open someplace using the old password?

\----------------

Thanks for all your great responses!

1. I definitely checked that it was not phishing. Very long convoluted URL's, but definitely amazon domain.

2. I did not use MFA on this account because I expected to create another account at some point, and my experience with Amazon is that sometimes you cannot have the same mobile number on more than one account MFA.

3. I did not want to believe that my email (albeit a low priority secondary one) was hacked, but it had to have been in order to confirm the password change, right? Chrome has a "compromised password" feature which is really amazing, and confirms that that email was "found in data breach". So I've got some email password work this weekend...

4. Out of curiosity, I will see if I can find anything in the CloudTrail. But as SelfDestructSep2020 recommends, I will probably "burn the account and make a new one".

&amp;#x200B;
## [7][Should you always use Control Tower, or does just using Organizations make sense in some situations?](https://www.reddit.com/r/aws/comments/jce5wr/should_you_always_use_control_tower_or_does_just/)
- url: https://www.reddit.com/r/aws/comments/jce5wr/should_you_always_use_control_tower_or_does_just/
---
Hi everybody,

I'm creating a greenfield multi-account setup and am trying to figure out the best tools for the job. A bit about my place of work:

* 5 engineers, 12 employees overall. We are growing however, and I'd like to be able to support dozens and maybe hundreds of engineers
* We have no regulatory requirements -- we don't store PII, health, or financial data or anything like that.
* I am the only DevOps guy, though I may hire 1-2 more in the next 12 months.

The tool that AWS recommends is Control Tower (CT), however I feel this may be a bit of overkill. Seems like it's intended for large enterprises within regulated environments. I don't see us using Guardrails heavily. SCPs seem useful but those come with Organizations as well. I also don't feel like CT is quite ready for the big time -- the fact that it doesn't support nested OUs is a minus.

Do you think that simply using Organizations to create accounts, and then using something like Terraform or Cloud CDK to deploy infra is enough? Or am I missing something with CT?

Thank you in advance.
## [8][Can I rely solely on snapshots for disaster recovery/restoration in production?](https://www.reddit.com/r/aws/comments/jcptn3/can_i_rely_solely_on_snapshots_for_disaster/)
- url: https://www.reddit.com/r/aws/comments/jcptn3/can_i_rely_solely_on_snapshots_for_disaster/
---
I'm currently building a multi-account disaster recovery plan. I wrote a script to do a logical pg\_dump and copy the file to S3 in another account.

I was curious if I could ditch the logical dumps and rely solely on cross-account snapshot sharing to remove the maintenance overhead. What's your prod DR system look like for RDS data? Do you think snapshots reliable enough to ditch logical dumps in a production app? Or should I do both?
## [9][For gpu ec2's, am I charged whenever I'm using the gpu's or at all times when the instance is up?](https://www.reddit.com/r/aws/comments/jco1uu/for_gpu_ec2s_am_i_charged_whenever_im_using_the/)
- url: https://www.reddit.com/r/aws/comments/jco1uu/for_gpu_ec2s_am_i_charged_whenever_im_using_the/
---
Beginner question, if I'm doing deep learning on a gpu ec2 instance like p2.xlarge, am I charged whenever the instance if up?

For example, if I launched the instance for 10 hours, but I did 8 hours of programming and only 2 hours of training that actually uses the gpu's, am I charged for the full 10 hours?

If so, how do I efficiently manage ec2 costs?
## [10][If I change the name of a field in a schema, how do I change the items currently in the database so their field is also renamed?](https://www.reddit.com/r/aws/comments/jcqdz5/if_i_change_the_name_of_a_field_in_a_schema_how/)
- url: https://www.reddit.com/r/aws/comments/jcqdz5/if_i_change_the_name_of_a_field_in_a_schema_how/
---

## [11][EBS Lifecycle Manager Tagging Question](https://www.reddit.com/r/aws/comments/jcdh0r/ebs_lifecycle_manager_tagging_question/)
- url: https://www.reddit.com/r/aws/comments/jcdh0r/ebs_lifecycle_manager_tagging_question/
---
I configured EBS Lifecycle Manager to run nightly backups using the default tagging settings which are:


instance-id  : $(instance-id)


timestamp :  $(timestamp)

The tags work great, but the name's are blank. It's not the end of the world, but from a more userfriendly standpoint, is there a way to customize the name field to something like &lt;date&gt; : &lt;instanceName&gt; : &lt;volumeName&gt;?
