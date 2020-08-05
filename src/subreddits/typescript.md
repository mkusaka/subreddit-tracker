# typescript
## [1][Who's hiring Typescript developers - August](https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/)
- url: https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/
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
## [2][Is it possible to define a type for a variable that does not allow undefined value ?](https://www.reddit.com/r/typescript/comments/i4565u/is_it_possible_to_define_a_type_for_a_variable/)
- url: https://www.reddit.com/r/typescript/comments/i4565u/is_it_possible_to_define_a_type_for_a_variable/
---
The question in the title .
## [3][Reducing TypeScript boilerplate in Next.js pages](https://www.reddit.com/r/typescript/comments/i3vw7j/reducing_typescript_boilerplate_in_nextjs_pages/)
- url: https://github.com/vriad/nxtk
---

## [4][is keyword, versus as](https://www.reddit.com/r/typescript/comments/i3slde/is_keyword_versus_as/)
- url: https://www.reddit.com/r/typescript/comments/i3slde/is_keyword_versus_as/
---
Normally I use the `as` keyword to assert a type to help the compiler read the code correctly. Sometimes I also use `as` as a type annotation (for others who read my code including myself).

There is also an `is` keyword. It appears to be another assertion. I have trouble telling what the difference is exactly; can someone clarify the distinction? I don't remember the last time I've used `is`.
## [5][Parenthesis Balancing using Monoids in TypeScript](https://www.reddit.com/r/typescript/comments/i3puyh/parenthesis_balancing_using_monoids_in_typescript/)
- url: https://medium.com//parenthesis-balancing-using-monoids-in-typescript-56e5c994e497?source=friends_link&amp;sk=3bc26cc47a80f3bf40c2da0f6221a1c6
---

## [6][Do popular API have public TypeScript interfaces for their data?](https://www.reddit.com/r/typescript/comments/i3mzur/do_popular_api_have_public_typescript_interfaces/)
- url: https://www.reddit.com/r/typescript/comments/i3mzur/do_popular_api_have_public_typescript_interfaces/
---
I'm making an app with the PokeApi, and I realize that if I want intellisense and type checking for my API responses, i'd have to create a whole massive interface. Do the makers of API's typically have a TS interface for their API, or is there anywhere else that might already? Thanks.
## [7][Aquarium - A minimalistic blockchain implementation](https://www.reddit.com/r/typescript/comments/i3k75l/aquarium_a_minimalistic_blockchain_implementation/)
- url: https://github.com/micheleriva/Aquarium
---

## [8][Express middleware - What is the best way to do undefined checking?](https://www.reddit.com/r/typescript/comments/i3m94q/express_middleware_what_is_the_best_way_to_do/)
- url: https://www.reddit.com/r/typescript/comments/i3m94q/express_middleware_what_is_the_best_way_to_do/
---
Might be a dumb question, but I have an express middleware that checks for the existence of various headers and sends an error response if they don't exist. Then in my route function I try to access some of those headers, but I am met with various typing problems where I would need to check for the existence of the headers again to make sure I'm not passing an undefined variable to a function.

So I guess my question is am I going to need to duplicate my header checks again, or is there any way to make typescript have scope of the middleware used and see that the checks have already been done? Or is having a middleware do this a silly thing to do in the first place?
## [9][Optional Generic Type of Class for Constructor Arguments](https://www.reddit.com/r/typescript/comments/i3pgw5/optional_generic_type_of_class_for_constructor/)
- url: https://www.reddit.com/r/typescript/comments/i3pgw5/optional_generic_type_of_class_for_constructor/
---
I'm trying to create the following base class:

```typescript
export class Event&lt;T = any&gt; {
  public data: T;

  constructor(data: T) {
    if (data) {
      this.data = data;
    }
  }
}
```

The problem is I may have events that when constructed I need an argument:
`new PostAddedEvent({ postId: string });`

Or events where I don't need them:
`new KernelShutdown()`

If I'm putting my constructor(data?: T), then for events that extend event like:
```typescript
class PostAddedEvent extends Event&lt;{ postId: string}&gt; {}
```

I can instantiate it without arguments (new PostAddedEvent()), which is something I don't want.

Is there an elegant way to have best of both worlds? Like a sort-of advanced type that makes the argument of the constructor optional?
## [10][Announcing the new TypeScript Website](https://www.reddit.com/r/typescript/comments/i3ojd3/announcing_the_new_typescript_website/)
- url: https://devblogs.microsoft.com/typescript/announcing-the-new-typescript-website/
---

## [11][Want to Interview Someone about Typescript for my Podcast](https://www.reddit.com/r/typescript/comments/i3qyif/want_to_interview_someone_about_typescript_for_my/)
- url: https://www.reddit.com/r/typescript/comments/i3qyif/want_to_interview_someone_about_typescript_for_my/
---
If you are an experienced Typescript programmer, I would love to have you appear on my podcast, Contentful { Creators } to chat about Typescript.  It would be sort of an introduction to Typescript.  Please have a look at my podcast and if you are interested, please reach out to me at [marcelo@themarcelolewin.com](mailto:marcelo@themarcelolewin.com)

[https://anchor.com/contentful-creators/](https://anchor.com/contentful-creators/)

Thanks!

\- Marcelo
