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
## [2][How can I type a decorator ?](https://www.reddit.com/r/typescript/comments/iwbbq7/how_can_i_type_a_decorator/)
- url: https://www.reddit.com/r/typescript/comments/iwbbq7/how_can_i_type_a_decorator/
---
[playground link of what I have done so far.](https://www.typescriptlang.org/play?#code/GYVwdgxgLglg9mABAQ2vMAeAKogpgDylzABMBnRAMXDQQD4AKYMALkSwEoWcBvAKACQAJ1xQQQpAwB0M5EIDmZDogC8dRPwECA9NrJwAtrkRkxwYIIC+Abj42gA)
## [3][Can we use TypeScript in place of JavaScript?](https://www.reddit.com/r/typescript/comments/iw8ukw/can_we_use_typescript_in_place_of_javascript/)
- url: https://www.reddit.com/r/typescript/comments/iw8ukw/can_we_use_typescript_in_place_of_javascript/
---
I have a confession to make, I am one of those guys who tried to learn JavaScript many times but failed miserably despite having coding skills in many other languages. Recently I heard about TypeScript and really liked its way of programming. Can I ask if we can replace JS and in its place use TS.

Lets say I have to built a simple calculator with HTML+JS+CSS only. Could I use TS to write the code,  then compile it to JS and use this for my above calculator project? Will TS code be able to interact with the DOM elements like native JS?
## [4][My first typescript project - OpenGrow: An open-source IoT plant care system for automated watering and data tracking!](https://www.reddit.com/r/typescript/comments/ivvb8u/my_first_typescript_project_opengrow_an/)
- url: https://www.reddit.com/r/typescript/comments/ivvb8u/my_first_typescript_project_opengrow_an/
---
Hello everyone,

I have recently completed my first typescript project! It is an open-source IoT plant care system called OpenGrow. The web-application component of the system is built using the MERN stack. I had a great experience using typescript as it helped avoid mistakes in sending/handling data packets sent across the web application.

[Fully assembled OpenGrow system with snapshots of web-app.](https://preview.redd.it/kgw4ddvko4o51.jpg?width=789&amp;format=pjpg&amp;auto=webp&amp;s=e57774134d56df68ac1da984c3b8c2ab0482e5fb)

With OpenGrow, you can set-up any of your plants to be automatically watered to maintain a configurable soil moisture level. The system also takes measurements of soil-moisture and ambient light every 5 minutes which the user can view from a responsive plot (made with Nivo graphs) on the web-app.

The STM32 microcontroller used at the edge is running FreeRTOS based firmware. The ESP8266 used for wi-fi communication runs on NodeMCU. This project also features its own MQTT broker which was made using NodeJs and the Aedes framework.

For those interested, I have made the project easy to set-up with readily available parts and a detailed guide!

Links:

[Github repository](https://github.com/esyywar/OpenGrow-IoT-Plant-Monitor)

[Medium article on the full system](https://medium.com/@rahuleswar_84677/opengrow-building-an-automated-iot-solution-for-plant-care-with-the-stm32-esp8266-and-mern-stack-cd4bb144326e)

The system has been working well for over a week in my home and I hope others will be able to replicate or modify it. This project touches on all the important components for a full IoT solution so I think it can also be a great learning resource for those interested in IoT!

If anyone encounters issues or finds problems, please raise an issue on the github repo and I will address it.
## [5][ORM or not](https://www.reddit.com/r/typescript/comments/iw2wgt/orm_or_not/)
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
## [6][@node-cool: The fastest way to write a node server with TypeScript](https://www.reddit.com/r/typescript/comments/iwb970/nodecool_the_fastest_way_to_write_a_node_server/)
- url: https://hacklone.github.io/node-cool?r=reddit-typescript
---

## [7][Is it possible to use webpack with a library that uses dependency injection?](https://www.reddit.com/r/typescript/comments/iwankx/is_it_possible_to_use_webpack_with_a_library_that/)
- url: https://www.reddit.com/r/typescript/comments/iwankx/is_it_possible_to_use_webpack_with_a_library_that/
---
Hi all,

I'm pretty new to webpack so apologies if this is an obvious answer. I'm currently trying to deploy a typescript lambda api to AWS using this library: [https://www.npmjs.com/package/ts-lambda-api](https://www.npmjs.com/package/ts-lambda-api)

The library uses Inversify for dependency injection: [https://github.com/inversify/InversifyJS](https://github.com/inversify/InversifyJS)

The library also advises building by just using zip but I'd like to use webpack if I can.

I currently have two files for my api, a main api.ts file and then a controller.ts file. Currently the api class finds the controllers using the same lines as in the ts-lambda-api docs:

    const controllersPath = [path.join(__dirname, "controllers")] 
    const app = new ApiLambdaApp(controllersPath, appConfig) 

As a result I need the controller directory to be available to the api.ts class once it's been bundled. Currently my webpack config looks like this:

    import path = require('path')
    import webpack = require('webpack')
    
    const config: webpack.Configuration = {
        module: {
            rules: [
                { test: /\.js\.map$/, use: "ignore-loader" },
                { test: /\.d\.ts$/, use: "ignore-loader" }
            ]
        },
        mode: 'production',
        entry: {
            my_api: ['./src/my-api.js'],
        },
        output: {
            path: path.resolve(__dirname, 'dist'),
            filename: '[name]/main.js',
            library: '[name]',
            libraryTarget: 'commonjs2'
        },
        target: 'node',
        externals: [/aws-sdk.*/, 'aws-lambda']
    };
    
    export default config

How can I add to it to allow my api access to the other classes it needs? The webpack docs have a page on multiple entry points ([https://webpack.js.org/concepts/entry-points/#multi-page-application](https://webpack.js.org/concepts/entry-points/#multi-page-application)) but that doesn't seem to be what I'm looking for as it gives me multiple bundles.

I can put multiple files into one bundle by doing something like:

    api: ['./src/api.js','./src/controllers/controller.js], 

Although this won't then work with the Injectable annotations. Anyone have any other ideas?
## [8][Having an error with a generic React component](https://www.reddit.com/r/typescript/comments/iw2ksw/having_an_error_with_a_generic_react_component/)
- url: https://www.reddit.com/r/typescript/comments/iw2ksw/having_an_error_with_a_generic_react_component/
---
Hi. Alright, so I'm in the process of refactoring some old code I wrote into TypeScript. This particular piece of code is a function that builds a React component. I MAY turn it into a functional component during the refactor, but there are a lot of factors in that decision, and that's a whole separate discussion.

Anyway, the purpose of the code snippet I'm going to provide is that it takes in a React component type plus some arguments, does a bunch of fancy setup, and then returns the component (plus a lot of other things, but again, I'm simplifying here). The React component has a generic type for its props to keep it in sync with some of the arguments being passed in. However, when I try to use the component in JSX, I get this error: 

    JSX element type doesn't have any construct or call signatures

Before I go any further, here is an extremely simplified version of my function:

    import React, { ElementType } from 'react';
    
    interface Args&lt;Props extends object&gt; {
    	// Arg properties go here
    }
    
    const componentBuilder = &lt;Props extends object&gt;(component: ElementType&lt;Props&gt;, args: Args&lt;Props&gt;) =&gt; {
    	// Skipping over logic
    
    	const Component: ElementType&lt;Props&gt; = component; // Just want the C to be capital for the JSX element. My OCD...
    
    	return (
    		&lt;div&gt;
    			&lt;Component /&gt;
    		&lt;/div&gt;
    	);
    };

So, when I try using Component like that in this example, is when I get the error. I don't understand why. After experimenting, I know it has something to do with the generic type. If I remove the generic parameter, everything works.

So, what can I do to make the generic parameter work? It's extremely important that it does. Thanks.
## [9][Array lookup in TypeScript should return "T | undefined" and not "T"](https://www.reddit.com/r/typescript/comments/ivlxdk/array_lookup_in_typescript_should_return_t/)
- url: https://www.reddit.com/r/typescript/comments/ivlxdk/array_lookup_in_typescript_should_return_t/
---
Is there a workaround for this? I found an issue report with no progress:

https://github.com/microsoft/TypeScript/issues/11122

&gt; const xs = [1,2,3];
&gt; const x5 = xs[5]; // type is number, expected number | undefined
&gt;
&gt; The type system doesn't know of care about the length of xs, so I would expect any lookup in the array to return T | undefined.

This has bitten me a few times and seems like an obvious hole. The best I've got is to define my own "safeArrayGet(arr, index)" function.
## [10][Bad pattern or not: Inline object literal for AS type assertion](https://www.reddit.com/r/typescript/comments/ivyw5o/bad_pattern_or_not_inline_object_literal_for_as/)
- url: https://www.reddit.com/r/typescript/comments/ivyw5o/bad_pattern_or_not_inline_object_literal_for_as/
---
    jsonwebtoken.sign((req.user as { user_id: string }).user_id, process.env.JWT_SECRET!);
    

It cleared the error. And I was a bit unsure about how much info the dB query would be returning in this example. I thought about taking the interface for the full user record and using Pick to assert that the user\_id was there at minimum. But then I wrote this instead. 

i'm curious if you guys think inlining a concrete type is a bad pattern in general or acceptable in these situatons.
## [11][code.store : A new TypeScript &amp; GraphQL backend as a service in beta.](https://www.reddit.com/r/typescript/comments/iv50px/codestore_a_new_typescript_graphql_backend_as_a/)
- url: https://www.reddit.com/r/typescript/comments/iv50px/codestore_a_new_typescript_graphql_backend_as_a/
---
Hi guys!

We've been working on this project for one year now. [code.store](https://code.store/) is a GraphQL + TypeScript back-end as a service with a focus on micro/macro-services re-use.  
· You can create scalable, serverless, back-end, and re-use them in different projects with a single command line.  
· We wanted to make it as simple as possible, so you can work on business logic, and we try to take care of everything else: database and code generation, automatic versioning, continuous integration, and deployments, backups and data migrations, and scaling.  
We see potential use cases:  
· Back-end for mobile applications, PWA, or any front-end apps.  
· Company-wide catalog of reusable, live microservices, that can be instantly deployed on new projects or applications.

It's our first public release, we would be happy with any feedback if you have time to play with it.

[https://code.store](https://code.store/)
