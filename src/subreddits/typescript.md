# typescript
## [1][Who's hiring Typescript developers - November](https://www.reddit.com/r/typescript/comments/jlsywo/whos_hiring_typescript_developers_november/)
- url: https://www.reddit.com/r/typescript/comments/jlsywo/whos_hiring_typescript_developers_november/
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
## [2][I made a VSCode plugin that does GoTo Definition faster than VSCode/TSServer](https://www.reddit.com/r/typescript/comments/jo3z86/i_made_a_vscode_plugin_that_does_goto_definition/)
- url: https://marketplace.visualstudio.com/items?itemName=verydanny.smart-goto
---

## [3][Short syntax for implicit type declaration with optional properties?](https://www.reddit.com/r/typescript/comments/joigd5/short_syntax_for_implicit_type_declaration_with/)
- url: https://www.reddit.com/r/typescript/comments/joigd5/short_syntax_for_implicit_type_declaration_with/
---
I am looking for the best/shortest way to initialize an object with values and declare its structure, including optional properties. E.g. I want to go from here:

    interface IPerson {
      name: string;
      phone?: string;
    }
    
    const dan: IPerson = {
      name: "Dan",
    };

to a solution that is similar to this, but even shorter:

    const dan = {
      name: "Dan",
      phone: undefined as undefined | string,
    };

Solutions I have come up with so far...

**A: Explicit "undefined" value.** Too verbose for my liking because it requires spelling out "undefined", and twice. A concise `?` style syntax would be nice, but is not supported here.

    const person = {
      name: "Dan",
      phone: undefined as undefined | string,
    };

**B: Helper function** makes code nicer to read and less verbose.

    export function optional&lt;T&gt;(init?: T): undefined | T {
      return init;
    }
    
    const dan = {
      name: "Dan",
      phone: optional&lt;string&gt;()
    }

Is there a nice, "in-language" solution to achieve a more concise syntax?
## [4][How can I declare a function declaration's type via a type alias?](https://www.reddit.com/r/typescript/comments/jo6x98/how_can_i_declare_a_function_declarations_type/)
- url: https://www.reddit.com/r/typescript/comments/jo6x98/how_can_i_declare_a_function_declarations_type/
---
If I have a function type alias like this for example:

    type F = &lt;T1, T2&gt;(param1: T1, param2: T2) =&gt; string;

Is it possible for me to declare a function declaration like this:

    function foo(param1, param2) {
      // ...
    }

to be of type `F`?

I know I can cast an expression like this:

    const foo = ((param1, param2) =&gt; {
      // ...
    }) as F;

but I want to keep the function as a declaration if possible.

Any recommendations?
## [5][Define class type. How to separate the types for instance from the types of the prototype?](https://www.reddit.com/r/typescript/comments/jo6zje/define_class_type_how_to_separate_the_types_for/)
- url: https://www.reddit.com/r/typescript/comments/jo6zje/define_class_type_how_to_separate_the_types_for/
---
How can I define the type of class so that when I use that type when I define the class there will be a separation between the instance type and the prototype instance?
Example :
```
interface IMyClass {
    //instance types
    a : string
    //prototype types
    b : () =&gt; void
    c : () =&gt; void
}
class myClass implements IMyClass {
    constructor() {
        this.a = "";
        //it has to lint error since c has to be on the prototype
        this.c = () =&gt; undefined
    }
    b() {
        return undefined
    }
    //it has to lint error because c is missing
}
```
## [6][import json possible while using esnext out?](https://www.reddit.com/r/typescript/comments/jo66pn/import_json_possible_while_using_esnext_out/)
- url: https://www.reddit.com/r/typescript/comments/jo66pn/import_json_possible_while_using_esnext_out/
---
Hi all!  Working on a thing that I want to use the latest everything possible, latest node, latest browsers, most advanced settings, etc.  Want everything to be *modern*.

I've got my typescript settings for module and target as "esnext", and it does generate code, but when I try to add an import of a json file, node throws me "Unknown file etension json for (filepath to json file)".

Is there some way to get there?  I wanted to do it via typescript import, because I want it to be both node and browser compatible.. 

I am willing to turn back my target settings if necessary, I guess...
## [7][how can i use git sparse checkout or git filter with DefinitelyTyped repo ?](https://www.reddit.com/r/typescript/comments/jnyhf1/how_can_i_use_git_sparse_checkout_or_git_filter/)
- url: https://www.reddit.com/r/typescript/comments/jnyhf1/how_can_i_use_git_sparse_checkout_or_git_filter/
---
I want to contribute to definitely typed but unfortunately I have low internet speed and cloning around 6000 types folders takes it forever.

I looked into git sparse checkout which helps to clone only specified folder with git. But I don't know which is the better way to clone repo.

Any help, suggestions or git magic tips is appreicated .
## [8][Need help. Using generic, derive the instance type of a provided class.](https://www.reddit.com/r/typescript/comments/jo1ha2/need_help_using_generic_derive_the_instance_type/)
- url: https://www.reddit.com/r/typescript/comments/jo1ha2/need_help_using_generic_derive_the_instance_type/
---
[ts playground example.](https://www.typescriptlang.org/play?noUnusedLocals=true&amp;noUnusedParameters=true#code/GYVwdgxgLglg9mABAJwKYHMYGcquQHgAVFUAPXMAEy0QG9ExUB3ACgDoOBDZdLALkTgA1mDhMwAbQC6ASgHDR4gNyIAvgD4WAfUQDaAKERHEAYQA2nLDQGEANIeMwwuZBFQAHKLsQt33gJJgOJyQqAAqAJ7uqETqMogAvOqIgcGhkdGx+qrxBsaIAPQFWHAAtqiIOCDAwNn6+miYOHgsecZF-ohMIV5hAMqIUHCInBBuniNgEYgQFlYORuaW1jNzNABicMNt+Uac3jjITugLuxAIhyDQcMgs+wKHx7mnu8ZQABbYbPsJI0ovxlUp1U9nyHS6PUG7wq7m4nHKLkGwwARhU4MBBlEKk40m5EOioRVZstTk4XOMvAJfPEknQAUY0FAQMgkO5gdkZEA)

Edit : [Here](https://www.typescriptlang.org/play?noUnusedLocals=true&amp;noUnusedParameters=true#code/GYVwdgxgLglg9mABAJwKYHMYGcquQHgEFFUAPXMAEy0XAGsw4B3MAbQF0AaASRPNSo04AIwBWqaAD4AFAH1EALkQBvAFCINiAMIAbAIZYaS5YjCom0gHTW9ydFiWEAlEu4BuRAF9O6zTDC4yBCoAA5QiojSIRHcTogAvJKI3KqecWqaiAD0WVhwALaoiDggwMCpqqpomDh40hmaObxMegGIACoAyohQcIh6EMFh-WAAnogQ+oa+GroGRhNTNABicH0NmRp6ETjI-ugzmxAIuyDQcMjS20q7++mHm5pQABbYltvx-W4Pmp6H3ocmogWm0XkUQrY9IVAj0+sIinBgD1RiEiv4cK1gohET1nkVJvNDv5AkNwkoonFEiofho0FAQMgkCF-qknEA) is the solution.
## [9][I made a tool to lookup IP address geo info and plot on a map! TypeScript frontend and backend. Link and source in comments.](https://www.reddit.com/r/typescript/comments/jno4sz/i_made_a_tool_to_lookup_ip_address_geo_info_and/)
- url: https://v.redd.it/u2ye4wgct4x51
---

## [10][Record with keys made of array of strings](https://www.reddit.com/r/typescript/comments/jo06bg/record_with_keys_made_of_array_of_strings/)
- url: https://www.reddit.com/r/typescript/comments/jo06bg/record_with_keys_made_of_array_of_strings/
---
I would like to be able to use an array of string to populate a record keys just like i would do with `Record&lt;keyof T, ...&gt;`

Is it possible ?
## [11][What cool tricks / unexpected good outcomes have you found with TypeScript recently?](https://www.reddit.com/r/typescript/comments/jndhdo/what_cool_tricks_unexpected_good_outcomes_have/)
- url: https://www.reddit.com/r/typescript/comments/jndhdo/what_cool_tricks_unexpected_good_outcomes_have/
---
I thought adding unnamed call signatures to interfaces was not possible, but it is. For an embarrassingly long time, I believed that writing this would result in errors:

    interface FooInterface {
      (): string;
       ...
     }

Because the parentheses have no name attached. For hours today I've been agonising over how to write a types declaration file for a library, where there is an interface of methods that each return said interface (i.e. for chaining), but at the end of the chain an entirely different object is returned and it is called with new parentheses, like so:

    const foo = a() // return type: FooInterface
      .b() // return type: FooInterface
      .c() // return type: FooInterface
      (); // return type: { ... }

I've been rewriting the types file over and over again, trying to navigate this problem, brushing off the error about `FooInterface` lacking a call signature since I believed that you can't just add an unlabelled call signature to an interface like that. But then along the way I found [this answer on SO](https://stackoverflow.com/a/57426632), where the author writes this:

    interface MyFunc {
      (arg: string): string //a call signature
      funcName: (arg: string) =&gt; string //express of function in object literal
    }

And in seeing the first example of a call signature, the bare-bones unlabelled one, I realised that there was a way to solve that problem, where there is the final set of parentheses which work as the interface's *own* call signature. So I put it into my file and it works, 100%. A nice trick that I don't think many people even consider since it feels like it should be illegal to have an unnamed method like that.

What cool tricks have you found lately, /r/typescript?
