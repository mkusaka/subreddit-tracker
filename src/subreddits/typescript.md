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
## [2][A React Router From Scratch in TypeScript](https://www.reddit.com/r/typescript/comments/f18o3l/a_react_router_from_scratch_in_typescript/)
- url: https://medium.com/@wim.jongeneel1/a-react-router-from-scratch-in-typescript-f0eec6ccb293?source=friends_link&amp;sk=4028f57135472edeb541ac13800055cb
---

## [3][Trouble groking infer](https://www.reddit.com/r/typescript/comments/f0roge/trouble_groking_infer/)
- url: https://www.reddit.com/r/typescript/comments/f0roge/trouble_groking_infer/
---
So I come from the Java world and I'm having some difficulty with Typescript's `infer` keyword. I'm playing around trying to infer the generic types of a nested data structure, but I can't seem to make the types line up, and I was hoping someone could show me the error of my ways. 

     type Tuple&lt;T&gt; = [Foo&lt;T&gt;, Bar&lt;T&gt;]
     type Tuples&lt;...&gt; = ReadonlyArray&lt;Tuple&lt;?&gt;&gt;

     const tuples = [ // inferred as Tuples&lt;string | number&gt; 
       ['one', 'two'], // inferred Tuple&lt;string&gt;
       [3, 4] // inferred Tuple&lt;number&gt;
     ]

This works fine if i specify the types explicitly, but I'm sure there's got to be a way to infer them.

edit: Adding a more detailed use case: 
edit2: [playground](https://www.typescriptlang.org/play/#code/C4TwDgpgBAYg9nAPAFQHxQLxQN5QG4CGANgFxTJQC+AUKJFAEIEBOK6WucwAFhM2RRrUA9MNgIoBAHYATRiyhSIEOT2gBnAgFtoAcwhLmASwDGUOtDAttEYH2rnw0ZAFcwRCG0xQA2vCRoADTyrGgAug4W5G4e6l5YAEoQBDJwUkQgAILMzAQgKDGeaKjU1CZp6sBQqQDKcDo8RlK63mwAFABmCGT+bMEARixkTKGoAJRkeHBGchjo2A5QS+VS6nAeAHREcLptFnAdUF1wG4REY4vLFesQWzt7TgdQg8wbXLzMF1SlUCuV5m53BAABLSGQeZitNBtDaw4CFdQCBFsMaYeaXLqQtp-Ko+Y4DFhhKBPeFA9SohZLKm1eq2bhNXb454sL5LGjfByiapwCDqKQAciq5S0YCMHk5YiiUhcWn6fCgRnUii4knU6iMuikBH6HnMcBCiEqxmaJV+FSqgwAXt5SUDQbIIW1Lj5cGcyPy0hB+VRgpw1PwoIKAO5wb2UMLBLlNDp8ZgqaJAw3AY26U1LF34YhkADMPpwxP9ZAALFQiVzEABaCvc3kCoX1UW6ghK-aHK0KpWuMlJlOmi4S8z0pUh5gAa3UDhxUAAji4AB42wEee3gviIaWy+UAHygRoZUFQTqpGbdgc9Yd9BY+7uAIbDEagztdWagucol-efGLpYcYyAA)
## [4][Typescript to C#?](https://www.reddit.com/r/typescript/comments/f0ho8j/typescript_to_c/)
- url: https://www.reddit.com/r/typescript/comments/f0ho8j/typescript_to_c/
---
I've tried searching, and I find various generators and such that will generate Typescript files from C#, but is there anything that does it the other way around? I.e. something that can generate C# types from Typescript? 🤔
## [5][The ultimate guide to create desktop apps for javascript entrepreneurs](https://www.reddit.com/r/typescript/comments/f0tbqf/the_ultimate_guide_to_create_desktop_apps_for/)
- url: https://medium.com/@merunasgrincalaitis/the-ultimate-guide-to-create-desktop-apps-for-javascript-entrepreneurs-4b2e1da0fe9c
---

## [6][Announcing TypeScript 3.8 RC](https://www.reddit.com/r/typescript/comments/f01vnf/announcing_typescript_38_rc/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-8-rc/
---

## [7][Advice on Transformation Utility library](https://www.reddit.com/r/typescript/comments/f0d3sm/advice_on_transformation_utility_library/)
- url: https://www.reddit.com/r/typescript/comments/f0d3sm/advice_on_transformation_utility_library/
---
Hello Everyone,

I have created a small library that is designed to transform data using a deduction map, this map contains an array of deductions, each deduction will use one or more values from the source and apply a list of transforms before storing it in the output at the given path.

&amp;#x200B;

Really interested in feedback and suggestions on how to improve the idea.

Thanks

&amp;#x200B;

Link: [https://github.com/robertpitt/deducer](https://github.com/robertpitt/deducer)

Example: [https://github.com/robertpitt/deducer/blob/master/test/index.test.ts#L94](https://github.com/robertpitt/deducer/blob/master/test/index.test.ts#L94)
## [8][Progress so far on my web browser for language-learning, LinguaBrowse, written in TypeScript using React Native](https://www.reddit.com/r/typescript/comments/f01it4/progress_so_far_on_my_web_browser_for/)
- url: https://twitter.com/LinguaBrowse/status/1225562067819749376?s=20
---

## [9][Could use a little help](https://www.reddit.com/r/typescript/comments/f09c0m/could_use_a_little_help/)
- url: https://www.reddit.com/r/typescript/comments/f09c0m/could_use_a_little_help/
---
Hi, everyone I'm somewhat kinda new to typescript but not oblivious to JavaScript, What I'm wanting to do (if possible) is wanting to combine my 3+ .ts files when it transpiles / compiles I would like for it to be in one .js file using the modular pattern.

similar to what sass kinda does with with multiply partials.

I saw a comment in the TSConfig file for "outfile" and I couldn't get it to work, I kept getting an error about lack of persmissions or something.
## [10][Is is possible to combine multiple NestJS Swagger decorators?](https://www.reddit.com/r/typescript/comments/f03w6o/is_is_possible_to_combine_multiple_nestjs_swagger/)
- url: https://www.reddit.com/r/typescript/comments/f03w6o/is_is_possible_to_combine_multiple_nestjs_swagger/
---
Hi. I'm using NestJS and Swagger. Before each GET router which returns result as a list, i add these decorators:

    @Get()
    @ApiQuery({ name: "page", required: false })
    @ApiQuery({ name: "limit", required: false })
    @ApiQuery({ name: "sort", required: false })
    @ApiQuery({ name: "order", required: false })

Is there any chance that i could combine those into one custom decorator, such as `@GetListOptions()`?
## [11][TypeScript Beta for using Optional Chaining](https://www.reddit.com/r/typescript/comments/ezy7ti/typescript_beta_for_using_optional_chaining/)
- url: https://www.reddit.com/r/typescript/comments/ezy7ti/typescript_beta_for_using_optional_chaining/
---
Hi there, i'd really like to use the cool new features like optional chaining..

I followed the npm i instructions from the blog

 [https://devblogs.microsoft.com/typescript/announcing-typescript-3-8-beta/](https://devblogs.microsoft.com/typescript/announcing-typescript-3-8-beta/) 

however after it downloaded, tsc -- version still says i'm on Version 3.6.4

And the compiler still complains about optional chaining... do you know how to fix this?

Thanks much!
