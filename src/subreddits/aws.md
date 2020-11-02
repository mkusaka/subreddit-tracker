# aws
## [1][re:Invent registration is now open](https://www.reddit.com/r/aws/comments/jkenu3/reinvent_registration_is_now_open/)
- url: https://register.virtual.awsevents.com/
---

## [2][In the Works – New AWS Region in Zurich, Switzerland](https://www.reddit.com/r/aws/comments/jmkkiv/in_the_works_new_aws_region_in_zurich_switzerland/)
- url: https://aws.amazon.com/blogs/aws/in-the-works-new-aws-region-in-zurich-switzerland/
---

## [3][Lambda not scaling with SQS. ConcurrentExecutions is always 1. SQS ApproximateNumberOfMessagesVisible = 60k](https://www.reddit.com/r/aws/comments/jmgwmv/lambda_not_scaling_with_sqs_concurrentexecutions/)
- url: https://www.reddit.com/r/aws/comments/jmgwmv/lambda_not_scaling_with_sqs_concurrentexecutions/
---
I set a lambda with a  
\- Memory: 1024MB  
\- SQS trigger, batch size = 10  
\- No concurrency limit. "Use unreserved account concurrency"  
\- SQS is a FIFO queue  
\- Lambda uses a VPC

Lambda has a success rate of 100% but ConcurrentExecutions never exceeds one (1). SQS ApproximateNumberOfMessagesVisible goes down slowly.

I thought with 60k messages on the queue it would start invoking lambdas up to 1k concurrently.

Do you guys have any ideas what's going on?

Thanks for your help.
## [4][AWS Cloud Support Interview coming up... super nervous :( and what does domain: generalist mean?](https://www.reddit.com/r/aws/comments/jml2mf/aws_cloud_support_interview_coming_up_super/)
- url: https://www.reddit.com/r/aws/comments/jml2mf/aws_cloud_support_interview_coming_up_super/
---
Hello kind folks of reddit,

I'm currently in the hiring process for the AWS Cloud Support Associate role and I passed the Online Assessment last week and my interview date is set for Nov 6 which is 4 days from now. In the email regarding the details of the interview, it says lists the domain for the interview as "generalist". My question is should I be worried about this? I've done a lot of reading online and people get specific domains for their interview I cannot find one person who's gotten generalist. I really hope this is not a bad thing. I put my top three interests as Big data, database, and Analytics at the end of my OA could I communicate to the interviewer to assess me for those teams? Also if anyone were to provide any tips or questions for the interview regarding any of those three teams or the generalist domain, I would really appreciate it!! I know the reddit community's got me :) 

Really appreciate all the help I can get, thanks in advance and I'll be sure to update you all after the interview

Thank you all so much!!
## [5][A method for managing backend application secrets using AWS S3 and CodeDeploy](https://www.reddit.com/r/aws/comments/jmjt5l/a_method_for_managing_backend_application_secrets/)
- url: https://riadrifai22.medium.com/managing-application-secrets-like-never-before-using-aws-s3-and-codedeploy-611d3121ecbe
---

## [6][Only 5 and half hours left to the HUMBLE BOOK BUNDLE: AWS, AZURE, GOOGLE, AND CLOUD SECURITY Get Amazon Web Services for Mobile Developers and more get it before it ends](https://www.reddit.com/r/aws/comments/jmnb2d/only_5_and_half_hours_left_to_the_humble_book/)
- url: https://www.humblebundle.com/books/aws-azure-google-and-cloud-security-books?partner=mfasula&amp;charity=2175662
---

## [7][Is it wise to use same application load balancer for API and UI? What are possible risks?](https://www.reddit.com/r/aws/comments/jmmbn2/is_it_wise_to_use_same_application_load_balancer/)
- url: https://www.reddit.com/r/aws/comments/jmmbn2/is_it_wise_to_use_same_application_load_balancer/
---
We can use host based routing in aws alb. So we can forward to api and ui target groups based on host headers.
But in this case is it vulnerable to use the same load balancer?
## [8][CHECK_NRPE: Socket timeout after 30 seconds || Nagios || Client](https://www.reddit.com/r/aws/comments/jmm506/check_nrpe_socket_timeout_after_30_seconds_nagios/)
- url: https://www.reddit.com/r/aws/comments/jmm506/check_nrpe_socket_timeout_after_30_seconds_nagios/
---
Able to do install Nagios Server on ubuntu but when doing for server getting this error on AWS. Go through many links of same error but not found the solution. 

[https://bobcares.com/blog/check\_nrpe-socket-timeout-after-10-seconds/](https://bobcares.com/blog/check_nrpe-socket-timeout-after-10-seconds/)

[https://serverfault.com/questions/605666/check-nrpe-socket-timeout-after-10-seconds-https-ssl/605738](https://serverfault.com/questions/605666/check-nrpe-socket-timeout-after-10-seconds-https-ssl/605738)

It would be great if anyone have an exact solution for it. 

For Nagios Server I go through below link:

 [https://www.itzgeek.com/how-tos/linux/ubuntu-how-tos/install-nagios-4-1-1-ubuntu-16-04.html](https://www.itzgeek.com/how-tos/linux/ubuntu-how-tos/install-nagios-4-1-1-ubuntu-16-04.html)

For Nagios Client followed below link:

[https://www.itzgeek.com/how-tos/linux/centos-how-tos/monitor-remote-linux-system-with-nagios-3.html](https://www.itzgeek.com/how-tos/linux/centos-how-tos/monitor-remote-linux-system-with-nagios-3.html)

&amp;#x200B;

[Host Error](https://preview.redd.it/45ka9ruiftw51.png?width=1166&amp;format=png&amp;auto=webp&amp;s=bfc786e08b04738bf750716e63dd39fafabd76dd)
## [9][Penetration testing API GW tips?](https://www.reddit.com/r/aws/comments/jm85on/penetration_testing_api_gw_tips/)
- url: https://www.reddit.com/r/aws/comments/jm85on/penetration_testing_api_gw_tips/
---
Ive recently been assigned a task that includes making a security review of an environment including an API Gateway pointing to a Lambda function, and while I have done Hack-the-boxes and CTFs in the past, im not entirely sure what to look for here in regards to misconfigurations of the API GW/Lambda which might be abuseable.   
Since there is no database or website behind the gateway, attacks like SQLi and XSS shouldnt be possible I guess and im considering if there could be some tampering of JWTs or revelation of too much information, but im not really sure how to attack this atm. I have already identified the possibility of Denial-of-wallet attacks, but as we have a ton of monitoring and alarms, i highly doubt this is of much risk.  
Anyway, to the question, do any of you guys do pentests on API GWs or Lambdas and have tips on what tools to use, what to look for etc? :)  
Im currently using various Kali tools like Burp and NMAP scripts, but im open to try out other frameworks as i see multiple API testing tools exists that i havent tried yet.

Any tips are appreciated as im not too experienced with security on AWS! \^\^
## [10][CWE/EventBridge doesn't receive IN_PROGRESS events from Codebuild Batch builds](https://www.reddit.com/r/aws/comments/jml5mg/cweeventbridge_doesnt_receive_in_progress_events/)
- url: https://www.reddit.com/r/aws/comments/jml5mg/cweeventbridge_doesnt_receive_in_progress_events/
---
Hi

Basically, see the problem in the title. For regular builds I receive all events I need. Here's my event pattern:

    {
      "detail": {
        "build-status": [
          "IN_PROGRESS",
          "SUCCEEDED",
          "FAILED",
          "STOPPED"
        ],
        "project-name": [
          "super-mega-codebuild-project"
        ]
      },
      "detail-type": [
        "CodeBuild Build State Change"
      ],
      "source": [
        "aws.codebuild"
      ]
    }

Is there something wrong with the pattern or it's an expected behaviour?
## [11][Schedule CSV export (to an email) of data from RDS](https://www.reddit.com/r/aws/comments/jmky8c/schedule_csv_export_to_an_email_of_data_from_rds/)
- url: https://www.reddit.com/r/aws/comments/jmky8c/schedule_csv_export_to_an_email_of_data_from_rds/
---
I have a requirement to email weekly data from RDS (PostGres) as CSV in an email attachment. What is the best way to do it. I know it can be done via lambda but is there any other good way to achieve this ? Any built in service may be ?
