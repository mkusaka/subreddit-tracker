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
## [2][TSConf:EU - Europe's TypeScript Conference](https://www.reddit.com/r/typescript/comments/f1msjk/tsconfeu_europes_typescript_conference/)
- url: https://tsconf.eu/
---

## [3][Weekend project: GraphQL-like queries and resolvers written in TypeScript](https://www.reddit.com/r/typescript/comments/f1fyar/weekend_project_graphqllike_queries_and_resolvers/)
- url: https://www.reddit.com/r/typescript/comments/f1fyar/weekend_project_graphqllike_queries_and_resolvers/
---
Hey folks, a month or two ago I asked myself, "what would it look like if I wanted to get GraphQL-like queries and resolvers in pure TypeScript?" I started hacking away and today I published the proof of concept under the title TypedQL. It allows you to define a graph query API and write queries against it. The real hero here is TypeScript Language Services, as seen in this [gif](https://imgur.com/fcUrI58).

Check it out [here](https://github.com/gregoryfabry/TypedQL); the README has some examples that show what the library can do. This is very much a proof of concept and there's definitely some black magic in the typings, but it can be downloaded and experimented with. I would love to hear your feedback, as I definitely think the concept has value and might be worth pursuing further.
## [4][typescript-eslint no implicit any rule?](https://www.reddit.com/r/typescript/comments/f1hayp/typescripteslint_no_implicit_any_rule/)
- url: https://www.reddit.com/r/typescript/comments/f1hayp/typescripteslint_no_implicit_any_rule/
---
Hi does anyone know of an [typescript-eslint](https://github.com/typescript-eslint/typescript-eslint) rule to disallow implicit any?

I know there is a rule to disallow *explicit* any, but I am wondering if there is a rule for disallowing implicit any?

Example:

    // instead of this:
    let s;
    
    // I want to force the use of a type here.. like this:
    let s: string;
    
    // disallowing explicit any solves this:
    let s: any;
    
    // but not this:
    let s;

Thank you!!
## [5][A React Router From Scratch in TypeScript](https://www.reddit.com/r/typescript/comments/f18o3l/a_react_router_from_scratch_in_typescript/)
- url: https://medium.com/@wim.jongeneel1/a-react-router-from-scratch-in-typescript-f0eec6ccb293?source=friends_link&amp;sk=4028f57135472edeb541ac13800055cb
---

## [6][Trouble groking infer](https://www.reddit.com/r/typescript/comments/f0roge/trouble_groking_infer/)
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
## [7][Typescript to C#?](https://www.reddit.com/r/typescript/comments/f0ho8j/typescript_to_c/)
- url: https://www.reddit.com/r/typescript/comments/f0ho8j/typescript_to_c/
---
I've tried searching, and I find various generators and such that will generate Typescript files from C#, but is there anything that does it the other way around? I.e. something that can generate C# types from Typescript? 🤔
## [8][The ultimate guide to create desktop apps for javascript entrepreneurs](https://www.reddit.com/r/typescript/comments/f0tbqf/the_ultimate_guide_to_create_desktop_apps_for/)
- url: https://medium.com/@merunasgrincalaitis/the-ultimate-guide-to-create-desktop-apps-for-javascript-entrepreneurs-4b2e1da0fe9c
---

## [9][Announcing TypeScript 3.8 RC](https://www.reddit.com/r/typescript/comments/f01vnf/announcing_typescript_38_rc/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-8-rc/
---

## [10][Advice on Transformation Utility library](https://www.reddit.com/r/typescript/comments/f0d3sm/advice_on_transformation_utility_library/)
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
## [11][Progress so far on my web browser for language-learning, LinguaBrowse, written in TypeScript using React Native](https://www.reddit.com/r/typescript/comments/f01it4/progress_so_far_on_my_web_browser_for/)
- url: https://twitter.com/LinguaBrowse/status/1225562067819749376?s=20
---

