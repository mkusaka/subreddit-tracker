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
## [2][Proper way to have @throws {FooBar} satisfy TSC and ESLINT?](https://www.reddit.com/r/typescript/comments/jh23jy/proper_way_to_have_throws_foobar_satisfy_tsc_and/)
- url: https://www.reddit.com/r/typescript/comments/jh23jy/proper_way_to_have_throws_foobar_satisfy_tsc_and/
---
Lets say I have the following in the comment block (along with the usual params, returns, etc) on a method on an interface,

 `@throws {FooBar} When you do the bad` ,

I get eslint complaining that  FooBar is undefined. The thing is this is an interface - I am not actually throwing FooBar here  - the implementation is separate.

If I import FooBar, eslint stops complaining - now I define FooBar. But now typescript is complaining - I am not actually using FooBar.

How do I get these two idiots to work together?
## [3][esfx - A collection of high quality packages written by a Senior SDE at Microsoft](https://www.reddit.com/r/typescript/comments/jgppge/esfx_a_collection_of_high_quality_packages/)
- url: https://www.reddit.com/r/typescript/comments/jgppge/esfx_a_collection_of_high_quality_packages/
---
I came across a TypeScript project called [esfx](https://github.com/esfx/esfx) a few weeks ago and only recently got the time to look closer at what it was. Brought to you by the developer behind the Metadata Reflection API ([`reflect-metadata`](https://github.com/rbuckton/reflect-metadata)) we all know and love, esfx contains a pretty significant assortment of useful constructs.

Treat yourself and take a look. Words cannot describe how extraordinary it is.
## [4][Is there any different on this two? var x = foo as any; var x: any = foo;](https://www.reddit.com/r/typescript/comments/jgi10h/is_there_any_different_on_this_two_var_x_foo_as/)
- url: https://www.reddit.com/r/typescript/comments/jgi10h/is_there_any_different_on_this_two_var_x_foo_as/
---

## [5][Property does not exist on type](https://www.reddit.com/r/typescript/comments/jgk0qu/property_does_not_exist_on_type/)
- url: https://www.reddit.com/r/typescript/comments/jgk0qu/property_does_not_exist_on_type/
---
Hi.

From my Axios request, my [res.data](https://res.data) looks like this: 

    {"count":1,"next":null,"previous":null,"results":[{"id":1,"name":"The Name","url":"http://somewhere"}]}

I'd like represent this response in Typescript, so I've done this:

    import axios, { AxiosResponse } from "axios";
    
    type Results = {
      results: Array&lt;{}&gt;;
    };
    
    type AxiosResponseData = {
      data: Results;
    };
    
    let res = &lt;AxiosResponse&gt; await axios.get(url, config);
    let data = &lt;type.AxiosResponseData&gt;res.data
    return data.results;

Typescript doesn't like the very last line, which this warning: `Property 'results' does not exist on type 'AxiosResponseData'`.

I'd really appreciate help figuring this out.  'AxiosResponseData'.
## [6][I need your advices](https://www.reddit.com/r/typescript/comments/jgkyug/i_need_your_advices/)
- url: https://www.reddit.com/r/typescript/comments/jgkyug/i_need_your_advices/
---
I've made a NodeJS library called Graceful Server months ago.

You can find it here : [https://github.com/gquittet/graceful-server](https://github.com/gquittet/graceful-server)

and here : [https://www.npmjs.com/package/@gquittet/graceful-server](https://www.npmjs.com/package/@gquittet/graceful-server)

&amp;#x200B;

What are its goals?

\- Accept HTTP connections on your API only when it's ready

\- Know precisely why your API crashed and shutdown

\- Close all HTTP connections and disconnect your API from all the data sources on shutdown (to avoid to keep opened ghost connections)

\- Give you a liveness and readiness endpoints (useful when you're using Kubernetes)

&amp;#x200B;

I want to know if the documentation is good.

Can you share idea on how to improve the documentation and features you want to have in this library?

Are you interested that I publish the roadmap?

&amp;#x200B;

This post is to thanks my 10k week downloads 💪

&amp;#x200B;

Let's make it evolve together ! 🚀
## [7][Which language for a blog website: Dart or TypeScript?](https://www.reddit.com/r/typescript/comments/jgmmg0/which_language_for_a_blog_website_dart_or/)
- url: https://www.reddit.com/r/typescript/comments/jgmmg0/which_language_for_a_blog_website_dart_or/
---
&gt; I know this is a biased source. But it's better than not asking...

&gt; [This question was originally posted][original] on r/dartlang.

[My blog][fanaro.com.br] is currently made with WordPress. But I've been wanting to revamp it with a more minimal design and pure code in the backend for quite a while now. The 2 reasons I've been avoiding this were: (1) my own experience as a developer and (2) the technologies I would use. My level of experience with HTML &amp; CSS is ok, but I have only minimal experience with JS and TS. While, with Dart, I think I'm already pretty advanced &amp;mdash; see [this browser extension][youtube_kbd_nav] I've created in the past month.

So I come here to ask if anyone would help me deciding if I could/should go with Dart for revamping my blog website. Is it on par with TS for this purpose? Or should I still go for TS?

From my perspective, most of the website's work will be either content (HTML and text) or the OO design, so, in the end, it wouldn't be too hard to simply rewrite it into another language anyway, though that would be super annoying &amp;mdash; and who knows if I'm gonna accidentally design something that uses a language-specific feature.

There is also Flutter to take into consideration. Even though I think HTML &amp; CSS are more appropriate for this type of static content, maybe Flutter could be the tipping point of the argument for Dart.

Since a blog website is usually very minimal anyway I could use it as an exercise for evaluating the basic differences between Dart and TS. But I don't know if I have the time or energy for that right now.


[fanaro.com.br]: https://fanaro.com.br
[original]: https://www.reddit.com/r/dartlang/comments/jfckso/blog_website_dart_or_typescript/
[youtube_kbd_nav]: https://github.com/FanaroEngineering/youtube_kbd_nav
## [8][Noob Question: jsconfig or tsconfig?](https://www.reddit.com/r/typescript/comments/jg1v1x/noob_question_jsconfig_or_tsconfig/)
- url: https://www.reddit.com/r/typescript/comments/jg1v1x/noob_question_jsconfig_or_tsconfig/
---
Hi, just getting started with typescript. Should I be using a jsconfig or tsconfig, or does it even matter?
## [9][How can I get the keys from a Map.keys() iterator as a type?](https://www.reddit.com/r/typescript/comments/jfzkor/how_can_i_get_the_keys_from_a_mapkeys_iterator_as/)
- url: https://www.reddit.com/r/typescript/comments/jfzkor/how_can_i_get_the_keys_from_a_mapkeys_iterator_as/
---
```
const entities = new Map&lt;string, []&gt;([
  ["hero", []],
  ["zombies", []],
  ["bullets", []],
  ["text", []],
]);

const mapIter = entities.keys();

for (const keys of mapIter) {
  console.log(keys); // hero, zombies, bullets text
}
```
I want to generate `type EntityKeys = "hero" | "zombies" | "bullets" | "text";` programatically. 

I want to use `entities.get("hero")` and have `"hero"` type-checked, which won't work with `Map&lt;string, []&gt;`. I can use `new Map&lt;EntityKeys, Entity[]&gt;` using `type EntityKeys =...` but then it is not dynamic and I have to create `type EntityKeys` manually.

Cheers!
## [10][Elegant way of creating multiple types from a single generic](https://www.reddit.com/r/typescript/comments/jg07c8/elegant_way_of_creating_multiple_types_from_a/)
- url: https://www.reddit.com/r/typescript/comments/jg07c8/elegant_way_of_creating_multiple_types_from_a/
---
Hello!

I'm tring to define some kind of "redux store" boilerplate using only ContextAPI

Here's what i have

```
import { Dispatch, SetStateAction, ReactNode } from 'react'

declare namespace ContextStore {

  export type StateAgent&lt;T&gt; = [T, Dispatch&lt;SetStateAction&lt;T&gt;&gt;]

  export type SetNewState&lt;T&gt; = (newValue: T) =&gt; void

  export type ContextHook&lt;T&gt; = [T, SetNewState&lt;T&gt;]

  export type ContextProviderProps = {
    children: ReactNode
  }
}

export = ContextStore
export as namespace ContextStore
```

I imagine there is a better way of outputting these 4 different types by passing only 1 generic argument.

These "T"s are actually the same type.

For example:
I create a `initialState = { state: true }` and then i use it to create all of the types above instead of having to import each one of them and passing `initialState` as generic, like so:

```
StateAgent&lt;typeof initialState&gt;
SetNewState&lt;typeof initialState&gt;
ContextHook&lt;typeof initialState&gt;
...
``` 

Maybe i can build a function that make these types avaiable with the generic built-in?

Thanks in advance
## [11][TypeScript JSX transform vs Babel JSX transform - pros and cons](https://www.reddit.com/r/typescript/comments/jffysy/typescript_jsx_transform_vs_babel_jsx_transform/)
- url: https://www.reddit.com/r/typescript/comments/jffysy/typescript_jsx_transform_vs_babel_jsx_transform/
---
I'm building a project with both TypeScript and Babel (and React) and I see that both TypeScript and Babel support the same JSX transform step. Either TypeScript can emit transformed JSX in JS code directly, or it can preserve the JSX in JSX files so Babel can handle it.

It looks like each compiler has their own spec compliant implementation of this transformation. I cannot find any details as to which one is better to use and why. There must be some differences (speed, reliability, spec compliance, support, bug fixes, type checking, etc) but I cannot seem to identify any.

Given the little information I have, it seems like Babel is the "canonical" JSX transformer as React specifically works with Babel to implement the spec. From the TS JSX PR I see that they are just trying to copy what Babel implemented. This leads me to believe that I should use the TS JSX preserve setting and let Babel handle it as it'll likely be more spec complainant and more stable.

Does anyone have any additional information to add here or know of anything which I should consider in making this decision? Thanks!

----

Edit: Follow up. After tons of very helpful comments I figured I'd share what I decided on. I am now using `babel` for all TypeScript compiling needs. `@babel/preset-env`, `@babel/preset-typescript` and `@babel/preset-react` specifically. To get builds to fail due to type errors and to see errors in the console during development, I'm using [fork-ts-checker-webpack-plugin](https://github.com/TypeStrong/fork-ts-checker-webpack-plugin). I now no longer need `tsc` for production or development at all and am only using it for debugging if I purely want to run a command to quickly see all type errors in the project.
