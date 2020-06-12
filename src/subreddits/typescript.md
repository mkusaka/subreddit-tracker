# typescript
## [1][Who's hiring Typescript developers - June](https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/)
- url: https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/
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
## [2][Arbitrary amount of generics in function call?](https://www.reddit.com/r/typescript/comments/h7ivji/arbitrary_amount_of_generics_in_function_call/)
- url: https://www.reddit.com/r/typescript/comments/h7ivji/arbitrary_amount_of_generics_in_function_call/
---
The mysql package allows you to make an SQL query with an arbitrary amount of statements. In my case, I've got a stored procedure with an OUT parameter:

    db.query('call getSalaries(1, 2, @average); select @average)'); // two statements

This would return an array where the first element is a Salary[] and the second element is an array of numbers (the average - though in this particular case, that array would contain a single number). In general, depending on the statements in the query and how many there are, I'm going to have an array of elements of different types, with corresponding length.

I would like to be able to call query in different ways. In general, if I call db.query&lt;A, B, C, ...etc&gt;(myQuery), then I know I'm going to get an array where the first element is A[], the second B[], the third C[], and so on. How would I go about modeling that?

I hope the question is clear. Thanks!
## [3][Node.js, Dependency Injection, Layered Architecture, and TDD: A Practical Example Part 1](https://www.reddit.com/r/typescript/comments/h13hr7/nodejs_dependency_injection_layered_architecture/)
- url: https://carlosgonzalez.dev/posts/node-js-di-layered-architecture-and-tdd-a-practical-example-part-1/
---

## [4][What is a valid use case for Interfaces with indexers in typescript?](https://www.reddit.com/r/typescript/comments/h7k8y7/what_is_a_valid_use_case_for_interfaces_with/)
- url: https://www.reddit.com/r/typescript/comments/h7k8y7/what_is_a_valid_use_case_for_interfaces_with/
---
I could not quite come up with a valid use case to justify why typescript supports interface with indexers. plain javascript provides you with indexing objects then why one would need to define an indexer in an interface.?
## [5][How to create a Conditional Type/Typeguard for this?](https://www.reddit.com/r/typescript/comments/h7dob0/how_to_create_a_conditional_typetypeguard_for_this/)
- url: https://www.reddit.com/r/typescript/comments/h7dob0/how_to_create_a_conditional_typetypeguard_for_this/
---
So I've got a data structure (`SingleValue`) that I want to be able to handle in 2 different ways. The first is when only one or no instances of it is needed, in which case the variable holding the value should be of type `SingleValue | null`. The other case is when many are needed, in which case the variable holding the value should be of type `SingleValue[]`.

How can I set up a single variable and a typeguard function for a type that is conditionally one of those two types based on a generic type argument.

This is what I have at the moment but it's not quite right:

```ts
type SingleValue = { /* ... */ };

type Value&lt;Multiple extends boolean&gt; = Multiple extends true
  ? SingleValue[]
  : SingleValue | null;

function isMultipleValue&lt;Multiple extends boolean&gt;(
 value: Value&lt;Multiple&gt;
): value is Value&lt;true&gt; {
 return Array.isArray(value);
}
```

The error I get is:

```
A type predicate's type must be assignable to its parameter's type.
  Type 'readonly SingleValue[]' is not assignable to type 'Value&lt;Multiple&gt;'.
```

(this error happens on the 3rd to last line)

Note: I cannot switch to just using an array for the one/none case due to how other parts of my project are set up and the implementation of 3rd party libraries. For the same reason `Multiple` but be a generic type argument.
## [6][How to warn on number to boolean coercion?](https://www.reddit.com/r/typescript/comments/h7aw3i/how_to_warn_on_number_to_boolean_coercion/)
- url: https://www.reddit.com/r/typescript/comments/h7aw3i/how_to_warn_on_number_to_boolean_coercion/
---
Hi, I was wondering if it's possible to get Typescript to warn whenever a possible number is coerced into a boolean. For example:

```
const x: number | null = 5;

if(x) {} // warn

const y = !x //warn
```

The reason for this is that I would like to avoid numbers potentially becoming falsey due to 0.
## [7][Duplication of code for ORMs](https://www.reddit.com/r/typescript/comments/h15lwb/duplication_of_code_for_orms/)
- url: https://www.reddit.com/r/typescript/comments/h15lwb/duplication_of_code_for_orms/
---
First of all, I am a newcomer to Typescript, and my previous projects involved C# and WPF.

On the C# side of things, I have always used Entity Framework as ORM, which was a mighty powerful ORM. It handled code-first as well as db-first approaches quite easily.

As far as I can tell, there is nothing in Typescript that handles db-first approach. But the bigger problem for me is duplication of code when using ORMs. Entity Framework could easily infer the database schema from the classes only. In Typescript, if I want to define a schema, I will have to do it twice, once as a plain old TS class, and secondly for the ORM. An example of such a code is [here](https://sequelize.org/master/manual/typescript.html).

I have found typegoose to do something akin to Entity Framework when it comes to ORM. TS-Mongoose looked promising until I ran into the issue [here](https://www.reddit.com/r/typescript/comments/h0e3ff/help_with_tsmongoose/).

Has anyone found a less tedious solution for Postgresql too?

&amp;#x200B;

EDIT: Found it. It's TypeORM.
## [8][Reliable typescript language parser.](https://www.reddit.com/r/typescript/comments/h11s6h/reliable_typescript_language_parser/)
- url: https://www.reddit.com/r/typescript/comments/h11s6h/reliable_typescript_language_parser/
---
We're trying to build a command line app to find orphan code in a larger typescript project.

Parsing the imports is an issue as they could be multiline.

Is there a parse that can take a .ts or .tsx file and parse out the imports reliably into an AST?
## [9][A lightweight yet powerful dependency injection framework using modern TypeScript and decorators](https://www.reddit.com/r/typescript/comments/h0ogq0/a_lightweight_yet_powerful_dependency_injection/)
- url: https://github.com/H1ghBre4k3r/dependory
---

## [10][Which of these TSLinter tools are most important for my offshore team to follow?](https://www.reddit.com/r/typescript/comments/h14yuh/which_of_these_tslinter_tools_are_most_important/)
- url: https://www.reddit.com/r/typescript/comments/h14yuh/which_of_these_tslinter_tools_are_most_important/
---
Hey guys,

My team had to turn off all of the following rules on TSLinter, but to improve standards over time I want to turn them on. Which of these rules are most important for code quality?

&amp;#x200B;

1. "jsx-wrap-multiline": false,
2. "jsx-no-lambda": false,
3. "jsx-self-close": false,
4. "jsx-no-string-ref": false,
5. "prefer-const": false,
6. "array-type": false,
7. "no-object-literal-type-assertion": false,
8. "one-variable-per-declaration": false,
9. "object-literal-shorthand": false,
10. "no-string-literal": false,
11. "no-implicit-dependencies": false,
12. "triple-equals": false,
13. "no-submodule-imports": false,
14. "no-console": false,
15. "no-empty": false,
16. "no-shadowed-variable": false,
17. "prefer-conditional-expression": false,
18. "forin": false,
19. "radix": false,
20. "space-within-parens": false,
21. "no-trailing-whitespace": false,
22. "whitespace":false,
23. "ban-types": false,
24. "comment-format": false,
25. "jsx-boolean-value": false,
26. "class-name": false,
27. "prefer-object-spread": false,
28. "variable-name": false,
29. "no-unused-expression": false,
30. "no-var-keyword": false,
31. "no-duplicate-imports": false,
32. "only-arrow-functions": false,
33. "no-this-assignment": false,
34. "max-classes-per-file": false,
35. "align": false
## [11][Foal TS - June release (version 1.9) - Auth with MongoDB/TypeORM, auth with REST APIs, handy shortcuts, CLI improvements](https://www.reddit.com/r/typescript/comments/h0tzhe/foal_ts_june_release_version_19_auth_with/)
- url: https://www.reddit.com/r/typescript/comments/h0tzhe/foal_ts_june_release_version_19_auth_with/
---
Foal TS framework version 1.9 is officially released!

[Foal TS - June release \(version 1.9\)](https://preview.redd.it/cbw3ao9078451.png?width=1250&amp;format=png&amp;auto=webp&amp;s=f1281bba1a03e446043f4cc901173ea9cf6a4016)

This version brings many improvements and fixes bugs.

1. **MongoDB can now be fully used with TypeORM**. The new `fetchMongoDBUser` allows you to retrieve the MongoDB user whether you use sessions or stateless JWTs.

[Authentication with JWT\/MongoDB\/TypeORM](https://preview.redd.it/st1vyyzi88451.png?width=363&amp;format=png&amp;auto=webp&amp;s=43047d0979f8707e848186f8142ea81109080ee4)

Docs: [https://foalts.gitbook.io/docs/topic-guides/authentication-and-access-control/quick-start](https://foalts.gitbook.io/docs/topic-guides/authentication-and-access-control/quick-start)

2. **\[Shortcut\] The request params and body** are passed as **second** and **third arguments** of the controllers.

[Quick access to the request body and params](https://preview.redd.it/y12146c998451.png?width=525&amp;format=png&amp;auto=webp&amp;s=1bd65f3567295334f29622b0346c1a3fe6f585d6)

Docs: [https://foalts.gitbook.io/docs/topic-guides/architecture/controllers#the-controller-method-arguments](https://foalts.gitbook.io/docs/topic-guides/architecture/controllers#the-controller-method-arguments)

3. \[CLI\] The CLI `generate` command has been extended to support subdirectories. You can now fully generate your controller (or service) architecture from the command line

[CLI commands](https://preview.redd.it/rbctunjr98451.png?width=649&amp;format=png&amp;auto=webp&amp;s=4f98058b17d5efce6efd22fbd12e0b4d87aeab26)

[Architecture generated](https://preview.redd.it/5vfxegbt98451.png?width=651&amp;format=png&amp;auto=webp&amp;s=3173ecd8a3925847c0f2f38cc572045345192821)

[app.controller.ts updated](https://preview.redd.it/hnjls0vq98451.png?width=645&amp;format=png&amp;auto=webp&amp;s=71cd1772b42804a0981cbdb8aa7ceae981d0fa30)

[api.controller.ts generated](https://preview.redd.it/gbiqvaos98451.png?width=644&amp;format=png&amp;auto=webp&amp;s=4ab3e7f480a60c528236976fce39170725c2e743)

Docs: [https://foalts.gitbook.io/docs/topic-guides/cli-and-development-environment/code-generation#create-a-controller](https://foalts.gitbook.io/docs/topic-guides/cli-and-development-environment/code-generation#create-a-controller)

Foal, in a few words, it's a **Node.js framework**:

* written in **TypeScript**
* provided **with batteries included** (Auth, OpenAPI, GraphQL, ORM, CLI, scripts, Cloud file storage, etc)
* and with a **simple and intuitive architecture** (no magic, no over-engineering).

And the must: it has more than 11,000 lines of documentation.

[https://foalts.org](https://foalts.org/)

[https://github.com/FoalTS/foal](https://github.com/FoalTS/foal)

[https://foalts.gitbook.io/docs/](https://foalts.gitbook.io/docs/)

\#TypeScript #JavaScript #NodeJS #FoalTS #CLI #Auth #Authentification #JWT
