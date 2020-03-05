# golang
## [1][dotGo 2020 has been cancelled](https://www.reddit.com/r/golang/comments/fdkbd1/dotgo_2020_has_been_cancelled/)
- url: https://twitter.com/dotGoEu/status/1234767083369639937
---

## [2][Unit testing CLI programs](https://www.reddit.com/r/golang/comments/fdrx7b/unit_testing_cli_programs/)
- url: https://medium.com/swlh/unit-testing-cli-programs-in-go-6275c85af2e7
---

## [3][Look what my colleague gave me. It's so cute! 🥰](https://www.reddit.com/r/golang/comments/fdab62/look_what_my_colleague_gave_me_its_so_cute/)
- url: https://i.redd.it/w44c9ga5emk41.jpg
---

## [4][How would you handle database interaction in unit tests?](https://www.reddit.com/r/golang/comments/fdvalp/how_would_you_handle_database_interaction_in_unit/)
- url: https://www.reddit.com/r/golang/comments/fdvalp/how_would_you_handle_database_interaction_in_unit/
---
Imagine you have started writing a new Go application, and you're about to write some unit tests that require database interaction. How would you do so? How would you store the environment configs for different environments?

I assume I need to use a separate database for tests and to run migrations for this database separately. 

Since previously I was writing applications in Ruby on Rails, I thought it would be best to have separate config/env files for each environment and have some handy commands that could be run on each environment outside of my app's logic.

What I managed to do, was storing my environment variables in two files under:

    config/.dev.env
    config/.test.env

Then using [migrate package](https://github.com/golang-migrate/migrate) for migrations, and putting migrate commands in the Makefile similar to below:

    ENV?=dev
    
    include config/.$(ENV).env
    
    export
    POSTGRES_URL=postgres://${POSTGRESQL_USER}:${POSTGRESQL_PASSWORD}@${POSTGRESQL_HOST}:${POSTGRESQL_PORT}/${POSTGRESQL_DB}?sslmode=disable
    
    # create the database
    .PHONY: db-create
    db-create:
    	@env "PGPASSWORD=${POSTGRESQL_PASSWORD}" psql -h "${POSTGRESQL_HOST}" -p "${POSTGRESQL_PORT}" \
    			-U "${POSTGRESQL_USER}" -w -c "SELECT 1 FROM pg_database WHERE datname = '${POSTGRESQL_DB}';" \
    			| grep 1 \
    			|| env "PGPASSWORD=${POSTGRESQL_PASSWORD}" psql -h "${POSTGRESQL_HOST}" -p "${POSTGRESQL_PORT}" \
    				-U "${POSTGRESQL_USER}" -w -c "create database \"${POSTGRESQL_DB}\";"
    
    # apply all new migrations. pass steps flag to limit number of migrations to run
    .PHONY: migrate-up
    migrate-up:
    	@migrate -database ${POSTGRES_URL} -path ./database/migrations up ${steps}

This allows me to run migrations for my test database with a simple: `make migrate-up ENV=test`

I have not found a similar approach looking at other go applications. Is there a better way?
## [5][Plik - OpenSource temporary file upload system ( wetransfer like ) in Go](https://www.reddit.com/r/golang/comments/fdv9ku/plik_opensource_temporary_file_upload_system/)
- url: https://www.reddit.com/r/golang/comments/fdv9ku/plik_opensource_temporary_file_upload_system/
---
We are pleased to release a first RC of Plik 1.3

Plik is an OpenSource temporary file upload system ( wetransfer like )

It is shipped with a cool UI and CLI.  
Files can be stored locally or in S3 or OpenStack Swift.  
Metadata can be stored locally ( sqlite3 ) or in a postgresql database.  
It supports authentication using local users or Goolgle API.  


Some cool features like one shot upload deleted after the first download, stream mode where the files are not stored on the server and piped directly from the uploader to the dowloader, basic auth protection,...

Give it a go ;)
## [6][Use or gRPC or json/REST for web services?](https://www.reddit.com/r/golang/comments/fdsy9h/use_or_grpc_or_jsonrest_for_web_services/)
- url: https://www.reddit.com/r/golang/comments/fdsy9h/use_or_grpc_or_jsonrest_for_web_services/
---
Hi, I am building a side project which requires to deploy an agent in customers docker clusters.

This agent will look for some events and push them to a backend server.

Right now, I have implemented 2 kinds of way to transmit those events:

* based on grpc (stream and non-stream)
* json/REST

In the case of a grpc stream, I would need to open the stream at the beginning of the program, keep the stream open and push events to the stream whenever they are coming. Would you see any drawback in keeping that connection open for the entire lifecycle of the process? Is the grpc client smart enough to maintain the connection open *no matter what*?

I can see some cons of using grpc for non-internal communication: mostly it requires that the connection is not intercepted by TLS terminating proxies which is not often the case. It feels like the advantages of using grpc (security, performance) in that case do not counter-balance the drawbacks I might face and I might be digging a rabbit hole.

What's your experience with grpc based web-services? would you recommend it in my use case?
## [7][Zip file with nested folders and architecture issue](https://www.reddit.com/r/golang/comments/fdv1bb/zip_file_with_nested_folders_and_architecture/)
- url: https://www.reddit.com/r/golang/comments/fdv1bb/zip_file_with_nested_folders_and_architecture/
---
I'm trying to build a web application which 
• Takes input from user
• Uses it to run a process which creates files
• Serves them to the user


Right now this is the flow of the app
• Takes some input from a html form
• Runs an external process (with os/exec) that uses user input as flags. 
• External process creates directories with files
• Zip the files I need 
• Serve them to the user

The way I separate files for each user is that for every request I create a random directory for each.
I then run the external program and set each user's directory as the output dir.

My issue is that this external program creates other directories and then puts actual output in the most nested so for every request I have :
RandomUserDir/X/output/*FilesINeedToZip

When I try to zip these files, the .zip I get will have the same structure when unzipped so the end user will see and have to go through all 3 directories before accessing his files which is not really great user experience.

Warning : the following feels like really bad practice

I tried to solve this by running os.Chdir and going into the output directory to zip the files and then zip which works - i don't get all dirs in the zip file - but if the user tries to make another request it won't work because the app will try to chdir in the directory it's already in

I solved this by having the handler chdir back to the original directory where data is stored and things kind of work now but when two users make a request simultaneously one of them will still get an error because of chdir.

This feels like pretty bad design so do you know what would be a better way to architecture this ?
 And maybe a way of not getting all directories when zipping but juste the files i need?

Any help would be appreciated thank you.
## [8][VSCode Rant](https://www.reddit.com/r/golang/comments/fdggl7/vscode_rant/)
- url: https://www.reddit.com/r/golang/comments/fdggl7/vscode_rant/
---
I need to rant because my frustration with Go and Vscode can no longer be contained.

&amp;#x200B;

I am primarily a Backend NodeJS Developer, but I love go. I have been advocating for years to use Go at my job. I have written some tools in Go that were adopted with success at my work. I have written small personal projects and had a blast doing it.

&amp;#x200B;

However, the state of gopls, and the vscode integration is just so terrible that I am no longer comfortable recommending Go to other developers who are primarily web based or from a Javascript background. Vscode intelisense and type inference is better for vanilla js than it is for go, a static language. The language server is constantly broken, and since I have updated Go to v1.14 I can't get the language server to work at all. My symbols are never loaded. No features work.

&amp;#x200B;

I understand that people will tell me to switch to GoLand, but I don't like JetBrain products. Nothing wrong with them but they're not for me, and regardless I am not going to try and convince a team of web developers to change their editors so they can pick up a new language. Just not going to happen.

&amp;#x200B;

I understand people debate go over "err != nil" or lack of generics, or issues that are language related. But having the tooling broken with arguably the most popular code editor is just unacceptable. 

&amp;#x200B;

Currently I have no pleasant environment to work with Go. I might as well download notepad++ and run a terminal instance next to it and depend entirely on the compiler messages.

&amp;#x200B;

As it stands, as much as I love Go, I think until the editor tooling gets better, I have no choice but to abandon my efforts to get other people to use it, because I myself am not enjoying my experience coding in my favourite language. Even my own personal adoption of it is decreasing over time just because of how piss poor the developer experience is now.

&amp;#x200B;

Rant over.

&amp;#x200B;

If anybody has a magic bullet / config to fix vscode and go I would love to hear it.

Otherwise, I'll check back in with the next few months hoping it is resolved, but I somehow doubt they won't just break it again immediately.
## [9][Is it normal to literally google everything?](https://www.reddit.com/r/golang/comments/fdtnit/is_it_normal_to_literally_google_everything/)
- url: https://www.reddit.com/r/golang/comments/fdtnit/is_it_normal_to_literally_google_everything/
---
I'm new, but I'm not new new, like 8 months new. I google literally anything. Even things that i did them 15-20 times before. Should i start memorizing or this is something normal?
## [10][Forward proxy with rate limiting.](https://www.reddit.com/r/golang/comments/fdth2t/forward_proxy_with_rate_limiting/)
- url: https://github.com/jamesmoriarty/goforward
---

