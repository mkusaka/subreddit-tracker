# golang
## [1][The Zen of Go by Dave Chaney](https://www.reddit.com/r/golang/comments/eydowm/the_zen_of_go_by_dave_chaney/)
- url: https://the-zen-of-go.netlify.com/
---

## [2][Golang homework interview challenge for Storj - Looking for some feedback](https://www.reddit.com/r/golang/comments/eyphsm/golang_homework_interview_challenge_for_storj/)
- url: https://www.reddit.com/r/golang/comments/eyphsm/golang_homework_interview_challenge_for_storj/
---
I've posted this question on code review in case you want some programming internet points [https://codereview.stackexchange.com/questions/236641/golang-homework-interview-challenge-for-storj](https://codereview.stackexchange.com/questions/236641/golang-homework-interview-challenge-for-storj) 

Basically I submitted this solution [https://github.com/Samyoul/storj-file-sender](https://github.com/Samyoul/storj-file-sender) in response to a code challenge from Storj as part of their interview process. Storj is the only company that I've interviewed with that has not given any kind of feedback at all. So perhaps you guys could give me some pointers.

Thank you and let the brutalities commence :D
## [3][Troubles in scanning a MS SQL Date type to a struct, time.Time adds back the time and I want only the date.](https://www.reddit.com/r/golang/comments/eynrk9/troubles_in_scanning_a_ms_sql_date_type_to_a/)
- url: https://www.reddit.com/r/golang/comments/eynrk9/troubles_in_scanning_a_ms_sql_date_type_to_a/
---
Hoping to be as clear as possible:

* I am using the driver [go-mssqldb](https://github.com/denisenkom/go-mssqldb) (SQL Server) alongside the popular [sqlx](https://github.com/jmoiron/sqlx) package;
* the SQL table has a column with the [date](https://docs.microsoft.com/en-us/sql/t-sql/data-types/date-transact-sql?view=sql-server-ver15) type which stores date in the YYYY-MM-DD  format;
* from the [go-mssqldb readme](https://github.com/denisenkom/go-mssqldb#parameter-types) I read that it supports the [civil.Date](https://github.com/golang-sql/civil) type, I supposed it was a two-way compatibility but seems it works only when writing to the DB and not when retrieving.

The SQL query which fetches the data is in the form of

    err := db.Mssql.Select(&amp;orders, "some query")

Where orders is a pointer to a struct in the form of:

    type Orders struct {
            ...
    	Date civil.Date  `json:"dateOrder" db:"date_order"`
    	...
    }

Which returns the error:

    sql: Scan error on column index 5, name \"date_order\": unsupported Scan, storing driver.Value type time.Time into type *civil.Date"

If I change the type to time.Time the Select works but for reasons I've been trying to investigate in days without much success the driver always handles the sql server `date` type as `time.Time` which in turn after the scanning adds to the retrieved value the time.

&amp;#x200B;

Expected result would be to obtain  YYYY-MM-DD, instead I obtain  YYY-MM-DDTHH:MM:SSZ

&amp;#x200B;

I cannot of course change the [go-mssqldb](https://github.com/denisenkom/go-mssqldb) driver I am using, nor sqlx, so the reasonable solution I am trying to find a proper destination type before going into a tunnel where I need to loop all the data retrieved and truncate the time.

&amp;#x200B;

**Possible solution #1 (dirty way):**

Scan the date column to a string and truncate it. I would honestly prefer a cleaner way to handle this.

**Possible solution #2 (tightly coupled)**

As of now the date data is shown in a frontend where the Jhon Doe of the situation could use JavaScript to manipulate the date like new ***Date***(dateObject).toLocaleDateString()

Very discouraged as it forces a behavior in other applications which should ignore this logic and just receive a date.
## [4][Rest API CRUD with golang and sqlite](https://www.reddit.com/r/golang/comments/eyouid/rest_api_crud_with_golang_and_sqlite/)
- url: https://github.com/lazarospsa/rest_go_sqlite
---

## [5][Golang basics - writing unit tests](https://www.reddit.com/r/golang/comments/eyokgb/golang_basics_writing_unit_tests/)
- url: https://blog.alexellis.io/golang-writing-unit-tests/
---

## [6][Gupdate: simple CLI tool to update Go to the latest version](https://www.reddit.com/r/golang/comments/eyki46/gupdate_simple_cli_tool_to_update_go_to_the/)
- url: https://www.reddit.com/r/golang/comments/eyki46/gupdate_simple_cli_tool_to_update_go_to_the/
---
Instead of downloading .tar.gz and extracting on each new release, I created Gupdate. Simple but it does the job. Ideas for improvement are welcome. (Linux only for now)  
[https://github.com/Surendrajat/Gupdate](https://github.com/Surendrajat/Gupdate)  


\* Not a GO version manager or something. Updates current installation to the latest version seamlessly and nothing else.
## [7][Master Go While Learning Containers](https://www.reddit.com/r/golang/comments/eyfp1a/master_go_while_learning_containers/)
- url: https://iximiuz.com/en/posts/master-go-while-learning-containers/?utm_medium=reddit&amp;utm_source=r_golang
---

## [8][Garnish - simple varnish implementation written in Go](https://www.reddit.com/r/golang/comments/eyoll2/garnish_simple_varnish_implementation_written_in/)
- url: https://developer20.com/garnish-simple-varnish-in-go/?utm_source=reddit&amp;utm_medium=link&amp;utm_campaign=garnish
---

## [9][Pet Projects are Fun, building the most over-engineered Go tool](https://www.reddit.com/r/golang/comments/eyobjm/pet_projects_are_fun_building_the_most/)
- url: https://www.reddit.com/r/golang/comments/eyobjm/pet_projects_are_fun_building_the_most/
---
Update your go.mod dependencies to their latest version by creating a parser to go.mod files.

Read the article:

[https://medium.com/securenative/pet-projects-are-fun-763ee39a732b?source=friends\_link&amp;sk=b4eea595cf3bdedb28dd29d71403a768](https://medium.com/securenative/pet-projects-are-fun-763ee39a732b?source=friends_link&amp;sk=b4eea595cf3bdedb28dd29d71403a768)

Go to Github:

[https://github.com/matang28/go-latest](https://github.com/matang28/go-latest)
## [10][How would you convince an archaic Bank to build its new backend in GO instead of JAVA?](https://www.reddit.com/r/golang/comments/eyo9qp/how_would_you_convince_an_archaic_bank_to_build/)
- url: https://www.reddit.com/r/golang/comments/eyo9qp/how_would_you_convince_an_archaic_bank_to_build/
---
I am currently in talks with an 80-year-old bank, trying to convince them to let us build their upcoming API in GO, rather than JAVA. We currently have more experience building the system in GO, and with a tight deadline in mind, I would rather invest time building the technology, rather than learning and planning it out in JAVA. I need some points that will be able to convince them. If you feel JAVA is still the right fit for building the banking infrastructure, I would lie to know your point of view aswell.
