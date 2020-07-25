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
## [2][Handling real-time communications - is Socket.IO still a good choice?](https://www.reddit.com/r/typescript/comments/hxhr1c/handling_realtime_communications_is_socketio/)
- url: https://www.reddit.com/r/typescript/comments/hxhr1c/handling_realtime_communications_is_socketio/
---
Back in 2014, I built my first React app using [Socket.IO](https://Socket.IO) because it contained a real-time communication component.  (It was basically Kahoot).  I've been thinking about revisiting real time communication. Chatrooms, basically, and then once I get that working, maybe build some party-game clones (like Cards Against Humanity or a clone of some of the Jackbox games).  At any rate, I'm just wondering if I'm overlooking some information on real time data. 

I know Firebase is often considered for real-time communication, but I'd kind of like to keep this independent of requiring 3rd party services - at least off the bat.
## [3][Strict mode with Typescript &amp; Create React App](https://www.reddit.com/r/typescript/comments/hxlz81/strict_mode_with_typescript_create_react_app/)
- url: https://www.reddit.com/r/typescript/comments/hxlz81/strict_mode_with_typescript_create_react_app/
---
CRA tries to hide away alot of its config files unless you eject. Is it possible to enable strict mode in Typescript without ejecting?
## [4][(babel plugin) TypeError: Path was expected to have a container](https://www.reddit.com/r/typescript/comments/hxk1ef/babel_plugin_typeerror_path_was_expected_to_have/)
- url: https://www.reddit.com/r/typescript/comments/hxk1ef/babel_plugin_typeerror_path_was_expected_to_have/
---
I am trying to convert https://github.com/camwiegert/typical/blob/master/typical.js into TypeScript.

In Vanilla JS, it looks like:

# typical.js

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

In TypeScript, it looks like:

# typical.ts

```ts
export async function type(node: HTMLElement, ...args: any[]): Promise&lt;void&gt; {
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

async function edit(node: HTMLElement, text: string): Promise&lt;void&gt; {
  const overlap = getOverlap(node.textContent, text);
  await perform(node, [
    ...deleter(node.textContent, overlap),
    ...writer(text, overlap),
  ]);
}

async function wait(ms: number): Promise&lt;void&gt; {
  await new Promise(resolve =&gt; setTimeout(resolve, ms));
}

async function perform(
  node: HTMLElement,
  edits: Iterable&lt;string | null&gt;,
  speed: number = 60
): Promise&lt;void&gt; {
  for (const op of editor(edits)) {
    op(node);
    await wait(speed + speed * (Math.random() - 0.5));
  }
}

export function* editor(
  edits: Iterable&lt;string | null&gt;
): Generator&lt;(node: any) =&gt; number, void, unknown&gt; {
  for (const edit of edits) {
    yield node =&gt; requestAnimationFrame(() =&gt; (node.textContent = edit));
  }
}

export function* writer(
  [...text]: string,
  startIndex: number = 0,
  endIndex: number = text.length
): Generator&lt;string, void, unknown&gt; {
  while (startIndex &lt; endIndex) {
    yield text.slice(0, ++startIndex).join('');
  }
}

export function* deleter(
  text: string | null,
  startIndex: number = 0
): Generator&lt;string, void, unknown&gt; {
  if (text === null) {
    yield '';
    return;
  }
  const [...chars] = text;
  let endIndex = chars.length;
  while (endIndex &gt; startIndex) {
    yield chars.slice(0, --endIndex).join('');
  }
}

export function getOverlap(start: string | null, [...end]: string): number {
  if (start === null) return 0;
  return [...start, NaN].findIndex((char, i) =&gt; end[i] !== char);
}
```

I am getting an error saying `Path was expected to have a container!` which I am not getting. It shows the following in the terminal:

```bash
1 | export async function type(node, ...args) {
  2 |     for (const arg of args) {
&gt; 3 |         switch (typeof arg) {
    |         ^
  4 |             case 'string':
  5 |                 await edit(node, arg);
  6 |                 break;
TypeError: ~/react-typical/src/typical.ts: Path was expected to have a container!
  1 | export async function type(node, ...args) {
  2 |     for (const arg of args) {
&gt; 3 |         switch (typeof arg) {
    |         ^
  4 |             case 'string':
  5 |                 await edit(node, arg);
  6 |                 break;
    at File.buildCodeFrameError (~/react-typical/node_modules/@babel/core/lib/transformation/file/file.js:248:12)
    at NodePath.buildCodeFrameError (~/react-typical/node_modules/@babel/traverse/lib/path/index.js:144:21)
    at checkPathValidity (~/react-typical/node_modules/babel-plugin-transform-async-to-promises/async-to-promises.js:1085:24)
    at relocateTail (~/react-typical/node_modules/babel-plugin-transform-async-to-promises/async-to-promises.js:1095:9)
    at Object.rewriteAwaitOrYieldPath (~/react-typical/node_modules/babel-plugin-transform-async-to-promises/async-to-promises.js:2685:21)
    at NodePath._call (~/react-typical/node_modules/@babel/traverse/lib/path/context.js:55:20)
    at NodePath.call (~/react-typical/node_modules/@babel/traverse/lib/path/context.js:42:17)
    at NodePath.visit (~/react-typical/node_modules/@babel/traverse/lib/path/context.js:90:31)
    at TraversalContext.visitQueue (~/react-typical/node_modules/@babel/traverse/lib/context.js:112:16)
    at TraversalContext.visitSingle (~/react-typical/node_modules/@babel/traverse/lib/context.js:84:19)
```

How do I solve this?

PS: I posted https://www.reddit.com/r/typescript/comments/hwzeo6/ts2488_type_string_null_must_have_a/ yesterday &amp; I did everything suggested there but couldn't solve it completely
## [5][Rest API app built with Deno and TypeScript](https://www.reddit.com/r/typescript/comments/hwzjdu/rest_api_app_built_with_deno_and_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hwzjdu/rest_api_app_built_with_deno_and_typescript/
---
Part of the reason many people are now paying attention to Deno is not only its ambitious goal to compete with Node.js, but also its built-in TypeScript support. Do you want to know what it's like to develop a REST API application with Deno. My friend [recounted](https://tsh.io/blog/deno-rest-api/) his experience with that. Thought you might find it interesting.
## [6][Generic type index signatures vs symbols](https://www.reddit.com/r/typescript/comments/hxj4bd/generic_type_index_signatures_vs_symbols/)
- url: https://www.reddit.com/r/typescript/comments/hxj4bd/generic_type_index_signatures_vs_symbols/
---
Trying to move the body type definition to its own interface, I converted it to a symbol to get it working:

    // Code Sample #1
    const keyIndex: unique symbol = Symbol();
    
    export default interface BodyInterface {
      [keyIndex]: {
         name: string,
         element: JSX.Element
       }
    }

Before it was like this:

    // Code Sample #2
    type navItemKeys = 'cover' | 'portfolio' | 'aboutContact';
    
    interface PropsShape &lt;keyNames extends string&gt; {
      body: {
        [keyIndex in keyNames]: {
          name: string,
          element: JSX.Element
        }
      }
    }
    
    // ------------------- Main Component
    export default (props: PropsShape&lt;navItemKeys&gt;) =&gt; {

I am not fully aware of what's going on here and hope someone can correct my thinking on the matter. My best guess is this:

Code Sample #2

* The `PropsShape` interface is taking the type argument `navItemKeys`.
* This argument `PropsShape` is an extension of the string type
* Moving into the interface type definition, the index signature is set to the generic type `keyNames`. Multiple properties of this type is inferred for a reason I don't fully understand
* I have another type `keyIndex` using the `in` keyword. It could have been set to `K` or anything else. It checks for a property of `keyNames`.
* I'm very unclear on what `keyIndex` actually is. Should I count it as just another generic type? If so, it seems to have sprung out of nowhere. Why was it not passed to the interface like this:

&amp;#x200B;

    interface PropsShape &lt;keyNames extends string, keyIndex&gt; {

Code Sample #1

* This one actually feels a little more sensible. A generic "Symbol' type is constructed and used as a generic index signature
* I believe the symbol type's main use case is exactly as a sort of generic index signature type
* Is this equivalent in functionality to Code Sample #2?
## [7][Help setting up a generic prop interface](https://www.reddit.com/r/typescript/comments/hx9g0x/help_setting_up_a_generic_prop_interface/)
- url: https://www.reddit.com/r/typescript/comments/hx9g0x/help_setting_up_a_generic_prop_interface/
---
Originally I hardcoded the expected property keys of the `body` prop:

    interface PropsShape {
      body: {
        cover: { name: string, element: JSX.Element },
        portfolio: { name: string, element: JSX.Element },
        aboutContact: { name: string, element: JSX.Element },
      }
    }
    
    export default (props: PropsShape) =&gt; {
      
      return (
        &lt;nav&gt;
          &lt;p&gt;placeholder&lt;/p&gt;
        &lt;/nav&gt;
      )
    }

I wanted to do away with hardcoded keys of the `body` prop, and thought a generic type as the property name would do it, but I'm not quite there. I need more practice with generics. What is the mistake below?

    interface navItemKeys {
      cover: string,
      portfolio: string,
      aboutContact: string,
    };
    
    interface PropsShape &lt;keyName&gt; {
      body: {
        [keyName as keyof navItemKeys]: { name: string, element: JSX.Element }
      }
    }
    
    export default (props: PropsShape&lt;keyof navItemKeys&gt;) =&gt; {
      
      return (
        &lt;nav&gt;
          &lt;p&gt;placeholder&lt;/p&gt;
        &lt;/nav&gt;
      )
    }
## [8][[OpenJS Foundation] Deno, a Secure Runtime for JavaScript and TypeScript - Ryan Dahl, Deno Land](https://www.reddit.com/r/typescript/comments/hx3jaz/openjs_foundation_deno_a_secure_runtime_for/)
- url: https://www.youtube.com/watch?v=doug6st5vAs
---

## [9][TypeDraft: Language is the New Framework](https://www.reddit.com/r/typescript/comments/hxgacb/typedraft_language_is_the_new_framework/)
- url: https://medium.com/@mistlog/typedraft-language-is-the-new-framework-84dfc433971f
---

## [10][How can I set a property as optional or not depending on a case?](https://www.reddit.com/r/typescript/comments/hx51e9/how_can_i_set_a_property_as_optional_or_not/)
- url: https://www.reddit.com/r/typescript/comments/hx51e9/how_can_i_set_a_property_as_optional_or_not/
---
Hi, I'd like to know what's best practice for this scenario which I'm sure is pretty common.

Let's say I have an entity called "Post" and the type is defined as the following:

``
type Post = {
  id: number,
  name: string,
  body: string,
  keywords: string
}
``

Now let's say in my API I have two methods: create and update and for these two methods I want to use DTOs for the create and update endpoints that look similar to the Post type, but exclude the id property, since that will be generated by the API. The two DTOs would be identical except that the Create DTO would make name and body required and keywords optional, where as the Update DTO, all properties would be optional since the api method will only update the properties it gets.

My question is should I create two separate DTOs, where the only difference is adding `?` on the name property? Or should I create a single SavePostDto that would look like this:

``
type SavePostDto = {
  name?: string,
  body?: string,
  keywords?: string
}
``

The issue with one DTO is that name and body should be required if we are creating the post. But creating two DTOs that are identical in properties seems annoying since you always have to remember to update both DTOs anytime you change the Post type. 

Is there another way to approach this that I'm not considering?
## [11][How to increase code coverage for mocha testing on basic Typescript example](https://www.reddit.com/r/typescript/comments/hwtsd3/how_to_increase_code_coverage_for_mocha_testing/)
- url: https://www.reddit.com/r/typescript/comments/hwtsd3/how_to_increase_code_coverage_for_mocha_testing/
---
I want to get a better understanding of how to test these kinds of functions in Typescript using mocha. In the StackOverflow post below I've outlined a basic small Typescript file, which I want to test and increase code coverage. Right now I'm not sure why the breakdown of the testing is the way that it is (50% Stmts | 100% Branch | 0% Funcs | 50% Lines) and how I would be able to get everything to 100%.

[https://stackoverflow.com/questions/63065938/how-to-increase-code-coverage-for-mocha-testing-on-basic-typescript-example](https://stackoverflow.com/questions/63065938/how-to-increase-code-coverage-for-mocha-testing-on-basic-typescript-example)

Any advice would greatly be appreciated!
