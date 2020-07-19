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
## [2][Any VSCode users? How to stop it from thinking type constructors are JSX code?](https://www.reddit.com/r/typescript/comments/htx1ez/any_vscode_users_how_to_stop_it_from_thinking/)
- url: https://www.reddit.com/r/typescript/comments/htx1ez/any_vscode_users_how_to_stop_it_from_thinking/
---
For example when I try to type:

    Promise&lt;string&gt;

The moment I type that closing angle bracket VSCode autocompletes with this:

    Promise&lt;string&gt;&lt;/string&gt;

This is in .ts files by the way. 

Has anyone figured out a way to get VSCode to get smarter about when it incorporates this autocompletion?

I'm sure there's a way to restrict it to .tsx files which I could implement on personal projects. But if dealing with a company project that uses JSX code in .ts files I don't think an extension level restriction would work. And even then, sometimes you want to use type constructors in files with JSX code.

If anyone else currently deals with this, the best practice I currently use, when I remember to, is to first type this out to prevent the autocompleting "closing tag":

    Promise&lt;&gt;
## [3][What's the weirdest, most odd-looking and/or most complex generic you've seen or made?](https://www.reddit.com/r/typescript/comments/htukg1/whats_the_weirdest_most_oddlooking_andor_most/)
- url: https://www.reddit.com/r/typescript/comments/htukg1/whats_the_weirdest_most_oddlooking_andor_most/
---

## [4][Finally I Developed SQL Query Generating App](https://www.reddit.com/r/typescript/comments/htw2nd/finally_i_developed_sql_query_generating_app/)
- url: https://youtu.be/K9uG63QSW-I
---

## [5][Is it acceptable to use switch(true) to match the first met condition?](https://www.reddit.com/r/typescript/comments/htqokr/is_it_acceptable_to_use_switchtrue_to_match_the/)
- url: https://www.reddit.com/r/typescript/comments/htqokr/is_it_acceptable_to_use_switchtrue_to_match_the/
---
I currently have the following switch statement in a TypeScript file:

    switch (true) {
        case reminder.repeat != ReminderRepeatType.ONCE: {
            this.reminders.repeating.push(reminder);
            break;
        }
        case reminder.date.toDateString() == today.toDateString(): {
            this.reminders.today.push(reminder);
            break;
        }
        case reminder.date &lt; today: {
            this.reminders.pending.push(reminder);
            break;
        }
        case reminder.date &gt; today: {
            this.reminders.past.push(reminder);
            break;
        }
    }

I only want each reminder object to appear in one array in order to prevent displaying a reminder twice.

Is this an acceptable way of using the switch statement? Is there a better way of doing this?
## [6][What is needed to get source maps working?](https://www.reddit.com/r/typescript/comments/htg5jx/what_is_needed_to_get_source_maps_working/)
- url: https://www.reddit.com/r/typescript/comments/htg5jx/what_is_needed_to_get_source_maps_working/
---
I''m using NDB which is like the google chrome debugger but for Node.js apps. I'm able to set breakpoints in compiled .js files but not source files in .ts. I see this notice in the app:

[ndb source map notice](https://preview.redd.it/czi5qwyx4mb51.jpg?width=1912&amp;format=pjpg&amp;auto=webp&amp;s=a43c92123fb5807a5d001bfa85e5a9c347f9d99d)

I don't see any additional config options in the settings menu. Anyone know what step I'm likely missing for my .ts file breakpoints to be ignored here? 

I am running the bottom most script in the bottom left panel, `node compiled/index.js -b`  I want it to map to my ts files in the `src` folder so I can figure out where the problems are.

    // tsconfig.json for this project
    {
       "compilerOptions": {
           "target" : "ES5",
           "module": "commonjs",
           "noImplicitAny": true,
           "removeComments": true,
           "sourceMap": true
           , "outDir": "compiled/"
           , "esModuleInterop": true
           , "strict": true
           , "typeRoots": ["./src/types", "node_modules/@types"]
       },
       "include": ["src/**/*"]
       , "exclude": ["node_modules", "compiled", "__tests__", "types"]
    }
## [7][How do I get around this error?](https://www.reddit.com/r/typescript/comments/htm7aj/how_do_i_get_around_this_error/)
- url: https://www.reddit.com/r/typescript/comments/htm7aj/how_do_i_get_around_this_error/
---
I'm working on an in memory expressJs typescript project. It is currenlty in memory because I haven't set up any sort of database to work with.  I'm trying to set up a delete request to the api and get this error:

S7053: Element implicitly has an 'any' type because expression of type 'any' can't be used to index type '{}'

Here is the weights object:

*let* weights = {};

&amp;#x200B;

Here is my delete request:

&amp;#x200B;

app.delete('weights/:id', (req , res) =&gt; {  


// This is where the error is  
weights\[req.body.id\] = *null*;  
res.status(200).send({message: 'Delete successful.'})  
});

&amp;#x200B;

I apologize for the lack of formatting, it won't work for me. What am I missing?
## [8][Promise &gt; Observable](https://www.reddit.com/r/typescript/comments/htu56a/promise_observable/)
- url: https://www.stackchief.com/blog/Observable%20vs%20Promise%20%7C%20When%20to%20use%20Promise
---

## [9][How to dynamically assign a property](https://www.reddit.com/r/typescript/comments/htfsm7/how_to_dynamically_assign_a_property/)
- url: https://www.reddit.com/r/typescript/comments/htfsm7/how_to_dynamically_assign_a_property/
---
So I'm working on a React/Typescript project, and I am trying to make this piece of code work:

    const inputChange = (event: ChangeEvent&lt;HTMLInputElement&gt;) =&gt; {
            setState((draft) =&gt; {
                draft.client[event.target.name] = event.target.value;
            });
        };

First, I'm using Immer and useImmer for my state setting, so the above code is safe and properly modifying immutable state.

Second, my main problem is in the square brackets. This function receives a ChangeEvent from an input element. The "name" property on that element matches one of the keys on the "client" object. I want to be able to assign the input's value to that property using this dynamic syntax.

The problem is TypeScript errors out because of issues trying to figure out the types here. I'm really not sure how to handle this properly. Help would be appreciated. Here is the error:

    Element implicitly has an 'any' type because expression of type 'string' can't be used to index type '{ accessTokenTimeoutSecs?: number | undefined; allowAuthCode?: boolean | undefined; allowClientCredentials?: boolean | undefined; allowPassword?: boolean | undefined; clientKey?: string | undefined; ... 4 more ...; refreshTokenTimeoutSecs?: number | undefined; }'.
      No index signature with a parameter of type 'string' was found on type '{ accessTokenTimeoutSecs?: number | undefined; allowAuthCode?: boolean | undefined; allowClientCredentials?: boolean | undefined; allowPassword?: boolean | undefined; clientKey?: string | undefined; ... 4 more ...; refreshTokenTimeoutSecs?: number | undefined; }'

Edit: More investigation, the problem seems twofold. First, knowing what property [event.target.name](https://event.target.name) that I'm passing into the square brackets matches. Second, ensuring that event.target.value, which has type 'string', is acceptable by that property. Again, help would be appreciated. Thanks.

Edit 2: Thank you to everyone for your help. I decided to post my solution here.

    // Constants for all the property names go here, along with arrays grouping them by type.
    
    type ClientStringProperty = typeof NAME | typeof CLIENT_KEY | typeof CLIENT_SECRET;
    type ClientNumberProperty = typeof ACCESS_TOKEN_TIMEOUT | typeof REFRESH_TOKEN_TIMEOUT;
    type ClientBooleanProperty = typeof ENABLED | typeof ALLOW_CLIENT_CREDS | typeof ALLOW_PASSWORD | typeof ALLOW_AUTH_CODE;
    
    const isStringProperty = (name: string): name is ClientStringProperty =&gt; {
        return STRING_PROPS.includes(name);
    };
    
    const isNumberProperty = (name: string): name is ClientNumberProperty =&gt; {
        return NUMBER_PROPS.includes(name);
    };
    
    const isBooleanProperty = (name: string): name is ClientBooleanProperty =&gt; {
        return BOOLEAN_PROPS.includes(name);
    };
    
    // Inside component
    const inputChange = (event: ChangeEvent&lt;HTMLInputElement&gt;) =&gt; {
            const { name, value, checked } = event.target;
            setState((draft) =&gt; {
                if (isStringProperty(name)) {
                    draft.client[name] = value;
                } else if (isNumberProperty(name)) {
                    draft.client[name] = value ? parseInt(value) : 0;
                } else if (isBooleanProperty(name)) {
                    draft.client[name] = checked;
                }
            });
        };

&amp;#x200B;
## [10][Excess property check question](https://www.reddit.com/r/typescript/comments/htbn6t/excess_property_check_question/)
- url: https://www.reddit.com/r/typescript/comments/htbn6t/excess_property_check_question/
---
https://www.typescriptlang.org/docs/handbook/interfaces.html#excess-property-checks

It says "If an object literal has any properties that the “target type” doesn’t have, you’ll get an error:". The error is only because of the typo of color/colour. So by excess property checking does it mean it typescript checks for what it thinks is an error, i.e. a typo error? Because in the previous example when it was

    interface LabeledValue {
        label: string;
    }
    
    function printLabel(labeledObj: LabeledValue) {
        console.log(labeledObj.label);
    }
    
    let myObj = {size: 10, label: "Size 10 Object"};
    printLabel(myObj);

"it’s only the shape that matters. If the object we pass to the function meets the requirements listed, then it’s allowed.", so it's allowed here even though there's no size. So what I'm trying to figure out is what's the criteria for the typescript amigo to judge? Possible typos and?
## [11][TypeScript developers interested in decentralization should be aware of Shardus.](https://www.reddit.com/r/typescript/comments/htmh0w/typescript_developers_interested_in/)
- url: https://medium.com/@Shardus/shardus-the-foundation-of-our-decentralized-future-976ae5106938?source=social.tw
---

