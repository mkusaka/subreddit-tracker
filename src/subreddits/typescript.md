# typescript
## [1][Who's hiring Typescript developers - January](https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/)
- url: https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][Distributing Pick&lt;T, K&gt;/Omit&lt;T,K&gt; over union types in TypeScript](https://www.reddit.com/r/typescript/comments/ejiudx/distributing_pickt_komittk_over_union_types_in/)
- url: https://davidgomes.com/pick-omit-over-union-types-in-typescript/
---

## [3][A short time learning Haskell can shortly IMPROVE your Typescript skills!](https://www.reddit.com/r/typescript/comments/ejllr2/a_short_time_learning_haskell_can_shortly_improve/)
- url: https://medium.com/@lironhazan/a-short-time-learning-haskell-can-shortly-improve-your-typescript-skills-523505900ac0?source=friends_link&amp;sk=daf9f13302e63282bc7664574873fa8c
---

## [4][Top-level await in TS 3.8 Nightly Build?](https://www.reddit.com/r/typescript/comments/ejgajz/toplevel_await_in_ts_38_nightly_build/)
- url: https://www.reddit.com/r/typescript/comments/ejgajz/toplevel_await_in_ts_38_nightly_build/
---
I'm trying to use top-level await in a TS 3.8. I installed the dev release from today (1/3/20). 

When writing typescript it continues to complain: "'await' expression is only allowed within an async function. (TS1308) when I run tsc.

Am I missing some step to allow top-level await?
## [5][Is this Guard class useful?](https://www.reddit.com/r/typescript/comments/ejq53p/is_this_guard_class_useful/)
- url: https://www.reddit.com/r/typescript/comments/ejq53p/is_this_guard_class_useful/
---
Hey -- 

I'm new to Typescript.  

I created this class:

 [https://github.com/NickHodges/TSGuard/blob/master/Guard.ts](https://github.com/NickHodges/TSGuard/blob/master/Guard.ts) 

is this useful in Typescript, or am I wasting my time?

I can, of course, expand on it.
## [6][Node CLI startup cache: any other projects like this?](https://www.reddit.com/r/typescript/comments/ejo4ez/node_cli_startup_cache_any_other_projects_like/)
- url: https://www.reddit.com/r/typescript/comments/ejo4ez/node_cli_startup_cache_any_other_projects_like/
---
I have an idea for a project to improve startup of CLI tools written in TypeScript / JavaScript via aggressive caching and lazy-loading of dependencies, without requiring pre-compilation or webpacking.  The idea is to hit a sweet-spot between CLI performance and developer experience.

* use `bundleDependencies` to lock all dependency versions.
* cache all `require.resolve()` results, all `fs.readFileSync()` related to loading modules
* cache all `ts.transpileModule()` results from `ts-node`.
* A bit of Proxy magic to lazy-load imports automatically, without code changes.
* Include v8-compile-cache

Effectively, code can be written to use ts-node in the simplest way possible.  At development time, this will "just work" without a compile step.  Production releases will automatically create a cache after the first startup.  For subsequent startups, a single cache file is loaded, after which everything is a cache hit in memory.  No more fs calls, no transpiler calls.

Anyone know of a project that already does this?

So far everything I've seen does not take this approach.  They either:

a) bundle it all into a new node binary.  
b) webpack, which can cause weird breakages, incompatibilities, and debugging issues.
## [7][Blogged Answers: Years in Review, 2018-2019 · Mark's Dev Blog - Story of redux and typescript](https://www.reddit.com/r/typescript/comments/ejc6g6/blogged_answers_years_in_review_20182019_marks/)
- url: https://blog.isquaredsoftware.com/2020/01/blogged-answers-years-in-review-2018-2019/
---

## [8][Generate type aliases from enums](https://www.reddit.com/r/typescript/comments/ej51kk/generate_type_aliases_from_enums/)
- url: https://www.reddit.com/r/typescript/comments/ej51kk/generate_type_aliases_from_enums/
---
    enum Birds {
        eagle,
        cardinal,
        falcon
    }
    
    type BirdTypes = keyof typeof Birds; // 'eagle' | 'cardinal' | 'falcon'

For the longest time I’ve been duplicating efforts when it came to having both an enum and type alias pointing to the same data, just learned about `typeof` which makes this a lot easier.
## [9][Looking for a library to generate images](https://www.reddit.com/r/typescript/comments/ej8y1k/looking_for_a_library_to_generate_images/)
- url: https://www.reddit.com/r/typescript/comments/ej8y1k/looking_for_a_library_to_generate_images/
---
Hi! I'm hoping this is the right place I can ask this question. I really like types and have gotten used to using typescript at work. I'm looking to work on a personal project where I generate high quality images programmatically (which in this case means pixel by pixel). Maybe if I actually get this off the ground it'll become clear what I mean.

Anyway, I'm seeing lots of libraries/npm packages online that talk about image manipulation (resizing, overlaying with text, etc) and ones that care about doing this on the web (like canvas). ~~If any of these are good for my use, it seems like it'd be at best by accident.~~ (Edit: To be clear, I don't think these are what I want but they are all I've been able to find.) Anyway, I'd like to be able to generate images that are of high enough quality to print. The point is looking the best over speed. Does anyone know of such a library?

Good typing would be a huge plus in my book.

Thanks so much for any advice!
## [10][Non-null assertion operator vs. Optional chaining](https://www.reddit.com/r/typescript/comments/ej924r/nonnull_assertion_operator_vs_optional_chaining/)
- url: https://www.reddit.com/r/typescript/comments/ej924r/nonnull_assertion_operator_vs_optional_chaining/
---
I was just wondering, is there any difference between the two when you're trying to access properties that may possibly be null or undefined? It seems to me optional chaining is the safer one of the two, though I've yet to use it a lot. I defaulted to using non-null assertion after dealing with async data in GraphQL and only just learned about optional chaining for accessing properties recently.
## [11][Are there any naming conventions for modules?](https://www.reddit.com/r/typescript/comments/ej6kip/are_there_any_naming_conventions_for_modules/)
- url: https://www.reddit.com/r/typescript/comments/ej6kip/are_there_any_naming_conventions_for_modules/
---
I have a module with a class in it.  The class is called `Guard` and I called the module `GuardClass.`

I'm less than satisfied with this.  Does anyone have any suggestions for naming conventions for modules?
