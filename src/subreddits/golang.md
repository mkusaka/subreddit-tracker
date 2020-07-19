# golang
## [1][Brian Kernighan: UNIX, C, AWK, AMPL, and Go Programming | AI Podcast #109 with Lex Fridman](https://www.reddit.com/r/golang/comments/htwr7e/brian_kernighan_unix_c_awk_ampl_and_go/)
- url: https://www.youtube.com/watch?v=O9upVbGSBFo
---

## [2][Static code analyzer for TODO comments, written in Go](https://www.reddit.com/r/golang/comments/htkvik/static_code_analyzer_for_todo_comments_written_in/)
- url: https://github.com/preslavmihaylov/todocheck
---

## [3][How do I make my prime program faster than JS?](https://www.reddit.com/r/golang/comments/hty9ij/how_do_i_make_my_prime_program_faster_than_js/)
- url: https://www.reddit.com/r/golang/comments/hty9ij/how_do_i_make_my_prime_program_faster_than_js/
---
Golang noob here.

I have a similar GO &amp; JS program to find the first N primes(naive). The golang code takes ~twice the time. What am I doing wrong?

**GO**
```
  const TARGET = 20000 // assume user input, unknown at compile.
  var count = 0
  var num = 2
  var primes []int

  for count &lt; TARGET {
    divisible := false
    for i := 2; i &lt; num; i++ {
      if (num % i == 0) {
        divisible = true 
        break
      }
    } 
    if (divisible == false) {
      primes = append(primes, num) // suspect doing this wrong
      count++
    }
    num++
  }
```

**Javscript**
```
let now = performance.now();
const TARGET = 20000;
let primes = [];
let num = 2;
let count = 0;

while (count &lt; TARGET) {
  let divisible = false
  for (let i = 2; i &lt; num; i++) {
    if (num % i == 0) {
      divisible = true;
      break;
    }
  }
  if (divisible == false) {
    primes.push(num);
    count++;
  }
  num++
}

```
## [4][cockroachdb/pebble -- RocksDB/LevelDB inspired key-value database in Go](https://www.reddit.com/r/golang/comments/htre6o/cockroachdbpebble_rocksdbleveldb_inspired/)
- url: https://github.com/cockroachdb/pebble
---

## [5][What is the stack should I use to build my go backend service (API)](https://www.reddit.com/r/golang/comments/hu051j/what_is_the_stack_should_i_use_to_build_my_go/)
- url: https://www.reddit.com/r/golang/comments/hu051j/what_is_the_stack_should_i_use_to_build_my_go/
---
As you know, there are a lot of frameworks for GO, database connectors, and stuff.

So my options are a bit cluttered as I don't know the use cases... So I'm posting here for guidance.
## [6][k8s pod scheduler a simple k8s scheduler to run long-running tasks with a gRPC Interface](https://www.reddit.com/r/golang/comments/htwpsq/k8s_pod_scheduler_a_simple_k8s_scheduler_to_run/)
- url: https://github.com/ahmagdy/k8s-pod-scheduler
---

## [7][Using Go build directives to optionally use new APIs in the standard library](https://www.reddit.com/r/golang/comments/htyzsx/using_go_build_directives_to_optionally_use_new/)
- url: https://utcc.utoronto.ca/~cks/space/blog/programming/GoBuildUsingNewAPIs
---

## [8][Module declares its path as, but was required as...](https://www.reddit.com/r/golang/comments/htyn0l/module_declares_its_path_as_but_was_required_as/)
- url: https://www.reddit.com/r/golang/comments/htyn0l/module_declares_its_path_as_but_was_required_as/
---
Trying to use argon2 with go.mod and getting the following error:

    go: github.com/golang/crypto/argon2: github.com/golang/crypto@v0.0.0-20200709230013-948cd5f35899: parsing go.mod:
    	module declares its path as: golang.org/x/crypto
    	        but was required as: github.com/golang/crypto

I've tried changing go.mod and the import address, but not having much luck. Any ideas?
## [9][reflink: very small package for btrfs/xfs copy-on-write file copy in native go](https://www.reddit.com/r/golang/comments/htfdhh/reflink_very_small_package_for_btrfsxfs/)
- url: https://github.com/KarpelesLab/reflink
---

## [10][Deploying Your Go Apps to AWS Elastic Beanstalk](https://www.reddit.com/r/golang/comments/htm1e0/deploying_your_go_apps_to_aws_elastic_beanstalk/)
- url: https://youtu.be/1WXJTlkf0S4
---

