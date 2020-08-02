# golang
## [1][[Q&amp;A] io/fs draft design](https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/
---
I posted a draft design today for new file system interfaces for Go.

Video: https://golang.org/s/draft-iofs-video

Design: https://golang.org/s/draft-iofs-design

Let's do the Q&amp;A about the design here in Reddit. My hope is that the threading support will help keep questions and answers matched.

Please start a new top-level comment for each new question.

See also the related [Q&amp;A for the //go:embed draft design](https://golang.org/s/draft-embed-reddit).
## [2][Design Draft: First Class Fuzzing](https://www.reddit.com/r/golang/comments/hvpr96/design_draft_first_class_fuzzing/)
- url: https://go.googlesource.com/proposal/+/refs/heads/master/design/40307-fuzzing.md
---

## [3][Check out my Go app that lets you split files into horcruxes](https://www.reddit.com/r/golang/comments/i26l3g/check_out_my_go_app_that_lets_you_split_files/)
- url: https://www.reddit.com/r/golang/comments/i26l3g/check_out_my_go_app_that_lets_you_split_files/
---
Not too long ago I made a program in Go called Horcrux which allows you to split a file into multiple 'horcruxes' (aka fragments) which can then be recombined to obtain the original file (as opposed to encrypting a single file with a password).  


Somebody urged me to create a GUI for it so I've done just that! I've used [https://fyne.io/](https://fyne.io/) and it's a pretty lightweight frontend. I'd prefer to use a native file select dialog but fyne currently doesn't support that. Let me know your thoughts :)  


[https://github.com/jesseduffield/horcrux-ui](https://github.com/jesseduffield/horcrux-ui)
## [4][grule-rule-engine version 1.5. 0 released. Grule is a Rule Engine library for the Golang programming language. Inspired by the acclaimed JBOSS Drools, done in a much simple manner.](https://www.reddit.com/r/golang/comments/i2634a/gruleruleengine_version_15_0_released_grule_is_a/)
- url: https://github.com/hyperjumptech/grule-rule-engine/blob/master/CHANGELOG.md#150---2020-08-02
---

## [5][todocheck v0.2.0 is live](https://www.reddit.com/r/golang/comments/i1wcgb/todocheck_v020_is_live/)
- url: https://github.com/preslavmihaylov/todocheck/releases/tag/v0.2.0
---

## [6][Using GitHub Actions to compile, sign, and notarize your MacOS Golang Binaries](https://www.reddit.com/r/golang/comments/i2a2fx/using_github_actions_to_compile_sign_and_notarize/)
- url: https://www.kencochrane.com/2020/08/01/build-and-sign-golang-binaries-for-macos-with-github-actions/
---

## [7][DDD vs. DB transactions. How to reconcile?](https://www.reddit.com/r/golang/comments/i1vy4s/ddd_vs_db_transactions_how_to_reconcile/)
- url: https://www.reddit.com/r/golang/comments/i1vy4s/ddd_vs_db_transactions_how_to_reconcile/
---
TL;DR

How can DDD (Domain Driven Design) architecture work with the need to use database transactions, or are they just simply incompatible?

---

I've been working on this small web app as a project to practice go web app dev with. I recently came across this [gophercon talk](https://www.youtube.com/watch?v=oL6JBUk6tj0) with Kat Zien where she describes Domain Driven Design, among other attempts at organizing a project. I thought why not try to apply the DDD principles to my practice project because it looks like a nice clean way to keep layers separated and flexible.

As a result I had my DB connection contained within "repositories", isolated away from the rest of the code base. The domain and infrastructure (specifically http) layers had no knowledge of which database (or simply persistence layer) was in use, only that it did it's job. Great. Separation of concerns.

Then, I ran into the need to use transactions which seems to turn DDD on its head due to the fact that accommodating a transaction (at the request level) forces it to be a cross cutting concern.

Basically what I did was kept the concept of the services and repositories in place but the repositories no longer have ownership of a database connection/implementation because the connection now gets passed to any repository operations that need it.

For instance, an http handler will kick off a transaction like so:

    storage.Transact(c.DB, func(ext storage.Ext) (err error) {
        // ...
        err = widgetService.UpdateWidget(ext, widgetID, widget)
        if err != nil {
            // ... oops, bad thing happened, rollback
            return err
        }
    })

The rollback and commit code is abstracted behind the `storage.Transact` function. The `storage.Ext` is an interface that will allow for use of an `*sqlx.DB` struct or `*sqlx.Tx`. That way I can run any service or repo function with either type of db struct, connection or transaction.

The `storage.Ext` interface is thus:

    type Ext interface {
        sqlx.Execer
        sqlx.Queryer
        Preparex(string) (*sqlx.Stmt, error)
    }

So now the persistence layer implementation is out in the open, not tucked away where the rest of the application can hum along with no knowledge of persistence implementation. I've lost the benefit of isolating it away from the rest of the app.

It occurred to me that I could move the transaction to be much closer to the repo layer and have each repository operation have it's own self contained transaction but that would lead to a lot of duplication of code. Many operations would not be reusable suddenly as they would need to be re-implemented any time they need to take place as part of a transaction if they were once self contained atomic operations of their own. I suppose one approach would be to have internal repo ops all take `storage.Ext` as first arg and the repo public interface would hide all this away.

Now I'm just thinking out loud stream of consciousness style... help me straighten my rudder out here. How would a gopher sort this out?
## [8][Go: How to Reduce Lock Contention with the Atomic Package](https://www.reddit.com/r/golang/comments/i1ryn9/go_how_to_reduce_lock_contention_with_the_atomic/)
- url: https://medium.com/a-journey-with-go/go-how-to-reduce-lock-contention-with-the-atomic-package-ba3b2664b549
---

## [9][pgx and migrations](https://www.reddit.com/r/golang/comments/i1zz5i/pgx_and_migrations/)
- url: https://www.reddit.com/r/golang/comments/i1zz5i/pgx_and_migrations/
---
Hey all,

Looking to switch over to using PGX natively.. e.g. no sql.db and such. Read several bits on it now and it seems like its a better option as I am only using PostgresQL (at least between various services that share the DB (but have separate DB instances for their services)). 

I have been using the golang-migrate/migrate for migration and as I am new to all this.. and nothing in production or even beta yet, really don't have a lot to migrate.. just a couple of table inserts to be able to get some code working like a root user to log in with, role types, etc. Eventually I assume I may have some breaking DB changes that migrate will hopefully help with and be of more value than it is currently for me. 

That said, is there a good solution, used by someone in production, tested, etc that works with PGX directly? 

Let me ask this.. given that migration is really just a series of files, in sequential order, of SQL commands.. is there anything wrong with just writing a bit of go code specifically for migration use, that I can run directly..without needing to depend on golang-migrate? It seems like especially for my own use which wont be a very complex or a huge amount of tables, easy enough to just write (and document) the SQL myself. From what I can tell.. the migrate .sql scripts are basically executed in sequential order. So just thinking if I wrote some code using pgx that executes scripts in one way, or the other (up or down), what is the major thing I am getting with something like golang-migrate that I would have to be concerned about doing myself?
## [10][Use Go to access QuickBooks desktop file?](https://www.reddit.com/r/golang/comments/i238ec/use_go_to_access_quickbooks_desktop_file/)
- url: https://www.reddit.com/r/golang/comments/i238ec/use_go_to_access_quickbooks_desktop_file/
---
I need to use Go to extract some employee/payroll info from a desktop Quickbooks file (on Windows of course)? Any suggestions on how to accomplish this?

&amp;#x200B;

PS: Would someone please inform the spam bot that this is not spam. Thanks.
## [11][Is there any existing queue management system in golang?](https://www.reddit.com/r/golang/comments/i229ae/is_there_any_existing_queue_management_system_in/)
- url: https://www.reddit.com/r/golang/comments/i229ae/is_there_any_existing_queue_management_system_in/
---
Is there any existing queue management system in golang from where I can see the

* stats of each queue,
* status of each tasks in the queue
* managing queue, workers and tasks
* migrating tasks from one queue to another queue

My use-case scenario:

I'm trying to implement SMS sending platform for multiple users where each users could have their own set of priority queues + flexible workers. The queue is transparent to user for each sms stats.

Users could view the tasks on queue and be able to Pause/Resume/Kill/Migrate them
## [12][GLab: A Gitlab CLI tool](https://www.reddit.com/r/golang/comments/i1lnqo/glab_a_gitlab_cli_tool/)
- url: https://www.reddit.com/r/golang/comments/i1lnqo/glab_a_gitlab_cli_tool/
---
Glab is a simple and elegant GitLab CLI tool written in Go (golang) for creating and managing issues, merge requests, and a lot more... Do [give it a try](https://github.com/profclems/glab). Your feedback is much appreciated.

[https://github.com/profclems/glab](https://github.com/profclems/glab)

https://gitlab.com/profclems/glab

&amp;#x200B;

https://preview.redd.it/cfcjpttwabe51.png?width=1024&amp;format=png&amp;auto=webp&amp;s=99d159203ce82dbddae57f732becea64b64c00d2
