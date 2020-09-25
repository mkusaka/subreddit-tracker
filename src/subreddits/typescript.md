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
## [2][Can anyone show me why react-typescript is good?](https://www.reddit.com/r/typescript/comments/izixhl/can_anyone_show_me_why_reacttypescript_is_good/)
- url: https://www.reddit.com/r/typescript/comments/izixhl/can_anyone_show_me_why_reacttypescript_is_good/
---
There is this company I am interested in working for. But they insist on using react in combination with typescript. 

I followed some tutorials about the combination of the two but I honestly feel like typescript is bloating the syntaxes more than it should. Some comments could also do the trick in my opinion. That is because react already offers 'a form' of structure to its code. 

This said, I am stil pretty inexperienced when it comes to developing bigger projects. That is why I would like to ask some people woth experience to show me some code examples where typoscript obviously make react code better. And/Or draw me a situation where it would obviously help. 

Thanks in advance!
## [3][Not quite what I was looking for](https://www.reddit.com/r/typescript/comments/iyur8c/not_quite_what_i_was_looking_for/)
- url: https://twitter.com/FarPixel/status/1309077814117904389
---

## [4][Passing a union to an overloaded function](https://www.reddit.com/r/typescript/comments/izgrsx/passing_a_union_to_an_overloaded_function/)
- url: https://www.reddit.com/r/typescript/comments/izgrsx/passing_a_union_to_an_overloaded_function/
---
The following code doesn't type-check unless the third function declaration is uncommented. I feel like I'm missing something here; shouldn't it be rather simple for the compiler to figure out the return type when the union is provided without me having to explicitly define it?

    type Foo = {type: "foo"};
    type Bar = {type: "bar"};

    declare function f(x: Foo): 1;
    declare function f(x: Bar): 2;
    // declare function f(x: Foo | Bar): 1 | 2;

    declare const input: Foo | Bar;
    f(input);

[Relevant TypeScript Playground](https://www.typescriptlang.org/play?#code/C4TwDgpgBAYg9nKBeKBvUkBcUBEAzBHAXwG4AoDaAIQEMAnZNS7HAI3uPLIBMIBjADb1oeAK4A7PsACWccVDwAKAB7Z4cAJTYAjOV6DhCiVNnylqqLTpaoAJnIB6B1H1C6I4zLkKVahFAAfS3obbUC7LldDPjkAZ2AoaXEwUWA-RCCrciUklOANEiA)
## [5][The first thing I do on any Typescript/Angular project](https://www.reddit.com/r/typescript/comments/izghe1/the_first_thing_i_do_on_any_typescriptangular/)
- url: https://medium.com/@ciesielskico/the-first-thing-i-do-on-any-typescript-angular-project-628d26ffd2a7?source=friends_link&amp;sk=5eed949a2c5e30c48709fbcafad1d00c
---

## [6][tsPEG: 2.0.0 now available! The TypeScript parser generator](https://www.reddit.com/r/typescript/comments/iz55ud/tspeg_200_now_available_the_typescript_parser/)
- url: https://www.npmjs.com/package/tspeg
---

## [7][types vs. interfaces](https://www.reddit.com/r/typescript/comments/iylqrr/types_vs_interfaces/)
- url: https://www.youtube.com/watch?v=IXAT3If0pGI
---

## [8][TS4.1: Is there a way to define this property path type that isn't excessively deep?](https://www.reddit.com/r/typescript/comments/iywewf/ts41_is_there_a_way_to_define_this_property_path/)
- url: https://www.reddit.com/r/typescript/comments/iywewf/ts41_is_there_a_way_to_define_this_property_path/
---
Experimenting with the upcoming TypeScript 4.1's template literal types, I tried to define a generic type that can check property paths.

Until TS 4.1, there was no possible way to type an expression such as `'foo.bar.baz'`, and you would have to settle for `string`. Now, with template literal types, I want to be able to type these property paths, and use them for things such as MongoDB queries and projection objects. For example:

`db.someCollection.find({ 'foo.bar.baz': { $exists: true } });`

This is the type that I came up with:

```ts
type PropsPath&lt;T&gt; =
    T extends object
        ? T extends any[]
            ? number
            : {
                [P in keyof T]: P | `${P}.${PropsPath&lt;T[P]&gt;}`
            }[keyof T]
        : '';
```

[Full example in TS Playground](https://www.typescriptlang.org/play?ts=4.1.0-dev.20200921#code/C4TwDgpgBACgTgezAZxgQ2ACwDwBUB8UAvAFBTlS5QQAewEAdgCbJQIBGAVhAMbBkVBAfkrU6jFlDQMQAbQC6AwcqgiGAVwC27CHCUqKALigBvfQeWyYUAJYMoAawggEAM0rzj1gD5QABgAkJjAAvgB0QfBIqBg4uFby+CF+5gYhsk4u7riKFsYA5PkA3CQkdvRwrmg80ADKCJoQuODQZoKuCAjGbQbsaHDGGtq6JYIh+n0AXsbsnQA2ENIl4yQ8CAzIwFBgXogo6FjY9Y3NkIREUPkdCMWlaxtbYIjcfDbrXv3ANmhz2ABKvAQcCY2Ci+1iRwaTRa+AANFAhjo4PhzqZ9FdOmE+nB8sYAIwkEJFIA)

Sadly, this type is considered "excessively deep or possibly infinite" by TS compiler. Is there any way to redefine it in a way that doesn't throw an error?
## [9][What’s Structural Typing And How Typescript Uses It To Its Benefit?](https://www.reddit.com/r/typescript/comments/iywrt5/whats_structural_typing_and_how_typescript_uses/)
- url: https://medium.com/@mahdighajary/whats-structural-typing-and-how-typescript-use-it-to-its-benefit-6cf491f1564?source=friends_link&amp;sk=e0265b7bf135cc7e5d912d9e74d57afb
---

## [10][Need some help with index signature and OR](https://www.reddit.com/r/typescript/comments/iyh08i/need_some_help_with_index_signature_and_or/)
- url: https://www.reddit.com/r/typescript/comments/iyh08i/need_some_help_with_index_signature_and_or/
---
Spent a lot of time trying to find an explanation of the following behavior:

    const stringKey = 'thisKeyDoesNotExist'
    const hashMap: {
        [key: string]: {
          valueNumber: number; 
          valueString: string;
        };
      } = {};
    
    
    // TS does not complain
    hashMap[stringKey] = hashMap[stringKey] || {
              valueString: stringKey,
              valueNumber: 'string!!!!!!!!!!!',
            };
    // TS does not complain
    hashMap[stringKey] = hashMap[stringKey] || "string";
    
    // Only with this syntax TS complains that there is an issue with `valueNumber`
    hashMap[stringKey] = hashMap[stringKey] ? hashMap[stringKey] : {
              valueString: stringKey,
              valueNumber: 'string!!!!!!!!!!!',
            };

Can somebody explain why the first 2 assignments are passing the TS checks?
## [11][Why does setting 'as const' allow a string to match a string union type?](https://www.reddit.com/r/typescript/comments/iy8b7t/why_does_setting_as_const_allow_a_string_to_match/)
- url: https://www.reddit.com/r/typescript/comments/iy8b7t/why_does_setting_as_const_allow_a_string_to_match/
---
I had to use this inside a fetch request, I consider it hacky but it works:

        mode: 'cors' as const, // allows string to match a string union type

I learned this here but never found an explanation for it. why does asserting the string as a const variable suddenly let it match a string union type (which I saw in the error popup)?
