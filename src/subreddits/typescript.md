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
## [2][A SQL database implemented purely in TypeScript type annotations.](https://www.reddit.com/r/typescript/comments/iww4hs/a_sql_database_implemented_purely_in_typescript/)
- url: https://github.com/codemix/ts-sql
---

## [3][Debug Visualizer A VS Code extension for visualizing data structures while debugging. Like the VS Code's watch view, but with rich visualizations of the watched value.](https://www.reddit.com/r/typescript/comments/iwpzbm/debug_visualizer_a_vs_code_extension_for/)
- url: https://marketplace.visualstudio.com/items?itemName=hediet.debug-visualizer&amp;1
---

## [4][Safe Lookup Table in typescript](https://www.reddit.com/r/typescript/comments/iwsfwl/safe_lookup_table_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/iwsfwl/safe_lookup_table_in_typescript/
---
Is it possible to create a lookup table in typescript with functions that accept different type arguments?
I'm have a function that returns  a discriminated union with a type property that determines the type of data it holds, I would like to then use the type property as a key to an object with functions that accept types that the discriminated union was formed from
For example

    type a = {
        type:"obj1";
        foo:string;
    }
    type b={
        type:"obj2";
        bar:number;
    }
    type union= a or b;
    const lookupTable={
        obj1:({foo}:a)=&gt;foo.toLowerCase();
        obj2:({bar}:b)=&gt;(bar**2).toString(); // all the functions return the same type
    }
    //That all works, my issue comes when attempting to do something like
    const unionObj:union = {type:"obj1", foo:"str"}; // my code returns this from a function
    const result = lookupTable[unionObj.type](unionObj);
    // complicated type errors cause it can't tell that unionObj has been narrowed down to the right type

Can anyone please help with this
## [5][Compiler - Extract the type of union type tuple](https://www.reddit.com/r/typescript/comments/iwkt0i/compiler_extract_the_type_of_union_type_tuple/)
- url: https://www.reddit.com/r/typescript/comments/iwkt0i/compiler_extract_the_type_of_union_type_tuple/
---
If I have an array of type `string[]` and I want to extract the type from the array, I must do: 

    const typeArgs = type.typeArguments();
    typeArgs[0]; // string

But what if I have a tuple of type \[number, string\] and want to get the type union of `string | number`?
## [6][Can we use TypeScript in place of JavaScript?](https://www.reddit.com/r/typescript/comments/iw8ukw/can_we_use_typescript_in_place_of_javascript/)
- url: https://www.reddit.com/r/typescript/comments/iw8ukw/can_we_use_typescript_in_place_of_javascript/
---
I have a confession to make, I am one of those guys who tried to learn JavaScript many times but failed miserably despite having coding skills in many other languages. Recently I heard about TypeScript and really liked its way of programming. Can I ask if we can replace JS and in its place use TS.

Lets say I have to built a simple calculator with HTML+JS+CSS only. Could I use TS to write the code,  then compile it to JS and use this for my above calculator project? Will TS code be able to interact with the DOM elements like native JS?
## [7][Anybody looking for TypeScript interns?](https://www.reddit.com/r/typescript/comments/iwrhg3/anybody_looking_for_typescript_interns/)
- url: https://www.reddit.com/r/typescript/comments/iwrhg3/anybody_looking_for_typescript_interns/
---
Currently in college for a cs degree, and I’d rather work with TypeScript than a legacy codebase.  
  
Feel like I have a very strong amount of TypeScript and web dev knowledge for a college student, so if your teams recruiting or you have resources for who is lmk!
## [8][How can I type a decorator ?](https://www.reddit.com/r/typescript/comments/iwbbq7/how_can_i_type_a_decorator/)
- url: https://www.reddit.com/r/typescript/comments/iwbbq7/how_can_i_type_a_decorator/
---
[playground link of what I have done so far.](https://www.typescriptlang.org/play?#code/GYVwdgxgLglg9mABAQ2vMAeAKogpgDylzABMBnRAMXDQQD4AKYMALkSwEoWcBvAKACQAJ1xQQQpAwB0M5EIDmZDogC8dRPwECA9NrJwAtrkRkxwYIIC+Abj42gA)
## [9][I created a Web Development Discord server for TypeScript!](https://www.reddit.com/r/typescript/comments/iwk6m9/i_created_a_web_development_discord_server_for/)
- url: https://www.reddit.com/r/typescript/comments/iwk6m9/i_created_a_web_development_discord_server_for/
---
Wonderful news!

&amp;#x200B;

I created my first Discord server and I'm so happy for what I've made. It's called "The Web Dev Heaven" it's a friendly community where web developers and designers are welcomed to talk about the latest technologies and languages. So we can all learn about different ( Databases, Frameworks, Adobe XD plugins, Algorithms, CMS's, etc. ) to improve our code. We are here for supporting our projects ( So if you are looking to gain traffic with your website or feedback on your project, we are here for you! ).

&amp;#x200B;

[https://discord.gg/uXE4E7n](https://discord.gg/uXE4E7n)

&amp;#x200B;

Thank you, if you join!  
Thank you, for the support!  
Thank you, for the upvote!
## [10][First TypeScript Projects: sentence-engine &amp; reduce-to-chunks](https://www.reddit.com/r/typescript/comments/iwgz2i/first_typescript_projects_sentenceengine/)
- url: https://www.reddit.com/r/typescript/comments/iwgz2i/first_typescript_projects_sentenceengine/
---
Hello hello! 

For some time now, I've been enjoying getting into TypeScript. I started by refactoring and expanding a previous project of mine, and now find myself wanting to use TS for most anything. However, I feel like there are a lot of simple things I just don't know about or need some pointers on how to get right, such as best practises around custom types and exporting them. 

I'd love to get some general feedback on the two projects I've been tinkering with so far, and also just wanna share in case either of them happen to be useful to anyone on here. :)

The project that I recently refactored from JS to TS is sentence-engine, a string templating engine that I try to make as versatile and configurable as possible.   
**sentence-engine:** [Repo](https://github.com/sindrekjr/sentence-engine) | [NPM](https://www.npmjs.com/package/sentence-engine)

This weekend I also put together a very small util function called reduce-to-chunks, which may not be as efficient as a for-loop but should offer some handy utility. This was provoked by my wanting a utility to concisely reduce arrays of dates into arrays of weeks.  
**reduce-to-chunks:** [Repo](https://github.com/sindrekjr/reduceToChunks) | [NPM](https://www.npmjs.com/package/reduce-to-chunks)
## [11][ORM or not](https://www.reddit.com/r/typescript/comments/iw2wgt/orm_or_not/)
- url: https://www.reddit.com/r/typescript/comments/iw2wgt/orm_or_not/
---
 Before I learned about the use of an [ORM](https://en.wikipedia.org/wiki/Object-relational_mapping) (Object-Relational Mapping), I handled all interactions with the database through either stored procedures or building a query and setting up parameters. It was not the most exciting part of development. When I was introduced to NHibernate, an ORM for .net based on Hibernate, an ORM for java, I jumped right in and started bringing it into my projects.

There are hundreds of different ORMs. The goal of most ORMs is to abstract interactions with the database, support multiple databases, and handle mapping. Some implementations also extend the base database functionality by adding caching, lazy loading, and bulk operations, to name a few. The most common denominator of these ORMs is that they generate SQL statements for you. Some do this through their own query language which could be string-based (HQL), or through methods that build the request internally.

In the past few years, I have completely moved away from using ORMs in my personal projects and when I have to work with ORMs, I am very cautious. I know SQL well enough to know what good SQL looks like. I’m also always having to work extra hard to ensure that the generated SQL performs well. ORMs can generate beautiful SQL, but it can also generate statements that are overly verbose and slow.

I’m not here to say that ORMs are evil and you must always stay away from them. I’m here to show the issues that using an ORM can bring to light. I want to start with the abstraction most ORMs provide. Abstraction is an amazing thing. ORMs use abstraction to provide a common API that can speak to different database flavors. The abstraction is what excited me at first. I thought being able to swap out the database with minimal changes to the code was amazing. I’ve used the ability to swap out a database, my use-case being that I wanted the ability to do unit testing against an in-memory database. It was pretty great, at least that was my opinion at the time.

Abstraction is not a bad thing. It can speed up development by dealing with most of the inner workings of database interactions. Where abstraction becomes a problem is that it is generating SQL for you. It is true that this opens the door for developers that are not comfortable with SQL, it forces developers to learn the ORM. Each ORM has a learning curve and each has its own API to interact with the database. This can cause developers to focus on the ORM, not the SQL it is generating. To be good at an ORM, I believe you not only need to understand the SQL it is producing, but you also must know how to use the API to get it to build the SQL you need it to produce. Most ORMs do allow you to write your own SQL, but that is typically not recommended, especially for the more opinionated ORM frameworks. There have been too many times that I have struggled to get ORMs to build SQL the way I need it. I typically find myself falling back to straight SQL for complex queries. ORMs are usually great at creating simple queries but when you need something more advanced, it can be tough to get what you need out of it. End of the day, my issue with the abstraction that ORMs provide is that they hide SQL. This can become a crutch for some people as they put off learning SQL and instead learn the ORM. If you need to switch to another language or system that uses a different ORM, you now need to learn the new ORM.

In my experience, to properly use most ORMs, you have to adapt your system to use it, which can become very invasive, requiring you to have to follow specific patterns to even use it. Some require you to use their own query language, others require you to build your objects with metadata. which is used to handle mappings. These are not bad things, but the more code you have to tailor to the ORM, the harder it is to replace it if you need to. Of course, you can abstract an ORM by isolating it and putting as much as you can behind interfaces but in my opinion, external dependencies such as a database should almost always be abstracted, especially if you want to bring in unit testing that does not actually hit the database.

I am a fan of almost anything that simplifies and speeds up development. Yes, ORMs do exactly that, but sometimes the way they do it can cause pain in the long run. I believe that anyone who uses an ORM should understand SQL well enough to understand the SQL that is generated. I also want to caution when using more advanced ORM features such as lazy-loading. These features can provide a lot of value but they can also cause unintended side-effects. I’ll never forget a bug that took too much time to track down. I was handling an object returned from NHibernate that had quite a few other table dependencies that were set up for lazy-loading. You could force the loading of the data at the time of the queries execution or you can let the ORM query the extra data when you need to use it. An example could be a User table. A user might have one or more roles. You would get a user object from the ORM and when you access that objects roles property, the ORM will execute an additional query to load that uses roles. This can be an awesome feature that can provide huge performance boosts. That said, it can also cause issues if your lazy-loading loads other lazy-loaded data. My bug was just that. Because of all of the lazy-loading, the state of my object and its properties were hard to debug. My issue ended up being that the state of the object didn’t always reflect the state of the database, especially as other processes were interacting with the same underlying data, changing the state of my dependencies as I accessed the lazy-loaded properties.

I’m taking my stab at creating an ORM for Typescript that is extremely lite and using its helpers is completely optional. I would say I’m about 60% done and I hope to finish it by October. I’m naming it nORM which stands for no ORM. when I’m ready, I’ll publish it to GitHub and create an NPM package.
