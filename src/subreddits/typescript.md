# typescript
## [1][Who's hiring Typescript developers - May](https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/)
- url: https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

Posting top level comments that aren't job postings, [that's a paddlin](https://i.imgur.com/FxMKfnY.jpg)

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][I wish TypeScript to support publishing npm module with "ts" files](https://www.reddit.com/r/typescript/comments/ggxwn9/i_wish_typescript_to_support_publishing_npm/)
- url: https://github.com/microsoft/TypeScript/issues/38452
---

## [3][[Showoff Saturday] PWA Store written in TypeScript](https://www.reddit.com/r/typescript/comments/ggnrby/showoff_saturday_pwa_store_written_in_typescript/)
- url: https://v.redd.it/sav4v1jh1tx41
---

## [4][Confused what to learn](https://www.reddit.com/r/typescript/comments/gh176y/confused_what_to_learn/)
- url: https://www.reddit.com/r/typescript/comments/gh176y/confused_what_to_learn/
---
Hello guys, i am really interested in front end development and i am a beginner,  i am really confused how to start learning and what to learn first 
Can you guys help me with that ?
Thanks
## [5][How should I type a function, which takes a function as an argument and returns a new function with the identical type.](https://www.reddit.com/r/typescript/comments/ggy2q6/how_should_i_type_a_function_which_takes_a/)
- url: https://www.reddit.com/r/typescript/comments/ggy2q6/how_should_i_type_a_function_which_takes_a/
---
Hey  there, I wonder if some typescript wiz can help me with this. I'm  writing a React hook basically which takes a function as its first  argument:

`type Callback&lt;T&gt; = (args: T) =&gt; Promise&lt;void&gt;`

It then returns an array containing an object, and a new function which takes the same args

`type ReturnType&lt;T&gt; = [State, (args: T) =&gt; Promise&lt;void&gt;]`

I'm then combining them for my main function definition:

`function myFunction&lt;T&gt;(cb: Callback&lt;T&gt;): ReturnType&lt;T&gt; { ... }`

Now,  this works fine for functions which take arguments. However I also want  it to work for functions which take no arguments, i.e.

`() =&gt; Promise&lt;void&gt;`

Currently it returns me the type:

`(args: unknown) =&gt; Promise&lt;void&gt;`

which doesnt match. Any help would be amazing!! &lt;3
## [6][Prepare your Angular interview with this quick and complete outline](https://www.reddit.com/r/typescript/comments/ggecix/prepare_your_angular_interview_with_this_quick/)
- url: https://medium.com//prepare-your-angular-interview-with-this-quick-and-complete-outline-1e9b5d761166?source=friends_link&amp;sk=8a4d78eadeeae8aa9b10a0f9e7c47f29
---

## [7][Developing polymorphism feature for TypeORM](https://www.reddit.com/r/typescript/comments/ggpj3y/developing_polymorphism_feature_for_typeorm/)
- url: https://www.reddit.com/r/typescript/comments/ggpj3y/developing_polymorphism_feature_for_typeorm/
---
Any one interested to work on developing polymorphism feature for TypeORM?

[https://github.com/pirate-lp/typeorm-polymorphism](https://github.com/pirate-lp/typeorm-polymorphism)

**Background**

I understand that we can hack our way around it and that there is no plan from original maintainer to develop the feature, but I really miss this functionality in the Node.JS community and as far as I searched, there is no functional library that supports this design. Even bookshelf.js runs into a lot of bugs and issues when it comes to working with polymorphic relationships!
## [8][Code Review Please: Modeling profit/loss of a time series of financial position state changes](https://www.reddit.com/r/typescript/comments/ggi6qs/code_review_please_modeling_profitloss_of_a_time/)
- url: https://codereview.stackexchange.com/questions/242009/model-profit-loss-of-a-time-series-of-financial-position-state-changes
---

## [9][Strongly-typed "array of function pointers"](https://www.reddit.com/r/typescript/comments/ggq60f/stronglytyped_array_of_function_pointers/)
- url: https://www.reddit.com/r/typescript/comments/ggq60f/stronglytyped_array_of_function_pointers/
---
**UPDATED: Solved using intersection instead of union**

I'm building the write-side of a CQRS from scratch for the purpose of learning, and what I am trying to type is the command handlers. Just to be clear, I really don't want an array of function pointers, I'm just speaking in a comfortable, an ancient tongue.  In JS, I would

    const handlers = {
      FooCreated: (payload) =&gt; { ...foo created handler },
      BarAssigned: (payload) =&gt; { ..bar assigned handler },
    }
    ...
    const handler = handlers[body.command];

I am seeking a TS solution with the fewest keystroke touches on the text "FooCreated", et. al.  I'll have hundreds of commands (and events, which will probably face the same issue).  What I have so far:

    type CommandMeta = {
      userId: string;
    };
    type Command =
      | {
          CreateFoo: {
            meta: CommandMeta;
            payload: {
              id: string;
              name: string;
              baz?: string;
            };
          };
        }
      | {
          AssignBar: {
            meta: CommandMeta;
            payload: { 
              id: string; 
              fizzbang: string; 
            };
          };
        };
    type CommandHandlers = {
      [C in keyof Command]: (meta: CommandMeta, payload: Command[C]) =&gt; Promise&lt;any&gt;
    };
    const handlers: CommandHandlers = {
      CreateFoo: (meta, payload) =&gt; {
        // meta &amp; payload is "any", I would like it to be typeof CommandMeta &amp; payload type for CreateFoo
      },
      // No complaint about missing "AssignBar" handler
    };

As noted in the commends in the command handler implementation, I am getting "any's' for the meta and payload parameters.  Furthermore, I am not getting any error for missing the "AssignBar" handler, and I can also put any old garbage name in there.

I'm not married to the above approach, so any quick boosts would be appreciated.  I'm not a big fan of spreading types across hundreds of files and typing (via keyboard, not TS) their names in lots and lots of places.  So I put a premium on an implementation that is less clean than most prefer.
## [10][How TS prevents your JS errors](https://www.reddit.com/r/typescript/comments/ggkuyj/how_ts_prevents_your_js_errors/)
- url: https://www.youtube.com/watch?v=dK2cnXp-xkc
---

## [11][How can I limit the strings used to index into an object?](https://www.reddit.com/r/typescript/comments/gg1r6z/how_can_i_limit_the_strings_used_to_index_into_an/)
- url: https://www.reddit.com/r/typescript/comments/gg1r6z/how_can_i_limit_the_strings_used_to_index_into_an/
---
Say I have a Movie type:

    interface Movie {
      title: string;
      year: number;
      studio: string;
      actors: string[];
      writers: string[];
    }
    
    const m: Movie = {
      title: "Silence of the Lambs",
      year: 1991,
      studio: "Strong Heart",
      actors: ["Jodie Foster", "Anthony Hopkins", "Ted Levine"],
      writers: ["Ted Tally"],
    };

and there are a few properties I want to iterate over, like

    for (let jobName of ["actors", "writers"]) {
      m[jobName].forEach(console.log);
    }

I can't do this unless I add  `[index: string]: string[] | string | number` to my interface and then assert `m[jobName] as string`, but then I can use *any* string to index in, and I don't want that.

In this case I could just do

    m.actors.forEach(console.log);
    m.writers.forEach(console.log);

but what if I'm not iterating over just two properties, what if I'm iterating over 50 and don't want pages and pages of manually specified iterations? This isn't exactly a real problem I'm having but I'd like to know what the correct response to this problem would be, because it's definitely something where I feel like I don't know the smart path.
