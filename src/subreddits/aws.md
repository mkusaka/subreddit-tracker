# aws
## [1][API error rates and latencies in Amazon Elastic Compute Cloud (Sydney)](https://www.reddit.com/r/aws/comments/eslenb/api_error_rates_and_latencies_in_amazon_elastic/)
- url: https://www.reddit.com/r/aws/comments/eslenb/api_error_rates_and_latencies_in_amazon_elastic/
---
I was getting following error when doing CLI operation today morning

&gt;An error occurred (InternalError) when calling the DescribeInstances operation (reached max retries: 4): An internal error has occurred

Next checked the status page and found that there was API Error and Latency error for EC2 service in Sydney region. 

&gt; 4:41 PM PST We are investigating increased API error rates and latencies in the AP-SOUTHEAST-2 Region. Connectivity to existing instances is not impacted.

One of my College rebooted a workspace and its still rebooting from past 45 minutes, but does not effect currently running instance or workspace.
## [2][Converting varbinary data and uploading to S3 produces corrupted xlsx file](https://www.reddit.com/r/aws/comments/ess8zd/converting_varbinary_data_and_uploading_to_s3/)
- url: https://www.reddit.com/r/aws/comments/ess8zd/converting_varbinary_data_and_uploading_to_s3/
---
I have a database that was previously used to store files converted to varbinary data. I am currently in the process of moving the files to S3. I've been able to convert pdf, img, doc, xls and most other file types, but when I try to convert an xlsx file it is always corrupted. I'm currently using the code below

*request.query(` select &lt;varbinarydata&gt; from &lt;table&gt; , (err, data) =&gt; {
            if (err) {
                mssql.close();
                throw (err);
            }
            else {
                var filename = &lt;DocumentNm&gt;
                var varbdatan = new Buffer(data.recordset[0].&lt;varbinarydata&gt;);     
                s3.putObject({
                    Bucket: &lt;S3 Bucket&gt;
                    Key: filename,
                    Body: varbdatan
                }, err =&gt; {
                    if (err) {
                        mssql.close();
                        throw (err);
                    }
                    else {
                        console.log('Data Successfully Inserted');
                        mssql.close();
                        callback(null, 1);
                    }
                });
            }
        });*
## [3][RDS DB hacked, what should I do?](https://www.reddit.com/r/aws/comments/esccbr/rds_db_hacked_what_should_i_do/)
- url: https://www.reddit.com/r/aws/comments/esccbr/rds_db_hacked_what_should_i_do/
---
My RDS database was hacked by bitcoin miners who left this message:

"To recover your lost Database and avoid leaking it: Send us 0.06 Bitcoin (BTC) to our Bitcoin address 1Mo24VYuZfZrDHw7GaGr8B6iZTMe8JbWw8 and contact us by Email with your Server IP or Domain name and a Proof of Payment. If you are unsure if we have your data, contact us and we will send you a proof. Your Database is downloaded and backed up on our servers. Backups that we have right now: \*\*\*, \*\*\*\*\*\* . If we dont receive your payment in the next 10 Days, we will make your database public or use them otherwise."

I already have a backup but I need to know how this happened and what to do to prevent it from happening again?

also who's fault is that? mine or aws?
## [4][Looks like Elastic Beanstalk has added support for Node 12](https://www.reddit.com/r/aws/comments/esl0ps/looks_like_elastic_beanstalk_has_added_support/)
- url: https://docs.aws.amazon.com/elasticbeanstalk/latest/relnotes/release-2020-01-21-linux.html
---

## [5][Having issues querying IOT Core from a different account, thinking it's a policy issue](https://www.reddit.com/r/aws/comments/esqyxf/having_issues_querying_iot_core_from_a_different/)
- url: https://www.reddit.com/r/aws/comments/esqyxf/having_issues_querying_iot_core_from_a_different/
---
I am working on a platform migration currently that requires some services to be housed in two seperate accounts.   
We have an IOT Core service in the first account that stores our IOT Devices and we are trying to programmatically grab the shadow of a given device from the second account but we are getting back "Forbidden Exception".

The workflow that should happen is as follows :-  


1. A Lambda Fn is called in the second account to confirm user ownership of IOT device with thingname
2. If ownership is confirmed the Lambda then calls the following function attempting to access IOT Core in first account :-  

3. What should happen is it returns the thingname but instead we are getting an error, "Forbidden Exception"
```
public static async iotTest(event) {
        try {
            const iotData = new AWS.IotData({endpoint: endPointFirstAccount});
            const params = {
                thingName: event.query.thingName
            };
            const data = await iotData.getThingShadow(params).promise();
            return data ? JSON.parse(data.payload) : null;
        } catch (error) {
            console.error('Error: ', error);
        }
    }
```
I think this is a policies issue but I'm not sure where to start. If anyone can help I'd be very grateful, thanks.
## [6][IAM Database Authentication For Amazon RDS In MySQL](https://www.reddit.com/r/aws/comments/esqipc/iam_database_authentication_for_amazon_rds_in/)
- url: https://www.ibexlabs.com/iam-database-authentication-for-amazon-rds-in-mysql/
---

## [7][AWS CLI credentials stored in plaintext? Anyone know a way to encrypt these or am I missing some aspect here?](https://www.reddit.com/r/aws/comments/esgq5u/aws_cli_credentials_stored_in_plaintext_anyone/)
- url: https://www.reddit.com/r/aws/comments/esgq5u/aws_cli_credentials_stored_in_plaintext_anyone/
---
Like the title says. Anyone know a way to store these in an encrypted state? Just seems odd that they're in a file called "credentials".  Not sure if I'm missing anything here and they're secured in another way. I tend to be a little nervous about leaving creds in text files. Thanks.
## [8][A nasty hack to enable using an API Gateway generated Javascript client SDK in Node.js](https://www.reddit.com/r/aws/comments/esmep0/a_nasty_hack_to_enable_using_an_api_gateway/)
- url: https://github.com/dmh2000/ApiGenNodeHack
---

## [9][RDS Certificate Warning Emails - What Am I Supposed To Do Here?](https://www.reddit.com/r/aws/comments/esiakt/rds_certificate_warning_emails_what_am_i_supposed/)
- url: https://www.reddit.com/r/aws/comments/esiakt/rds_certificate_warning_emails_what_am_i_supposed/
---
I keep getting these emails that start with: 

&gt;We previously sent a communication in early October to update your RDS SSL/TLS certificates by October 31, 2019. We have extended the dates and now request that you act before February 5, 2020 to avoid interruption of your applications that use Secure Sockets Layer (SSL) or Transport Layer Security (TLS) to connect to your RDS and Aurora database instances. Note that this new date is only 4 weeks before the actual Certificate Authority (CA) expiration on March 5, 2020. Because our own deployments, testing, and scanning to validate all RDS instances are ready for the expiry must take place during the final 4 weeks, the February 5th date cannot be further extended.
 
&gt;You are receiving this message because you have an Amazon RDS database instance(s) that requires action.

I have no idea what this means.  I am tempted to just go in, backup my db, delete the RDS instance, launch a new one, and go on with my life.  I do not independently manage any certs for db connectivity in my apps.

Would that solve it?
## [10][How do you check which specific load balancer a Elastic Beanstalk app is using?](https://www.reddit.com/r/aws/comments/esk541/how_do_you_check_which_specific_load_balancer_a/)
- url: https://www.reddit.com/r/aws/comments/esk541/how_do_you_check_which_specific_load_balancer_a/
---
I am trying to figure out how to check which specific load balancer a specific elastic beanstalk app is using, is there an easy way to do this?
