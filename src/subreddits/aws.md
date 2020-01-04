# aws
## [1][Why do i have to create a NAT Gateway per AZ but only one Internet Gateway?](https://www.reddit.com/r/aws/comments/ejk8nm/why_do_i_have_to_create_a_nat_gateway_per_az_but/)
- url: https://www.reddit.com/r/aws/comments/ejk8nm/why_do_i_have_to_create_a_nat_gateway_per_az_but/
---
I'm a bit confused about this whole thing. It seems for every AZ with a public/private subnet setup, i need a NAT gateway so the private instances can talk to the internet. But all the public subnets only need a single internet gateway. Why is this?
## [2][AWS EC2 free tier + WordPress + Cloudfront + AWS SSL Certificate + Namecheap domain = AAAAAH](https://www.reddit.com/r/aws/comments/ejp4md/aws_ec2_free_tier_wordpress_cloudfront_aws_ssl/)
- url: https://www.reddit.com/r/aws/comments/ejp4md/aws_ec2_free_tier_wordpress_cloudfront_aws_ssl/
---
Hi everyone,

I am a total noob to this, so please bare with me here. I am trying to set up a blog to start writing a little bit and I created a free AWS EC2 instance today and installed Wordpress on it. I also have a TLD with Namecheap, that I only used as an email domain so far (with [Outlook.com](https://Outlook.com) Premium). I also got the A record and CNAME records set up properly, so that the URLs correctly forward to the WordPress installation on AWS (and my emails still work) - this involves an elastic IP and the respective DNS entries. Next, I created an SSL certificate with AWS and got it successfully validated through AWS DNS validation. So far so good. Only a few hours of work for someone who needs to google every step - at least I learned a lot :-).

Now the last (at least for now it seems to me like that) is to activate the SSL certificate with my domain to make sure my blog uses HTTPS. So far all browsers show it as HTTP and "not secure" and when I enter the URL as HTTPS I get the usual message that the connection is not private.

What I have understood is that I need to run the connection between the EC2 instance and the SSL through CloudFront, so I set that up and also created a CloudFront distribution, linked the instance as well as the alternate domain names and the SSL certificate, **but the website is still shown as not secure.** Safari shows in the certificate details "Certificate generated at boot time", so it seems the instance does not pull the correct certificate - for whatever reason :).

I am sure I am missing a ton of details that I need to provide so that you guys can help me, but I would really appreciate some guidance here.

&amp;#x200B;

Edit:

* The instance is in the Ohio zone and the certificate in the Virginia zone - in case that make it worse
* I have not done anything with IAM or security groups so far
## [3][Secrets management](https://www.reddit.com/r/aws/comments/ejpr8c/secrets_management/)
- url: https://www.reddit.com/r/aws/comments/ejpr8c/secrets_management/
---
Mobycast started a two part podcast series on secrets management that goes into detail on both AWS Secrets Manager and AWS Systems Manager Parameter store. Good information for how to keep secrets exposure limited to the code that needs access to it and only at runtime.  [https://mobycast.fm/episode/93-psst-secrets-handling-for-cloud-native-apps-part-1/](https://mobycast.fm/episode/93-psst-secrets-handling-for-cloud-native-apps-part-1/)
## [4][ELI5: DynamoDB vs RDS](https://www.reddit.com/r/aws/comments/ejo1g4/eli5_dynamodb_vs_rds/)
- url: https://www.reddit.com/r/aws/comments/ejo1g4/eli5_dynamodb_vs_rds/
---
I know DynamoDB is NoSQL while RDS is relational(SQL)

It seems like all the functions you would want out of RDS(relations thought you have to do more loads) you could do with DynamoDB

With DynamoDB you would have the scalability of NoSQL

When would you want to use RDS over DynamoDB?

It also appears to be more expensive to use RDS over DynamoDB because you're paying a set free(&gt;$600) per month whereas in DynamoDB you're paying per usage(WCU and RCU)
## [5][Nesting security groups?](https://www.reddit.com/r/aws/comments/ejr2lz/nesting_security_groups/)
- url: https://www.reddit.com/r/aws/comments/ejr2lz/nesting_security_groups/
---
So I'm basically trying to block origin for anything other than our CDN, and our own public IP addresses. Problem is, we have many public IPs (many locations).  

My original plan was to have 6 distinct security groups for our public CIDRs, and an extra one for our CDN CIDRs. I did not know that I could only have 5 security groups per ELB.

I then tried nesting the public CIDR groups into the main CDN CIDR sg, but it seems that nesting is only used to allow instances and not actual CIDRs or IPs.

&amp;#x200B;

Any recomendations on how to achieve this? My company has 160 public IPs, plus another 32 from our CDN.

&amp;#x200B;

Thanks!

  
V.
## [6][How to manage roles when using SAML or OIDC sso?](https://www.reddit.com/r/aws/comments/ejnfv8/how_to_manage_roles_when_using_saml_or_oidc_sso/)
- url: https://www.reddit.com/r/aws/comments/ejnfv8/how_to_manage_roles_when_using_saml_or_oidc_sso/
---
Currently we have IAM users for each person who needs access to AWS. We assign permissions to groups, and then users to those groups as they make sense. Some of those might be department groups, others might be project or other special groups.

With SAML, the user has to pick from a list of roles that they have permission to assume, and use the permissions granted to that role. They can switch to another role later if they need to.

With users and groups, the user has the sum of all permissions on their groups (and on the user, if any). With roles, they get one set at a time. 

Say Role A has permission to read from S3 Bucket A, and Role B has permission to write to Bucket B. Copying an object from A to B can still be done, but it isn't nearly as straight forward. You end up swapping roles a lot.

I see a couple possible options:

* Duplicate the permissions in a whole bunch of roles. For users in Dept X and Project Y, create the DeptX\_ProjectY role. And the DeptX\_ProejctY\_ProjectZ role, and so on.
* Calculate the permissions of each user, and create a role for each user. Programatically, hopefully.

Are there tools to help with this? Is there some other solution that I'm missing?
## [7][codebuild buildspec.yml: is it possible to substitute an environment variable to describe an SSM parameter name? example inside](https://www.reddit.com/r/aws/comments/ejkkww/codebuild_buildspecyml_is_it_possible_to/)
- url: https://www.reddit.com/r/aws/comments/ejkkww/codebuild_buildspecyml_is_it_possible_to/
---
is this possible?

    env:
      parameter-store:
        ENV_PARAM_NAME: /$ENVIRONMENT/param-name
      secrets-manager:
        ENV_SECRET_NAME: $ENVIRONMENT/secret-name:SECRET_KEY

where i could externally set `$ENVIRONMENT` to say `dev` or `prod`?

the use case is a buildspec that is identical for both dev and prod with the exception of the parameter / secret name prefixes

i have tried all of the approaches i could think of including: `"$ENVIRONMENT"`, `${ENVIRONMENT}`, `${$ENVIRONMENT}}`, `${${ENVIRONMENT}}`, `${env:$ENVIRONMENT}` 

but none seem to work
## [8][AWS Shield Vs CloudFlare](https://www.reddit.com/r/aws/comments/ejhewd/aws_shield_vs_cloudflare/)
- url: https://www.reddit.com/r/aws/comments/ejhewd/aws_shield_vs_cloudflare/
---
Can anyone explain the difference to me. What are some of the pros and cons of each?
## [9][Trying the free aws and invoice..](https://www.reddit.com/r/aws/comments/ejkmbx/trying_the_free_aws_and_invoice/)
- url: https://www.reddit.com/r/aws/comments/ejkmbx/trying_the_free_aws_and_invoice/
---
hi everyone,

i was trying free aws. We got a $ 68 invioce. but I was using it for free. no warning was given to me. Why this invoice?

i don't have enough money to pay. please help me :( what can I do ???
## [10][AWS Client VPN Pricing - Disassociate to save money](https://www.reddit.com/r/aws/comments/ejkky9/aws_client_vpn_pricing_disassociate_to_save_money/)
- url: https://www.reddit.com/r/aws/comments/ejkky9/aws_client_vpn_pricing_disassociate_to_save_money/
---
I'm using an AWS Client VPN Endpoint as a temporary remote management access point.   I don't need this VPN on all the time so I'm trying to avoid unneeded charges.

Based on the documentation ( https://aws.amazon.com/vpn/pricing/ ) it appears you are charged for association time as well as connection time.

If I simply disconnect from the VPN and also Disassociate any subnets when I'm done for the day, I should be free of any additional charges.  Then later when I need it again, just Associate and I'm ready to go.  Is this correct?

Any downsides to this that I may be overlooking, besides just having to wait for association time?
