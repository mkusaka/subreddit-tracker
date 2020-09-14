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
## [2][Effective Typescript | A Non-Technical Guide](https://www.reddit.com/r/typescript/comments/is5d0m/effective_typescript_a_nontechnical_guide/)
- url: https://matthewmullin.io/effective-typescript/
---

## [3][Reactive formControl default dependant on other formControl](https://www.reddit.com/r/typescript/comments/isied7/reactive_formcontrol_default_dependant_on_other/)
- url: https://www.reddit.com/r/typescript/comments/isied7/reactive_formcontrol_default_dependant_on_other/
---
Hello,

I have two dropdown menus. The first control:

`&lt;select formControlName="colorName"&gt;&lt;option *ngFor="let color of sortedColor" value="color.value"&gt;{{`[`color.name`](https://color.name)`}}&lt;/option&gt;&lt;/select&gt;`

The second control:

`&lt;select formControlName="colorDescription"&gt;`

//I am supposed to do a mutable pre-selection of the description here, you can change the description, but this dropdown is suposed to match the color name from the previous dropdown at first`&lt;option *ngFor="let description of sortedColor" value="color.value"&gt;{{color.description }}&lt;/option&gt;&lt;/select&gt;`

I am basically accessing the same Enum and need to match up and I do so by subscribing to the valueChanges of my formControl "colorName", but I don't know how to retrieve the matching color Enum from the formControl and comparing it to my list "sortedColor" and creating the default value for my second formControl

&amp;#x200B;

Thanks for taking the time to read
## [4][How do I set minimum and maximum elements for an array?](https://www.reddit.com/r/typescript/comments/isi7dn/how_do_i_set_minimum_and_maximum_elements_for_an/)
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
## [5][Any solid barebone React, Redux, axios application?](https://www.reddit.com/r/typescript/comments/is8te2/any_solid_barebone_react_redux_axios_application/)
- url: https://www.reddit.com/r/typescript/comments/is8te2/any_solid_barebone_react_redux_axios_application/
---
Any solid barebone React, Typescript, Redux, axios application? I am looking for a skeleton that's easy to understand so I can get functional as quickly as possible. Is there anything good you could recommend?
## [6][Total newbie wanting to learn](https://www.reddit.com/r/typescript/comments/is7dnb/total_newbie_wanting_to_learn/)
- url: https://www.reddit.com/r/typescript/comments/is7dnb/total_newbie_wanting_to_learn/
---
I am a Linux admin with my coding background being bash and yaml gor Ansible playbooks. The team I just joined is transitioning to Amazon's Cloud Developer Kit for Infrastructure as Code using Typescript.

So where is a good source to learn TS from the ground up that does not assume a JS background?
## [7][How did you get really good at typescript?](https://www.reddit.com/r/typescript/comments/is6i00/how_did_you_get_really_good_at_typescript/)
- url: https://www.reddit.com/r/typescript/comments/is6i00/how_did_you_get_really_good_at_typescript/
---
I’ve used typescript for about two years every day at work. I am very comfortable with 99% of typescript, but sometimes I look at package typing (e.g. vue 3) and it is clear that some people are at a completely different level. 

When it comes to generics I feel I can get by, but not write anything too complicates. Is there a good resource (book, course, video, site) to help with this? Any good sites to practice it on?
## [8][Pipe and compose](https://www.reddit.com/r/typescript/comments/irxcjg/pipe_and_compose/)
- url: https://www.reddit.com/r/typescript/comments/irxcjg/pipe_and_compose/
---
How would you type a compose function , pipe function in Typescript  ?

&amp;#x200B;

`const compose = (...fns) =&gt; (x) =&gt; fns.reduceRight((acc, fn) =&gt; fn(acc), x);`

`const pipe = (...fns) =&gt; (x) =&gt; fns.reduce((acc, fn) =&gt; fn(acc), x);`
## [9][Good resources for learning more deep metaprogramming of typescript?](https://www.reddit.com/r/typescript/comments/irvlhf/good_resources_for_learning_more_deep/)
- url: https://www.reddit.com/r/typescript/comments/irvlhf/good_resources_for_learning_more_deep/
---
I want to improve my typescript skills by learning more complex, advanced techniques with typescript to develop better APIs etc for parts of the app. Any good resources such as sites or blogs?

Edit: I'll read the Typescript handbook. Surprisingly haven't read the whole thing all this time...
## [10][Typescript type definition for role property with Identity Server 4 token auth](https://www.reddit.com/r/typescript/comments/irwt8d/typescript_type_definition_for_role_property_with/)
- url: https://www.reddit.com/r/typescript/comments/irwt8d/typescript_type_definition_for_role_property_with/
---
I'm wondering if anyone's encountered this before....but using the token auth implementation of Identity Server 4, the role property can either be a string or an array of strings, coming back, depending on whether or not a user has more than one role:

https://github.com/IdentityServer/IdentityServer4/issues/2468 

I'm wondering if anyone's encountered this and come up with a good workaround, rather than my dirty `Array.isArray()` check on the role property?
## [11][From Rust to TypeScript](https://www.reddit.com/r/typescript/comments/irhluf/from_rust_to_typescript/)
- url: https://valand.dev/blog/post/from-rust-to-typescript
---

