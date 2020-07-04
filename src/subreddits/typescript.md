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
## [2][Typescript and mutually exclusive variables ?](https://www.reddit.com/r/typescript/comments/hl1af0/typescript_and_mutually_exclusive_variables/)
- url: https://www.reddit.com/r/typescript/comments/hl1af0/typescript_and_mutually_exclusive_variables/
---
I know that there are Boolean types for true or false. However I would like to have a value pairs, or some might call them interdependent polarities. The world around us is full of that. Like temperature = is warm or is cold. Or age = young vs old. Light = dark | bright. I was thinking about two methods in an object that will set to true or false the corresponding properties. But again, these are propreties and not actual variables. Or maybe there is some other way? It must be simpler than I think it is
## [3][WHERE to put my types/interfaces?](https://www.reddit.com/r/typescript/comments/hknqf4/where_to_put_my_typesinterfaces/)
- url: https://www.reddit.com/r/typescript/comments/hknqf4/where_to_put_my_typesinterfaces/
---
Hey,  
So, I am working on my first TS project and I want to make sure I start off correctly.   
What is the general concensus in terms of architecture? This is a website, not a library.  


1. Do you have your types/interfaces in a seperate file?
   1. If so, do you name it with a '\*.d.ts" or is that ONLY for libraries/modules to be used outside the codebase?
2. Do you put your types IN the same file that is being used? If you have to re-use that type to you just export it? (won't that get a bit disorganized after a while?)
3. Do you have a "types" folder that you define by feature?

&amp;#8203;

    /src
      /src/products
         --- product.tsx
         --- product.types.ts // these can be used in this feature or wherever?

Basically, what I am asking is HOW do you structure your typings that is robust &amp; expandable?
## [4][Why isn't filter kicking out undefined from this array?](https://www.reddit.com/r/typescript/comments/hkndzx/why_isnt_filter_kicking_out_undefined_from_this/)
- url: https://www.reddit.com/r/typescript/comments/hkndzx/why_isnt_filter_kicking_out_undefined_from_this/
---
       protected validateAndSetSentenceCandidates(candidates: string[]) {
          const passedAndUndefined: Array&lt;Sentence | undefined&gt; = candidates
             .map(candidate =&gt; {
                if (candidate.length &gt;= 2) {
                   return new Sentence(candidate);
                }
                // else undefined returned implicitly
             });
    
          const passedTest: Sentence[] = passedAndUndefined
             .filter((item: Sentence | undefined) =&gt; item !== undefined);
          // .filter((item: Sentence | undefined) =&gt; typeof item !== "undefined");
    
    */
    const passedTest: Sentence[]
    Type '(Sentence | undefined)[]' is not assignable to type 'Sentence[]'.
      Type 'Sentence | undefined' is not assignable to type 'Sentence'.
        Type 'undefined' is not assignable to type 'Sentence'.ts(2322)
    /*

Both attempts to filter them out above failed
## [5][How to tell TS a certain path is unreachable?](https://www.reddit.com/r/typescript/comments/hkpin3/how_to_tell_ts_a_certain_path_is_unreachable/)
- url: https://www.reddit.com/r/typescript/comments/hkpin3/how_to_tell_ts_a_certain_path_is_unreachable/
---
Without the final (hacky) return statement on the empty string, Typescript demands that the return type include `void`. But the `while` loop should never exit without returning a string.

Is there any way to tell TS this? Or is this hack necessary?

       protected getValidInput(configStep: WizardSteps): string {
          while (true) {
             const rawInput: string = configStep.prompt();
    
             // check for an exit value
             this.exitCheck(rawInput);
    
             const inputIsValid: boolean = configStep.validateInput(rawInput);
    
             if (inputIsValid) {
                return rawInput;
             } else {
                console.log(configStep.invalidInputMessage)
             }
          }
    
          // unreachable code, needed for TS compiler
          return "";
       }
## [6][Valuable compiler options](https://www.reddit.com/r/typescript/comments/hkhdur/valuable_compiler_options/)
- url: https://www.reddit.com/r/typescript/comments/hkhdur/valuable_compiler_options/
---
After forgetting `strictNullChecks` I've forgotten how helpful it is to constantly have to account for `null | undefined` during my initial drafting phase.

Any others you guys would promote for most projects? I also have `noImplicitAny`, and of course the `esModuleInterop`
## [7][Cannot invoke an object which is possibly 'undefined'.ts(2722)](https://www.reddit.com/r/typescript/comments/hklt5q/cannot_invoke_an_object_which_is_possibly/)
- url: https://www.reddit.com/r/typescript/comments/hklt5q/cannot_invoke_an_object_which_is_possibly/
---
Only the Not Null assertion operator (!) on validateFile worked below, not any of the other attempts at constraining null/undefined:

          const configData = this.steps.reduce(
             (accumulator: Partial&lt;ConfigData&gt;, currentStepInstance: NonNullable&lt;WizardSteps&gt;): Partial&lt;ConfigData&gt; =&gt; {
    
             if (currentStepInstance.needsFileValidation &amp;&amp; currentStepInstance.hasOwnProperty("validateFile")) { // also tried "in" keyword
    
    // (parameter) currentStepInstance: WizardSteps
    // Cannot invoke an object which is possibly 'undefined'.ts(2722)
                currentStepInstance.validateFile!();
             }

In general I try to avoid the ! operator as it isn't always honest. Is it necessary here or is there a better way?

The interface:

    export default interface WizardSteps {
       readonly hasSaveableData: boolean
       , readonly needsFileValidation: boolean
       , readonly invalidInputMessage?: string
       , readonly configDataKey?: string
    
       , explain(): void
       , prompt(): string
       , validateInput(userInput?: string): boolean
       , validateFile?(filePath?: string): boolean
    }

&amp;#x200B;
## [8][TS2464: A computed property name must be of type 'string', 'number', 'symbol', or 'any'. 54 [x.y] : z](https://www.reddit.com/r/typescript/comments/hkiodn/ts2464_a_computed_property_name_must_be_of_type/)
- url: https://www.reddit.com/r/typescript/comments/hkiodn/ts2464_a_computed_property_name_must_be_of_type/
---
    // TS2464: A computed property name must be of type 'string', 'number', 'symbol', or 'any'.
    
    54  [currentStepInstance.configDataKey] : validatedUserInput

I'm trying to use -- I think it's called "scope resolution" -- here on the key name, since property access requires dot notation. The key name is a string. 

The below syntax didn't work, how should it be written?

             if (currentStepInstance.hasSaveableData) {
                Object.assign(returnObject, {
                   [currentStepInstance.configDataKey: string] : validatedUserInput
                });
             }
## [9][How do I de-structure express request.body without "Unsafe assignment of an any value." error?](https://www.reddit.com/r/typescript/comments/hk1mp8/how_do_i_destructure_express_requestbody_without/)
- url: https://www.reddit.com/r/typescript/comments/hk1mp8/how_do_i_destructure_express_requestbody_without/
---
I am very new to typescript and when vscode threw that error after doing some search I mixed few stuffs from the results and came up with this solution.

`interface Body {`  
   `name: string;`  
   `email: string;`  
   `password: string;`  
 `}`  
 `const body: Body = request.body as Body;`  
 `const { name, email, password } = body;`

How can I improve this as it seems a lot of code for de-structuring.
## [10][How to deal with array of different class instances in typescript](https://www.reddit.com/r/typescript/comments/hk2z2a/how_to_deal_with_array_of_different_class/)
- url: https://www.reddit.com/r/typescript/comments/hk2z2a/how_to_deal_with_array_of_different_class/
---
So my code looks like this:

An abstract class named Component with nothing inside which all component inherit from.

An abstract generic base Entity class functioning as an Array abstraction with methods to add/get instances of classes,

A Entity class which extends BaseEntity with all the getters, one for each component, the reason I created it is because I must always force the type like &lt;Health&gt; this.getComponent(Health), I want to get rid of it, is there a way to type check BaseEntity so that when I do  this.getComponent(Health), the compiler knows its Health instance?

&amp;#x200B;

    export class Component {}
    export class Health extends Component { ............}
    export class Position extends Component { ............. }
    
    .....................
    export abstract class BaseEntity&lt;R&gt; {
        private readonly entity: Array&lt;R&gt; = [];
    
        protected getComponent&lt;T extends { new(...args: any[]): R }&gt;(component: T): R {
            return this.entity.find(
                currentComponent =&gt; currentComponent instanceof component
            );     
        } 
    
        addComponent(component: R): R {
            const index = this.entity.indexOf(component)
    
            if (index === -1) {
                 this.entity.push(component)
    
                 return component
            }
        }
    }
    
    .....................
    export abstract class Entity extends BaseEntity&lt;Component | Entity&gt; {
        get getHealth(): Health {
            return &lt;Health&gt; this.getComponent(Health)
        }
    
        get getPosition(): Position {
            return &lt;Position&gt; this.getComponent(Position)
        }
    }
    export class Player extends Entity {}
    export class Monster extends Entity {}
## [11][Type 'string' is not assignable to type '"X" | "Y" | "Z"'](https://www.reddit.com/r/typescript/comments/hjvw3n/type_string_is_not_assignable_to_type_x_y_z/)
- url: https://www.reddit.com/r/typescript/comments/hjvw3n/type_string_is_not_assignable_to_type_x_y_z/
---
Recently I had to use the `ReturnType` generic type constructor to build a type for me, that was the return shape of an obsurely documented method in Google Text To Speech's API:

    protected async fetchAndWriteAudio(options: {
          input : { text : string }
          , voice : { languageCode : string, ssmlGender: voiceGender }
          , audioConfig : { 
             audioEncoding : "AUDIO_ENCODING_UNSPECIFIED" | "LINEAR16" | "MP3" | "OGG_OPUS" }
       }, fileNameAndPath: string)
          : Promise&lt;ReturnType&lt;typeof TextToSpeechClient.prototype.synthesizeSpeech&gt;&gt; {

But now the title error shows on the error-commented lines below. I thought a specific string value, depicted as a quad union from the `ReturnType` constructor, should match one of those specific strings. Instead, none match. Why is that and what is the proper fix?

       public async makeSentenceAudio(subFolderPath: string): Promise&lt;void&gt; {
          // setup the foreign and english text string
          const foreignSentenceText = this.textTranslate(
             this.sentence.englishVersion
             , translationDirection.toForeign);
          const englishText = this.sentence.englishVersion;
    
          
          /**** Generate Audios ****/
          // define audio options (including text)
          const sharedOptions = {
             input: { text : "" }
             , voice: { languageCode : "", ssmlGender : voiceGender.male }
             , audioConfig : {audioEncoding : "OGG_OPUS"}
          }
    
          const foreignAudioOptions = {
             ...sharedOptions
             , input: {text: ""}
          }
    
          const englishAudioOptions = {
             ...sharedOptions
             , input: {text: ""}
          }
    
          this.fetchAndWriteAudio(foreignAudioOptions, "placeholder");
          this.fetchAndWriteAudio(englishAudioOptions, "placeholder");

.

    // the error message. There was also one on englishAudioOptions
    
    const foreignAudioOptions: {
        input: {
            text: string;
        };
        voice: {
            languageCode: string;
            ssmlGender: voiceGender;
        };
        audioConfig: {
            audioEncoding: string;
        };
    }
    Argument of type '{ input: { text: string; }; voice: { languageCode: string; ssmlGender: voiceGender; }; audioConfig: { audioEncoding: string; }; }' is not assignable to parameter of type '{ input: { text: string; }; voice: { languageCode: string; ssmlGender: voiceGender; }; audioConfig: { audioEncoding: "OGG_OPUS" | "AUDIO_ENCODING_UNSPECIFIED" | "LINEAR16" | "MP3"; }; }'.
    
      The types of 'audioConfig.audioEncoding' are incompatible between these types.
        Type 'string' is not assignable to type '"OGG_OPUS" | "AUDIO_ENCODING_UNSPECIFIED" | "LINEAR16" | "MP3"'. ts(2345)
