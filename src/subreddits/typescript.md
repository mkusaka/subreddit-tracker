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
## [2][Not quite what I was looking for](https://www.reddit.com/r/typescript/comments/iyur8c/not_quite_what_i_was_looking_for/)
- url: https://twitter.com/FarPixel/status/1309077814117904389
---

## [3][types vs. interfaces](https://www.reddit.com/r/typescript/comments/iylqrr/types_vs_interfaces/)
- url: https://www.youtube.com/watch?v=IXAT3If0pGI
---

## [4][What’s Structural Typing And How Typescript Uses It To Its Benefit?](https://www.reddit.com/r/typescript/comments/iywrt5/whats_structural_typing_and_how_typescript_uses/)
- url: https://medium.com/@mahdighajary/whats-structural-typing-and-how-typescript-use-it-to-its-benefit-6cf491f1564?source=friends_link&amp;sk=e0265b7bf135cc7e5d912d9e74d57afb
---

## [5][TS4.1: Is there a way to define this property path type that isn't excessively deep?](https://www.reddit.com/r/typescript/comments/iywewf/ts41_is_there_a_way_to_define_this_property_path/)
- url: https://www.reddit.com/r/typescript/comments/iywewf/ts41_is_there_a_way_to_define_this_property_path/
---
Experimenting with the upcoming TypeScript 4.1's template literal types, I tried to define a generic type that can check property paths.

Until TS 4.1, there was no possible way to type an expression such as 'foo.bar.baz', and you would have to settle for string. Now, with template literal types, I want to be able to type these property paths, and use them for things such as MongoDB queries and projection objects. For example:

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
## [6][What’s Structural Typing And How Typescript Use It To Its Benefit?](https://www.reddit.com/r/typescript/comments/iywe8t/whats_structural_typing_and_how_typescript_use_it/)
- url: https://medium.com/@mahdighajary/whats-structural-typing-and-how-typescript-use-it-to-its-benefit-6cf491f1564?source=friends_link&amp;sk=e0265b7bf135cc7e5d912d9e74d57afb
---

## [7][Need some help with index signature and OR](https://www.reddit.com/r/typescript/comments/iyh08i/need_some_help_with_index_signature_and_or/)
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
## [8][Why does setting 'as const' allow a string to match a string union type?](https://www.reddit.com/r/typescript/comments/iy8b7t/why_does_setting_as_const_allow_a_string_to_match/)
- url: https://www.reddit.com/r/typescript/comments/iy8b7t/why_does_setting_as_const_allow_a_string_to_match/
---
I had to use this inside a fetch request, I consider it hacky but it works:

        mode: 'cors' as const, // allows string to match a string union type

I learned this here but never found an explanation for it. why does asserting the string as a const variable suddenly let it match a string union type (which I saw in the error popup)?
## [9][Discriminated unions with functions](https://www.reddit.com/r/typescript/comments/iyc8fj/discriminated_unions_with_functions/)
- url: https://www.reddit.com/r/typescript/comments/iyc8fj/discriminated_unions_with_functions/
---
I am working on trying to narrow some typing in a library that needs to take a user defined function. I only want to support a narrow scope of function definitions, and I'm hoping to get help from the compiler. I was able to strip things down to the smallest example of the issue I'm seeing here:

    type ValidateString = (value: string) =&gt; boolean;
    type ValidateNumber = (value: number) =&gt; boolean;
    type ValidationType&lt;T extends Validations&gt; = 
        T extends ValidateString ? string 
        : T extends ValidateNumber ? number
        : never
    type Validations = ValidateString|ValidateNumber
    function validate&lt;T extends Validations&gt;(validation: T, value: ValidationType&lt;T&gt;) {
        return validation(value) // &lt;-- Error
    }

The error explained under `value`:

    Argument of type 'ValidationType&lt;T&gt;' is not assignable to parameter of type 'never'.
      Type 'string | (T extends ValidateNumber ? number : never)' is not assignable to type 'never'.
        Type 'string' is not assignable to type 'never'.
          Type 'string | number' is not assignable to type 'never'.
            Type 'string' is not assignable to type 'never'.(2345)

Ultimately this is part of a bigger library where the library will receive a "schema" with attributes. Each attribute could optionally have a "validation" property, which would include a user defined function to validate the attribute. If you know of any examples in other applications that does this, or something similar, that would also be helpful.
## [10][[Question] Letting HR people using my fullstack app](https://www.reddit.com/r/typescript/comments/iycznf/question_letting_hr_people_using_my_fullstack_app/)
- url: https://www.reddit.com/r/typescript/comments/iycznf/question_letting_hr_people_using_my_fullstack_app/
---
**TLDR:** Im making a full-stack app (some kind of a very simple CRM system). I want to show to HR people how my app looks like from inside but in the same time people in the company I work at use this app. How to let HR people use my app in a secure and easy way without showing them real data?

***Full story:***  
I'm building a full-stack app for the small company that I am currently working at. It will be some kind of a very simple CRM.

Employees have to manage their sales. This app will help them. It provides CRUD operations with those sales, show statistics and many more nice little features.  
This app will be really used by the company so it will store sensitive data in the database.

I want to show how it works from inside, but I know that HR people won't clone and run the repo, so I need a feature like "LOGIN AS GUEST".  
Guest can look around whole app, but I can't show them any data inserted by my company.

So I have a few ideas (sorted from best one to worst one in my opinion) and I need your advice:

1. Make a completely different deployment for hr people and different one for the company
2. Make a seperate connection for guest users, they will connect to completely different DB where they can CRUD but they will connect with the same server (same deployment).
3. Don't make guest user feature, just record a short nice video how things look, maybe put it on the landing page with some informations.
4. Allow guests to use the app, connect them to the same database my company use but check in backend if user has role of a guest they cannot CRUD anything or they get their own tables in the database and they can CRUD. I don't feel safe when thinking about this approach

I really want to give a good view on that app as I feel it will be a game-changer for my career and maybe I will land my first job as developer.

Please let me know what you think, maybe you've got some other ideas?
## [11][Wrapping Utility Types with typeof Internally](https://www.reddit.com/r/typescript/comments/iygdqq/wrapping_utility_types_with_typeof_internally/)
- url: https://www.reddit.com/r/typescript/comments/iygdqq/wrapping_utility_types_with_typeof_internally/
---
Struggling to figure something out and now I'm  questioning if it's even possible. In a generator functions, the yield keyword still isn't well supported so I'm trying to wrap the utility types I've been using.

Instead of:

    const data: ReturnType&lt;typeof myFunction&gt; = yield myFunction();

Something like:

    type YieldReturn&lt;F&gt; = ReturnType&lt;typeof F&gt;;
    
    const data: YieldReturn&lt;myFunction&gt; = yield myFunction();
