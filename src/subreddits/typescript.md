# typescript
## [1][Who's hiring Typescript developers - February](https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/)
- url: https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/
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
## [2][Webpack/TypeScript/React starter kit as of 2020](https://www.reddit.com/r/typescript/comments/f866km/webpacktypescriptreact_starter_kit_as_of_2020/)
- url: https://www.reddit.com/r/typescript/comments/f866km/webpacktypescriptreact_starter_kit_as_of_2020/
---
Hey everyone. I know that there are tons of TypeScript starter kits but decided to share mine.

[https://krasimirtsonev.com/blog/article/beginning](https://krasimirtsonev.com/blog/article/beginning)

I did it because it's the very minimum setup and can be installed super quickly with just `npx beginning`.
## [3][Variance question within typescript generics.](https://www.reddit.com/r/typescript/comments/f81r51/variance_question_within_typescript_generics/)
- url: https://www.reddit.com/r/typescript/comments/f81r51/variance_question_within_typescript_generics/
---
As I continue down the typescript rabbit hole, I've stumbled across something that's puzzling me. I'm used to the `extends` keyword, in Java, signaling that a generic type is covariant. `List&lt;T extends Number&gt;`, for example, would let me know that I can safely get `Number`s from the `List`, but the list might contain any type of `Number`. 

I've run across something that appears to be backwards. Here is a [playground](https://www.typescriptlang.org/v2/en/play?#code/MYewdgzgLgBAJgQygmBeGBvAUDGAiACwFMAbEkPALkx13wDcESBXIq-AdxACcS48ANLQC+Q3HgBmICtWx0GTVuzwAjBN0EihwrAHpd+YmQowAPvikUsUAJ4AHIjADSYEBzBOiNiGhgBrLxAJGFsHIPgkBCw9AzhmAFsVGABLMCgibgkEYCIsIgAPOx5YVPTM7McAERAAZShmCWC5GAlUuAAhG08bGAAeABUYAvSwOB9oblSAcwA+GAAKAJtqfoBKagnpgG4sHRj4ECIIMAByWC5uPwB+EIJknwQ7B3UfKBAYFUdQNO4ERkmEGkUmAYINUtAiAg4DBwqB-slAVABDASMkAjAAJIwLjMPhDQpEYBQAB0eUKxRgwBICAgPmqdQaEgx8TsJBSLJIRHiRDSdNq9UaNFwrVGnW6iy81Bcbg8Xgg6xgmzAUyFdG4RHq3BBiGQAG0lgBdYmMFi5XA6Pb6eAJJKlDJZHJkorcEppe0VGD0gUSACCUwQqVVIo6XS8fUGwx5Y2crnc3Qgcwly1BCqVUx2loMFz8ECuTopVJpfIZjT9AbAzNZ7NZXJ5UGL3qDbTFXiTGygk2VqY701VuHVmpBSxEuyAA) link with what's puzzling me. Could someone tell me how I'm reading this wrong? Thanks.

edit: updated [playground](https://www.typescriptlang.org/v2/en/play?#code/MYewdgzgLgBAJgQygmBeGBvAUDGAiACwFMAbEkPALkx13wDcESBXIq-AdxACcS48ANLQC+Q3HgBmICtWx0GTVuzwAjBN0EihwrAHpd+YmQowAPvikUsUAJ4AHIjADSYEBzBOiNiGhgBrLxAJGFsHIPgkBGt7RxAoYm4AZShuAEswAHMfdAByI3Icsxgcrl44HKw9AzhmAFsVGHSoIm4JBGAiLCIADzseWCaWto6YABEQZOYJYLkYCXS4ACEbTxsYAB4AFRge5rA4H2g0zIA+GAAKAJtqTYBKaiP0jIBuLB0q+BAiCDAc2FK-AB+EIEVI+BB2BzqHxQEAwFSOUBgFIIRhpBDIxpgGDbdLQIgIOAwcKgNGpDFQAQwEipAIwACSMC4zD4O16RGAUAAdF1ev0YMASAgID5xpNpvTanYSI0pSQiLUiMjRRMoFMZrR5vtlqsLldqC43B4vBB7jBHpkaPJuEQ1dxsYhkABtK4AXS5jBYnVwOne+nNUFSZE+31+-x4fl5fW4sEFwpV4okkulACZZdKFUqoAm1dMrXMFjqvHqvNQ4glkscsmaLRl87gbXbsVcRG9Kv7yy0sc1Wu1Oj1owNkUM+2NVeqAIIZBDpfNapYrYtbNl7A7OVzuVYQM6XUs4mspJ6vP0GAE+FTMWDxMHmogK8-tPwcdQHKP8uMiseJqczsDJmWpHKmbKl+uYarg85FjYu7XAGVYHlW9YwI2zD2v4XitieITcDYTwwAg06zhwqTxN2I4jBiRIfhAWCDL2IyLOo+YIJWTwPIemTHpU1EwIxABe6byoqIGMdwzGsZa6B4BAICKtemR4G2tHDvRjjjAAsjwRCJvmcAgAA6hGy67Eqa6iSc5yeqwNxmpsXEfNePhngWYBEFy7lYDxGlaYm-66QZEaWYoRDUPxZr8UhKFoVZ3owL6QA) link with more examples.
## [4][A naive approach to functional programming using TypeScript](https://www.reddit.com/r/typescript/comments/f7pr4k/a_naive_approach_to_functional_programming_using/)
- url: https://github.com/alohawav/naive_functional_programming
---

## [5][Calling a function generically by name... not sure why this doesn't work](https://www.reddit.com/r/typescript/comments/f7ua54/calling_a_function_generically_by_name_not_sure/)
- url: https://www.reddit.com/r/typescript/comments/f7ua54/calling_a_function_generically_by_name_not_sure/
---
https://www.typescriptlang.org/play/index.html?target=99&amp;module=1&amp;ts=3.8-Beta#code/JYOwLgpgTgZghgYwgAgGIHt3IN4ChkHIBGcUAXMgBQAeFIArgLZHQCUyAvAHzIDOYUUAHN8hEgC8KlAJ4V+gkEPbdkDZtFwBfXLhj0QCMMHQhkMADwBpZBGqQQAE17IA1hGnoYaTF0oxMFBjoADTIjBBgABboDhSWoQB0SQAOpHCMvBQACmnhkFC85kEA2pYAulysFABu6MAOOKIE-ujFedEOZZRJCalQ6bysANxaQA

Any ideas why this doesn't work? Reproducing the code below:

```
interface Foo {
  bar: (x: number) =&gt; string
  baz: (y: string) =&gt; number
}

function f&lt;K extends keyof Foo&gt;(
  foo: Foo, 
  method: K, 
  ...params: Parameters&lt;Foo[K]&gt;): void {
    foo[method](...params);
}
```

The TypeScript error is: `Expected 1 arguments, but got 0 or more.`

Should be fairly obvious what I'm trying to do here, but I want to be able to pass a `Foo` object in and call it's method `method` with the correct parameters.
## [6][How not to trick TypeScript compiler and not be tricked by it](https://www.reddit.com/r/typescript/comments/f7c0x4/how_not_to_trick_typescript_compiler_and_not_be/)
- url: https://indepth.dev/how-not-to-trick-typescript-compiler-and-not-be-tricked-by-it/
---

## [7][redturn.ts - Event-based/Queueing redis mutex lock in typescript](https://www.reddit.com/r/typescript/comments/f7inmh/redturnts_eventbasedqueueing_redis_mutex_lock_in/)
- url: https://github.com/kevinwilson541/redturn.ts
---

## [8][Location.state typing with React-Router](https://www.reddit.com/r/typescript/comments/f7fk08/locationstate_typing_with_reactrouter/)
- url: https://www.reddit.com/r/typescript/comments/f7fk08/locationstate_typing_with_reactrouter/
---
Can anyone help me out with using react-router with Typescript? See this link: [https://stackoverflow.com/questions/60343598/react-typescript-how-add-types-to-location-state-when-passed-in-a-route](https://stackoverflow.com/questions/60343598/react-typescript-how-add-types-to-location-state-when-passed-in-a-route)
## [9][Announcing TypeScript 3.8](https://www.reddit.com/r/typescript/comments/f71hpa/announcing_typescript_38/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-8/
---

## [10][Typing the Technical Interview in TypeScript: solve an interview with no runtime code!](https://www.reddit.com/r/typescript/comments/f6v4np/typing_the_technical_interview_in_typescript/)
- url: https://gal.hagever.com/posts/typing-the-technical-interview-in-typescript/
---

## [11][AdminBro v2 beta is out!!!](https://www.reddit.com/r/typescript/comments/f6szkd/adminbro_v2_beta_is_out/)
- url: https://www.reddit.com/r/typescript/comments/f6szkd/adminbro_v2_beta_is_out/
---
AdminBro is an automated admin interface for Node.js apps based on React, written in TypeScript = Django admin for node on steroids !!!

Now, we almost finish the entire scope for v2.0 and we released it along with the demo version: 

Demo page can be found here:

[https://admin-bro-example-app-staging.herokuapp.com/admin](https://admin-bro-example-app-staging.herokuapp.com/admin)

Github repo is here: [https://github.com/SoftwareBrothers/admin-bro](https://github.com/SoftwareBrothers/admin-bro)

and you can install it using: 

```
yarn add admin-bro@2.0.0-beta.26
```

In this version: 

* there is an entire Design System written specially for AdminBro
* i18n support
* hooks support
* the entire UI change to Drawer'like exploring resources

Now we are working on the documentation which (WiP) can be fond here:

[https://softwarebrothers.github.io/admin-bro-dev/v2/](https://softwarebrothers.github.io/admin-bro-dev/v2/)

Let me know what do you think - and star the repo if you like the idea :)
