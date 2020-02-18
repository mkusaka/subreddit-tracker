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
## [2][How does your typescript code sound? The app creates a melody from the JS or TS code you enter](https://www.reddit.com/r/typescript/comments/f5q86v/how_does_your_typescript_code_sound_the_app/)
- url: https://soundcode.now.sh/
---

## [3][Any strategies to make typescript launch faster?](https://www.reddit.com/r/typescript/comments/f5oacn/any_strategies_to_make_typescript_launch_faster/)
- url: https://www.reddit.com/r/typescript/comments/f5oacn/any_strategies_to_make_typescript_launch_faster/
---
I'm using ts-node and jest. Launching a test takes forever. Typescript takes about 1-10 seconds more to launch than simple node. There doesn't seem to be any feature to tell me which files are taking the longest so there's no way to optimize. My project is rather large, hundreds of files, 100k lines+. Some files are required dynamically, which might cause a problem but no way to really know how much of a problem it is.
## [4][I still don't entirely understand decorators in TypeScript?](https://www.reddit.com/r/typescript/comments/f5e59h/i_still_dont_entirely_understand_decorators_in/)
- url: https://www.reddit.com/r/typescript/comments/f5e59h/i_still_dont_entirely_understand_decorators_in/
---
That's really my question, I just don't understand what they are exactly. That doesn't mean I can't use them, but don't understand what function they serve in my code?   Thanks in advance.
## [5][proxy-watcher - A function that recursively watches an object for mutations via Proxies and tells you which paths changed](https://www.reddit.com/r/typescript/comments/f5he2l/proxywatcher_a_function_that_recursively_watches/)
- url: https://github.com/fabiospampinato/proxy-watcher
---

## [6][decoupling a complex set of generic interfaces from its much simpler TypeScript code](https://www.reddit.com/r/typescript/comments/f5of1v/decoupling_a_complex_set_of_generic_interfaces/)
- url: https://www.reddit.com/r/typescript/comments/f5of1v/decoupling_a_complex_set_of_generic_interfaces/
---
I am writing a TypeScript library, and I am trying to add generics to the public-facing functions in order to the strengthen their type checking and improve IDE autosuggestions. The problem is that the generics end up needing to be passed all throughout the codebase because some of the public-facing functions are also used internally. And the generic types are also a bit complex, with multiple parameters, ternary logic, lookups, and references to other custom generic types. I don't want to rewrite every function with these new types. The public-facing functions need them, the internals do not.

I want to write the generic types so that they behave as the library's interface but do not affect any of the existing internal code. Like a layer that just sits in front, completely decoupled, much like how type declaration files are for JS files. Can a TypeScript library have its own declaration files? How would you suggest I solve this?
## [7][Is a generic deep flatten function possible in TS?](https://www.reddit.com/r/typescript/comments/f5d1ud/is_a_generic_deep_flatten_function_possible_in_ts/)
- url: https://www.reddit.com/r/typescript/comments/f5d1ud/is_a_generic_deep_flatten_function_possible_in_ts/
---
Is it possible to write a generic TS function for recursively flattening an arbitrarily nested array?

For example:

    const input = [1,[2,[3,[4,[5],6],7],8],9];
    const output = flattenDeep(input); // [1, 2, 3, 4, 5, 6, 7, 8, 9]

This is pretty simple to write in JS

    function flattenDeep(array) {
      const result = [];
      for(const value of array) {
        if(Array.isArray(value)) {
          result.push(...flattenDeep(value));
        } else {
          result.push(value);
        }
      }
      return result;
    }

But is there a way to add generic types to this, so it will infer the input/output types in the example above?
## [8][No compile check for generics?](https://www.reddit.com/r/typescript/comments/f5cu8o/no_compile_check_for_generics/)
- url: https://www.reddit.com/r/typescript/comments/f5cu8o/no_compile_check_for_generics/
---
Given the following code:

&gt;interface Contract&lt;T&gt; {  
}  
class Deal&lt;D&gt; implements Contract&lt;D&gt; {  
}  
class Agreement&lt;A&gt; implements Contract&lt;A&gt; {  
}  


I would have expected this not to compile:

&amp;#x200B;

&gt;let deal:Contract&lt;number&gt; = new Deal&lt;number&gt;()  
let agreement:Contract&lt;string&gt; = new Agreement&lt;string&gt;()  
&gt;  
&gt;  
// expects this to not compile  
agreement = deal;

or this

&gt;let deal:Deal&lt;number&gt; = new Deal&lt;number&gt;()  
let agreement:Agreement&lt;string&gt; = new Agreement&lt;string&gt;()  
&gt;  
&gt;  
// expects this to not compile  
agreement = deal;

But they both compile!

Here is a playground link [https://www.typescriptlang.org/play/index.html?ssl=17&amp;ssc=18&amp;pln=17&amp;pc=26#code/JYOwLgpgTgZghgYwgAgMIHtxUWAPAFQD5kBvAKAF8yyEAbOAZweQBEI5bcXjgBbAB1oReEcMwxYcXYuSo16TZAEEA5lAjDReJTwFCRYtJjDYE2mZWpCwyACbtaALgkmpIAK68ARtGIBeZBAIAHdWB1wPb18ACgBKMmtkODUNAzBnY1M8BhNQFX9AkOUUzXBcHKg8wjjqAHpa5AgAD34IM2YwAAtgDvRA9BsEdAFgITJk9VKbAPsOAG46hsTZpzYOCM8fKAKg0LXOSK3q+PrkRInUrUdVSbTy3JB85ADd4tute8rH48XkU+bWu1kF0esC+uBkEMRmNThcps87A45kA](https://www.typescriptlang.org/play/index.html?ssl=17&amp;ssc=18&amp;pln=17&amp;pc=26#code/JYOwLgpgTgZghgYwgAgMIHtxUWAPAFQD5kBvAKAF8yyEAbOAZweQBEI5bcXjgBbAB1oReEcMwxYcXYuSo16TZAEEA5lAjDReJTwFCRYtJjDYE2mZWpCwyACbtaALgkmpIAK68ARtGIBeZBAIAHdWB1wPb18ACgBKMmtkODUNAzBnY1M8BhNQFX9AkOUUzXBcHKg8wjjqAHpa5AgAD34IM2YwAAtgDvRA9BsEdAFgITJk9VKbAPsOAG46hsTZpzYOCM8fKAKg0LXOSK3q+PrkRInUrUdVSbTy3JB85ADd4tute8rH48XkU+bWu1kF0esC+uBkEMRmNThcps87A45kA)

I would have expected that a Generic&lt;A&gt; is not the same as Generic&lt;B&gt; and the compiler should not allow assigning one to the other. What am I missing here?
## [9][swc is a typescript / javascript compiler. It's 20x faster than babel.](https://www.reddit.com/r/typescript/comments/f53z29/swc_is_a_typescript_javascript_compiler_its_20x/)
- url: https://swc-project.github.io/
---

## [10][TypeScript Generics. Discussing naming conventions](https://www.reddit.com/r/typescript/comments/f58ox4/typescript_generics_discussing_naming_conventions/)
- url: https://wanago.io/2020/02/17/typescript-generics-discussing-naming-conventions/
---

## [11][Typescript Namespace Tricks](https://www.reddit.com/r/typescript/comments/f5g906/typescript_namespace_tricks/)
- url: https://www.reddit.com/r/typescript/comments/f5g906/typescript_namespace_tricks/
---
Hello! I wanted to share a blog post I wrote, [Typescript Namespace Tricks](https://www.gatlin.io/blog/post/typescript-namespace-tricks?utm_source=reddit&amp;utm_medium=link&amp;utm_campaign=1-blog-post), about ways to use Typescript Namespaces that I wish I knew back when I first started writing Typescript. I hope you enjoy :D
