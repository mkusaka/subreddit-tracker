# typescript
## [1][Who's hiring Typescript developers - October](https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/)
- url: https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/
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
## [2][How to get intuitively comfortable with TypeScript?](https://www.reddit.com/r/typescript/comments/j8vwzf/how_to_get_intuitively_comfortable_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/j8vwzf/how_to_get_intuitively_comfortable_with_typescript/
---
I've seen the docs and have played around with TS in toy projects, in an attempt to add it to my repertoire. I come from a background where my projects didn't use TS at all. How does one develop an intuitive sense of using TS for production?

While I realize the significance, I find it hard to bridge the gap between my amateur knowledge and being production-ready as the rest of my skillset.

I keep having to double-check what to type something while I'm implementing, sometimes resorting to online searches especially for imported libraries. There's also the pain of not knowing for sure if I typed it the right way. Is it normal for it to be time-consuming at the start?
## [3][Vue and TypeScript](https://www.reddit.com/r/typescript/comments/j91mz2/vue_and_typescript/)
- url: https://www.reddit.com/r/typescript/comments/j91mz2/vue_and_typescript/
---
I am totally new to typescript i just wanted to do some side project to learn a bit more of it with vue. But i cant even configure my first app i am having problems with vue imports. I have made shim.d.ts but nothing is fixing the import error on "import App from './App.vue'" ....

Is there any straightforward tutorial how to set up my vue app with typescript , or if anyone could give me suggestion how to fix this and avoid these kind of errors in future id be really thankful.
## [4][Generic type for type conversion](https://www.reddit.com/r/typescript/comments/j93t98/generic_type_for_type_conversion/)
- url: https://www.reddit.com/r/typescript/comments/j93t98/generic_type_for_type_conversion/
---
Hi,

I have some custom types like

    type MyType = {
        id: number;
        title: string;
        checked: boolean;
    }

I am trying to make a generic type (if possible) that will convert any boolean property in the type to binary (0|1).

&amp;#x200B;

I was thinking of something like the following

    type BoToBiFormat = {[x:string]:any}
    type BoolToBinary&lt;T extends BoToBiFormat&gt; = {
        [x in keyof T]: T[x] === boolean ? 0|1 : T[x];
    }

but that obviously didn't work. Is there a way to make such a type?
## [5][How to view JSDoc (or at least signature) for untyped JS dependencies? Tried both WebStrorm and VisualCode](https://www.reddit.com/r/typescript/comments/j8hj0y/how_to_view_jsdoc_or_at_least_signature_for/)
- url: https://i.redd.it/2q0surn3j8s51.png
---

## [6][Bug in type definitions for HOF?](https://www.reddit.com/r/typescript/comments/j8xnfz/bug_in_type_definitions_for_hof/)
- url: https://www.reddit.com/r/typescript/comments/j8xnfz/bug_in_type_definitions_for_hof/
---
Hello, I'm currently writing a little Axios-like wrapper around the Fetch API for the size benefits. To do this, I enlist a higher order function that differentiates body-less methods like GET, DELETE, OPTIONS, etc. from POST, PUT, etc. 

Here's my code (roughly)


```
type BodylessMethods = 'get' | 'delete' | 'options' | 'head';
type BodyMethods = 'post' | 'put' | 'patch';
const isOfBodylessMethod = (method: string): method is BodylessMethods =&gt; ['get', 'delete', 'options', 'head'].includes(method);

...

private request = (method: BodyMethods | BodylessMethods) =&gt; {
  return isOfBodylessMethod(method)
    ? (resource: string, params?: Record&lt;string, unknown&gt;) =&gt; {
        ...
      }
    : (resource: string, data?: Record&lt;string, unknown&gt;, params?: Record&lt;string, unknown&gt;) =&gt; {
        ...
      };
};

public get = this.request('get');
```

The problem I'm running into is that, for some reason the bodyless methods (`.get()`, `.delete()`, etc) throw an error if I make `data` a non-optional parameter in that second function, and I don't quite know why. I'm given `TS2554: Expected 2-3 arguments, but got 1.` Why would 2 arguments be expected? The first function has 1 required, 1 optional, the second 2 required and 1 optional. Surely 1-3 arguments should be expected?

Now, this can easily be fixed by keeping `data` optional, and should likely be solved that way as the spec doesn't require a body with POST, PUT, etc., but I'd like to know why this issue is popping up. I don't understand where it is coming from.
## [7][Best TypeScript video tutorial for beginners or intermediate (free or paid) in 2020?](https://www.reddit.com/r/typescript/comments/j8qe9c/best_typescript_video_tutorial_for_beginners_or/)
- url: https://www.reddit.com/r/typescript/comments/j8qe9c/best_typescript_video_tutorial_for_beginners_or/
---

## [8][Self casting from json?](https://www.reddit.com/r/typescript/comments/j8vjk1/self_casting_from_json/)
- url: https://www.reddit.com/r/typescript/comments/j8vjk1/self_casting_from_json/
---
I have a function that accepts a class object. If I pass a json inside, typescript seems to automatically cast the json object into the object declared in the parameter. Is this reliable? Or should I map it and pass it in?
## [9][Typescript project starter: npm-ready, plus testing and CLI templates](https://www.reddit.com/r/typescript/comments/j88pwf/typescript_project_starter_npmready_plus_testing/)
- url: https://github.com/bscotch/typescript-template
---

## [10][Jeff and Dom Make A Game - a coworker and his friend have been making a multiplayer game in TypeScript the past 5 months and streaming the whole thing](https://www.reddit.com/r/typescript/comments/j8b8nv/jeff_and_dom_make_a_game_a_coworker_and_his/)
- url: https://youtu.be/z8e_cUsJyxM
---

## [11][If interviewing at a JS-only place, would you say you like Typescript because it makes your code more type-safe?](https://www.reddit.com/r/typescript/comments/j7x355/if_interviewing_at_a_jsonly_place_would_you_say/)
- url: https://www.reddit.com/r/typescript/comments/j7x355/if_interviewing_at_a_jsonly_place_would_you_say/
---
Something like 75% of all errors are TypeErrors and all my portfolio items showcase mid level Typescript even though by experience I'm a junior. It's the one technology I feel strongly about because in the JS days I spent 30% of my time building the project, 70% of the time figuring out where the interpreter's dynamic type inference is throwing an error.

Do JS only companies take offense if you start raving about how much TS has reinforced your code and sense of clarity on what is being returned where? Should I downplay it if I know they're not big on TS? 

Intuition says yes especially given my lack of corporate experience (3 years learning/practicing at home though). PS: I know the best answer is to find a company that has TS but in these times I can't be picky.
