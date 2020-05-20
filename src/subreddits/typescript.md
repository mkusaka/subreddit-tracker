# typescript
## [1][Who's hiring Typescript developers - May](https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/)
- url: https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/
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
## [2][Using JSDoc tags to test functions [Prototype]](https://www.reddit.com/r/typescript/comments/gn6sit/using_jsdoc_tags_to_test_functions_prototype/)
- url: https://i.redd.it/8gqa8evpfvz41.gif
---

## [3][Fcaljs - extensive math expression evaluator library](https://www.reddit.com/r/typescript/comments/gn7ves/fcaljs_extensive_math_expression_evaluator_library/)
- url: https://www.npmjs.com/package/fcal
---

## [4][Deno: A Simple Guide](https://www.reddit.com/r/typescript/comments/gnadh0/deno_a_simple_guide/)
- url: https://www.martinmck.com/posts/deno-a-simple-guide/
---

## [5][The Great Dropbox CoffeeScript to Typescript Migration](https://www.reddit.com/r/typescript/comments/gmll5p/the_great_dropbox_coffeescript_to_typescript/)
- url: https://dropbox.tech/frontend/the-great-coffeescript-to-typescript-migration-of-2017
---

## [6][Destructuring and sum types](https://www.reddit.com/r/typescript/comments/gn7oz6/destructuring_and_sum_types/)
- url: https://www.reddit.com/r/typescript/comments/gn7oz6/destructuring_and_sum_types/
---
I had typed array with an object with one optional field. Then I decided that instead of an optional prop what I really want is an array of objects not containing that property or all the elements containing such property, so I switched to a sum type.

&amp;#x200B;

The problem is that this makes destructuring of such prop impossible.

Here is an example of the sum type:

&amp;#x200B;

```js

type tab&lt;T&gt; =

| { label: string }

| { label: string, value: T }

```



This is what I can not do:


```
    tabs.map(({ label, badge, value }, idx) =&gt; {

        callSomeFun(value || idx)

    })
```


This works for flow, but not for typescript:


```
    tabs.map(({ label, badge, ...tab }, idx) =&gt; {

        const value = tab.value || idx

        callSomeFun(value)

    })
```


Also because of this, some other parts of the code get an inferred type of mixed, which in runtime is not correct because I'm just "or-ing" them:



const \[activeTab, setActiveTab\] = useState(tabs\[0\].value || initialTab)



activeTab becomes of mixed type



Is there a better way of doing this while ensuring the type safety of having the prop on all items or no one?
## [7][Pass Paramus to server side props](https://www.reddit.com/r/typescript/comments/gmub97/pass_paramus_to_server_side_props/)
- url: https://www.reddit.com/r/typescript/comments/gmub97/pass_paramus_to_server_side_props/
---
I wonder if anyone knows if it is possible to send arguments to getServersideProps in Next Js.

At the end of my URL I have a hard codded page number, that fetches different movies per page, but I want to have it dynamic with getServersideProps, is it possible?


export const getServerSideProps: GetServerSideProps = async (
  ctx: GetServerSidePropsContext&lt;ParsedUrlQuery&gt;
) =&gt; {
  console.log(ctx.query)
  const response = await fetch(
    `https://api.themoviedb.org/3/movie/upcoming?api_key=${process.env.API_KEY}&amp;language=en-US&amp;page=1`
  )

  const data = await response.json()

  return {
    props: {
      data: data.results,
    },
  }
}
## [8][A GitHub template featuring automation that makes it pretty darn easy to publish and maintain high-standard NPM modules.](https://www.reddit.com/r/typescript/comments/gma48m/a_github_template_featuring_automation_that_makes/)
- url: https://github.com/garronej/ts_ci
---

## [9][Is ES5 dead yet?](https://www.reddit.com/r/typescript/comments/gmr38p/is_es5_dead_yet/)
- url: https://www.reddit.com/r/typescript/comments/gmr38p/is_es5_dead_yet/
---
I've been packaging all my modules lately with ES2015 as the target.

Is there any reason to even think about ES5 anymore?

I mean we're all moving WASM now anyway right? :P
## [10][Hegel: An advanced static type checker for JavaScript](https://www.reddit.com/r/typescript/comments/gm7ut0/hegel_an_advanced_static_type_checker_for/)
- url: https://twitter.com/devongovett/status/1262069817655259136
---

## [11][When are parenthesis not redundant?](https://www.reddit.com/r/typescript/comments/gmhh6c/when_are_parenthesis_not_redundant/)
- url: https://www.reddit.com/r/typescript/comments/gmhh6c/when_are_parenthesis_not_redundant/
---
Context: I'm working on adding intersection types + parenthesis support to [typecheck.macro](https://github.com/vedantroy/typecheck.macro).

It seems like parenthesis are redundant in most cases. Examples:

(A | B)\[\] = Array&lt;A | B&gt;

(A) = A

(A | B) | (C &amp; D) = A | B | C &amp; D because Typescript implicitly does parenthesize for you

It seems like the only place parenthesis affect the order of "type operations" is:

A &amp; (B | C ) =  A &amp; B | A &amp; C. 

Can someone correct me if I am wrong?
