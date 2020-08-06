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
## [2][Strict null checks . Are they really worth it ?](https://www.reddit.com/r/typescript/comments/i4q827/strict_null_checks_are_they_really_worth_it/)
- url: https://www.reddit.com/r/typescript/comments/i4q827/strict_null_checks_are_they_really_worth_it/
---
Since I enabled strict null check option in tsconfig I have faced some linting errors that make me thing that strict null checks are really not worth it .

For me to satisfy the linter I have to spent more time to make my code less readable by using `as` a lot of times . Sometimes there so many `as` that I have to use , that I even think adding an unnecessary if clause so all the linting will get satisfied without the need of `as` .

But then again I have almost a single day experience with strict null checks .

What is your experience with strict null checks ? Are they more harm than good or the opposite ?

Edit : [Here](https://www.reddit.com/r/typescript/comments/i4q827/strict_null_checks_are_they_really_worth_it/g0jvh1z?utm_source=share&amp;utm_medium=web2x) is something that I was missing and will make me not need `as` and `if` .
## [3][Can a class reference itself? If so, should it?](https://www.reddit.com/r/typescript/comments/i4e6do/can_a_class_reference_itself_if_so_should_it/)
- url: https://www.reddit.com/r/typescript/comments/i4e6do/can_a_class_reference_itself_if_so_should_it/
---
Title says class, but I mean type.

The context is a model of a navigation dropdown menu where you may have many sub-level items.

The model:

    export type LinksObj = {
        linkText: string;
        linkExternalUrl?: string;
        linkTarget?: string | '_name';
        linkIcon?: string;
        ariaLabel?: string;
        subLinkItems?: LinksObj[];
      }

The possibilities:`LinksObj.subLinkItems.subLinkItems.subLinkItems.subLinkItems.linkText;`

This seems to be working just fine, but should I be doing this? I can't see any issues with it, but it seems odd.
## [4][Is it possible to define a type for a variable that does not allow undefined value ?](https://www.reddit.com/r/typescript/comments/i4565u/is_it_possible_to_define_a_type_for_a_variable/)
- url: https://www.reddit.com/r/typescript/comments/i4565u/is_it_possible_to_define_a_type_for_a_variable/
---
The question in the title .
## [5][Is there such an tsconfig option such as exclude from compiling but not from linting ?](https://www.reddit.com/r/typescript/comments/i4aisl/is_there_such_an_tsconfig_option_such_as_exclude/)
- url: https://www.reddit.com/r/typescript/comments/i4aisl/is_there_such_an_tsconfig_option_such_as_exclude/
---
I am asking because I want to lint my test files with the same file that I lint my src files . But I do not want to compile my test files . Exclude option excludes the test files from both linting and compiling .

Should I go for two different tsconfig files ?

How will  I make sure that they have the same compiler options ?

Edit : I am using ts-jest for tests . It automatically compiles the tests file behind the scenes .
## [6][Gary Bernhardt - TypeScript and Testing | Full Stack Radio](https://www.reddit.com/r/typescript/comments/i4bof4/gary_bernhardt_typescript_and_testing_full_stack/)
- url: https://www.fullstackradio.com/episodes/144
---

## [7][Reducing TypeScript boilerplate in Next.js pages](https://www.reddit.com/r/typescript/comments/i3vw7j/reducing_typescript_boilerplate_in_nextjs_pages/)
- url: https://github.com/vriad/nxtk
---

## [8][is keyword, versus as](https://www.reddit.com/r/typescript/comments/i3slde/is_keyword_versus_as/)
- url: https://www.reddit.com/r/typescript/comments/i3slde/is_keyword_versus_as/
---
Normally I use the `as` keyword to assert a type to help the compiler read the code correctly. Sometimes I also use `as` as a type annotation (for others who read my code including myself).

There is also an `is` keyword. It appears to be another assertion. I have trouble telling what the difference is exactly; can someone clarify the distinction? I don't remember the last time I've used `is`.
## [9][Announcing the new TypeScript Website](https://www.reddit.com/r/typescript/comments/i3ojd3/announcing_the_new_typescript_website/)
- url: https://devblogs.microsoft.com/typescript/announcing-the-new-typescript-website/
---

## [10][Parenthesis Balancing using Monoids in TypeScript](https://www.reddit.com/r/typescript/comments/i3puyh/parenthesis_balancing_using_monoids_in_typescript/)
- url: https://medium.com//parenthesis-balancing-using-monoids-in-typescript-56e5c994e497?source=friends_link&amp;sk=3bc26cc47a80f3bf40c2da0f6221a1c6
---

## [11][Do popular API have public TypeScript interfaces for their data?](https://www.reddit.com/r/typescript/comments/i3mzur/do_popular_api_have_public_typescript_interfaces/)
- url: https://www.reddit.com/r/typescript/comments/i3mzur/do_popular_api_have_public_typescript_interfaces/
---
I'm making an app with the PokeApi, and I realize that if I want intellisense and type checking for my API responses, i'd have to create a whole massive interface. Do the makers of API's typically have a TS interface for their API, or is there anywhere else that might already? Thanks.
