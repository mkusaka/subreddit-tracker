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
## [2][Must I always extend protected members to be protected?](https://www.reddit.com/r/typescript/comments/h9g59c/must_i_always_extend_protected_members_to_be/)
- url: https://www.reddit.com/r/typescript/comments/h9g59c/must_i_always_extend_protected_members_to_be/
---
I have something like this:

    abstract class Base {
        protected description: string = ``;
    }
    
    class Subclass extends Base {
        private description: string = ``;
    
    }
    
    /* class Subclass
    Class 'Subclass' incorrectly extends base class 'Base'.
      Property 'description' is private in type 'Subclass' but not in type 'Base'.ts(2415) */

I don't plan to extend `Subclass` so I thought marking `description` as `protected` would not be necessary or very clear to others. I therefore thought `description` in `Subclass` should be marked `private`, but only the `protected` keyword avoids errors.

Can anyone help me understand why `protected` is needed here?

Is the case simply that any time you extend a class you must identify overwritten members with the `protected` keyword?
## [3][A class-based fixtures generator powered by TypeScript and decorators](https://www.reddit.com/r/typescript/comments/h91zp9/a_classbased_fixtures_generator_powered_by/)
- url: https://github.com/CyriacBr/class-fixtures-factory
---

## [4][When to export interfaces?](https://www.reddit.com/r/typescript/comments/h9bhbc/when_to_export_interfaces/)
- url: https://www.reddit.com/r/typescript/comments/h9bhbc/when_to_export_interfaces/
---
Hi.  If I have a file, `Foo.tx` which defines `interface IFoo` and an export default class `export default class Foo implements IFoo`; when should I `export { IFoo }`.

Are there any advantages to exporting the interface too?

---

**Edit: explanation**

I have a bunch of TypeScript files which are for a module.  Will exporting interfaces help with IntelliSense?  Will not exporting interfaces for the module (though I am exporting classes) strip away something helpful from those using the node module I am making?
## [5][Types and autocomplete IntelliSense](https://www.reddit.com/r/typescript/comments/h8tyod/types_and_autocomplete_intellisense/)
- url: https://www.reddit.com/r/typescript/comments/h8tyod/types_and_autocomplete_intellisense/
---
Hi.  I have a TypeScript project which I want to put out as npm module.

Currently the TypeScript source code is in `src` directory.
```
src
| - Foo
       | - Foo.ts  // -&gt; export default class Foo
| - Bar
        | - Bar.ts  // -&gt; export default class Bar
| - index.ts // -&gt; export { Foo, Bar }
```
With the current `tsconfig.json`, types are emitted in the `types` directory, and compiled JavaScript code is in `dist`.

But when I create a `main.js` file which imports `Foo` from `dist/index`; IntelliSense does not suggest any methods I had defined for `Foo`.  How do I make it work?
## [6][React: Do I really have to declare every properties inside props?](https://www.reddit.com/r/typescript/comments/h96wp3/react_do_i_really_have_to_declare_every/)
- url: https://www.reddit.com/r/typescript/comments/h96wp3/react_do_i_really_have_to_declare_every/
---
Using vscode.

A component have many props, eg. - `props.navigation`, `props.showMenu`, etc. But typescript is giving me error saying `navigation does not exist in Object` and squiggly red underlines `navigation`. I can resolve this by clicking Quick Fix and then declare property 'navigation'. This adds navigation in the Object interface.

Do I really have to do this for every property inside props?
## [7][Rosebox 0.2.3](https://www.reddit.com/r/typescript/comments/h8pev6/rosebox_023/)
- url: https://www.reddit.com/r/typescript/comments/h8pev6/rosebox_023/
---
[Rosebox](https://www.rosebox.dev/) is an active project with the ultimate goal to build an independent and complete JS/TS styling framework with an emphasis on writing declarative, reliable, and conflict-free style-code.
## [8][Question regarding a dynamic key to get values from an object.](https://www.reddit.com/r/typescript/comments/h8ukgv/question_regarding_a_dynamic_key_to_get_values/)
- url: https://www.reddit.com/r/typescript/comments/h8ukgv/question_regarding_a_dynamic_key_to_get_values/
---
So I will try to explain my issue:  


"account" is an object containing the usual stuff like name/password/etc..

In my updateAccount function I have a parameter "data" which can contain the Id and whatever fields needs updating. 

&amp;#x200B;

Now, here arises my struggle with typescript, in normal javascript I would simply do:

&amp;#x200B;

Object.keys(data).forEach(key =&gt; {

if(account\[key\]) { this is a valid key to update }

});

This allows me to only update what I need and filter out anything unneeded.  


Yet... typescript seems to be annoyed as hell when you try do this claiming that strings are not a valid child of an account and what not. (the code runs it just gives TS warnings)

Looking into it... it should be fixed by doing:

Object.keys(data).forEach(&lt;T extends keyof AccountModel&gt;(key: T) =&gt; {  


But then it nags about key and value being incompatible... Can anyone shed some light on how you are supposed to do this?
## [9][Why does this work?](https://www.reddit.com/r/typescript/comments/h89nb2/why_does_this_work/)
- url: https://www.reddit.com/r/typescript/comments/h89nb2/why_does_this_work/
---
I'm having trouble understanding this odd behavior...

```
interface Foo {
    n? : number;
    s: string;
}

const foo = { n : 1 } as Foo;  // OK

```

If  `s` here isn't optional, then why does this work? Isn't this an error? Shouldn't we be assigning `s` to something?
## [10][webpack and conflicting types with identical versions in imported library.](https://www.reddit.com/r/typescript/comments/h8dqld/webpack_and_conflicting_types_with_identical/)
- url: https://www.reddit.com/r/typescript/comments/h8dqld/webpack_and_conflicting_types_with_identical/
---
I'm really stuck on this issue... 

If I have a react components library that I want to isolate I can't just setup a devDependency on @types/react because when I import it  I get a conflict saying that types were found in two locations.  

This is super silly because they're the same types and same version.

I *could* setup peerDependencies, I think, but then I can't have unit tests in the main code or use that as a module by itself.

It ends up causing a LOT of hassle and now I'm basically dead in the water until I completely refactor my  components/modules into separate npm projects.

Is there a workaround here?
## [11][How can I get ts specific errors , for all my files , and log them in the terminal without building ?](https://www.reddit.com/r/typescript/comments/h8dfyq/how_can_i_get_ts_specific_errors_for_all_my_files/)
- url: https://www.reddit.com/r/typescript/comments/h8dfyq/how_can_i_get_ts_specific_errors_for_all_my_files/
---
I am taking about the errors which are provided when you tsc . Nut I want to get them without building .
