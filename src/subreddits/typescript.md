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
## [2][Best resources for learning typescript](https://www.reddit.com/r/typescript/comments/gnvyd3/best_resources_for_learning_typescript/)
- url: https://www.reddit.com/r/typescript/comments/gnvyd3/best_resources_for_learning_typescript/
---
As a longtime JS dev who now accepts TS as my lord and saviour - what are the best resources for 'deeply' learning TS fundamentals?

Are there any equivalent resources like 'Eloquent Javascript', 'You don't know JS yet' etc. for TS?

Video series/course suggestions also welcome.

Thanks!
## [3][Typing some complex lodash stuff](https://www.reddit.com/r/typescript/comments/gnsseg/typing_some_complex_lodash_stuff/)
- url: https://www.reddit.com/r/typescript/comments/gnsseg/typing_some_complex_lodash_stuff/
---
I was trying to type return value of \_.[defaultsDeep](https://lodash.com/docs/4.17.15#defaultsDeep)

Heres the return type for 2 arguments

`type Primitive = string | number | boolean | null`  
`type DeepDefault&lt;T, K&gt; = T extends Primitive ? T : Omit&lt;T, keyof K&gt; &amp; {`  
  `[Key in keyof K]: Key extends keyof T ? DeepDefault&lt;T[Key], K[Key]&gt; : K[Key]`  
`} ;`  
Heres the link to [playground](https://www.typescriptlang.org/play/?ssl=5&amp;ssc=4&amp;pln=1&amp;pc=1#code/C4TwDgpgBACgTgSwLYOAgbtAvFAzsRAOwHMoAfKQgVyQCMI5ypaB7FgGwgENCnr32AKEGhIUACIQIYSQDMuVdsAA8AFQA0UANIA+KDlVQIAD2ARCAE1yxEKNJigB+KIYBcUAPJ21mgNYQQFlltPQAyKABvQSgoAG0tAKgEXn9A4K0AXXcEkCNTcysoVKCXJwkpGQh5RRVVeICMzS16kAy9bJaMwQBfKABuYVFoFloAKwBGfUjoqC53cfUZ2nd8ImJFmIh3KJiY2QQ4fHcCKggZ7pnZbZm9g6OoAHJEYgALYAfznsHwYbGAJimOygAGN5hsoBYVgRkusZltprs8BBgSxLO55OxcGcYhc9tdEftDsB3A8AO5wVHED44r4iH5QACSHjAYC4U0k0jkCiUyhGE00fL+OkEKMI+CgLBZc0ZzNZgJm0oWS3cACJVjCVeDQVAlTFIVA1dCSJq4fjdoT7icIOCYljRfqMVjzuCrgjzXdiY9nm9qVALhcgA)   
The problem is typing the general case of N parameters   
I know i can just add overloads for up to \~10 parameters for example like so  
`DeepDefaultFor3&lt;T,K,J&gt; = DeepDefault&lt;T, DeepDefault&lt;K,J&gt;&gt;`  
but that doesn't seem right  
Does any one have ideas about hot one would type  \_.[defaultsDeep](https://lodash.com/docs/4.17.15#defaultsDeep) of N parameters  when N is unknown?
## [4][Using JSDoc tags to test functions [Prototype]](https://www.reddit.com/r/typescript/comments/gn6sit/using_jsdoc_tags_to_test_functions_prototype/)
- url: https://i.redd.it/8gqa8evpfvz41.gif
---

## [5][ESLint - What are the main differences between "eslint-config-airbnb-typescript" and "eslint-config-standard-with-typescript" conventions?](https://www.reddit.com/r/typescript/comments/gney1v/eslint_what_are_the_main_differences_between/)
- url: https://www.reddit.com/r/typescript/comments/gney1v/eslint_what_are_the_main_differences_between/
---
Section [https://github.com/typescript-eslint/typescript-eslint/blob/master/docs/getting-started/linting/README.md#eslint-configs](https://github.com/typescript-eslint/typescript-eslint/blob/master/docs/getting-started/linting/README.md#eslint-configs) lists two list of ESLint rules:

* Airbnb's ESLint config - [https://www.npmjs.com/package/eslint-config-airbnb-typescript](https://www.npmjs.com/package/eslint-config-airbnb-typescript)
* Standard - [https://www.npmjs.com/package/eslint-config-standard-with-typescript](https://www.npmjs.com/package/eslint-config-standard-with-typescript)

What are the main differences between them?
## [6][[question] Which type of dependency should @types/foobar belong to?](https://www.reddit.com/r/typescript/comments/gni4gr/question_which_type_of_dependency_should/)
- url: https://www.reddit.com/r/typescript/comments/gni4gr/question_which_type_of_dependency_should/
---
Hi. Consider a TypeScript project which uses `foo` and `bar` dependencies.  Should the corresponding `@types/foo` and `@types/bar` belong to `devDependencies` XOR `dependencies`?

The `create-react-app` boiler template lists `@types/*` in `dependencies`.  But I just don't know where to put it.

Does it really matter where they are?
## [7][Trying to gather Geolocation speed data over an interval](https://www.reddit.com/r/typescript/comments/gngvfz/trying_to_gather_geolocation_speed_data_over_an/)
- url: https://www.reddit.com/r/typescript/comments/gngvfz/trying_to_gather_geolocation_speed_data_over_an/
---
I have an angular ionic 5 app running on a cordova android emulator. What I need to do is use the ionic Geolocation plugin to find the speed of the user. I need to record the user's speed every second, and make sure it falls in the range of 0.7 &lt; x &lt; 2.5 m/s. If it does not, continue recording every second until it does. I need to record 10 valid speed measurements using this method. Then determine the average speed from that set of data. Here is a link to my Stackoverflow [post](https://stackoverflow.com/questions/61920194/finding-speed-of-an-interval-of-10-seconds), and here is my typescript file:

`import { Component } from '@angular/core';`

`import { Geolocation } from '@ionic-native/geolocation/ngx';`



`u/Component({`

`selector: 'app-home',`

`templateUrl: 'home.page.html',`

`styleUrls: ['home.page.scss'],`

`})`

`export class HomePage {`

`framework = 'At exact boarding time';`



`constructor(private geolocation: Geolocation) {}`



`latitude: any = 0; //latitude`

`longitude: any = 0; //longitude`

`speed: any = 1; //speed`



`choices = {`

`timeout: 10000,` 

`enableHighAccuracy: true,` 

`maximumAge: 3600`

`};`



`average_walking_speed(){`

`this.geolocation.getCurrentPosition(this.choices).then((resp) =&gt; {`

`this.speed = resp.coords.speed;`

`console.log(this.speed);`

`}).catch((error) =&gt; {`

`console.log('Error getting speed', error);`

`});`

`}`



`// use geolocation to get user's device coordinates`

`getCurrentCoordinates() {`

`this.geolocation.getCurrentPosition().then((resp) =&gt; {`

`this.latitude = resp.coords.latitude;`

`this.longitude = resp.coords.longitude;`

`this.speed = resp.coords.speed;`

`}).catch((error) =&gt; {`

`console.log('Error getting location', error);`

`});`



`document.getElementById("updater").innerHTML = ' ' + this.speed +" m/s"`

`}`    

`}`
## [8][Vanilla library types](https://www.reddit.com/r/typescript/comments/gncmsy/vanilla_library_types/)
- url: https://www.reddit.com/r/typescript/comments/gncmsy/vanilla_library_types/
---
So I've installed the '@types/\[library\]' for what I need. I can't figure out how to leverage that .d.ts file so I can use those types in my own code. Is that what it is supposed to do? For example I'm using express and install it's types. I pass the req object to a function and I want to typecheck that parameter when writing the function for expresses req object but I don't know how to do it.
## [9][myzod v1.0.0 is out](https://www.reddit.com/r/typescript/comments/gncnh6/myzod_v100_is_out/)
- url: https://www.reddit.com/r/typescript/comments/gncnh6/myzod_v100_is_out/
---
Today I released version 1.0.0 of [myzod](https://www.npmjs.com/package/myzod)! 

I just want to thank everybody who has given feedback and participated over the past month or two.

Hope you guys find it useful.
## [10][Fcaljs - extensive math expression evaluator library](https://www.reddit.com/r/typescript/comments/gn7ves/fcaljs_extensive_math_expression_evaluator_library/)
- url: https://www.npmjs.com/package/fcal
---

## [11][Destructuring and sum types](https://www.reddit.com/r/typescript/comments/gn7oz6/destructuring_and_sum_types/)
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
