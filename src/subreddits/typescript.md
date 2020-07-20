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
## [2][How to Infer Types from Schema](https://www.reddit.com/r/typescript/comments/hub8rd/how_to_infer_types_from_schema/)
- url: https://www.reddit.com/r/typescript/comments/hub8rd/how_to_infer_types_from_schema/
---
Hi all. I've recently started using typescript more heavily, and  I'm working a generic form-validation library. One of the things I'd like to be able to do is pass in a data model, and be able to get type inference/intellisense from the keys from that object. I've not learned all the details of how to use generics yet, which I know will be key for this question, so any pointers towards what I am missing would be greatly appreciated.

A simplified example:

```
interface SpecificSchemaInterface {
    fields: Record&lt;String, any&gt;
}

interface SpecificSchemaInterface {
    fields: {
      firstName: string;
      lastName: string;
      age: number;
    }
}

const schema = {
  fields: {
    firstName: 'Bob,
    lastName: 'Bobson',
    age: 25,
  },
};

function createValidatorFromSchema(s: SpecificSchemaInterface) {
  // ???
}

const validator: Validator&lt;SpecificSchemaInterface&gt; = createValidatorFromSchema(schema)

// Then validator would have an interface like:
interface Validator {
  rules: {
    firstName(firstName: string): boolean,
    lastName(lastName: string): boolean,
    age(age: number): boolean, // note that the type is inferred here from the schema
  }
}
```

My questions are that then
1) I probably want something like `Validator&lt;SpecificSchemaInterface&gt;`?
2) I need to be able to both extract the object keys, and the types from the schema in a reusable way inside my new `Validator`

I hope this isn't too vague, but if there is some terminology I should be looking into (beyond just "look at generics") or guides that go into this sort of thing, that would be very helpful.
## [3][Am I generating `index.d.ts` file for an `npm` module `@camwiegert/typical` correctly?](https://www.reddit.com/r/typescript/comments/hugbmu/am_i_generating_indexdts_file_for_an_npm_module/)
- url: https://www.reddit.com/r/typescript/comments/hugbmu/am_i_generating_indexdts_file_for_an_npm_module/
---
I am a noob at TS &amp; certainly don't know how to write `index.d.ts`. Guided by some auto-completion inside VSCode &amp; some of my existing knowledge, I managed to cobble this together.

I am writing types for `@camwiegert/typical` npm module. Source code could be found at https://github.com/camwiegert/typical

Here is the main file `typical.js` for brevity:

```js
export async function type(node, ...args) {
    for (const arg of args) {
        switch (typeof arg) {
            case 'string':
                await edit(node, arg);
                break;
            case 'number':
                await wait(arg);
                break;
            case 'function':
                await arg(node, ...args);
                break;
            default:
                await arg;
        }
    }
}

async function edit(node, text) {
    const overlap = getOverlap(node.textContent, text);
    await perform(node, [...deleter(node.textContent, overlap), ...writer(text, overlap)]);
}

async function wait(ms) {
    await new Promise(resolve =&gt; setTimeout(resolve, ms));
}

async function perform(node, edits, speed = 60) {
    for (const op of editor(edits)) {
        op(node);
        await wait(speed + speed * (Math.random() - 0.5));
    }
}

export function* editor(edits) {
    for (const edit of edits) {
        yield (node) =&gt; requestAnimationFrame(() =&gt; node.textContent = edit);
    }
}

export function* writer([...text], startIndex = 0, endIndex = text.length) {
    while (startIndex &lt; endIndex) {
        yield text.slice(0, ++startIndex).join('');
    }
}

export function* deleter([...text], startIndex = 0, endIndex = text.length) {
    while (endIndex &gt; startIndex) {
        yield text.slice(0, --endIndex).join('');
    }
}

export function getOverlap(start, [...end]) {
    return [...start, NaN].findIndex((char, i) =&gt; end[i] !== char);
}
```

I want to write types for this file. So I made `src/@types/@camwiegert/typical/index.d.ts` like:

```ts
declare module '@camwiegert/typical' {
	export function type(node: HTMLElement, ...args: any[]): Promise&lt;void&gt;

	declare function edit(node: HTMLElement, text: string): Promise&lt;void&gt;

	declare function wait(ms: number): Promise&lt;void&gt;

	declare function perform(node: HTMLElement, edits: any, speed?: number): Promise&lt;void&gt;

	export function editor(edits: any): Generator&lt;(node: any) =&gt; number, void, unknown&gt;

	export function writer(
		[...text]: Iterable&lt;any&gt;,
		startIndex?: number,
		endIndex?: number
	): Generator&lt;string, void, unknown&gt;

	export function deleter(
		[...text]: Iterable&lt;any&gt;,
		startIndex?: number,
		endIndex?: number
	): Generator&lt;string, void, unknown&gt;

	export function getOverlap(start: any, [...end]: Iterable&lt;any&gt;): number
}
```

Idk if I converted it correctly or not. Can anyone check if it's correct?
## [4][Custom hook with typescript?](https://www.reddit.com/r/typescript/comments/hub5ci/custom_hook_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hub5ci/custom_hook_with_typescript/
---
Hey guys, I'm pretty stuck as to how to get around using :any here. Any ideas?

&amp;#x200B;

    import { useState } from 'react';
    
    export const useFormFields = (initialState: any) =&gt; {
      const [fields, setValues] = useState(initialState);
    
      return [
        fields,
        (event: React.ChangeEvent&lt;HTMLInputElement&gt;) =&gt; {
          setValues({
            ...fields,
            [event.target.name]: event.target.value,
          });
        },
      ];
    };

I've tried using interfaces and types and then adding it in, but ...fields expects an object and the whole thing breaks.
## [5][A directory structure for React projects](https://www.reddit.com/r/typescript/comments/hui3yt/a_directory_structure_for_react_projects/)
- url: https://medium.com//a-directory-structure-for-react-projects-99fb084c61f1?source=friends_link&amp;sk=e5c0c00d0e03e6661ec7e710f84eebfb
---

## [6][React Router Is Not Working With Typescript?](https://www.reddit.com/r/typescript/comments/hufp62/react_router_is_not_working_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hufp62/react_router_is_not_working_with_typescript/
---
When the image is clicked in **Header Component via &lt;Link&gt;**, React is not rendering **Content Component**

Visual representation of the problem:

[https://gyazo.com/bb14af9cb3f0870107ec775a328e60c0](https://gyazo.com/bb14af9cb3f0870107ec775a328e60c0)

&amp;#x200B;

I'm not sure what the issue is, my previous react projects works fine with this setup, but here's what i have

&amp;#x200B;

    import { Route, Router, Switch } from "react-router-dom";
    import history from "../history";
    
    const App: React.FC&lt;{}&gt; = () =&gt; {
        return (
            &lt;React.Fragment&gt;
                &lt;Router history={history}&gt;
                    &lt;Header /&gt;
                    &lt;Switch&gt;
                        &lt;Route path="/" exact component={Content} /&gt;
                    &lt;/Switch&gt;
                &lt;/Router&gt;
            &lt;/React.Fragment&gt;
        );
    };
    
    export default App;

**history.tsx**

    import { createBrowserHistory } from "history";
    export default createBrowserHistory();

&amp;#x200B;

**Header**

    const Header: React.FC&lt;HeaderProps&gt; = (props) =&gt; {
        //const history = useHistory();
        return (
            &lt;nav&gt;
                &lt;Link to="/"&gt;
                    &lt;img
                        className="logo"
                        src={logo}
                        alt="pixar-logo"
                    /&gt;
                &lt;/Link&gt;
    ...
    export default Header;

**Content**

    const Content: React.FC&lt;{}&gt; = () =&gt; {
        return (
            &lt;div&gt;
                &lt;div className="contentContainer"&gt;
                    &lt;h2 className="quote"&gt;
    ...
    
    export default Content;
## [7][Any VSCode users? How to stop it from thinking type constructors are JSX code?](https://www.reddit.com/r/typescript/comments/htx1ez/any_vscode_users_how_to_stop_it_from_thinking/)
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
## [8][What's the weirdest, most odd-looking and/or most complex generic you've seen or made?](https://www.reddit.com/r/typescript/comments/htukg1/whats_the_weirdest_most_oddlooking_andor_most/)
- url: https://www.reddit.com/r/typescript/comments/htukg1/whats_the_weirdest_most_oddlooking_andor_most/
---

## [9][Finally I Developed SQL Query Generating App](https://www.reddit.com/r/typescript/comments/htw2nd/finally_i_developed_sql_query_generating_app/)
- url: https://youtu.be/K9uG63QSW-I
---

## [10][Is it acceptable to use switch(true) to match the first met condition?](https://www.reddit.com/r/typescript/comments/htqokr/is_it_acceptable_to_use_switchtrue_to_match_the/)
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
## [11][What is needed to get source maps working?](https://www.reddit.com/r/typescript/comments/htg5jx/what_is_needed_to_get_source_maps_working/)
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
