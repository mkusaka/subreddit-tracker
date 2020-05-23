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
## [2][TypeScript taught me something today](https://www.reddit.com/r/typescript/comments/goovef/typescript_taught_me_something_today/)
- url: https://www.reddit.com/r/typescript/comments/goovef/typescript_taught_me_something_today/
---
So I'm working on a form today, and when I was coding one particular select input, I noticed I needed two different pieces of information from whatever I was selecting. I tried a few different things but nothing was working. As I was noodling around, I hovered over the value attribute in the option tag, and TS told me that the value can be a string\[\] as well as a string or a number. I guess I would have eventually solved the problem, but TypeScript taught me how to get multiple values and from there it was just a matter of of hovering instead of reading a lot of documentation. Thank you Typescript!
## [3][query in Optional parameter: typescript](https://www.reddit.com/r/typescript/comments/gp3t7g/query_in_optional_parameter_typescript/)
- url: https://www.reddit.com/r/typescript/comments/gp3t7g/query_in_optional_parameter_typescript/
---
Hello All,

I  have a query

topic: optional parameter

I want to call a specific optional parameter? is there a way to do that, without sending all  the optional parameter. For eg:- 

in the below example I want to send value only for z variable without sending x and y value ,

something like

add( z: 5);

function add (x ?:  number,  y ?: number, z ?: number){

}

is there a way to do it?

thanks in advance
## [4][How to use TypeScript Airbnb configuration?](https://www.reddit.com/r/typescript/comments/gorkz8/how_to_use_typescript_airbnb_configuration/)
- url: https://www.reddit.com/r/typescript/comments/gorkz8/how_to_use_typescript_airbnb_configuration/
---
I try to use ESLint Airbnb configuration for several days and I don't know how.
## [5][What to learn?](https://www.reddit.com/r/typescript/comments/gorl4d/what_to_learn/)
- url: https://www.reddit.com/r/typescript/comments/gorl4d/what_to_learn/
---
I am trying to get into web development and just finished learning HTML and CSS. I have used a bit of javascript here and there but never learnt it.

My question is should I learn Javascript or start learning Typescript. I know Typescript is a superset of javascript, but is there any advantage or disadvantage of trying to learn one over other. 

I know how to code in C++/ python and coming from a strongly typed language, Typescript feels like the way to go.
## [6][Generic not working for types imported via jsdoc comment , in js files , but works via typescript types defined in jsdoc comment . How to fix that ?](https://www.reddit.com/r/typescript/comments/goj3yj/generic_not_working_for_types_imported_via_jsdoc/)
- url: https://i.redd.it/r9j7js42gb051.png
---

## [7][Demo: TypeScript 3.8 Private Fields](https://www.reddit.com/r/typescript/comments/gonmxn/demo_typescript_38_private_fields/)
- url: https://www.youtube.com/watch?v=c0QMj8x0-Mk
---

## [8][How many of you need to validate cross-file types?](https://www.reddit.com/r/typescript/comments/goos0o/how_many_of_you_need_to_validate_crossfile_types/)
- url: https://www.reddit.com/r/typescript/comments/goos0o/how_many_of_you_need_to_validate_crossfile_types/
---
Hi! Author of [typecheck.macro](https://github.com/vedantroy/typecheck.macro) here. I've recently run into an issue, where due to the way Babel works, I don't think it will be possible to validate a type that spans multiple files.

Example: An interface that extends an imported type. Of course you can still generate validators for a type that remains constrained to a single file, and then export those validators to other files.

So, my question is. How big of an issue is it that you can't validate cross-file types? Is this problematic or not?
## [9][Creating an open source team management app. Server and client are both typed and use Prisma, Nexus, Next and Apollo. Would love your feedbacks, and contributions!](https://www.reddit.com/r/typescript/comments/goau56/creating_an_open_source_team_management_app/)
- url: https://github.com/troup-io
---

## [10][Is a function a valid void type?](https://www.reddit.com/r/typescript/comments/go4ahi/is_a_function_a_valid_void_type/)
- url: https://www.reddit.com/r/typescript/comments/go4ahi/is_a_function_a_valid_void_type/
---
While I was refactoring some code, I changed some currying functions to simple functions. Obviously I started with the interfaces and I realized this strange thing:

    interface FuncInterface {
      fi: (id: number) =&gt; void;
    }
    const func = (id: number) =&gt; () =&gt; {
      console.log(id);
    };
    const implementation: FuncInterface = {
      fi: func,
    };

TS doesn't complain although `func` is returning a function.

Why is this happening?
## [11][Confused by convention of using the word "Intrinsic"](https://www.reddit.com/r/typescript/comments/gohl2r/confused_by_convention_of_using_the_word_intrinsic/)
- url: https://www.reddit.com/r/typescript/comments/gohl2r/confused_by_convention_of_using_the_word_intrinsic/
---
**Problem:**

I understand  what intrinsic means generally, but I'm  wondering if there is some concrete meaning it has in the context of Typescript.

**Background:**

React has [JSX.IntrinsicElements](https://www.typescriptlang.org/docs/handbook/jsx.html#intrinsic-elements), Which is described as:

\*something intrinsic to the environment (browser)\*e.g.  native elements (e.g. strings like "div", "span") as opposed to components

~~But TypeScript when I  look through the~~ [~~typescript source~~](https://github.com/microsoft/TypeScript/blob/master/tests/baselines/reference/intrinsics.js)~~, I can't tell if Intrinsic simply denotes:- inheritance (e.g. "base type"),   or- environment feature (e.g. "div")~~  (actually thats just test of jsx.intrtinsic\*)

This is further muddied, as I wonder "Intrinsic" is alluding to something  like an [intrinsic function](https://en.wikipedia.org/wiki/Intrinsic_function)

*a function whose implementation is handled specially by the* [*compiler*](https://en.wikipedia.org/wiki/Compiler)*.*

Can anyone put me on the right track?

**EDIT**: IMHO Intrinsic is simply *something intrinsic to the environment*
