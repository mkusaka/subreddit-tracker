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
## [2][Blogged Answers: Years in Review, 2018-2019 · Mark's Dev Blog - Story of redux and typescript](https://www.reddit.com/r/typescript/comments/ejc6g6/blogged_answers_years_in_review_20182019_marks/)
- url: https://blog.isquaredsoftware.com/2020/01/blogged-answers-years-in-review-2018-2019/
---

## [3][Top-level await in TS 3.8 Nightly Build?](https://www.reddit.com/r/typescript/comments/ejgajz/toplevel_await_in_ts_38_nightly_build/)
- url: https://www.reddit.com/r/typescript/comments/ejgajz/toplevel_await_in_ts_38_nightly_build/
---
I'm trying to use top-level await in a TS 3.8. I installed the dev release from today (1/3/20). 

When writing typescript it continues to complain: "'await' expression is only allowed within an async function. (TS1308) when I run tsc.

Am I missing some step to allow top-level await?
## [4][Generate type aliases from enums](https://www.reddit.com/r/typescript/comments/ej51kk/generate_type_aliases_from_enums/)
- url: https://www.reddit.com/r/typescript/comments/ej51kk/generate_type_aliases_from_enums/
---
    enum Birds {
        eagle,
        cardinal,
        falcon
    }
    
    type BirdTypes = keyof typeof Birds; // 'eagle' | 'cardinal' | 'falcon'

For the longest time I’ve been duplicating efforts when it came to having both an enum and type alias pointing to the same data, just learned about `typeof` which makes this a lot easier.
## [5][Looking for a library to generate images](https://www.reddit.com/r/typescript/comments/ej8y1k/looking_for_a_library_to_generate_images/)
- url: https://www.reddit.com/r/typescript/comments/ej8y1k/looking_for_a_library_to_generate_images/
---
Hi! I'm hoping this is the right place I can ask this question. I really like types and have gotten used to using typescript at work. I'm looking to work on a personal project where I generate high quality images programmatically (which in this case means pixel by pixel). Maybe if I actually get this off the ground it'll become clear what I mean.

Anyway, I'm seeing lots of libraries/npm packages online that talk about image manipulation (resizing, overlaying with text, etc) and ones that care about doing this on the web (like canvas). ~~If any of these are good for my use, it seems like it'd be at best by accident.~~ (Edit: To be clear, I don't think these are what I want but they are all I've been able to find.) Anyway, I'd like to be able to generate images that are of high enough quality to print. The point is looking the best over speed. Does anyone know of such a library?

Good typing would be a huge plus in my book.

Thanks so much for any advice!
## [6][Non-null assertion operator vs. Optional chaining](https://www.reddit.com/r/typescript/comments/ej924r/nonnull_assertion_operator_vs_optional_chaining/)
- url: https://www.reddit.com/r/typescript/comments/ej924r/nonnull_assertion_operator_vs_optional_chaining/
---
I was just wondering, is there any difference between the two when you're trying to access properties that may possibly be null or undefined? It seems to me optional chaining is the safer one of the two, though I've yet to use it a lot. I defaulted to using non-null assertion after dealing with async data in GraphQL and only just learned about optional chaining for accessing properties recently.
## [7][Are there any naming conventions for modules?](https://www.reddit.com/r/typescript/comments/ej6kip/are_there_any_naming_conventions_for_modules/)
- url: https://www.reddit.com/r/typescript/comments/ej6kip/are_there_any_naming_conventions_for_modules/
---
I have a module with a class in it.  The class is called `Guard` and I called the module `GuardClass.`

I'm less than satisfied with this.  Does anyone have any suggestions for naming conventions for modules?
## [8][Take control of unexpected data at runtime with TypeScript](https://www.reddit.com/r/typescript/comments/eiyigo/take_control_of_unexpected_data_at_runtime_with/)
- url: https://matiasklemola.com/typescript-runtime-safety
---

## [9][Events in vscode extension](https://www.reddit.com/r/typescript/comments/ej41v5/events_in_vscode_extension/)
- url: https://www.reddit.com/r/typescript/comments/ej41v5/events_in_vscode_extension/
---
Hello everybody,

I have a question which is specific to vscode API and I am not sure where to ask for advice.

Specifically - I need to emit event in my extension in one window, so another window would get that event. As a simple example - it might be a scrolling event. When one document got scrolled, extension creates an event, so the same extension in another window gets notified and scrolls the document correspondingly.

What subreddit would be a good place to ask for help with it? I saw some vscode subreddits, but those do not look like dev communities.
## [10][This is probably a dumb question, but can I compile js files as if it were typescript if I have type defs?](https://www.reddit.com/r/typescript/comments/eiwken/this_is_probably_a_dumb_question_but_can_i/)
- url: https://www.reddit.com/r/typescript/comments/eiwken/this_is_probably_a_dumb_question_but_can_i/
---
So basically, I want to bundle and compile my projects from typescript using tsickle and closure compiler, however some libraries I'm using are not written in typescript, but javascript. They do have type defs. I could just compile them separately with simple\_optimization level if I have to, but would prefer to compile and bundle everything together with the closure compiler using advanced\_optimization.
## [11][Anyone use typescript with prop-types to support JS and TS consumers of a component library?](https://www.reddit.com/r/typescript/comments/eizuu9/anyone_use_typescript_with_proptypes_to_support/)
- url: https://www.reddit.com/r/typescript/comments/eizuu9/anyone_use_typescript_with_proptypes_to_support/
---
Our component library has JS and TS consumers so we’d need to support both. Most of the examples I see using the InferProps symbol are a little simple/contrived. I was wondering if anyone has experience using it on a larger component library. 

I’ve also seen the Babel plugin, and would be curious to hear any opinions on it as well.
