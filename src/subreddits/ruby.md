# ruby
## [1][Comparison of approaches to implementing multitenancy in Rails apps](https://www.reddit.com/r/ruby/comments/gsqfux/comparison_of_approaches_to_implementing/)
- url: https://blog.arkency.com/comparison-of-approaches-to-multitenancy-in-rails-apps/
---

## [2][Take the 2020 Ruby on Rails Survey](https://www.reddit.com/r/ruby/comments/gscsd8/take_the_2020_ruby_on_rails_survey/)
- url: https://rails-hosting.com/
---

## [3][This blog is a web page (using Jekyll in 2020)](https://www.reddit.com/r/ruby/comments/gsfa66/this_blog_is_a_web_page_using_jekyll_in_2020/)
- url: https://www.reddit.com/r/ruby/comments/gsfa66/this_blog_is_a_web_page_using_jekyll_in_2020/
---
[This is not strictly Ruby-related, but y'all are my people so I thought you might appreciate this.]

I just threw some notes together about how I'm thinking about my personal site in the JAM stack era. A lot of people seem to be trying out Gatsby or Next.js on their personal blogs, so I thought it would be fun to use techniques from Rails land to improve the performance of my "boring" Jekyll blog so that it feels like a single-page app with minimal JavaScript. Spoiler alert: I'm using Turbolinks.

[https://www.joshuawood.net/this-blog-is-a-web-page](https://www.joshuawood.net/this-blog-is-a-web-page)

Next up I'm planning to add purgecss. I'm also looking for a way to bundle a subset of Font Awesome icons via CSS (maybe with an SVG sprite, I guess).

I'm still exploring new optimizations and progressive enhancement techniques--what do you think of this approach?
## [4][Questions about resources for learning Ruby](https://www.reddit.com/r/ruby/comments/gsfsox/questions_about_resources_for_learning_ruby/)
- url: https://www.reddit.com/r/ruby/comments/gsfsox/questions_about_resources_for_learning_ruby/
---
Hey everyone! Sorry if this kind of post is unwelcome, but I would love your thoughts on the best way to learn and practice Ruby (on Rails) rapidly. I am starting an internship on Monday, and I need to get as good with Ruby as possible before then. I misread an email from my manager last month and I spent the past month learning Java, whereas I am actually going to be using Ruby.

I am very familiar with programming concepts from a number of languages (C, Java, Python, some JavaScript), so I would like to avoid tutorials/documentation like Codecademy that start at the very basics. However, I don't have enough time to do a deep dive like I did with the Oracle Java documentation. Ideally, I can get a grip on the syntax and semantics, and also get some experience actually building in it.

Thank you in advance! I'm an idiot and would be very grateful of any suggestions lol
## [5][Stack Overflow Developer Survey 2020](https://www.reddit.com/r/ruby/comments/gs2zjc/stack_overflow_developer_survey_2020/)
- url: https://insights.stackoverflow.com/survey/2020#technology-how-technologies-are-connected
---

## [6][How can I capture console input when debugging?](https://www.reddit.com/r/ruby/comments/gsfd0w/how_can_i_capture_console_input_when_debugging/)
- url: https://www.reddit.com/r/ruby/comments/gsfd0w/how_can_i_capture_console_input_when_debugging/
---
The default debug terminal in VSCode apparently doesn't let me read from STDIN and I have no idea how to solve this problem. Is there any way I can debug Ruby scripts with breakpoints and all the fancy stuff and be able to read console input?
## [7][The Ruby Blend Episode 14: StimulusReflex, BridgetownRB, RailsBytes, AppLocale.dev, and more!](https://www.reddit.com/r/ruby/comments/gs65w5/the_ruby_blend_episode_14_stimulusreflex/)
- url: https://fireside.fm/s/ouBAUjGy+7QL69WzH
---

## [8][SVN2GIT, which runs on Ruby, fails when trying to migrate a project. Please help](https://www.reddit.com/r/ruby/comments/gs8t1o/svn2git_which_runs_on_ruby_fails_when_trying_to/)
- url: https://www.reddit.com/r/ruby/comments/gs8t1o/svn2git_which_runs_on_ruby_fails_when_trying_to/
---
I tried running svn2git and I keep getting this after I enter my password. Does anyone know how to solve this? Ruby seems to be the problem here.

"#&lt;Thread:0x00007fc6ed8171b0@/Library/Ruby/Gems/2.6.0/gems/svn2git-2.4.0/lib/svn2git/migration.rb:431 run&gt; terminated with exception (report\_on\_exception is true):

**Traceback** (most recent call last):

2: from /Library/Ruby/Gems/2.6.0/gems/svn2git-2.4.0/lib/svn2git/migration.rb:432:in \`block (2 levels) in run\_command'

1: from /Library/Ruby/Gems/2.6.0/gems/svn2git-2.4.0/lib/svn2git/migration.rb:432:in \`loop'

/Library/Ruby/Gems/2.6.0/gems/svn2git-2.4.0/lib/svn2git/migration.rb:438:in \`block (3 levels) in run\_command': **undefined local variable or method \`stdin' for #&lt;Svn2Git::Migration:0x00007fc6ed0dc788&gt; (NameError)**

**Did you mean?** **String"**

&amp;#x200B;

Thank you in advance!
## [9][Rails 6.1 adds support for signed ids to Active Record](https://www.reddit.com/r/ruby/comments/grmymc/rails_61_adds_support_for_signed_ids_to_active/)
- url: https://blog.saeloun.com/2020/05/20/rails-6-1-adds-support-for-signed-ids-to-active-record.html
---

## [10][[Help] Trying to scrape website with Ruby](https://www.reddit.com/r/ruby/comments/gs6iav/help_trying_to_scrape_website_with_ruby/)
- url: https://www.reddit.com/r/ruby/comments/gs6iav/help_trying_to_scrape_website_with_ruby/
---
Greetings,

I'm trying different scripting methods to automate this task and am failing. Can someone see what is wrong with this script and why it's not pulling in links? I'm attempting to scrape this website for all filenext.com links and generate a report.

    require 'open-uri'
    require 'base64'
    require 'net/http'
    require 'net/https'
    require 'uri'
    require 'pstore'
    require 'nokogiri'
    require 'terminal-table'
    
    ROOT_URI = "https://www.gfxtra31.com/"
    
    if !File.exist?("gfxtra31_progress.pstore")
      print("Looks like this is your first run. Writing file to track progress...")
      progress = PStore.new("gfxtra31_progress.pstore")
      progress.transaction do
        progress[:categories] = [
          {:complete =&gt; false, :name=&gt;"3d-models-files-addons/", :pagecount=&gt;588, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"adobe-after-effects/", :pagecount=&gt;477, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"artimages/", :pagecount=&gt;143, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"celebrities/", :pagecount=&gt;17, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"ebooks/", :pagecount=&gt;1448, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"fashion/", :pagecount=&gt;38, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"fonts/", :pagecount=&gt;343, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"footages/", :pagecount=&gt;428, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"only-for-you/", :pagecount=&gt;1242, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"adobe-indesign/", :pagecount=&gt;251, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"inspirational/", :pagecount=&gt;1, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"magazines/", :pagecount=&gt;2271, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"adobe-photoshop/", :pagecount=&gt;3963, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"plug-ins-add-ons/", :pagecount=&gt;100, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"portfolios/", :pagecount=&gt;38, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"presentation-templates/", :pagecount=&gt;174, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"scraps/", :pagecount=&gt;423, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"scripts/", :pagecount=&gt;578, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"software/", :pagecount=&gt;3496, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"sounds/", :pagecount=&gt;450, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"stock-image-cd/", :pagecount=&gt;2889, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"themes/", :pagecount=&gt;1437, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"tutorials/", :pagecount=&gt;1884, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"vectors/", :pagecount=&gt;4384, :next_page=&gt;1},
          {:complete =&gt; false, :name=&gt;"wallpapers/", :pagecount=&gt;330, :next_page=&gt;1}
        ]
        progress[:posts] = []
        progress[:links] = []
      end
      print(" OK!\n")
    end
    
    PROGRESS = PStore.new("gfxtra31_progress.pstore")
    
    def get_categories
      categories = []
      DATA.readlines.map(&amp;:chomp).each do |category|
        page = Nokogiri::HTML(open(ROOT_URI + category).read)
        page_links = page.css('.navigation a')
        pagecount = page_links.length &gt; 1 ? page_links[-2].text.to_i : 1
        categories &lt;&lt; { name: category, pagecount: pagecount, next_page: 1 }
      end
      categories
    end
    
    def get_posts(page_url)
      Nokogiri::HTML(open(page_url).read).css('.baslik a').map { |e| e['href'] }
    end
    
    def get_url(post_url)
      doc = Nokogiri::HTML(open(post_url).read)
      obscured = doc.at('a[rel="nofollow"]')['href'].rpartition('?url=')[-1]
      return obscured if obscured.include?("/")
      decoded = Base64.decode64(obscured)
      if decoded.include?("filenext")
        return decoded.rpartition('.html')[0..1].join('')
      end
      url = URI.parse(URI.escape(decoded))
      res = Net::HTTP.start(url.host, url.port) { |http| http.get(url.request_uri) }
      sleep(rand(3..15))
      res['location']
    end
    
    def report_progress
      rows = []
      PROGRESS.transaction(true) do
        PROGRESS[:categories].each do |cat|
          rows &lt;&lt; [ cat[:name].tr("-/", " "), "#{cat[:next_page] - 1} / #{cat[:pagecount]}" ]
        end
      end
      table = Terminal::Table.new headings: ["Category", "Pages Crawled"], rows: rows
      puts table
    end
    
    def fetch_next_post_set
      PROGRESS.transaction do
        cat = PROGRESS[:categories].find { |c| !c[:complete] }
        return false if cat == nil
        link = ROOT_URI + cat[:name] + "page/" + cat[:next_page].to_s + "/"
        new_batch = get_posts(link)
        PROGRESS[:posts].concat(new_batch)
        cat[:next_page] = cat[:next_page] + 1
        if cat[:next_page] &gt; cat[:pagecount]
          cat[:complete] = true
        end
        puts("Fetched #{new_batch.length} posts from #{cat[:name]}.")
      end
      true
    end
    
    def translate_all_posts_to_links(output_file)
      PROGRESS.transaction do
        while PROGRESS[:posts].length &gt; 0
          link = get_url(PROGRESS[:posts].pop)
          PROGRESS[:links] &lt;&lt; link
          puts("\t#{link}")
          output_file.puts(link)
          output_file.flush
          exit if Object.const_defined?(:Ocra)
          sleep(rand(3..15))
        end
      end
    end
    
    report_progress
    
    ongoing = true
    File.open("filenext_links.txt", "a") do |f|
      while ongoing
        ongoing = fetch_next_post_set
        sleep(rand(3..15))
        translate_all_posts_to_links(f)
      end
    end
