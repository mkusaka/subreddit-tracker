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
## [2][I’m creating an open source team management app - Troup. I’d love your opinion on it!](https://www.reddit.com/r/typescript/comments/hgjlq0/im_creating_an_open_source_team_management_app/)
- url: https://www.reddit.com/r/typescript/comments/hgjlq0/im_creating_an_open_source_team_management_app/
---
Hey guys,

I had posted my project here before and got some quality feedback and some people got in touch with me. It was absolutely wonderful. 

## Background

I've worked in quite a few companies (startups and giants), and have found a lingering problem. Everywhere, our team and teams in general spent a lot of time setting up and configuring a process for every app. That's a LOT of apps in most places.

Ex: 
- Trello for issue tracking on a per-team basis (frontend/backend/etc..) 
- GitHub issues for common issue tracking across various team
- Google Docs for any kind of documentation etc.

I soon got very frustrated having to keep switching between so many apps, and copy-pasting links on Slack/email. This happened 2 years ago and is what inspired me to start this project. I even created a frontend-only view of it and showed it to a lot of people I had worked with and to some people across various roles in the startup ecosystem. 

## Premise

The premise of the project is ensure that teams do not need to spend humongous amounts of time setting things up in an intricate and convoluted app. The purpose of Troup is simple:
- Provide an already-established process for your team without setup.
- Make information available at a central place to reduce hoops when trying to get a status update.
- Make work extremely transparent and visible across the board for all the teams.

The way I hope to achieve this is by creating a deeply-linked and content-first approach of a fusion between GitHub issues and GitHub projects. That actually served as an inspiration to this all. 

Users will create projects and boards. Projects can be about anything. Think of it as a repository. Boards are for the teams. Think of it as a Trello board for each team.

Within each project, users can create tickets. Tickets are essentially issues. On each ticket, you can assign it to multiple boards. Each board you assign a ticket to, a task will be created in the board, with a link to the ticket. On the ticket, you will see a list of all the boards it is associated with, and the status on each board.

A task within a board can be only in one of 5 columns (states). 
- On Hold (-1)
- Todo (0)
- In Progress (1)
- In Review (2)
- Done (3)

This is the part where Troup enforces a strict flow. And I fixed on this after talking to a lot of people across various competencies and domains.

The goal is to ensure any team (not just developers) can just create a board and start working on achieving their goals.

Troup will also offer first-class integrations with Slack, GitHub, Gitlab, Bitbucket, Invision, Figma, Google Docs, etc.

The plan is to have an entire ecosystem around it.

Going forward, there will also be other features like:
- Milestones
- OKRs
- Sprints
To name a few.

Apps that compete in this space are: Asana, Monday, Avaza. But Troup is different in the sense, that it has the power of content and focuses on it much much more than the other apps. The other apps provide you a LOT of freedom, which I've personally seen cause hindrances, miscommunications and tensions in teams where time is critical.

--

I’ve been at it on and off for a while, but over the past months I’ve been focusing heavily on it. 

This is my first foray into Typescript at such a scale. I’m absolutely enjoying it and learning a lot everyday. 

## Frontend

The frontend is made on NextJS with Apollo Client and Styled Components. There’s not much there but I’ve done quite a bit of work to get contextual data and route management. 

https://github.com/troup-io/troup-client

## Backend

Built on top of Nexus and Prisma with Postgres, this is where majority of the work has been done. With simple custom RBAC and data fetching, modeling with a view of splitting code into services. 

https://github.com/troup-io/troup-server

I’d absolutely love your feedback and any tips/tricks/improvements as I want to complete this project but also have the code be really beautiful. 

Thanks!

**Edit:** Updated the repo link for the client.

**Edit:** Added background and premise to give a bit more context.
## [3][Passing a function as a parameter to another function](https://www.reddit.com/r/typescript/comments/hgg8lw/passing_a_function_as_a_parameter_to_another/)
- url: https://www.reddit.com/r/typescript/comments/hgg8lw/passing_a_function_as_a_parameter_to_another/
---
I'm brand new to TypeScript/Angular so forgive my lack of clarity in my explanation but I'm having trouble figuring out if this scenario is possible.

I have a button, when pressed calls a function that generates a pop-up for confirmation to perform a web service call. If you press "Ok", it'll call the web service, deleting a record from a database, all this done using an Observable. I also have an Update button that needs the same confirmation pop-up so I thought I'd reuse the same confirmActionDialog() function I have for deleting records. My hope was, instead of passing a string or number to determine what course of action to take with an if statement  like this: 

    if (param === 'delete') { 
        //doDelete 
    } else if (param === 'update') { 
        //doUpdate 
    }

I could pass, from the HTML, the web service functions, like this:

    &lt;button name="save-button" 
            (click)="confirmActionDialog(() =&gt; saveObject())"&gt;
                Save
     &lt;/button&gt; 

But nothing I try seems to work. Is it possible? What would the parameter to my 
    confirmActionDialog()

function need to be in order to do something like this? Again, I'm brand new to this world so if I need to provide more detail, please let me know. Thanks in advance!
## [4][Understanding TypeScript enum | Const and Ambient Enum](https://www.reddit.com/r/typescript/comments/hgogw4/understanding_typescript_enum_const_and_ambient/)
- url: https://appdividend.com/2020/06/27/understanding-typescript-enum/
---

## [5][[sci-fi]My Story](https://www.reddit.com/r/typescript/comments/hgpyp0/scifimy_story/)
- url: https://www.reddit.com/r/typescript/comments/hgpyp0/scifimy_story/
---
My story really began 2 years ago while working as a biochem intern at my local uni. I was updating the department webpage and had some styling notes written on a Post-It. A coworker startled me, causing me to knock over my note and a Petri dish containing *acidyminococcus* bacteria. I picked everything up and went back to work. We had known that this bacteria possessed an enzyme called cpf1 that could be used to edit genes. What we definitely did not know is that if the bacteria consumed code written for the web, it would affect that editing.  

Almost immediately after my mishap, the color in my face began to change into the exact same shade of aqua blue as the header on our department webpage. Terrified but incredibly curious, I grabbed another post-it and scribbled some CSS to change the style back to inherit.  I set it in another dish with the bacteria, waited a few seconds, then wiped the post-it on my face. My face color returned to normal!!!

Fast forward a few months, and the world had completely changed. I was the head of a new company, HumanCSS, and designer styling was now available for people to patch their own genetics.
Want your features to be more bold? We could do that. Want your muscles to grow when you flex? We had you covered. Want every other hair to have a glowing effect? Yep, we could do that too. My Post-it/bacteria accident had given birth to the greatest genetic discovery of our times and I did what any intelligent human would--I found a way to monetize it.

Before you judge me, know that my company had also discovered some bugs in the entire process. Not allowing the bacteria to consume the entire code could cause half of your body to be transparent or your legs to _literally have no width_. We worked hard to test all of our "Sheets" as we called them. The safe ones were marketed and the unsafe ones were set aside to be fixed.

R&amp;D had just discovered we could give basic bits of javascript code to the bacteria. The results were unexpected. We stored them with our unsafe sheets.

Someone broke into our lab. They stole the unsafe ones, not knowing what they were doing. Now there are people walking around in circles or randomly just...stopping. That would be bad enough, except that the genetic code...is communicable. Apparently RNA is the greatest flash drive ever, and copies of bad code are transferring now from person to person, replicating itself every broken line by broken line. People are dying. People are becoming _undefined_ and it is all my fault. Maybe if we'd used TypeScript instead...
## [6][I need a repo/github example that shows a project with BOTH js/tsx files?](https://www.reddit.com/r/typescript/comments/hgppi7/i_need_a_repogithub_example_that_shows_a_project/)
- url: https://www.reddit.com/r/typescript/comments/hgppi7/i_need_a_repogithub_example_that_shows_a_project/
---
Basically, I have a legacy site in ".js" (react). I am moving it over to ".tsx", BUT I get sooo many errors when I change any file names to ".tsx". Its not only type errors, but I literally get errors on something simple like:  


     const App =&gt; {
       return &lt;div&gt;I am App home &lt;/div&gt;
     }

  My compiler shows that "I" "am" etc.. are variables and the &lt;div&gt; is generating errors. It's all a mess. So, I am thinking my babelrc or perhaps my webpack is jacked. I need to be able to leave the ".js" files alone and have them build as usual, and just add up date files with ".tsx" and slowly move to typescript. See anything wrong, or do you have an example of a repo with BOTH being used?  


 Here are my configs:  


     babelrc
      -----------
      {
      "presets": [
          "@babel/preset-react",
          "@babel/preset-env",
          "@babel/preset-typescript"
      ],
      "plugins": [
        "@babel/plugin-proposal-optional-chaining",
        "@babel/plugin-proposal-nullish-coalescing-operator",
        "@babel/plugin-proposal-class-properties",
        "@babel/plugin-syntax-dynamic-import",
        "@babel/proposal-object-rest-spread",
        ["module-resolver", {
          "root": ["./src"],
          "alias": {
            "src": "./src",
            "utils": "./src/utils/",
            "components": "./src/components",
            "features": "./src/features",
          }
        }]
      ]
    }
    

 webpack:

     const path = require('path');
     const webpack = require('webpack');
     const {CleanWebpackPlugin} = require('clean-webpack-plugin');
     const HtmlWebpackPlugin = require('html-webpack-plugin');
     const MiniCssExtractPlugin = require('mini-css-extract-plugin');
    
     const isDevelopment = process.env.NODE_ENV === 'development';
    
     module.exports = {
       entry: ['./node_modules/regenerator-runtime/runtime.js', './src/index.js'],
       module: {
         rules: [
           {
            test: /\.js?$/,
            exclude: /node_modules/,
            use: ['babel-loader'],
          },
          {
            test: /\.module\.s(a|c)ss$/,
            loader: [
              isDevelopment ? 'style-loader' : MiniCssExtractPlugin.loader,
              {
                loader: 'css-loader',
                options: {
                  modules: true,
                  sourceMap: isDevelopment,
                },
              },
              {
                loader: 'sass-loader',
                options: {
                  sourceMap: isDevelopment,
                },
              },
            ],
          },
          {
            test: /\.s?(css)$/,
            exclude: [/node_modules/, /styles/],
            use: [
              isDevelopment ? 'style-loader' : MiniCssExtractPlugin.loader,
              {
                loader: 'css-loader',
                query: {
                  modules: {
                    localIdentName: '[path][name]__[local]--[hash:base64:5]',
                  },
                  importLoaders: true,
                  sourceMap: true,
                },
              },
              'css-modules-typescript-loader',
              'postcss-loader',
              'sass-loader',
            ],
          },
          {
            test: /\.s?(css)$/,
            include: [/styles/],
            use: [
              isDevelopment ? 'style-loader' : MiniCssExtractPlugin.loader,
              {
                loader: 'css-loader',
                query: {
                  modules: false,
                  importLoaders: true,
                  sourceMap: true,
                },
              },
              'postcss-loader',
              'sass-loader',
            ],
          },
          {
            test: /\.(eot|otf|ttf|woff|woff2)$/,
            use: 'file-loader?name=[name]_[hash:base64:5].[ext]',
          },
          {
            test: /\.(jpg|png|gif|svg)$/,
            use: 'file-loader',
          },
          {
            test: /\.(ts|tsx)?$/,
            include: path.resolve(__dirname, 'src'),
            use: ['ts-loader'],
          },
        ],
      },
      resolve: {
        modules: ['src', 'node_modules'],
        extensions: ['*', '.js', '.jsx', '.ts', '.tsx']
      },
      plugins: [
        new CleanWebpackPlugin(),
        new HtmlWebpackPlugin({
          title: 'My Site',
          template: './src/index.html',
        }),
        new MiniCssExtractPlugin({
          filename: isDevelopment ? '[name].css' : '[name].[hash].css',
          chunkFilename: isDevelopment ? '[id].css' : '[id].[hash].css',
        }),
      ],
      output: {
        path: path.join(__dirname, '../dist'),
        chunkFilename: '[name].[hash].chunk.js',
        filename: '[name].[hash].js',
      },
    };
    

 tsconfig

      {
      "compilerOptions": {
        // because we're compiling with Babel instead of TypeScript
        // and we're only using TypeScript for type checking, we'll set "noEmit"
        // to true so running type checking doesn't generate any files.
        "strict": false,
        "noEmit": true,
        "forceConsistentCasingInFileNames": true,
        "baseUrl": "./src",
        "jsx": "preserve",
        "allowJs": true,
        "lib": ["dom", "es2018"],
        "paths": {
           "test-utils": ["./tools/tests/test-utils"]
         }
      },
      "include": ["./src/**/*.ts", "./src/**/*.tsx", "*.ts", "*.tsx", "./@types"],
      "exclude": [
        "./src/**/__tests__/**/*.ts",
        "./src/**/__tests__/**/*.js",
        "node_modules",
        "build",
        "dist",
        "./src/**/*.js",
        "./src/**/*.jsx"
      ]
    }
## [7][Mapping An Interface (Array)](https://www.reddit.com/r/typescript/comments/hgk3lc/mapping_an_interface_array/)
- url: https://www.reddit.com/r/typescript/comments/hgk3lc/mapping_an_interface_array/
---
Hi. I'm pretty new to `TypeScript` but have experience in other typed languages. I'm trying to do this:

```
interface Foo {
  foo: string;
}

interface Foos {
  [index: string]: Foo[];
}

let response: Foos;

response = requestFoo.payload.map((p) =&gt; ({
   foo: p.foo_field
} as Foos))
```

But I'm getting a type error that `Type Foo[] is not assignable to type Foos`. What am I doing wrong?
## [8][TypeScript React with Material-ui's makeStyles/withStyles/createStyles](https://www.reddit.com/r/typescript/comments/hggp5k/typescript_react_with_materialuis/)
- url: https://www.reddit.com/r/typescript/comments/hggp5k/typescript_react_with_materialuis/
---
I'm coming to react from the angular world, and I'm having a bit of trouble making sense of how react uses material-ui's makeStyles/withStyles/createStyles functionality. 

I'm working on a component, that calls another component, that calls another component. 

The very bottom component, makes use of the `withStyles` factory and `createStyles`:

    
    interface Props {
      value: ReactNode;
    
      color?: string;
      size?: number | string;
      align?: 'left' | 'center' | 'right';
    }
    
    const styles = createStyles({
      root: {
        lineHeight: '1em',
        fontSize: (props: Props) =&gt; props.size ?? '1em',
        textAlign: (props: Props) =&gt; props.align ?? 'center',
      },
      value: {
        color: (props: Props) =&gt; props.color ?? 'inherit',
      },
      caption: {
        fontSize: '0.675em',
        verticalAlign: 'middle',
        display: 'flex',
      },
    });
    
    export type CaptionedValueProps = Props &amp; WithStyles&lt;typeof styles&gt;;
    
    const CaptionedValueComponent: React.FC&lt;CaptionedValueProps&gt; = ({ value, classes, children }) =&gt; (
      &lt;div className={classes.root}&gt;
        &lt;strong className={classes.value}&gt;{value}&lt;/strong&gt;
        &lt;div className={classes.caption}&gt;{children}&lt;/div&gt;
      &lt;/div&gt;
    );
    
    export const CaptionedValue = memo(withStyles(styles)(CaptionedValueComponent));
    

and because of the the `withStyles` factory call, at this very bottom level component, i was told to make the component at the very top make use of the `classes` prop, passing it into the middle component, and have the middle component emulate this^ component, by invoking the `withStyles`  factory method.

My stab at the middle component is this:

    interface Props {
      current: number;
      previous?: number;
      formatter?: ValueFormatter&lt;number&gt;;
      size?: string | number;
      align?: 'left' | 'center' | 'right';
      colorCurrent?: boolean;
      invertColor?: boolean;
    }
    
    interface StyleProps {
      colorCurrent?: boolean;
      invertColor?: boolean;
      previous?: number;
      delta?: number;
    }
    
    const styles = createStyles({
      value: {
        color: ({ colorCurrent, invertColor, previous, delta }: StyleProps) =&gt;
          colorCurrent ? valueDeltaColor(previous ?? 0, delta ?? 0, invertColor ?? false) : undefined,
      },
      delta: {
        color: ({ previous, delta, invertColor }: StyleProps) =&gt;
          valueDeltaColor(previous ?? 0, delta ?? 0, invertColor ?? false),
        display: 'inline-flex',
        flexWrap: 'nowrap',
        alignItems: 'center',
      },
    });
    
    const useStyles = makeStyles(() =&gt; styles);
    
    export type ValueComparisonProps = Props &amp; WithStyles&lt;typeof styles&gt;;
    
    const ValueComparisonComponent: React.FC&lt;ValueComparisonProps&gt; = ({
      current,
      previous,
      formatter = NumberFormatters.decimal,
      size,
      align = 'right' as CaptionedValueProps['align'],
      colorCurrent = false,
      invertColor = false,
      classes,
    }) =&gt; {
      const delta = !previous ? 0 : (current - previous) / previous;
      const hasChange = previous !== 0 &amp;&amp; delta !== 0;
      const stylesHook = useStyles({
        colorCurrent,
        invertColor,
        previous,
        delta,
      });
    
      if (!previous) {
        return (
          &lt;CaptionedValue value={formatter`${current}`} size={size} align={align}&gt;
            &amp;nbsp; {/* required for reason */}
          &lt;/CaptionedValue&gt;
        );
      }
      //CaptionedValue does not appreciate being given a delta, so nuke it
      if (classes) delete classes.delta;
    
      return (
        &lt;CaptionedValue
          classes={classes || { value: styles.value }}
          value={formatter`${current}`}
          size={size}
          align={align}
        &gt;
          {hasChange &amp;&amp; (
            &lt;span className={stylesHook.delta}&gt;
              {delta &gt; 0 ? &lt;ArrowUpward fontSize="inherit" /&gt; : &lt;ArrowDownward fontSize="inherit" /&gt;}
              {NumberFormatters.percent`${delta}`}
            &lt;/span&gt;
          )}
          &amp;nbsp;{formatter`(${previous})`}
        &lt;/CaptionedValue&gt;
      );
    };
    
    export const ValueComparison = memo(withStyles(styles)(ValueComparisonComponent));
    
    

  It all works, but I don't know that this is the kosher way to do it...especially because of `classes={classes || { value: styles.value }}` i had to do `if (classes) delete classes.delta` as the `CaptionedValue` component complained about having  `.delta` passed to it.
## [9][Introduction React Native, Typescript, Redux, Rxjs, Node.js, Mongo, Mongoose, Nexus Aurora Platform](https://www.reddit.com/r/typescript/comments/hg5g8z/introduction_react_native_typescript_redux_rxjs/)
- url: https://www.youtube.com/watch?v=06wsg2f76hQ
---

## [10][How to properly package a module for type definitions](https://www.reddit.com/r/typescript/comments/hg810t/how_to_properly_package_a_module_for_type/)
- url: https://www.reddit.com/r/typescript/comments/hg810t/how_to_properly_package_a_module_for_type/
---
I'm trying to create a module for type definitions to be shared across multiple typescript projects (mostly angular, but also typescript cloud functions). I was able to create a simple package with just a single index.d.ts file with all my types exported. However I've come across the issue of wanting a union type that I also need to access the members of in an array. What I did before I tried creating the shared module was the following:

`export const UnionTypeList= [ 'type1', 'type2'] as const;`

`export interface UnionType = typeof UnionTypeList[number]`

However if I try declare a const array in the index.d.ts file it won't let me. I've moved it to a separate ts file and imported that into the definition file, but now Angular is telling me that file is not being included in typescript compilation despite me adding it to the .tsconfig files array.

Is there any easier way I should be doing this? And please don't just tell me to use a monorepo, that's not the answer I'm looking for.
## [11][Requiring property on a type based on the presence of another property.](https://www.reddit.com/r/typescript/comments/hg52p4/requiring_property_on_a_type_based_on_the/)
- url: https://www.reddit.com/r/typescript/comments/hg52p4/requiring_property_on_a_type_based_on_the/
---
So, I am not really sure what to search to find out more about this concept, and therefore I apologise in advance if this is a duplicate, but I have been stuck for a day now and need some help.

I am using TS 3.6.4 in a React project, and I am using an interface to validate props passed to a component.

So, let's say I have an interface called IButton as follows:

    interface IButton {
      label: string;
      iconName?: string;
      iconColor?: string;
    }

Is there a way I can only require `iconColor` be supplied only if `iconName` is present, and if `iconName` is not present, then have both required to be undefined?

For some context, the reason I want to do this is because the button may not always have an icon, but when an icon is being used, the color for said icon is required alongside it.
