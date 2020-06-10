# typescript
## [1][Who's hiring Typescript developers - June](https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/)
- url: https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/
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
## [2][What's a good charting library that I can use with TypeStyle &amp; React?](https://www.reddit.com/r/typescript/comments/h02b26/whats_a_good_charting_library_that_i_can_use_with/)
- url: https://www.reddit.com/r/typescript/comments/h02b26/whats_a_good_charting_library_that_i_can_use_with/
---
My understanding is that most charting libraries rely on traditional style rules, so wouldn't play well with TypeStyle (or styled-components) - is that accurate? If so, what recommendations do people have?
## [3][Help with return values out of `reduce` method](https://www.reddit.com/r/typescript/comments/h06rlf/help_with_return_values_out_of_reduce_method/)
- url: https://www.reddit.com/r/typescript/comments/h06rlf/help_with_return_values_out_of_reduce_method/
---
I have an interface `ConfigData` that identifies a couple of key/value pairs (`languageCode: string`, and `numberOfRepeats: number`).

Not every step object being reduced will contribute to the accumulator. Of 5 steps passing through, only 2 will add a k/v pair.

I want `const configData` in the example below to receive a `ConfigData` object after the reduce method evaluates. But my code is incorrect as is in (at least) two places:

       run(): ConfigData {
    
          const configData = this.steps.reduce((accumulator: Partial&lt;ConfigData&gt;, currentStep: WizardSteps): ConfigData =&gt; {
             // instruct and ask for feedback
             currentStep.explain();
             const rawUserResponse: string = currentStep.prompt();
    
             if (currentStep.needsInputValidation) {
                const validatedInput: string = currentStep.validate(rawUserResponse);
                
                return {
                   [currentStep.configDataKey] : validatedInput
    // const validatedInput: string
    // Type '{ [x: string]: string; }' is missing the following 
    // properties from type 'ConfigData': languageCode, 
    // numberOfRepeats ts(2739)
                };
             }
    
             if (currentStep.needsFileValidation) {
                // blocks until file is validated
                currentStep.validate();
             }
    
    
          }, {});
    
          return configData;
    // const configData: Partial&lt;ConfigData&gt;
    // Type 'Partial&lt;ConfigData&gt;' is not assignable to type 'ConfigData'.
    // Property 'languageCode' is optional in type 'Partial&lt;ConfigData&gt;' 
    // but required in type 'ConfigData'. ts(2322)
       }

I'm having a hard time with both of these. The second one is at least comprehensible but I'm lost on how to fix both. Anyone know how?
## [4][Does typescript works with pnpm?](https://www.reddit.com/r/typescript/comments/h043pc/does_typescript_works_with_pnpm/)
- url: https://www.reddit.com/r/typescript/comments/h043pc/does_typescript_works_with_pnpm/
---
Pnpm is a package manager that uses symlinks for node modules. Yesterday I started learning typescript but the code doesn't runs with pnpm. Using the npm instead of pnpm works. I stumbled upon few stackoverflow and github questions but none had the answer.
## [5][Creating typescript libraries / packages.](https://www.reddit.com/r/typescript/comments/gzu67w/creating_typescript_libraries_packages/)
- url: https://www.reddit.com/r/typescript/comments/gzu67w/creating_typescript_libraries_packages/
---
When setting up a new TS lib or package, what are your preferred methods to get this done simply and effectively. Recently each method I have tried (webpack, tsdx, etc) I always seem to hit some config issues which mean I am not writing actual code for half a day.

I would love to hear this groups methods for setting up packages painlessly which can be published to NPM etc.

Thanks!
## [6][[HEP] Type checking mixed arrays](https://www.reddit.com/r/typescript/comments/h0797r/hep_type_checking_mixed_arrays/)
- url: https://www.reddit.com/r/typescript/comments/h0797r/hep_type_checking_mixed_arrays/
---
How I am supposed to type check this function?

    type Component = Position | Health | Speed | Combat |  Parent
    
    export class Entity {
    	protected components = [];
    
            addComponent&lt;T&gt;(component: T): T {
    		this.components.push(component)
    
    		return component
    	}
    
    	getComponents(...components) {
    		return components.map((component) =&gt; this.getComponent(component));
    	}
    }

&amp;#x200B;

This is how I want it to be called. As for the components and retrieve them in the order you called them. There are more components and they can be in different order.

    const [healthTarget, combatTarget, parentTarget] = target.getComponents(
        Health.prototype, Combat.prototype, Parent.prototype
    )

&amp;#x200B;
## [7][Article: How to Use Attribute Directives to Avoid Repetition in Angular Templates](https://www.reddit.com/r/typescript/comments/h074m8/article_how_to_use_attribute_directives_to_avoid/)
- url: https://volosoft.com/blog/attribute-directives-to-avoid-repetition-in-angular-templates
---

## [8][Can all JS Code be expressed in TS without the use of "any"? (Example)](https://www.reddit.com/r/typescript/comments/h03n06/can_all_js_code_be_expressed_in_ts_without_the/)
- url: https://www.reddit.com/r/typescript/comments/h03n06/can_all_js_code_be_expressed_in_ts_without_the/
---
Hi everyone. I'm really struggling trying to figure out how to write this code in TypeScript. I have ESLint setup to warn against all uses of "any" and want to avoid "any" bleeding throughout my code and making TypeScript useless. After spending 4 hours on this one function I have to ask, are there just some things in JavaScript which you cannot express in TypeScript without using "any"?

Here's what I'm working on:

```
async function depaginate (apiMethod, params) {
    const result = {};
    do {
        const res = await apiMethod(params).promise();
        for (const key in res) {
            if (Array.isArray(result[key])) {
                result[key].push(...res[key]);
            } else {
                result[key] = res[key];
            }
        }
        params.NextToken = res.NextToken;
    } while (typeof params.NextToken === "string");
    return result;
}
```

The `depaginate` function can take any function as the first argument, an object as the second argument and builds a result combining the arrays of all keys in the API response.

The issue here is that `apiMethod` can be one of thousands of functions and `params` can be one of thousands of types so I cannot explicitly type them.

I tried using Generics for the arguments, but still ran into tons of issues. Can anyone help?
## [9][Public method or prototype?](https://www.reddit.com/r/typescript/comments/h05opn/public_method_or_prototype/)
- url: https://www.reddit.com/r/typescript/comments/h05opn/public_method_or_prototype/
---
Hello.  Consider this `class`:

```ts
class Dog {
  private name: string;
  constructor (name: string) {
    this.name = name;
  }
  // 1. public method
  public walk() {
    return `${this.name} went for a walk today.`
  }
}

// 2. prototype inheritance
Dog.prototype.walk = function () {
  return `${this.name} went for a walk today.` 
}
```
Are `1` and `2` the same?  What are the advantages and disadvantages of using one over the other?

And should I implement `.toString()` method as `public` method or using `Dog.prototype.toString`?
## [10][Appwrite 0.6 Comes with Full TypeScript Support](https://www.reddit.com/r/typescript/comments/gzo6tj/appwrite_06_comes_with_full_typescript_support/)
- url: https://medium.com//appwrite-0-6-comes-with-full-typescript-support-b207caa47a1a?source=friends_link&amp;sk=265033b4520c260d0d6a8757985bff09
---

## [11][Fence your TypeScript, for saner project boundaries](https://www.reddit.com/r/typescript/comments/gzqprm/fence_your_typescript_for_saner_project_boundaries/)
- url: https://zohaib.me/fence-your-typescript/
---

