# aws
## [1][CloudWatch now offers composite alarms. Great for reducing alarm fatigue and triggering scale down actions](https://www.reddit.com/r/aws/comments/fe6s3y/cloudwatch_now_offers_composite_alarms_great_for/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/03/amazon-cloudwatch-now-allows-you-to-combine-multiple-alarms/
---

## [2][Using EKS encryption provider support for defense-in-depth](https://www.reddit.com/r/aws/comments/fe2ie5/using_eks_encryption_provider_support_for/)
- url: https://aws.amazon.com/blogs/containers/using-eks-encryption-provider-support-for-defense-in-depth/
---

## [3][550 Permission Denied... again!](https://www.reddit.com/r/aws/comments/fecuno/550_permission_denied_again/)
- url: https://www.reddit.com/r/aws/comments/fecuno/550_permission_denied_again/
---
Hello!

I got a new network yesterday, and all of a sudden, I can't access my server via FileZilla. Yesterday with my old network, it worked flawlessly. Now with the new one, I get the message below (connecting with my old network doesn't work anymore either, it gives the same message):

Status:	Connecting to \[my aws address\]

Status:	Connection established, waiting for welcome message...

Status:	Plain FTP is insecure. Please switch to FTP over TLS.

Status:	Server does not support non-ASCII characters.

Status:	Logged in

Status:	Retrieving directory listing...

Command:	PWD

Response:	257 "/var/www/\[mysite\]" is the current directory

Command:	TYPE I

Response:	200 Switching to Binary mode.

Command:	PORT 192,168,0,174,196,83

Response:	500 Illegal PORT command.

Command:	PASV

Response:	550 Permission denied.

Error:	Failed to retrieve directory listing
## [4][Tampermonkey script to show only the Services you care about/use in the Web Console](https://www.reddit.com/r/aws/comments/fe6url/tampermonkey_script_to_show_only_the_services_you/)
- url: https://www.reddit.com/r/aws/comments/fe6url/tampermonkey_script_to_show_only_the_services_you/
---
Tampermonkey -&gt; Dashboard -&gt; Utilities -&gt; Install from URL

[https://gist.githubusercontent.com/rdkls/f0f544c73d359798e51d07f27a7ef520/raw/a1a802343ff0790244630b2702ac3eda0f49ad51/aws-web-console-service-menu-customize.js](https://gist.githubusercontent.com/rdkls/f0f544c73d359798e51d07f27a7ef520/raw/a1a802343ff0790244630b2702ac3eda0f49ad51/aws-web-console-service-menu-customize.js)

Results:

https://preview.redd.it/6t3z14nhtyk41.png?width=1027&amp;format=png&amp;auto=webp&amp;s=fd0a676d4788e67b4753e6de6cb63af7459325ae
## [5][AWS MUMBAI Cancelled due to COVID-19](https://www.reddit.com/r/aws/comments/fdtrgc/aws_mumbai_cancelled_due_to_covid19/)
- url: https://www.reddit.com/r/aws/comments/fdtrgc/aws_mumbai_cancelled_due_to_covid19/
---
Just got an mail that the event is cancelled.
## [6][Looking for a diagram Tool which will highly specific infrastructure when hovering mouse over different sections .](https://www.reddit.com/r/aws/comments/febwag/looking_for_a_diagram_tool_which_will_highly/)
- url: https://www.reddit.com/r/aws/comments/febwag/looking_for_a_diagram_tool_which_will_highly/
---
I have the idea in my head about such a tool where different IT infrastructure can be built into a document/diagram such as lucidchart. Hovering the mouse over a section with headings would grey out non related sections with just the desired section becoming highlighted. Hopefully that makes sense. Any suggestions would be warmly greeted.
## [7][Best practice for API and Lambda](https://www.reddit.com/r/aws/comments/fe4gn0/best_practice_for_api_and_lambda/)
- url: https://www.reddit.com/r/aws/comments/fe4gn0/best_practice_for_api_and_lambda/
---
Hello everyone. I'm wondering what's the best practice if you build a AWS REST API that uses different routes. Let's assume we have 3 routes. One is for product and one for user. Basically example.com/product and example.com/user. Now I also need a GET request for a specific product. The third route is something like example.com/user?id=1


I see three different options.


1) I create a lambda for example.com/product, for example.com/product?id=1 and for example.com/user.
Pro: each lambda has only one responsibility
Contra: I get a lot lambdas for bigger sites and a lot repetition because the database model, connection parameters etc. are the same, only the SQL queries are different.


2) I handle everything inside 1 lambda. 
Pro: I have one lambda per service. 
Contra: I have multiple responsibilities, the code can be very long, hard to test and messy.


3) create a Lambda for the user logic and a lambda for the product logic.
Pro: I have well defined domains and I don't repeat myself too much. I can reuse most of my code. I don't have so many lambdas.
Contra: I might have several responsibilities and the code can become very long.

My gut feeling says that the third option is the right one but I also like the first approach. I just dislike the repetition on approach one and the fact that I might have 20 lambdas for one application.

Additionally:
Are there best practices for the naming of lambdas?
## [8][What is wrong with my payload](https://www.reddit.com/r/aws/comments/fe54ow/what_is_wrong_with_my_payload/)
- url: https://www.reddit.com/r/aws/comments/fe54ow/what_is_wrong_with_my_payload/
---
Here is my Code for an update function:

    
     (
     "operation": "update",
      "tableName": "Test1",
      "payload": {
        "TableName": "Test1",
        "IndexName": "ID",
        "Key": {
          "ID": {
            "S": "aosdjndidj"
          },
          "Region": {
            "N": "1"
          }
        },
        "ExpressionAttributeNames": {
          "#R": "Region"
        },
        "UpdateExpression": "Set #R = :r",
        "ExpressionAttributeValues": {
          ":r": 2
        },
        "ReturnValues": "UPDATED_NEW"
      }
    }

It returns the error:  


    {
      "errorType": "ValidationException",
      "errorMessage": "The provided key element does not match the schema",
## [9][Locking down a resource to a single lambda function](https://www.reddit.com/r/aws/comments/feb0cm/locking_down_a_resource_to_a_single_lambda/)
- url: https://www.reddit.com/r/aws/comments/feb0cm/locking_down_a_resource_to_a_single_lambda/
---
I have an elasticsearch domain and am trying to work out how i would craft the resource policy on the ES Domain to only allow a single lambda to access the es:\* Actions.

I have attempted to set the principle of the to the ARN of the lambda function such that

`"Principal": {:`  
`"AWS": " -lambda ARN-"`  
`},`  
`"Action": "es:*",`

&amp;#x200B;

However this dosn't work; any advice on how to do this would be awesome;

cheers
## [10][Question: How do you guys implement Firewall in your AWS Ecosystem?](https://www.reddit.com/r/aws/comments/fe6z7u/question_how_do_you_guys_implement_firewall_in/)
- url: https://www.reddit.com/r/aws/comments/fe6z7u/question_how_do_you_guys_implement_firewall_in/
---
We have a requirement to implement Layer 7 protection for our resources, as I understand AWS has AWS WAF. We have apis running on EKS with ELB. ELB is not supported as a resource endpoint to be attached on WAF. We can always map our endpoint on Cloudfront, which I think may incur additional cost on the stack, no?

Part of the requirement, we need to show that we will get alerts if there's bad activities going on.

Looking forward to hear your setup and recommendations. Thank you :)
