# typescript
## [1][Who's hiring Typescript developers - September](https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/)
- url: https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/
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
## [2][How to add custom types to the TypeScript project with examples](https://www.reddit.com/r/typescript/comments/it5n5p/how_to_add_custom_types_to_the_typescript_project/)
- url: https://drag13.io/posts/custom-typings/index.html
---

## [3][I need help with setting dynamic types](https://www.reddit.com/r/typescript/comments/isq6ed/i_need_help_with_setting_dynamic_types/)
- url: https://www.reddit.com/r/typescript/comments/isq6ed/i_need_help_with_setting_dynamic_types/
---
Hey everyone, I'm trying to type a map like class and I'm not sure if it's possible to do it the way I want it to work. Essentially what I want is this:

```
factory.define('user', UserModel, { /* ... default values for the UserModel */ } )

factory.create('user', { /* ... UserModel fields to override */ }): UserModel
```

In the `define` method you specify a key string, a type it matches to and some default values. Then in `create` TS understand that the `'user'` key means that the second parameter must be of type `UserModel`, and it returns a `UserModel`.

I don't want to have to define all the keys (`'user'` in this case) as a type ahead of time, I want to define the methods such that typescript understands that the `define` parameters are setting keys and values. Anyone know if this is possible, and if so, how?

Many thanks for your help!

Edit 1: updated to a better explaining example.
## [4][Small SaaS looking to add Developer Cofounder](https://www.reddit.com/r/typescript/comments/ist6rr/small_saas_looking_to_add_developer_cofounder/)
- url: https://www.reddit.com/r/typescript/comments/ist6rr/small_saas_looking_to_add_developer_cofounder/
---
Hi group. This is the small business we're building one customer at a time - [https://podsquad.app/](https://podsquad.app/). We're looking to add another developer cofounder to the team. Frontend is React Native and backend is an Express JS server running on Heroku using Parse for storage. Please reach out of interested!
## [5][Effective Typescript | A Non-Technical Guide](https://www.reddit.com/r/typescript/comments/is5d0m/effective_typescript_a_nontechnical_guide/)
- url: https://matthewmullin.io/effective-typescript/
---

## [6][How do I set minimum and maximum elements for an array?](https://www.reddit.com/r/typescript/comments/isi7dn/how_do_i_set_minimum_and_maximum_elements_for_an/)
- url: https://www.reddit.com/r/typescript/comments/isi7dn/how_do_i_set_minimum_and_maximum_elements_for_an/
---
basically, I want to limit the length of the `coordinates` parameter in the `getDirection` function.

how do I force the caller to pass an array of `LatLng` with a minimum length of 2 and a maximum elements of 25?

  
please take a look at my codes below

&amp;#x200B;

    interface LatLng {
     lat: number
     lng: number
    }
    
    const getDirection = async (coordinates: LatLng[]) {
    
    }

Iv'e seen tuple examples but it looks like it's only for setting a minimum or fixed length of elements.
## [7][Total newbie wanting to learn](https://www.reddit.com/r/typescript/comments/is7dnb/total_newbie_wanting_to_learn/)
- url: https://www.reddit.com/r/typescript/comments/is7dnb/total_newbie_wanting_to_learn/
---
I am a Linux admin with my coding background being bash and yaml gor Ansible playbooks. The team I just joined is transitioning to Amazon's Cloud Developer Kit for Infrastructure as Code using Typescript.

So where is a good source to learn TS from the ground up that does not assume a JS background?
## [8][How did you get really good at typescript?](https://www.reddit.com/r/typescript/comments/is6i00/how_did_you_get_really_good_at_typescript/)
- url: https://www.reddit.com/r/typescript/comments/is6i00/how_did_you_get_really_good_at_typescript/
---
I’ve used typescript for about two years every day at work. I am very comfortable with 99% of typescript, but sometimes I look at package typing (e.g. vue 3) and it is clear that some people are at a completely different level. 

When it comes to generics I feel I can get by, but not write anything too complicates. Is there a good resource (book, course, video, site) to help with this? Any good sites to practice it on?
## [9][Any solid barebone React, Redux, axios application?](https://www.reddit.com/r/typescript/comments/is8te2/any_solid_barebone_react_redux_axios_application/)
- url: https://www.reddit.com/r/typescript/comments/is8te2/any_solid_barebone_react_redux_axios_application/
---
Any solid barebone React, Typescript, Redux, axios application? I am looking for a skeleton that's easy to understand so I can get functional as quickly as possible. Is there anything good you could recommend?
## [10][Pipe and compose](https://www.reddit.com/r/typescript/comments/irxcjg/pipe_and_compose/)
- url: https://www.reddit.com/r/typescript/comments/irxcjg/pipe_and_compose/
---
How would you type a compose function , pipe function in Typescript  ?

&amp;#x200B;

`const compose = (...fns) =&gt; (x) =&gt; fns.reduceRight((acc, fn) =&gt; fn(acc), x);`

`const pipe = (...fns) =&gt; (x) =&gt; fns.reduce((acc, fn) =&gt; fn(acc), x);`
## [11][Good resources for learning more deep metaprogramming of typescript?](https://www.reddit.com/r/typescript/comments/irvlhf/good_resources_for_learning_more_deep/)
- url: https://www.reddit.com/r/typescript/comments/irvlhf/good_resources_for_learning_more_deep/
---
I want to improve my typescript skills by learning more complex, advanced techniques with typescript to develop better APIs etc for parts of the app. Any good resources such as sites or blogs?

Edit: I'll read the Typescript handbook. Surprisingly haven't read the whole thing all this time...
