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
## [2][What `{type: T}` does here?](https://www.reddit.com/r/typescript/comments/imwrt7/what_type_t_does_here/)
- url: https://www.reddit.com/r/typescript/comments/imwrt7/what_type_t_does_here/
---
\`\`\`  
type LookUp&lt;U, T extends string&gt; = { \[K in T\]: U extends { type: T } ? U : never }\[T\]  
\`\`\`  


I don't understand this helper type at all, because I don't know what {type: T} does here, also, this is not in the docs. So could you please explain this to me, thank you!
## [3][Help with Typescript Compiler internals](https://www.reddit.com/r/typescript/comments/impxj1/help_with_typescript_compiler_internals/)
- url: https://www.reddit.com/r/typescript/comments/impxj1/help_with_typescript_compiler_internals/
---
EDIT: 

Perhaps mentioning the database side was a distraction. I was just trying to provide some context on the usage, but my real project is to add a new type to the language, similar to the `any` type, but which is consistent and allows outputting the final inferred type at compile time. What I'm envisioning is a fetch like function that returns `Promise&lt;inferred&gt;`. The user interacts with this variable of the `inferred` type as normal, and at compile time the type checker goes through and decides what the user expects the value to be. For example:

    declare const add(x: number, y: number): number
    declare const x: inferred
    let y: number = x.bar
    add(5, y)
    x.foo.substr(5,2)

At the end of this, the compiler should infer `x` to have type `{ bar: number, foo: string }`. Notice that in no place did the user have to use a type annotation (other than the variable declaration)

ORIGINAL:

I am a computer science senior, working on an independent study project with a professor at my university. This professor is currently working on a paper on database querying which allows specifying the shape of the returned data. Since I had an interest in programming language theory, we decided to try and integrate this with the TypeScript language.  

The idea is for a type to represent the result of a database query, and then usage of a variable with that type can be tracked by TypeScript's type checker to determine what shape the user expects the data to be in at runtime. The compilation process would then output this type (probably as a TypeScript interface or JSON schema), and send it along with the query to the database, which would then transform data to the correct shape, or send an error if transformation isn't possible. Of course, the type system should ensure usage of a variable is consistent, i.e. the user shouldn't treat the variable as a string at one point and then later in the program use it as a number. The type system should also ensure that all usage of the type is compatible with deserialized JSON, as that is all that can be provided by the query function.

I have already started work on a [fork of TypeScript](https://github.com/WhiteAbeLincoln/TypeScript), but I have recently gotten stalled since the language is so large and there doesn't seem to be any publicly available documentation on the internal workings of the checker.ts file. My professor is unable to help since he doesn't have much experience with typescript or PLT.

Rather than struggling on by myself, I figured I could reach out and ask for help. I've already tried emailing members of the TS team, but failed to get any response. 

Does this community have any experience with the typescript compiler, or advice on implementing this, or know who I can contact who could guide me in the right direction? Anything would help, but some documentation on the checker.ts file would be especially useful, or some kind of guide on adding new types to the checker, as was done previously with the unknown type.
## [4][Fullstack types + runtime validation?](https://www.reddit.com/r/typescript/comments/imz9ng/fullstack_types_runtime_validation/)
- url: https://www.reddit.com/r/typescript/comments/imz9ng/fullstack_types_runtime_validation/
---
I want to find a nice library that supports runtime validation and typescript types as well. I would love for it to be serializable so we can consume it from a client and have the same types and validations there.

I know that graphql has some libs that generates types and allows the client to generate types from a gql schema, but I want this to work with vanilla typescript without graphql. Any advices on where to start?
## [5][design patterns recommendation](https://www.reddit.com/r/typescript/comments/imlmr9/design_patterns_recommendation/)
- url: https://www.reddit.com/r/typescript/comments/imlmr9/design_patterns_recommendation/
---
can you recommend me resources for design patterns? thank you.
## [6][First impressions of Prisma](https://www.reddit.com/r/typescript/comments/imapom/first_impressions_of_prisma/)
- url: https://www.reddit.com/r/typescript/comments/imapom/first_impressions_of_prisma/
---
I've been messing around in Prisma for some hours now. The syntax for defining your  schema is compact, all of the tables are defined in a single schema file.  This schema file can either be used to create the tables in the DB, or you can introspect an existing DB to generate the models for the schema.

The syntax for queries is easy to remember and autocompletion makes fields in the queries discoverable.

I guess I would just encourage you to try it out and see how you feel. You can add it to an existing project and start running queries in like 10 minutes (by introspecting your DB).

Most of my experience is TypeORM in NestJS. I was sort of discouraged from doing much on the back end because of the complexity of TypeORM. Now I feel like I can build an application that I am willing to maintain.

I remember being frustrated waiting for Prisma 2 to come out and frustrated that Prisma 1 wasn't something that I could use moving forward after going through a lengthy tutorial on Prisma 1. Regardless, I am glad the creators had ambitions for a better product.

I hope this project is well maintained moving forward because I want to keep using it. Which is why I am sort of evangelizing on here. I want to keep using it because of the syntax and autocompletion.
## [7][Are there any plans on adding CTFE or Const Expressions?](https://www.reddit.com/r/typescript/comments/imttia/are_there_any_plans_on_adding_ctfe_or_const/)
- url: https://www.reddit.com/r/typescript/comments/imttia/are_there_any_plans_on_adding_ctfe_or_const/
---
I did some searching around and I couldn't find any conversations related to this within TypeScript, but I feel I must just not be knowing where to look.  
*CTFE = Compile time function execution*  
[https://tour.dlang.org/tour/en/gems/compile-time-function-evaluation-ctfe](https://tour.dlang.org/tour/en/gems/compile-time-function-evaluation-ctfe)

Rust Const Eval: (I got the title wrong, should be Const Evalutions):  
[https://doc.rust-lang.org/reference/const\_eval.html](https://doc.rust-lang.org/reference/const_eval.html)
## [8][How do I get started?](https://www.reddit.com/r/typescript/comments/imsvg3/how_do_i_get_started/)
- url: https://www.reddit.com/r/typescript/comments/imsvg3/how_do_i_get_started/
---
I'm looking for advice and resources to start learning typescript. Also any suggestion would be more than welcome
## [9][How do I extend a class prototype with TypeScript type checking?](https://www.reddit.com/r/typescript/comments/imdds1/how_do_i_extend_a_class_prototype_with_typescript/)
- url: https://stackoverflow.com/questions/63739143/how-do-i-extend-a-class-prototype-with-typescript-type-checking
---

## [10][Deploy Friday: E21 JavaScript News - Typescript 4.0 and more](https://www.reddit.com/r/typescript/comments/imgfr0/deploy_friday_e21_javascript_news_typescript_40/)
- url: https://www.youtube.com/watch?v=A2jNLyzl794&amp;feature=youtu.be
---

## [11][TypeScript ORM with no query builder, supporting full SQL queries](https://www.reddit.com/r/typescript/comments/iltfzj/typescript_orm_with_no_query_builder_supporting/)
- url: https://github.com/Seb-C/kiss-orm#kiss-orm
---

