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
## [2][Guides for truly understanding generics in TS](https://www.reddit.com/r/typescript/comments/josnlh/guides_for_truly_understanding_generics_in_ts/)
- url: https://www.reddit.com/r/typescript/comments/josnlh/guides_for_truly_understanding_generics_in_ts/
---
I come from a background where I never had to deal with generics, so they're still quite a foreign concept to me.

Does anyone know a good guide/tutorial on generics, from zero to advanced? I really want to understand this part of TS so I can have a more complete experience with it.

Thanks
## [3][How do I add proper typing to this function with Object.entries](https://www.reddit.com/r/typescript/comments/joxmf9/how_do_i_add_proper_typing_to_this_function_with/)
- url: https://www.reddit.com/r/typescript/comments/joxmf9/how_do_i_add_proper_typing_to_this_function_with/
---
Here is the [demo](https://www.typescriptlang.org/play?#code/JYOwLgpgTgZghgYwgAgLIE8DCiAWKDeAUMsnAFzJEkkBGFV1JCFAzmFKAOYDaAugNzFGyACat2XQYwC+U5LMLTCCAPYg2yMDi4sAKioBKEALYqAbigC8ybgCIwhk+YgBGWwBpk9x6YsAmWwFlNQ0EXAgKDGwEPGRrBnJKIVp6ZOpmG28jX1cPLwcABSgIFmgLN08sp3887yKSsogA3nc0kjEvYChigBsIMzhwWzSlEiUFQjB0AAcUVDhp2ZFdKEGWGBUoYwAeXWQIAA9IEBEWSmkAPjikkm4AaWRQZABrCHQVGGRdXgB+SIWlis1hstrt7rwrgAfZAAChh0zgq2MLAounBAEo4ldBuh0YIJjAAK4gBBgYBqTSrCxQUrbAwXGFhGIRZAGTzsYGbYwAeRoACsIKT-osIMtVuoQTt6ZiGMUwISoCBkLyBaSAHQQcAcEowjkSrkqwVgdFq4oiQlIGFmi0QIH6rYAORUIggniZeDxikIhAA9D7kDgVAB3UQqZAASU0MxQWmAZ0AvBuAcR3guowMhrUg7esuU6XddGeFc66bHrs1s7m92eKy8YAGpwHqEiC8TGWK4MVSpygl6uSivoXgUEDOlDSa7uiBFuSdjQgCBB6J4euNqxpKazD6UzlbZdNuKWay2IkkslqYbCZA-Lf2usNpsw4cu9FtZAUDnU0q63tc3fFx8QT0hDlBUlQYEg1Qgici1aRhuFLPs3kHZA5wXcJfyEBQJnfaBaSicIGQnTwElSRg6BuYQMgfEdWyuf81RgYAekgKA4TYDgQE4GjkAAQljDi9B8Zw1VABBGxdFgYTYrh0WfGQMMUT0gA)

I am trying to type this function 
```js
const reduceTransformNode = (cacheNode, [transformKey, transformValue]) =&gt; {
  const { [transformKey]: node } = cacheNode;
  const newCacheValue =
    typeof transformValue === "function"
      ? transformValue(node)
      : traverse(transformValue, node);

  return {
    ...cacheNode,
    [transformKey]: newCacheValue
  };
};
```

I cannot seem to solve it because there seems to be a circular dependency between `traverse` and `reduceTransformNode` 

This is one [solution](https://www.typescriptlang.org/play?#code/JYOwLgpgTgZghgYwgAgLIE8DCiAWKDeAUMsnAFzJEkkBGFV1JCFAzmFKAOYDaAugDTFGyACat2XQYwC+Q2bIQB7EG2RgcXFgBVFAJQgBbRQDcUAXmTcA5GD2GTEAIxX+yG3aOmATFd4BuQiUVMGQEXAgKDGwEPGQLBnJKIVp6ZGTqZkt3fU8nFzdbAAUoCBZoU2dXbPtvfPdi0vKIHwE04VEKACJgKBKAGwhjOHBOtuppKWR5QkIwdAAHFFQ4ecWRLShhlhhFKAMAHi1kCAAPSBARFkppAD44pJJuAGlkUGQAawh0RRhkLV4APyRFZrDZbHZ7Q7PXh3AA+yAAFAj5nBNgYWBQtNCAJRxO7DdDYwiyQgwACuIAQYGAymQJREZKQYJUEIMADlFCIIIdjmcIBcrvhLJ90OIOCBOLwKASpjcEWEYhAOVzMa5uOxwbsDE8vq4NSytQA1OB9MkQKWWNjiziuAm8bGYh6hZSqIXqzYGvY69AWkCclDSe4KvDKiABJgukIgCAAd2ieGNpvM6Tmix+ag92yNJrNcTMFk65Mp1OUnXSJABGc1e0TZoRfq5RPaFA1pigZQR+qzNZzEFcDYg2ICQhKYDJUBATpIADpZ8Glf7Jo8u6zvb7Y-GILWIHIAtJhzMi1SaZPW9AyvtdHL52RdHrM6yAPI0ABWECpwNWEHWD61l5uDrILoTqjuOk7Pm+VLTvyEilJ2v57BB75gNi070oyEAIuhTIIeyi6hOEQ7EjMQSqAY6DMt2BhIR+aAgt+lGsvsUThHc8RCIkDApFOGQUPW-q4mYdwDtOMDAH0kBQAiVp4sgACE6iaDoOQONOoAIKaXIsNJ7C4k24xyMRgSRnSpRkhJ9xnu23IsYq17hK45GMVqNEocZKiKAM059IonBYWZElEoQQA) I found that would work but not ideal. 

```ts
function reduceTransformNode&lt;T extends { [key: string]: any }&gt;(cacheNode: T, [transformKey, transformValue]: [string, any]): T {
  const { [transformKey]: node } = cacheNode;
  const newCacheValue =
    typeof transformValue === "function"
      ? transformValue(node)
      : traverse(transformValue, node);

  return {
    ...cacheNode,
    [transformKey]: newCacheValue
  };
};
```

I don't want to have `any` in the code. Can anyone give this trick TypeScript problem a try
## [4][Help: Mapped Type modifiers are being lost](https://www.reddit.com/r/typescript/comments/joyh92/help_mapped_type_modifiers_are_being_lost/)
- url: https://www.reddit.com/r/typescript/comments/joyh92/help_mapped_type_modifiers_are_being_lost/
---
I'm trying to make a mapped type based on two given types. I want the mapped type to have all keys that two given types have but respect and modifiers on those keys.

My first attempt at such a type was:

    type Merge1&lt;T1, T2&gt; = {
        [K in keyof T1 | keyof T2]: SomeType;
    }

But the modifiers obviously become lost in this.

I then tried:

    type Merge2&lt;T1, T2&gt; = {
        [K in keyof (T1 &amp; T2)]: SomeType;
    }

And this works in almost every single case. Yay for almost.

The case were it doesn't work is with const assertions and I have no idea why. I initially thought the reason why was because `"a" &amp; "b"` is `never` but `string &amp; number` should also be `never` and it works fine in this case.

[Here's a playground link](https://www.typescriptlang.org/play?#code/PQ18ZXTUFAgAQBUCeAHApogkgcgFtEAXAJ1QEsA7AcxIHtEATTAYwBsBDUzAOjjjEM2ALKZSNTAEYAPMikAaFACYAfIgC8iAN5xE+xAG0A0omqIA1plT0AZiimIAPpet2VAXQBciAMr0CTDQsAG44AF8BAxJhRDEJTGU5RRV1LV1o6JMzKlcbewAKeUQAMhUASm8-AKDhMOjIhFhmltbmpsQAEQBXAgJUFFiAWjNiPABnRCp6YjMCdHpSYk4qYn5BWP9A4Ox0xE4fbqoLaYB3XPCwprab25aO5Exx4mo6HfGBISxEAA1NHTgAEgeJwmPQqOwBrZ6PQfM9SK8FHoDAAjbhwsiIiJXL7YACa-10gOhsKmvRR4iR0TRpAA-BiEbQkZdPrEfqxwc9CUCQWCIVCYT4AEScFGsIVU1HoxBCli2CXY1nfPEcqhc9JAkk+ACskv0NPpMpoAAsKAqWdc7lbraAHk9iB8OiJ6EwKLZxJNuNh2PRnvxcSh7VIAIL-eKSWQ-JR41RhAOPZ5SABCYfEEZk7M5xGjquesYESAA6osLJMOaQeKxiJD-bEE8RlKGtOHEhno-mOp16E8qHhZqcS7XvvXlCnm2nW5m1dnECqs7HEEh0o0iyWy4tK9XUEPsCOAMKphJJKfPJQn4j5+P25SdQ+SJJz6c5+dhIA) so you can see what I'm talking about.

Anyone have any idea on how I can fix my type? Or is this a bug in TypeScript itself?

Edit: Probably not the most accurate title now that I think about it, but oh well. If you're reading this then it got the job done.
## [5][Strict Mode without config file?](https://www.reddit.com/r/typescript/comments/jotkaz/strict_mode_without_config_file/)
- url: https://www.reddit.com/r/typescript/comments/jotkaz/strict_mode_without_config_file/
---
In VSCode, can I force a TypeScript file that exists outside of TypeScript project to be interpreted in strict mode?
## [6][How do you initialize an "empty" map?](https://www.reddit.com/r/typescript/comments/jopxz2/how_do_you_initialize_an_empty_map/)
- url: https://www.reddit.com/r/typescript/comments/jopxz2/how_do_you_initialize_an_empty_map/
---
I have a type `type MyMap = Map&lt;string, string&gt;` .

* When I attempt to implement this type as `const defaultMap:MyMap = new Map()` , I get an error `Only a void function can be called with the 'new' keyword.ts(2350)`
* I changed it to `Map()` and I get `Type 'Map&lt;unknown, unknown&gt;' is missing the following properties from type 'Map&lt;string, string&gt;': [Symbol.iterator], [Symbol.toStringTag]ts(2739)`
* I changed it to `Map&lt;string, string&gt;()` and I got the same error as above
## [7][Short syntax for implicit type declaration with optional properties?](https://www.reddit.com/r/typescript/comments/joigd5/short_syntax_for_implicit_type_declaration_with/)
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
## [8][I made a VSCode plugin that does GoTo Definition faster than VSCode/TSServer](https://www.reddit.com/r/typescript/comments/jo3z86/i_made_a_vscode_plugin_that_does_goto_definition/)
- url: https://marketplace.visualstudio.com/items?itemName=verydanny.smart-goto
---

## [9][Define class type. How to separate the types for instance from the types of the prototype?](https://www.reddit.com/r/typescript/comments/jo6zje/define_class_type_how_to_separate_the_types_for/)
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
## [10][How can I declare a function declaration's type via a type alias?](https://www.reddit.com/r/typescript/comments/jo6x98/how_can_i_declare_a_function_declarations_type/)
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
## [11][import json possible while using esnext out?](https://www.reddit.com/r/typescript/comments/jo66pn/import_json_possible_while_using_esnext_out/)
- url: https://www.reddit.com/r/typescript/comments/jo66pn/import_json_possible_while_using_esnext_out/
---
Hi all!  Working on a thing that I want to use the latest everything possible, latest node, latest browsers, most advanced settings, etc.  Want everything to be *modern*.

I've got my typescript settings for module and target as "esnext", and it does generate code, but when I try to add an import of a json file, node throws me "Unknown file etension json for (filepath to json file)".

Is there some way to get there?  I wanted to do it via typescript import, because I want it to be both node and browser compatible.. 

I am willing to turn back my target settings if necessary, I guess...
