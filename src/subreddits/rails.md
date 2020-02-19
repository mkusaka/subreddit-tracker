# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/f2r9mk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/f2r9mk/personal_projects_show_off_your_own_project_andor/
---
In this thread you can showcase your personal pet project to other redditors.

Need help with a specific problem or just wanna have some extra eyeballs on your code? Ask away!

A suggested format to get you started:

1. **Name of your project**
2. **A short description**
3. **Application stack**
4. **Link to Live app**
5. **Link to GitHub**
6. **You experience level**
7. **Other information or areas that you would like advice on**

 

^(Many thanks to Kritnc for getting the ball rolling.)
## [3][Ruby 2.7 removes taint checking mechanism](https://www.reddit.com/r/rails/comments/f67vxa/ruby_27_removes_taint_checking_mechanism/)
- url: https://blog.saeloun.com/2020/02/18/ruby-2-7-access-and-setting-of-safe-warned-will-become-global-variable
---

## [4][Installing ActiveAdmin (and sassc) breaks Bootstrap asset compilation](https://www.reddit.com/r/rails/comments/f69bw7/installing_activeadmin_and_sassc_breaks_bootstrap/)
- url: https://www.reddit.com/r/rails/comments/f69bw7/installing_activeadmin_and_sassc_breaks_bootstrap/
---
Hello!

I've been struggling for a few days to get to the root of this problem but I finally have a clearer idea:  
\- I have a simple app that using Bootstrap, installed through webpacker but I still use the asset pipeline for CSS. (I need to compile Boostrap myself because i'm overriding a lot of the sass variables.)  
\- I installed the ActiveAdmin gem, setup my admin, no problem in dev but \`assets:precompile\` fails in production env with a vague sass compilation error:

    &gt;&gt; c/svg%3e") #FEFEFE no-repeat center right 3.25rem/add(0.675em, 1rem) add(0.6
    /Users/oiorain/.rvm/gems/ruby-2.6.5/gems/sassc-2.2.1/lib/sassc/engine.rb:49:in `render'

The error is somewhere in the bootstrap SASS which hasn't changed. Very weird...

It took me a while to pin this down but turns out ActiveAdmin adds the \`sassc-rails\` gem which is now used to compile the assets. The bootstrap-sass gem has been updated to be compatible with sassc but it seems like the yarn package is not?

I don't know where to go from here.   
Logically, It seems I should be moving towards having both CSS and JS handled by webpacker but it seems impossible to have activeadmin and the bootstrap yarn package.   
I'm kind of failing to understand why rails-sassc wouldnt be able to compile something that the regular rails sass compiler can. There are subtleties that escape me here. Would any be able to suggest a way forward? Does anyone has had this kind of struggle between gems and webpacker libs?
## [5][want to build an app that primarily communicates via sms](https://www.reddit.com/r/rails/comments/f61v06/want_to_build_an_app_that_primarily_communicates/)
- url: https://www.reddit.com/r/rails/comments/f61v06/want_to_build_an_app_that_primarily_communicates/
---
can someone point me in the right direction on creating a rails app that users would sign up with a phone number and then the app would engage with users solely via sms?
## [6][Rails: How to group by nested association?](https://www.reddit.com/r/rails/comments/f60w5t/rails_how_to_group_by_nested_association/)
- url: https://www.reddit.com/r/rails/comments/f60w5t/rails_how_to_group_by_nested_association/
---
I have a Follow model with `user_id` and `track_id`. The Track model has an `artist_id` field.

What I want to do is count which artists have the most followers, but since users follow "tracks" and not "artists", I need to figure out how to do a count through the tracks.

So, what I was thinking was to do some sort of group by on a nested association. i.e. Group the Follow records by "track -&gt; artist_id", somehow.

Then I could count the number of users for each.

Is that even possible? Is there more info that would be useful here?
## [7][Help Authoring Tool (HAT) Software Market 2020-2026 Report Analysis of Major Players:MadCap, Help+Manual, Adobe RoboHelp, HelpSmith, ClickHelp, Dr.Explain](https://www.reddit.com/r/rails/comments/f69g3l/help_authoring_tool_hat_software_market_20202026/)
- url: https://www.reddit.com/r/rails/comments/f69g3l/help_authoring_tool_hat_software_market_20202026/
---
ReportsandMarkets.com adds “**Global Help Authoring Tool (HAT) Software Market 2019 by Company, Regions, Type and Application, Forecast to 2026**” new reports to its research database. The report spread across 100 to 140 pages with tables and figures in it.

This report studies the Help Authoring Tool (HAT) Software Market with many aspects of the industry like the market size, market status, market trends and forecast, the report also provides brief information of the competitors and the specific growth opportunities with key market drivers. Find the complete Help Authoring Tool (HAT) Software Market analysis segmented by companies, region, type and applications in the report.

**Market Segment by Companies, this report covers: –** MadCap, Help+Manual, Adobe RoboHelp, HelpSmith, ClickHelp, Dr.Explain, HelpNDoc, HelpStudio, WebWorks ePublisher, Daux.io, Document360, GenHelp, HelpScribble, and WinCHM 

**COMPLETE REPORT Help Authoring Tool (HAT) Software @** [**https://www.reportsandmarkets.com/reports/global-help-authoring-tool-hat-software-market-report-2020?utm\_source=reddit&amp;utm\_medium=47**](https://www.reportsandmarkets.com/reports/global-help-authoring-tool-hat-software-market-report-2020?utm_source=reddit&amp;utm_medium=47) 

**Help Authoring Tool (HAT) Software Market** continues to evolve and expand in terms of the number of companies, products, and applications that illustrates the growth perspectives. The report also covers the list of Product range and Applications with SWOT analysis, CAGR value, further adding the essential business analytics. Help Authoring Tool (HAT) Software Market research analysis identifies the latest trends and primary factors responsible for market growth enabling the Organizations to flourish with much exposure to the markets.

**Market Segment by Regions, regional analysis covers**

Asia-Pacific

North America

Europe

South America

Middle East &amp; Africa

**Research objectives****:**   

To study and analyze the global Help Authoring Tool (HAT) Software market size by key regions/countries, product type and application, history data from 2015 to 2019, and forecast to 2026.

To understand the structure of Help Authoring Tool (HAT) Software market by identifying its various sub segments.

Focuses on the key global Help Authoring Tool (HAT) Software players, to define, describe and analyze the value, market share, market competition landscape, SWOT analysis and development plans in next few years.

To analyze the Help Authoring Tool (HAT) Software  with respect to individual growth trends, future prospects, and their contribution to the total market.

To share detailed information about the key factors influencing the growth of the market (growth potential, opportunities, drivers, industry-specific challenges and risks).

To project the size of Help Authoring Tool (HAT) Software submarkets, with respect to key regions (along with their respective key countries).

To analyze competitive developments such as expansions, agreements, new product launches and acquisitions in the market.

To strategically profile the key players and comprehensively analyze their growth strategies.

**REQUEST EXCLUSIVE FREE SAMPLE OF THIS REPORT @** [**https://www.reportsandmarkets.com/sample-request/global-help-authoring-tool-hat-software-market-report-2020?utm\_source=reddit&amp;utm\_medium=47**](https://www.reportsandmarkets.com/sample-request/global-help-authoring-tool-hat-software-market-report-2020?utm_source=reddit&amp;utm_medium=47) 

**Key Developments in the Help Authoring Tool (HAT) Software** **Market**

To describe Help Authoring Tool (HAT) Software Introduction, product type and application, market overview, market analysis by countries, market opportunities, market risk, market driving force;

To analyze the manufacturers of Help Authoring Tool (HAT) Software,with profile, main business, news, sales, price, revenue and market share in 2019 and 2020;

To display the competitive situation among the top manufacturers in Global, with sales, revenue and market share in 2019 and 2020;

To show the market by type and application, with sales, price, revenue, market share and growth rate by type and application, from 2015 to 2020;

To analyze the key countries by manufacturers, Type and Application, covering North America, Europe, Asia Pacific, Middle-East and South America, with sales, revenue and market share by manufacturers, types and applications;

Help Authoring Tool (HAT) Software market forecast, by countries, type and application, with sales, price, revenue and growth rate forecast, from 2020 to 2026;

To analyze the manufacturing cost, key raw materials and manufacturing process etc.

To analyze the industrial chain,sourcing strategy and downstream end users (buyers);

To describe Help Authoring Tool (HAT) Software sales channel, distributors, traders, dealers etc.

To describe Help Authoring Tool (HAT) Software Research Findings and Conclusion, Appendix, methodology and data source

The Help Authoring Tool (HAT) Software Market research report completely covers the vital statistics of the capacity, production, value, cost/profit, supply/demand import/export, further divided by company and country, and by application/type for best possible updated data representation in the figures, tables, pie chart, and graphs. These data representations provide predictive data regarding the future estimations for convincing market growth. The detailed and comprehensive knowledge about our publishers makes us out of the box in case of market analysis.

**FOR MORE ENQUIRY OF Help Authoring Tool (HAT) Software @** [**https://www.reportsandmarkets.com/enquiry/global-help-authoring-tool-hat-software-market-report-2020?utm\_source=reddit&amp;utm\_medium=47**](https://www.reportsandmarkets.com/enquiry/global-help-authoring-tool-hat-software-market-report-2020?utm_source=reddit&amp;utm_medium=47)

In this study, the years considered to estimate the market size of Help Authoring Tool (HAT) Software are as follows:

History Year: 2015-2019

Base Year: 2019

Estimated Year: 2020

Forecast Year 2020 to 2026

This report includes the estimation of market size for value (million USD) and volume (M Units). Both top-down and bottom-up approaches have been used to estimate and validate the market size of Help Authoring Tool (HAT) Software market, to estimate the size of various other dependent submarkets in the overall market. Key players in the market have been identified through secondary research, and their market shares have been determined through primary and secondary research. All percentage shares, splits, and breakdowns have been determined using secondary sources and verified primary sources.

**Table of Contents**

Chapter 1 Overview of Help Authoring Tool (HAT) Software 

Chapter 2 Global Market Status and Forecast by Regions

Chapter 3 Global Market Status and Forecast by Types

Chapter 4 Global Market Status and Forecast by Downstream Industry

Chapter 5 Market Driving Factor Analysis of Help Authoring Tool (HAT) Software

Chapter 6 Help Authoring Tool (HAT) Software Market Competition Status by Major Manufacturers

Chapter 7 Help Authoring Tool (HAT) Software Major Manufacturers Introduction and Market Data

Chapter 8 Upstream and Downstream Market Analysis of Help Authoring Tool (HAT) Software 

Chapter 9 Cost and Gross Margin Analysis of Help Authoring Tool (HAT) Software 

Chapter 10 Marketing Status Analysis of Help Authoring Tool (HAT) Software 

Chapter 11 Report Conclusion

Chapter 12 Research Methodology and Reference

**About Us****:**

Reports and Markets is not just another company in this domain but is a part of a veteran group called Algoro Research Consultants Pvt. Ltd. It offers premium progressive statistical surveying, market research reports, analysis &amp; forecast data for a wide range of sectors both for the government and private agencies all across the world. The database of the company is updated on a daily basis. Our database contains a variety of industry verticals that include: Food Beverage, Automotive, Chemicals and Energy, IT &amp; Telecom, Consumer, Healthcare, and many more. Each and every report goes through the appropriate research methodology, Checked from the professionals and analysts.

**Contact Us****:**

Sanjay Jain

Manager – Partner Relations &amp; International Marketing

[www.reportsandmarkets.com](http://www.reportsandmarkets.com/)

[info@reportsandmarkets.com](mailto:info@reportsandmarkets.com)

Ph: +1-352-353-0818 (US)
## [8][Confused on how to run a Rails 6 app using Webpack in production mode](https://www.reddit.com/r/rails/comments/f5tibt/confused_on_how_to_run_a_rails_6_app_using/)
- url: https://www.reddit.com/r/rails/comments/f5tibt/confused_on_how_to_run_a_rails_6_app_using/
---
I've built an app in Rails 6 with a frontend using React via webpack. During development, it was as simple as \`rails s\` and \`bin/webpack-dev-server\`. After getting this app to deploy to Heroku, I immediately started having problems with React components working properly and getting errors like: 

    react-dom.production.min.js:4638 TypeError: Cannot convert undefined or null to object
        at hasOwnProperty (&lt;anonymous&gt;)
        at Modal.js:21
        at Array.forEach (&lt;anonymous&gt;)
        at Modal.js:20
        at t.n.render (Modal.js:302)
        at Qi (react-dom.production.min.js:4243)
        at Ji (react-dom.production.min.js:4234)
        at wc (react-dom.production.min.js:6676)
        at yu (react-dom.production.min.js:5650)
        at Mu (react-dom.production.min.js:5639)
    pc @ react-dom.production.min.js:4638
    react-dom.production.min.js:2837 Uncaught TypeError: Cannot convert undefined or null to object
        at hasOwnProperty (&lt;anonymous&gt;)
        at Modal.js:21
        at Array.forEach (&lt;anonymous&gt;)
        at Modal.js:20
        at t.n.render (Modal.js:302)
        at Qi (react-dom.production.min.js:4243)
        at Ji (react-dom.production.min.js:4234)
        at wc (react-dom.production.min.js:6676)
        at yu (react-dom.production.min.js:5650)
        at Mu (react-dom.production.min.js:5639) 

Someone told me to make sure the app works in production mode locally before deploying to Heroku. I know how to run \`rails s -p production\` but how do I start \`bin/webpack-dev-server\` and I'm not quite sure how to configure my app for production in general.
## [9][Fortunately, finished Hartl's Rails tutorial. Feel weak on DB/SQL side, virtually no JS work except a small book, and there is Capstone Rails Tutorials(free). Any suggestion what/where to study next?](https://www.reddit.com/r/rails/comments/f65j6p/fortunately_finished_hartls_rails_tutorial_feel/)
- url: https://www.reddit.com/r/rails/comments/f65j6p/fortunately_finished_hartls_rails_tutorial_feel/
---
I do feel a bit weak on the DB/SQL side.

I was going to continue as a basic primer on to Learn Enough's JS tutorial, and perhaps move on from there to other resources (You Don't Know JS Yet, Odin Project JS course, etc)

There is also Capstone Rails Tutorials(free at https://tutorials.railsapps.org/), for getting better at Rails I would of course presume. 

Any suggestion one what course of study to go next based on the above, and/or where?
## [10][A way for a Rails application to interact with a CardDAV service?](https://www.reddit.com/r/rails/comments/f5ta1k/a_way_for_a_rails_application_to_interact_with_a/)
- url: https://www.reddit.com/r/rails/comments/f5ta1k/a_way_for_a_rails_application_to_interact_with_a/
---
I am updating this legacy application to Rails 6 and would like to add a way to export users' contact info (names, e-mail, phone) to administrators smartphones. 

This will enable administrators to see caller IDs and make their work with call center and messaging applications a lot easier as right now they have to go to a control panel and look up users by their phone numbers.

I've seen some hacks for this, eg admins would dock their phone and run an app that syncs over an API or some sort of contact import process and would rather not go this way, admins are not always very computer-savvy or simply don't have enough time, so I'd prefer have their smartphones import contacts the same way they do with Gmail accounts for instance.

Does anyone know a way to get this done? Any pointers will be much appreciated.

Thank you very much!
## [11][Question About Railties, rails repo and rails gems](https://www.reddit.com/r/rails/comments/f5i1am/question_about_railties_rails_repo_and_rails_gems/)
- url: https://www.reddit.com/r/rails/comments/f5i1am/question_about_railties_rails_repo_and_rails_gems/
---
I am having a hard time understanding the rails repo. Active record, railties and all others(action view, etc) are in a monorepo. rails/rails. 

They are all separate gems as well? Can someone explain this to me? Each action/active is a gem &amp; available on rubygems. 

Or is the rails repo /activerecord boilerplate to connect to the actual gem? 

Tldr: what's the difference between activerecord gem and the folder that sits in rails repo titled "activerecord"? 

Thanks!

Edit: it looks like the individual gems (activerecord) is not on rubygems. Now I am more confused. When I run "bundler list" it treats each one as a separate gem. Why is this?
## [12][How do you remove a user's personally identifiable information but retain user activity?](https://www.reddit.com/r/rails/comments/f5lafp/how_do_you_remove_a_users_personally_identifiable/)
- url: https://www.reddit.com/r/rails/comments/f5lafp/how_do_you_remove_a_users_personally_identifiable/
---
I'd like to remove Jane Doe's email and full name from the system but retain Jane D.'s activity -- all comments, tasks etc., similar to how Basecamp functions.

Has any of you solved this problem and how did you do it?

Thank you!
