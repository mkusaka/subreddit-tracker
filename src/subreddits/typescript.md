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
## [2][Building a game with TypeScript. Drawing Grid 4/5](https://www.reddit.com/r/typescript/comments/il512l/building_a_game_with_typescript_drawing_grid_45/)
- url: https://medium.com/@gregsolo/building-a-game-with-typescript-iii-drawing-grid-4-5-398af1dd638d?sk=49c92b3604c33e633bf6dc4a1e2846ed
---

## [3][What to expect from Typescript 4.0 and more](https://www.reddit.com/r/typescript/comments/il2uaw/what_to_expect_from_typescript_40_and_more/)
- url: https://medium.com/@metodieff.stefan/what-is-new-in-typescript-4-0-and-more-6c5fa72fa1db?source=friends_link&amp;sk=649d4059694facf4f13223b3831da308
---

## [4][Declaration of types in a separate file](https://www.reddit.com/r/typescript/comments/ikruk6/declaration_of_types_in_a_separate_file/)
- url: https://www.reddit.com/r/typescript/comments/ikruk6/declaration_of_types_in_a_separate_file/
---
Hi everyone!

I have a module which export a class with some methods. Methods returns a complex data.

For example `SomeClass.ts`:

`// some imports`

`export type ComplexType1 = … // generated from json. Contains a lot of subtypes, enums etc.`

`export type ComplexType2= …`

`// and so on`

`class SomeClass {`

`method1 = (): ComplexType1 =&gt; {…}`

`method2 = (): ComplexType1 =&gt; {…}`

`// etc`

`}`

This structure is repeated in several files.

My problem is that type declarations takes several screens. I trying just copy-paste types into SomeClass.d.ts but types became unavailable inside SomeClass.ts

How I can move types in separate file?
## [5][How to define a prop type for an object that is generic](https://www.reddit.com/r/typescript/comments/ikvnvh/how_to_define_a_prop_type_for_an_object_that_is/)
- url: https://www.reddit.com/r/typescript/comments/ikvnvh/how_to_define_a_prop_type_for_an_object_that_is/
---
So I have a React component with maybe some excessive generics. I've got a component that takes in an array of rules, and a separate prop that is an object of arguments to pass to each rule. It's a pattern that I've adopted because it avoids creating a closure over the arguments so the rules can be re-evaluated on each render (since the props being passed get updated).

Anyway, my main problem is my PropTypes. I want to still use PropTypes because this code is for a library that will also be used in JavaScript projects, not just TypeScript ones. I just can't get these to work together, the moment I do it breaks compilation.

I know it has something to do with the fact that I'm typing "ruleProps" as an object, whereas it's generic in the TypeScript interface. I'll also point out that it's the "rules" prop which breaks, because the arguments expected by the function no longer match the type.

Hoping for guidance on how to do this. Thanks.

    export interface Rule&lt;RuleProps extends object&gt; {
        allow: (ruleProps?: RuleProps) =&gt; boolean;
        redirect: string;
    }
    
    interface Props&lt;CompProps extends object, RuleProps extends object&gt; {
        rules?: Array&lt;Rule&lt;RuleProps&gt;&gt;;
        ruleProps?: RuleProps;
    }
    
    ProtectedRoute.propTypes = {
        rules: PropTypes.arrayOf(PropTypes.shape({
            allow: PropTypes.func.isRequired,
            redirect: PropTypes.string.isRequired
        })),
        ruleProps: PropTypes.object
    };

Edit: The error I'm getting is:

    TS2322: Type 'Rule&lt;RuleProps&gt;[] | undefined' is not assignable to type 'Rule&lt;object&gt;[] | undefined'.   Type 'Rule&lt;RuleProps&gt;[]' is not assignable to type 'Rule&lt;object&gt;[]'.     Type 'Rule&lt;RuleProps&gt;' is not assignable to type 'Rule&lt;object&gt;'.       Property 'allow' is missing in type '{}' but required in type 'RuleProps'.

This error occurs in a unit test when I try to actually use the component and pass values to these props. If I remove the PropTypes, everything works just fine.
## [6][If you are using TypeDoc, I created a theme with particles background. Hope you like it.](https://www.reddit.com/r/typescript/comments/ikofcd/if_you_are_using_typedoc_i_created_a_theme_with/)
- url: https://github.com/tsparticles/typedoc-particles-theme
---

## [7][TypeScript Jesus Anders Hejlsberg has been working on some fun new features: Template string types and mapped type `as` clauses](https://www.reddit.com/r/typescript/comments/ikc3h4/typescript_jesus_anders_hejlsberg_has_been/)
- url: https://github.com/microsoft/TypeScript/pull/40336
---

## [8][learning typescript](https://www.reddit.com/r/typescript/comments/il0ysf/learning_typescript/)
- url: https://www.reddit.com/r/typescript/comments/il0ysf/learning_typescript/
---
so, i really wanna learn typescript, for discord.js and maybe a little bit of game dev and website dev, but i honestly have zero clue on how to get started, i learn by doing, and not just watching videos because then i can't remember things, any tips? i would really appreciate the help
## [9][TypeScript tips and tricks for beginners: Part1 - Declarations](https://www.reddit.com/r/typescript/comments/ikqd8n/typescript_tips_and_tricks_for_beginners_part1/)
- url: https://drag13.io/posts/typescript-tips-tricks-declarations/index.html
---

## [10][Type representing valid integer indexes of a const array?](https://www.reddit.com/r/typescript/comments/ikom9c/type_representing_valid_integer_indexes_of_a/)
- url: https://www.reddit.com/r/typescript/comments/ikom9c/type_representing_valid_integer_indexes_of_a/
---
Say I have a readonly array declared as:

    const arr = [‘this’, ‘is’, ‘a’, ‘test’] as const;


I now want to declare a type representing all valid index integers for this readonly array. So, I want:

    type indexes = 0|1|2|3;

But I want the above type computed. Assume there are dozens of indexes. Is there any way to represent this without manually typing out the union of integers?
## [11][Why does undefined "disappear" from my optional props?](https://www.reddit.com/r/typescript/comments/ikl0pb/why_does_undefined_disappear_from_my_optional/)
- url: https://www.reddit.com/r/typescript/comments/ikl0pb/why_does_undefined_disappear_from_my_optional/
---
If I put the following into the [Typescript Playground](https://www.typescriptlang.org/play?#code/JYOwLgpgTgZghgYwgAgApQPYAcDOyDeAUMssACYD8AXMiAK4C2ARtANzG1wMTXI5hRQAc3YBfQoQgAPLBihhkZCPDoAbBTDogEYYBhDIAKtLAAKfKTIAaTt2Sia6bDgCUNfoJBCCHKBDB0UAYABgAk+OQOyOEgXBCiwWKEQA), and hover over `id` or `name` I get the correct type followed by ` | undefined`. However, if I do that in my project in VS Code, `undefined` is no longer there. Why does `undefined` disappear in my project?

It's problematic because type checking "further down" becomes incorrect which also affects ESLint rules. It for example wants to remove the `?` in `id?.toString()` because of "Unnecessary optional chain on a non-nullish value.". But in the interface, it clearly _is_ optional.

What am I missing here?

```
interface Props {
  id?: number;
  name?: string;
}

export default function Text({ id, name }: Props): string {
  return `${id}: ${name}`;
}

```
