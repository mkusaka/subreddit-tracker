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
## [2][A SQL database implemented purely in TypeScript type annotations.](https://www.reddit.com/r/typescript/comments/izymbn/a_sql_database_implemented_purely_in_typescript/)
- url: https://i.redd.it/8yy97gfqyep51.png
---

## [3][VS Code generate explicit type annotation from inferred type](https://www.reddit.com/r/typescript/comments/izyayk/vs_code_generate_explicit_type_annotation_from/)
- url: https://www.reddit.com/r/typescript/comments/izyayk/vs_code_generate_explicit_type_annotation_from/
---
Hey everyone! I've been missing this feature for a long time and there was no solution, so yesterday I made an extension for VS Code that generates explicit types. Give it a go if you want :)  
[https://marketplace.visualstudio.com/items?itemName=nick-lvov-dev.typescript-explicit-types](https://marketplace.visualstudio.com/items?itemName=nick-lvov-dev.typescript-explicit-types)
## [4][Genetic Algorithm applied to the Infinite Monkey Theorem](https://www.reddit.com/r/typescript/comments/izpcrt/genetic_algorithm_applied_to_the_infinite_monkey/)
- url: https://github.com/f0lg0/geneticAlgorithm-TS
---

## [5][Union with never: Worth it?](https://www.reddit.com/r/typescript/comments/j02y64/union_with_never_worth_it/)
- url: https://www.reddit.com/r/typescript/comments/j02y64/union_with_never_worth_it/
---
Say I have a function that *can* (but does not always) error:

    const fn = (code:number): string | never =&gt; {
        if (code &gt;= 400) {
            throw Error(Error code: ${code});
        }
        return Code: ${code}; 
    }

&amp;#x200B;

Is it considered good style/helpful/accurate to include "never" in the return type? Or is never reserved for functions that by design-never return.
## [6][TypeScript — Excess Property Checks](https://www.reddit.com/r/typescript/comments/j028h2/typescript_excess_property_checks/)
- url: https://medium.com/@tal.ohana.x/typescript-excess-property-checks-6ffe5584f450
---

## [7][I made some unofficial TSX and JSX logos if you like logos like me](https://www.reddit.com/r/typescript/comments/j028bn/i_made_some_unofficial_tsx_and_jsx_logos_if_you/)
- url: https://github.com/Protectator/jsx-tsx-logos
---

## [8][Types of Apps that can be built with Angular Framework](https://www.reddit.com/r/typescript/comments/j03f9m/types_of_apps_that_can_be_built_with_angular/)
- url: https://tekkiwebsolutions.com/blog/angular-framework-apps/
---

## [9][Confusing situation with typing multiple recursive union types](https://www.reddit.com/r/typescript/comments/izwd5j/confusing_situation_with_typing_multiple/)
- url: https://www.reddit.com/r/typescript/comments/izwd5j/confusing_situation_with_typing_multiple/
---
I'm working on a lambda calculus interpreter out of *The Implementation of Functional Programming Languages* and I'm running into some trouble getting the type safety I'm looking for.

In the basic lambda calculus, an expression can be one of three things:

    type Expression =
      | Variable
      | Abstraction&lt;Expression&gt;
      | Application&lt;Expression&gt;

    interface Variable {
      nodeType: "var";
      id: string;
    }

    interface Abstraction&lt;E&gt; {
      nodeType: "abs";
      argument: Variable;
      body: E;
    }

    interface Application&lt;E&gt; {
      nodeType: "app";
      left: E;
      right: E;
    }

This works fairly well like you'd expect. Sometimes the lambda calculus AST is extended with the notion of variable definition, called a Let binding, and I'll call that "language" AugmentedExpression

    type AugmentedExpression =
      | Variable
      | Abstraction&lt;AugmentedExpression&gt;
      | Application&lt;AugmentedExpression&gt;
      | Let&lt;AugmentedExpression&gt;

    interface Let&lt;E&gt; {
      nodeType: "let";
      varName: Variable;
      equals: E;
      inExpr: E;
    }

This works as well, on its own. The issue is, I want to be able to write a function that can take either an `Expression` or an `AugmentedExpression` and return that same type. For instance, if I write a function that takes two Expressions and returns an Application of one to the other, I know that the return is also of type Expression. 

    function app(left: Expression, right: Expression): Expression {
      return { nodeType: 'app', left, right };
    }

That's fine. I could also write exactly the same code and substitute AugmentedExpression for Expression, and it would work just fine. If I try to write the types such that it guarantees that if you pass in an Expression you get an Expression back, and if you pass in an AugmentedExpression you get an AugmentedExpression back, it gets more difficult.

Initially I tried:

    type Language = Expression | AugmentedExpression
    function app&lt;L extends Language&gt;(left: L, right: R): L 

That doesn't work, you have to either cast the result to unknown and then back to L, or you have to specify the return type as Application&lt;L&gt;, which kind of makes sense, and works for that situation.

But sometimes the return is more complicated, consider repeated application:

    function appN(...args: Expression): Expression {
      if(args.length == 2)
        return app(...args);
      else
        return appN(app(args[0], args[1]), ...args.slice(2))
     }

Here the same thing is true: if you pass in Expressions, you get an Expression, if you sub it for AugmentedExpression it works exactly the same way, returns AugmentedExpression. But I can't make this type check. I want to be able to write it as something like `appN&lt;L extends Language&gt;(...args: L[]) : L`, but the type checker does not like it one bit. 

Any ideas on how I can express this in Typescript types? Without getting super unwieldy, which all of the solutions I've come up with so far have been. I'm guessing that the issue comes down to the fact that what I'm trying to do is somewhere between nominal and structural typing.
## [10][Can someone help me out](https://www.reddit.com/r/typescript/comments/iztqsw/can_someone_help_me_out/)
- url: https://www.reddit.com/r/typescript/comments/iztqsw/can_someone_help_me_out/
---
Hi,

I'm having the following 'problem': [https://stackoverflow.com/questions/64059374/why-isnt-the-argument-assignable-only-when-we-dont-know-exactly-what-type-from/64059656#64059656](https://stackoverflow.com/questions/64059374/why-isnt-the-argument-assignable-only-when-we-dont-know-exactly-what-type-from/64059656#64059656)

Can someone explain this behaviour to me?

Thanks
## [11][Exporting third-party type for consumption](https://www.reddit.com/r/typescript/comments/izqfax/exporting_thirdparty_type_for_consumption/)
- url: https://www.reddit.com/r/typescript/comments/izqfax/exporting_thirdparty_type_for_consumption/
---
Not sure how best to ask this question, so please bear with me :)  


I've got an npm package that I'm building/publishing which derives a type from a third-party package:

&amp;#x200B;

`import { GridChildComponentProps } from 'react-window'`

`interface OwnCellProps {`

  `pageSize: number;`

`}`

`export type CellProps = OwnCellProps &amp; GridChildComponentProps;`

&amp;#x200B;

Now when my consuming project uses this npm package it gets \`CellProps\` as type \`any\`. When I click through the references, I see that my declaration file has a similar line:

`export declare type CellProps = OwnCellProps &amp; GridChildComponentProps;`

OwnCellProps has the correct type, because it's defined above, but GridChildComponentProps has no type.  Likely because it is not defined.  There is the same import line at the top of the file, but \`@types/react-window\` is not installed in the consuming project.

So my question is: how can the consuming project get the types from an npm package's third-party type dependency?

I have \`@types/react-window\` listed as a devDependency, but the consuming project does not install it.  I'm wondering if there's a way to include the third-party type in the declaration file created in my npm package.  

Any help is appreciated!
