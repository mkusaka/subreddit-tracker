# aws
## [1][S3 Path Deprecation Plan Updated](https://www.reddit.com/r/aws/comments/iyd84u/s3_path_deprecation_plan_updated/)
- url: https://aws.amazon.com/blogs/aws/amazon-s3-path-deprecation-plan-the-rest-of-the-story/
---

## [2][I'm trying to define how the AWS docs are so bad..](https://www.reddit.com/r/aws/comments/j129il/im_trying_to_define_how_the_aws_docs_are_so_bad/)
- url: https://www.reddit.com/r/aws/comments/j129il/im_trying_to_define_how_the_aws_docs_are_so_bad/
---
But I can't put it down to any one thing. They're too verbose, I can't actually find out what I have to do and a lot of things dealing with the console aren't there anymore.

Is it just me or is this some of the most difficult documentation to force your way through? Not because it's necessarily complicated, it just never seems to get to the point.

But its like more than that. I hate it. Why do I hate their documentation so much?
## [3][How can AWS security group changes be applied immediately?](https://www.reddit.com/r/aws/comments/j1c20c/how_can_aws_security_group_changes_be_applied/)
- url: https://www.reddit.com/r/aws/comments/j1c20c/how_can_aws_security_group_changes_be_applied/
---
Why it doesn't take time to propagate across different AZs?
## [4][Advice on using RDS with db-migrate](https://www.reddit.com/r/aws/comments/j1bwy2/advice_on_using_rds_with_dbmigrate/)
- url: https://www.reddit.com/r/aws/comments/j1bwy2/advice_on_using_rds_with_dbmigrate/
---
I have set up a RDS instance and need to make it publicly accessible to make it play nice with my db-migrate process that is called during my deployment and locally during testing.   


However I understand that making an RDS instance public is bad practice. I now want to know if I set my RDS instance back to being publicly inaccessible, how would I trigger my build process against it to execute the database migration scripts.
## [5][Session Manager](https://www.reddit.com/r/aws/comments/j1bjty/session_manager/)
- url: https://www.reddit.com/r/aws/comments/j1bjty/session_manager/
---
I am staring to use Session Manager.  I understand that it will allow us to manage resources centrally. Anyone have a bad experience or advise to offer?
## [6][How do I make sure I don't accidentally create a large bill on AWS?](https://www.reddit.com/r/aws/comments/j0tyyh/how_do_i_make_sure_i_dont_accidentally_create_a/)
- url: https://www.reddit.com/r/aws/comments/j0tyyh/how_do_i_make_sure_i_dont_accidentally_create_a/
---
Basically title. I'm looking to set up some of my applications and compile processes on AWS, but read about all the horror stories of people being charged an arm and a leg because of a misconfiguration or misunderstanding of pricing. I've also heard of cases where legitimate usage suddenly blew up (a lot of users storming a website, for example) and the bill shot up exponentially.

I know there are options to set up billing alerts, but are there ways to tell AWS to cut off all instances given a certain price limit? Just want to avoid any chance of things mentioned above happening. Thanks!
## [7][Using S3 as a storage in Angular app](https://www.reddit.com/r/aws/comments/j1ax6u/using_s3_as_a_storage_in_angular_app/)
- url: https://www.reddit.com/r/aws/comments/j1ax6u/using_s3_as_a_storage_in_angular_app/
---
I am developing a multiuser web app using Angular for frontend + PHP for backend. I want to use S3 as a storage of users' files.

Each user should have the abilty to operate with his files (list all files / download / upload), but must not see other users' files - security is very important. Eventualy, I should introduce user groups and file sharing, so I was thinking of create IAM user for each user in my app.

I have a conceptual question should I communicate with S3 on frontend (Amplify?) or backend (AWS SDK for PHP)?
## [8][Can lambda connect to RDS using IAM authentication?](https://www.reddit.com/r/aws/comments/j1aluk/can_lambda_connect_to_rds_using_iam_authentication/)
- url: https://www.reddit.com/r/aws/comments/j1aluk/can_lambda_connect_to_rds_using_iam_authentication/
---
Hello there,

  
I am using a lambda function to run a sql script on a Postgre RDS DB. At the moment I am connecting to DB using username and password. Is it possible to connect to DB using RDS IAM Authentication?

Code for lambda is written in .NET core.

Thanks in advance.
## [9][Public SQL server to S3 parquet files: Best practice?](https://www.reddit.com/r/aws/comments/j1abon/public_sql_server_to_s3_parquet_files_best/)
- url: https://www.reddit.com/r/aws/comments/j1abon/public_sql_server_to_s3_parquet_files_best/
---
# The Scenario 
There is a publicly accessible SQL database, with data going back several years. 
Each day, new data are appended in the form of 1 minute snapshots of several sensors.

**Each day, I would like to download yesterdays data, and save it as a daily parquet file to an s3 bucket.**

# My currently solution
Use AWS Lambda with python 3.7, and a pandas and pyodbc layer to give me access to those modules.
The function runs a query on the server, then saves that data in parquet format to the S3 bucket.
Code is below.
I plan on adding in an SNS topic that gets pushed to in the event the function fails, so I can get an email letting me know if it's failed.

It does seem to work, but as I am very very new to all of this, and I'm not even sure if Lambda functions are the best place to do this or whether I should be using EC2 instances isntead. I wanted to ask **Is there a better way of doing this and is there anything I should watch for? Several stackoverflow posts suggest lambda might auto-retry on fails continuously, which i'd like to avoid!**

Thank you for being patient with an AWS newbie!

best,

Toast


        BASESQLQUERY = "SELECT * FROM TABLE"


	def getStartAndEndDates():
		""" Return yesterdays and todays dates as strings """
		startDate = datetime.now() - timedelta(3)
		endDate = datetime.now() - timedelta(2)
		datesAsStrings = [date.strftime('%Y-%m-%d') for date in [startDate, endDate]]
		return datesAsStrings 


	def runSQLQuery(serverAddress, 
				databaseName,
				username,
				password,
				datesAsStrings):
		""" Download yesterdays data from the database """
		with pyodbc.connect('DRIVER={ODBC Driver 17 for SQL Server};SERVER='+serverAddress+';DATABASE='+ databaseName +';UID='+username+';PWD='+ password) as conn:
			yesterday = datesAsStrings[0]
			today = datesAsStrings[1]
			fullSQLquery = BASESQLQUERY + f"WHERE TimeStamp BETWEEN '{yesterday}' AND '{today}';"
			dataReturnedFromQuery = pd.read_sql_query(fullSQLquery, conn)
		return dataReturnedFromQuery



	def lambda_handler(event, context):
                """Download yesterdays SQL data and save it as a parquet file in S3"""

		datesAsStrings = getStartAndEndDates()
		startDate, endDate = datesAsStrings

		logging.info(f'Downloading data from {startDate}.')
		try:
			logging.debug(f'Running SQL Query')
			dataReturnedFromQuery = runSQLQuery(serverAddress=SERVER_ADDRESS,
										    databaseName=DATABASE_NAME,
										    username=USERNAME,
										    password=PASSWORD,
										    datesAsStrings=datesAsStrings)
			logging.debug(f'Completed SQL Query')
			
			filename= startDate.replace('-','') + '.parquet'

			wr.s3.to_parquet(
				dataReturnedFromQuery ,
				f"s3://{BUCKET_NAME}/{filename}")
		except:
			logging.info(f'Failed to download data from {startDate}.')
			raise
		
		logging.info(f'Successfully downloaded data from {startDate}.')
		return {
			'statusCode': 200,
			'body': "Download Successfull"
		}
## [10][why sagemaker use jupyter notebooks?](https://www.reddit.com/r/aws/comments/j19ugd/why_sagemaker_use_jupyter_notebooks/)
- url: https://www.reddit.com/r/aws/comments/j19ugd/why_sagemaker_use_jupyter_notebooks/
---
I was just trying to understand how SageMaker is working from a high level point of view.

It seems that SageMaker launch JupyterNotebooks. Why? 

I'm new to JupyterNotebooks too. After a quick search it seems to me that JupyterNotebook is a kind of "environment" for several languages....it seems a nice web application and you can do nice things, but I don't really understand what it has to do with a programming environment to run machine learning models...

It is an alternative to build a Machine Learning systems on EC2 installing all the environment yourself?

But still, why a notebook application?

Really I miss something. Maybe some high level idea about that will enlighten me
## [11][Service agreement on getting shut down](https://www.reddit.com/r/aws/comments/j19twk/service_agreement_on_getting_shut_down/)
- url: https://www.reddit.com/r/aws/comments/j19twk/service_agreement_on_getting_shut_down/
---
What level of service agreement does AWS provide for this type of scenario where American government asks Amazon to stop doing business with a certain country? Would the account holders of that particular country be blocked and their accounts be frozen so that they can no longer access to servers or data on AWS cloud?
