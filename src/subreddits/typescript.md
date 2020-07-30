# typescript
## [1][Who's hiring Typescript developers - July](https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/)
- url: https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/
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
## [2][Dependency Injection: inversifyJS vs tsyringe](https://www.reddit.com/r/typescript/comments/i0jdmn/dependency_injection_inversifyjs_vs_tsyringe/)
- url: https://www.reddit.com/r/typescript/comments/i0jdmn/dependency_injection_inversifyjs_vs_tsyringe/
---
 I am new to DI.

I could not find a good comparison of the two frameworks. On the first look tsyringe looks more elegant to me. It seems to auto register classes. In inversifyJS you seem to have to bind every class to the container first before it couls be injected? On the other hanf inversifyJS is much more popular (but is also older). Why is this?

Any PROs &amp; CONs?

Thank you!
## [3][Rosebox 0.3.0](https://www.reddit.com/r/typescript/comments/i0l1x3/rosebox_030/)
- url: https://www.reddit.com/r/typescript/comments/i0l1x3/rosebox_030/
---
What's new in RB 0.3.0:

\- A new category of properties like (e.g. `paddingX`, `paddingY`).

\- New types for existing properties (e.g. `paddingObject`, `marginObject`, `animationObject)`

\- New properties including the animation properties. Overall, RB 0.3.0 has support for 111 properties.

\- `RBStyle` (or the deprecated `RoseboxProperties`) now supports any implemented CSS property by adding  `[x: string]: any` to it.

I'd also take the chance to share this brief article that serves as an introduction to RB:

[https://levelup.gitconnected.com/the-rosebox-project-30b01e7b3a12?source=your\_stories\_page---------------------------](https://levelup.gitconnected.com/the-rosebox-project-30b01e7b3a12?source=your_stories_page---------------------------) 

&amp;#x200B;

Website: [https://www.rosebox.dev/](https://www.rosebox.dev/)Github: [https://github.com/hugonteifeh/rosebox](https://github.com/hugonteifeh/rosebox)

Discord:  [https://discord.com/invite/bd7BHf2](https://discord.com/invite/bd7BHf2)
## [4][Running ts-node vs compiled JS files](https://www.reddit.com/r/typescript/comments/i07561/running_tsnode_vs_compiled_js_files/)
- url: https://www.reddit.com/r/typescript/comments/i07561/running_tsnode_vs_compiled_js_files/
---
New TS user here and while setting up a dev workflow, I am confused between the two for local development:  

1. Use `ts-node` to watch for changes and run `*.ts` files without transpiling to JS.
2. Transpile to JS first and then run the `*.js` files.

Option 1 looks more natural to me but I wouldn't know for sure if my transpiled JS files work as expected. Option 2 would make sure my transpiled JS files work as expected but that would require transpiling on EVERY file change. This question was asked [last year](https://www.reddit.com/r/typescript/comments/9o1zzp/tsnode_vs_compiling_for_development/) as well but I couldn't find a satisfactory answer.

Thanks!
## [5][Help with mutating functions](https://www.reddit.com/r/typescript/comments/i06arp/help_with_mutating_functions/)
- url: https://www.reddit.com/r/typescript/comments/i06arp/help_with_mutating_functions/
---
I'm having some difficulty wrapping my head around what I'm trying to do using Typescript with a function that mutates data.  

Let's say I have a person entity who can be either sitting or standing.  I want to make a function that will make him "stand up".  The function should only work on sitting entities and not allow you to pass an entity that is already standing.  After the function has done its work, it should assert that the entity is now standing, and after the function has been run allow you to continue with code that operates on standing entities.

Here's an example of the kind of thing I'm trying to do.  The error is obviously that it ends up in the `never` type, because the intersection of an entity that was sitting, but is now standing, is `never`

    interface Person {
      name: string;
      position: 'SITTING' | 'STANDING';
    }

    interface SittingPerson extends Person {
      position: 'SITTING';
    }

    interface StandingPerson extends Person {
      position: 'STANDING';
    }

    const bob: SittingPerson = { name: 'Bob', position: 'SITTING' };

    function standUp(p: SittingPerson): asserts p is StandingPerson {
      p.position = 'STANDING';
    }

    standUp(bob);

On the line:

    function standUp(p: SittingPerson): asserts p is StandingPerson 

there is an error that SittingPerson is not assignable to StandingPerson.  Obviously this is true, but what I'd like to be able to say to Typescript is:

`"Make your assertion based on the value at the end of the function, not at the start"`

I'm aware of this code immediately starts working if I simply make `bob` a `Person` instead of a `SittingPerson`, however the point is that I explicitly know which state he is in already, I don't want to define him as the less specific `Person`.  I don't want people to be able to invoke `standUp` on a person who may already be standing.  I want it to be as clear and concise as possible.

I'm wondering if I'm going about this the wrong way, and if there's a better way to do what I'm trying to do.

The reason it's using mutation rather than just returning a new person entity, is that it's in the context of a game, so these actions could potentially be happening countless times in a short period, so mutation is preferred vs creating new objects.

Thanks for any help
## [6][Is it possible to default export a type , and default export an interface without name ?](https://www.reddit.com/r/typescript/comments/hzzxti/is_it_possible_to_default_export_a_type_and/)
- url: https://www.reddit.com/r/typescript/comments/hzzxti/is_it_possible_to_default_export_a_type_and/
---
Questions in the title .

Edit :

My bad for not wording it properly .

I am not looking at default exporting two things .

I am looking on being able to default export a type . Also I am looking on being able to default export an interface without name .

There will always be one default export in each module .

I just had two different and not related question not worded properly .

Edit : 

Problem solved : [Here](https://www.reddit.com/r/typescript/comments/hzzxti/is_it_possible_to_default_export_a_type_and/fzmr3t1?utm_source=share&amp;utm_medium=web2x) is how to default export a type . [Here](https://www.reddit.com/r/typescript/comments/hzzxti/is_it_possible_to_default_export_a_type_and/fzmnme2?utm_source=share&amp;utm_medium=web2x) is how to default export an interface (you have to use a name on the definition of the interface ).
## [7][How to unit test private object variables mocha chai for typescript project?](https://www.reddit.com/r/typescript/comments/i08u9i/how_to_unit_test_private_object_variables_mocha/)
- url: https://www.reddit.com/r/typescript/comments/i08u9i/how_to_unit_test_private_object_variables_mocha/
---
After researching this I found the best response would be to test the private variables of an object against the prototype of the object. Effectively then inside the test creating a prototype and accessing the variables in that way. When I do this, however, it seems that all variables are not null but are undefined. How am I able to test the private variables, so that they remain defined? For reference, this is what I'm doing now:

[https://stackoverflow.com/questions/63162403/how-to-unit-test-private-variables-mocha-chai-for-typescript-project](https://stackoverflow.com/questions/63162403/how-to-unit-test-private-variables-mocha-chai-for-typescript-project)
## [8][How to resolve Error: error:0909006C:PEM routines:get_name:no start line?](https://www.reddit.com/r/typescript/comments/i069cn/how_to_resolve_error_error0909006cpem_routinesget/)
- url: https://www.reddit.com/r/typescript/comments/i069cn/how_to_resolve_error_error0909006cpem_routinesget/
---
I have a Typescript project where I read in `AmazonRootCA1.pem` for fully functioning `AwsIotRestApi`. That is, when I run the API it seems to be working just fine. I want to test the API with mocha. When I attempt to do so I get the following error:

    Error: error:0909006C:PEM routines:get_name:no start line

`AmazonRootCA1.pem`, the file which is being read is correctly formatted, so I'm not entirely sure how to go about solving this. Would anyone know how to resolve this error?
## [9][Any ideas to refactor this class any further? (Stong use of `.matchAll`)](https://www.reddit.com/r/typescript/comments/i020ss/any_ideas_to_refactor_this_class_any_further/)
- url: https://www.reddit.com/r/typescript/comments/i020ss/any_ideas_to_refactor_this_class_any_further/
---
I have been experimenting with using \`Array.from\` and \`reduce\` to build the output file instead of the \`for of\`, but keep on running into syntax problems. Any input? Would be great to get a mini code review.

The purpose of the program is to let you copy/paste YouTube chapters from creators' videos and make a chapter file that can be used with .mkv and .mp4 videos. Perfect for videos downloaded with \`youtube-dl\`.

I had the whole thing as a function before, but have been reading about domain-driven design and now think it belongs in a class. The 'chapter file' is definitely a 'domain concept'

[https://www.typescriptlang.org/play/?ssl=1&amp;ssc=1&amp;pln=47&amp;pc=1#code/MYewdgzgLgBAJgQygmBeGADACgJxAKwFNhYBaGAZQFccA3QgTxgQAcWBYAKFJ97-4GkugCmJA8H8wAFAAYAjAC4ArAHY5UgEwBKGAElIyADb6YIBAGsYAMxwIAtoQDuIHOYRg4MCIShUWMAEYIEACWwB6EdOEi4tJqqgDMcgAsMlq60AiG8IRg4JZOzFRQIDZIIWERODA4hOk4sLkWQfqEMMAAFq4A5oRRkuqqAGzxWgCigU1MJWAI3TD0OMHgEMxuWSzZcNnAQTW9MaoAHINaAPI4na5BAF4teIW7nGJ9sVIAnHJxA1oAwtVILQBBQptb7gKB4QzhFbuUBgRqdGgtAAyIE6QTAMAAZDAAEqENHQKF3KAPJ77GQJBSJLQULw+VrgMDEKBBXL2IJQNowACy4E6IAAIgAhPb9GTvRJSH5-EkwACqnkqNhAmyMrncQRsLGadjAsGqBPBpXAotikqSmh0Wp12Vg+lR6NNikSchkXx0egyRjgYOytGhrRlLQAdH7LE1CBguFxgPpAstvh0WCScAAxCMwADeXBguZgLBwQVo-xg+nRNQAKiBuUh2mgYAB6CQAfgAPCy7AA+AA6cEzagAvnJe-2hyPBxpu5OqFJ1FIW632qwUxWOc1O8GAFQaBudADcObzBaLJaXyfCEDkHnB6M6AG0ALr1x8Hzh5-OF4uyxrNABytkIK9oELMB92jN881hYCqBIJwJHRFhCnTP8AKAm9QIAGhgBCkIjNCQM6LRswg99cyCCxJAAQk5IIIGDMsmQgKsaygdpgxJaB4LARCoGQwgNCIw9SNIzk8HsGAACJfxAVokxTbC4ScEoWQaEAqDcCTX2E3MByEkS2lo4Mf0If87HrHDeIjUzCFfPT32aWAzxTX8qBsPwoXQKQtO0hyYGU9pCDgJFy2WdALL44N-LaAFDAkGi6IYytq1rNoNG84SLHyCRfMS4wKKiwLgsYwSSO03NfMzGAO0ILCnPCVcoGaGAB3rRLg06O4WAgSj0rKqqDLouqFmDRCIDaCQ7L63MMG+AAJAEsArEYcQAEkzIaXLc8IB1QNbqoHYNZykbswDmhaltW9a5PCTb3JwAdfwBbkRl2q7l3qtdCAHKNSrKtLJp0vTdNKtSoB4viJBKsrqm8HAMXi4Mhro-AQHRCQJJOiT-tK4H31B8GrIAyGswBqo6ThzA9oGozCbsAcAH0kYwXrgeBmMlkc660wzdAmXExN3u55oJtKgByQgAA9bG1Qh2IlqBRYwvTEGQLhsfZyAQGaejUQkIaIvx3DhYErggA](https://www.typescriptlang.org/play/?ssl=1&amp;ssc=1&amp;pln=47&amp;pc=1#code/MYewdgzgLgBAJgQygmBeGADACgJxAKwFNhYBaGAZQFccA3QgTxgQAcWBYAKFJ97-4GkugCmJA8H8wAFAAYAjAC4ArAHY5UgEwBKGAElIyADb6YIBAGsYAMxwIAtoQDuIHOYRg4MCIShUWMAEYIEACWwB6EdOEi4tJqqgDMcgAsMlq60AiG8IRg4JZOzFRQIDZIIWERODA4hOk4sLkWQfqEMMAAFq4A5oRRkuqqAGzxWgCigU1MJWAI3TD0OMHgEMxuWSzZcNnAQTW9MaoAHINaAPI4na5BAF4teIW7nGJ9sVIAnHJxA1oAwtVILQBBQptb7gKB4QzhFbuUBgRqdGgtAAyIE6QTAMAAZDAAEqENHQKF3KAPJ77GQJBSJLQULw+VrgMDEKBBXL2IJQNowACy4E6IAAIgAhPb9GTvRJSH5-EkwACqnkqNhAmyMrncQRsLGadjAsGqBPBpXAotikqSmh0Wp12Vg+lR6NNikSchkXx0egyRjgYOytGhrRlLQAdH7LE1CBguFxgPpAstvh0WCScAAxCMwADeXBguZgLBwQVo-xg+nRNQAKiBuUh2mgYAB6CQAfgAPCy7AA+AA6cEzagAvnJe-2hyPBxpu5OqFJ1FIW632qwUxWOc1O8GAFQaBudADcObzBaLJaXyfCEDkHnB6M6AG0ALr1x8Hzh5-OF4uyxrNABytkIK9oELMB92jN881hYCqBIJwJHRFhCnTP8AKAm9QIAGhgBCkIjNCQM6LRswg99cyCCxJAAQk5IIIGDMsmQgKsaygdpgxJaB4LARCoGQwgNCIw9SNIzk8HsGAACJfxAVokxTbC4ScEoWQaEAqDcCTX2E3MByEkS2lo4Mf0If87HrHDeIjUzCFfPT32aWAzxTX8qBsPwoXQKQtO0hyYGU9pCDgJFy2WdALL44N-LaAFDAkGi6IYytq1rNoNG84SLHyCRfMS4wKKiwLgsYwSSO03NfMzGAO0ILCnPCVcoGaGAB3rRLg06O4WAgSj0rKqqDLouqFmDRCIDaCQ7L63MMG+AAJAEsArEYcQAEkzIaXLc8IB1QNbqoHYNZykbswDmhaltW9a5PCTb3JwAdfwBbkRl2q7l3qtdCAHKNSrKtLJp0vTdNKtSoB4viJBKsrqm8HAMXi4Mhro-AQHRCQJJOiT-tK4H31B8GrIAyGswBqo6ThzA9oGozCbsAcAH0kYwXrgeBmMlkc660wzdAmXExN3u55oJtKgByQgAA9bG1Qh2IlqBRYwvTEGQLhsfZyAQGaejUQkIaIvx3DhYErggA)
## [10][Add property to a type](https://www.reddit.com/r/typescript/comments/hzvwp6/add_property_to_a_type/)
- url: https://www.reddit.com/r/typescript/comments/hzvwp6/add_property_to_a_type/
---
Hi,
I am starting out with typescript and have been stuck on this problem for some time.

I have an object type whose keys are in an enum. So I made it the following way.

```
type MyType = {
    [x in myenum]: string;
}
```

This works. However I need to add one more property to it, say `id:number` I have tried using the extends method as well as directly adding it. But it did not work.

Any suggestions to make it work?
## [11][Would you guys advise a junior who's put a lot of time into Typescript to hold out for a TS job or pick up whatever comes along?](https://www.reddit.com/r/typescript/comments/hzn9h5/would_you_guys_advise_a_junior_whos_put_a_lot_of/)
- url: https://www.reddit.com/r/typescript/comments/hzn9h5/would_you_guys_advise_a_junior_whos_put_a_lot_of/
---
My situation is I've been studying programming intensely for 3 years now. JS was my first language, then I did a Wordpress plugin in PHP, and now I'm working with Typescript and React. I also know Express and Postgres.

I'm not afraid of a job that is focused on say Rails or Laravel development, I can mostly read Python (dabbled a bit in it early). So I'm confident that i could jump into any Ruby or Python job and be quite capable for an entry level developer. I just would prefer to develop my muscles using TS utility types, is/as, generics that take function arguments etc.

I'm moving to Atlanta which is an easier job market at the bottom that the top 3 where getting any offer at all can be a prize for a junior.

What do you guys think, take anything that comes along or try to get a TS job? I'm confident in my skils and think I could surely get a React or Angular job, I'm less certain about the quantity of jobs that use TS in Atlanta or the wisdom of wanting to work in that specific tech stack.

PS: To clarify, since TS can be on both ends of the stack, I'd be happy if I got to use TS on either end. No complaints about TS/React with a Rails API for example.
