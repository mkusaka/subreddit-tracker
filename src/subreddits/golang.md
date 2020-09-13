# golang
## [1][We are the Go Time podcast. Ask us anything. (AMA)](https://www.reddit.com/r/golang/comments/io94yi/we_are_the_go_time_podcast_ask_us_anything_ama/)
- url: https://www.reddit.com/r/golang/comments/io94yi/we_are_the_go_time_podcast_ask_us_anything_ama/
---
Hi everyone! I'm Jon Calhoun, one of the panelists on the Go Time podcast. For those of you unfamiliar - it is a Go podcast that we record live every Tuesday at 3pm ET. We usually have a guest or two on each episode, and discuss a specific topic (defer, testing, databases, infra, etc). You can check it out here: &lt;https://changelog.com/gotime&gt;

This coming episode we want to try something a little different - we want to make a Q&amp;A episode. There are two reasons for doing this:

1. We are hoping this inspires topics for future episodes.
2. We want a venue to discuss questions that don't really fit into an entire episode on their own.

To make this happen we would like everyone here to post any Go-related questions that you would like us to discuss on air, and we will try to get to as many as possible. I'll also try to type up answers here while we discuss them on the episode.

We will be answering questions live tomorrow, Tuesday, Sep 8. We will repeat questions on air, and since we record live you can join in on the Gophers Slack to ask follow-up questions or to elaborate on questions.
## [2][Converts 'go mod graph' output into Graphviz's DOT language](https://www.reddit.com/r/golang/comments/irupxd/converts_go_mod_graph_output_into_graphvizs_dot/)
- url: https://www.reddit.com/r/golang/comments/irupxd/converts_go_mod_graph_output_into_graphvizs_dot/
---
[https://github.com/lucasepe/modgv](https://github.com/lucasepe/modgv)

&amp;#x200B;

Graphically view the dependencies of your latest project.

https://preview.redd.it/9t2h8xumnvm51.jpg?width=1280&amp;format=pjpg&amp;auto=webp&amp;s=fa3b692eb4ddfb50f5a3192f7f820f68acc36a05
## [3][Playwright for Go: a browser automation library to control Chromium, Firefox and WebKit with a single API.](https://www.reddit.com/r/golang/comments/irvizx/playwright_for_go_a_browser_automation_library_to/)
- url: https://github.com/mxschmitt/playwright-go
---

## [4][Go Discord Bot to Auto-mute players in Among Us](https://www.reddit.com/r/golang/comments/irpja2/go_discord_bot_to_automute_players_in_among_us/)
- url: https://github.com/denverquane/amongusdiscord
---

## [5][Asynq v0.12 released with Redis Cluster support](https://www.reddit.com/r/golang/comments/irrwts/asynq_v012_released_with_redis_cluster_support/)
- url: https://www.reddit.com/r/golang/comments/irrwts/asynq_v012_released_with_redis_cluster_support/
---
[Asynq](https://github.com/hibiken/asynq) is a Redis-backed task queue library in Go.

The library has always supported connecting to a single Redis instance or connecting to Redis using Sentinels.

With v0.12 release, the library now [supports Redis Cluster](https://github.com/hibiken/asynq/wiki/Redis-Cluster). It shards task data by queue, so the application developer has control over how the data are split. It also comes with the CLI support, just append `--cluster` flag to run the commands against Redis cluster.

&amp;#x200B;

[Overview of Asynq with Redis cluster](https://preview.redd.it/4bi41s2xgum51.png?width=811&amp;format=png&amp;auto=webp&amp;s=342dda7d88e3c70c75a644f102cfad6c5bec826b)

Please check it out and feedbacks are appreciated! Thank you :)
## [6][Go gRPC Backend Playground](https://www.reddit.com/r/golang/comments/irwro2/go_grpc_backend_playground/)
- url: https://backend-playground.transcendent.app/
---

## [7][Leaf – General purpose hot-reloader for all projects](https://www.reddit.com/r/golang/comments/irgv18/leaf_general_purpose_hotreloader_for_all_projects/)
- url: https://github.com/vrongmeal/leaf
---

## [8][The go/ast package is seriously cool. Let's build an automatic document generator for a NATS microservice!](https://www.reddit.com/r/golang/comments/irfkf9/the_goast_package_is_seriously_cool_lets_build_an/)
- url: https://medium.com/@riptidedata/cool-stuff-with-gos-ast-package-pt-1-981460cddcd7
---

## [9][Casbin-Forum: Next-generation forum software based on React + Golang](https://www.reddit.com/r/golang/comments/irplrs/casbinforum_nextgeneration_forum_software_based/)
- url: https://github.com/casbin/casbin-forum
---

## [10][I had previously uploaded a video where I built a load balancer in Go. It got a good response. So, I made a follow-up video where I add a "health check" functionality to it. Do check it out and let me know what you think](https://www.reddit.com/r/golang/comments/iryd9w/i_had_previously_uploaded_a_video_where_i_built_a/)
- url: https://youtu.be/r9mcmZEhD9Q
---

## [11][Using CSV as a DB](https://www.reddit.com/r/golang/comments/irpenw/using_csv_as_a_db/)
- url: https://www.reddit.com/r/golang/comments/irpenw/using_csv_as_a_db/
---
I'm writing a small daemon for personal use -- 1 user -- and I'd like to retain the ability to manually edit the data with my editor. I'm not yet sure how many rows, but I doubt it'll be more than 1,000. Might have a few "tables" (CSV files) with cross-referenced IDs (probably just mirroring line numbers).

I'm wondering if anyone has experience / insights using CSV for this. I came across a project that allowed querying CSV with SQL -- that's overkill for my use case. And I'll also be writing to the file -- updating rows.

I'm wondering if there are known limits to this approach. E.g., does it start to get perceptibly slow at 100,000 or 1,000,000 rows? Are field updates tedious because you have to scan newlines, then commas -- and then deal with insertions?

Any recommended literature for working with the filesystem like this?

Still just exploring the idea -- apologies for the lack of detail.
