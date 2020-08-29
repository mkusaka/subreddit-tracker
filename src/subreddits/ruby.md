# ruby
## [1][I've always struggled with rails schema migrations, so I made my own TUI for it!](https://www.reddit.com/r/ruby/comments/iiipbx/ive_always_struggled_with_rails_schema_migrations/)
- url: https://github.com/jesseduffield/lazy_migrate
---

## [2][How environment check works in Ruby on Rails?](https://www.reddit.com/r/ruby/comments/iirt1i/how_environment_check_works_in_ruby_on_rails/)
- url: https://medium.com/rubycademy/how-environment-check-works-in-ruby-on-rails-4cfbd0434605
---

## [3][Linting Ruby](https://www.reddit.com/r/ruby/comments/iifp3k/linting_ruby/)
- url: https://kevindeisz.com/2020/08/28/linting-ruby.html
---

## [4][How to Build Your Own Gem in Ruby](https://www.reddit.com/r/ruby/comments/ii8g4i/how_to_build_your_own_gem_in_ruby/)
- url: https://blog.engineyard.com/how-to-build-your-own-gem-in-ruby
---

## [5][Newbie - Looking for private messaging and in-app notifications addons. Please help.](https://www.reddit.com/r/ruby/comments/iidz2z/newbie_looking_for_private_messaging_and_inapp/)
- url: https://www.reddit.com/r/ruby/comments/iidz2z/newbie_looking_for_private_messaging_and_inapp/
---
I have a foreign team building me an online community using ruby and react.  I need to add two features that I hope are available as an addon or plugin.

1. Private messaging system - This will be like sending emails or private notes to each other.  Would love in "in box" with messages received, read and sent.
2. In-app notifications for when users reply to your posts in our custom forum.  So, if someone replies to your post, you'd get an in-app notification badge like the facebook badges.

Does anyone know of any addons or plugins that could help satisfy these requirements?
## [6][Super Bombinhas - gameplay videos](https://www.reddit.com/r/ruby/comments/ii5m7h/super_bombinhas_gameplay_videos/)
- url: https://www.reddit.com/r/ruby/comments/ii5m7h/super_bombinhas_gameplay_videos/
---
Hi!

I posted some gameplay videos of my [open source platformer game](https://victords.itch.io/super-bombinhas) on YouTube, please check it out: [https://www.youtube.com/channel/UCDF29vao7Xur\_mTS4NFRviA/videos](https://www.youtube.com/channel/UCDF29vao7Xur_mTS4NFRviA/videos)

The videos started with "NEW" are from the current version of the project. There's also a video from a past version from many years ago.

Thanks!
## [7][Dependency Graphs in Ruby - Rewind Backups](https://www.reddit.com/r/ruby/comments/ii83l5/dependency_graphs_in_ruby_rewind_backups/)
- url: https://rewind.io/blog/dependency-graphs-in-ruby/
---

## [8][[help a novice] Refacotring my netbox wrapper](https://www.reddit.com/r/ruby/comments/ii9m89/help_a_novice_refacotring_my_netbox_wrapper/)
- url: https://www.reddit.com/r/ruby/comments/ii9m89/help_a_novice_refacotring_my_netbox_wrapper/
---
So, I try to get my feet wet again with a script that helps me automating the creation of a dhcp config.

Now I am at a point where everything works as expected. I need to get my head around API pagination but as I am currently only handling \~100 hosts that will be an issue for future KervyN (poor guy).

I want to refactor the code in the following terms:

* readability
* ABC lint error (I've read about the basic concept and there was a medium article on "Use dictionaries for that, but I don't know how to do it in this case)
* Maybe working with class(es) (if it makes sense). I don't know if it makes sense here, or if the code would benefit from it.

So here is the code:

    #!/usr/bin/env ruby
    
    require 'CGI'
    require 'http'
    require 'resolv'
    
    AUTH_HEADER = { Authorization: 'Token xxxxxxxxxxxxxxxxxxxxxxxx' }.freeze
    API_BASEURL = 'https://netbox/api/dcim/'.freeze
    
    @hosts = {}
    
    def request_api(endpoint, parameters = {})
      query = ''
      # if parameters are set we itereate over it and escape the values
      unless parameters.empty?
        query = ('?' + parameters.collect { |k, v| "#{CGI.escape(k.to_s)}=#{CGI.escape(v)}" }.join('&amp;'))
      end
    
      HTTP.headers(AUTH_HEADER).get(API_BASEURL + endpoint + '/' + query).parse
    end
    
    def sites
      statuses = %w[active planned]
    
      sites = request_api('sites')['results']
      sites.each { |site| @hosts[site['slug'].to_sym] = {} if statuses.include?(site['status']['value']) }
    end
    
    def hosts_from_site(site)
      types = %w[storage cpu]
      types.each do |type|
        request_api('devices', { site: site, name__isw: type, status: 'active' })['results'].map do |host|
          @hosts[site.to_sym][host['id']] = { name: host['name'], type: type, macaddresses: [] }
        end
      end
    end
    
    def mac_addresses_for_host_ids
      @hosts.map do |site, hosts|
        hosts.keys.each do |id|
          request_api('interfaces', { device_id: id.to_s })['results'].each do |interface|
            @hosts[site.to_sym][id][:macaddresses] &lt;&lt; interface['mac_address']&amp;.downcase if interface['name'] =~ /\Ae[tn][ho0]\d/i &amp;&amp; !interface['mac_address'].nil?
          end
        end
      end
    end
    
    def parse_logical_location
      @hosts.map do |site, hosts|
        hosts.map { |hostid, host| @hosts[site][hostid][:logcal_location] = host[:name].split('.').last }
      end
    end
    
    def ip_address_for_hostname
      @hosts.map do |site, hosts|
        hosts.map { |hostid, host| @hosts[site][hostid][:address] = Resolv::getaddress(host[:name] + '.example.org') }
      end
    end
    
    def build_shit
      sites
      @hosts.keys.each { |k| hosts_from_site(k.to_s) }
      mac_addresses_for_host_ids
      parse_logical_location
      ip_address_for_hostname
    
      @hosts.map do |_site, hosts|
        hosts.map { |_hostid, host| puts "dhcp-host=#{host[:macaddresses].join(',')},#{host[:name].split('.').first},#{host[:address]},set:netboot,set:known"  }
      end
    end

the "@hosts" hash looks like this in the end

    &gt;&gt; pp @hosts
    {:location1=&gt;
      {19=&gt;
        {:name=&gt;"storage1.fra1",
         :type=&gt;"storage",
         :macaddresses=&gt;["CC:CC:CC:CC:CC:CC", "DD:DD:DD:DD:DD:DD"],
         :logcal_location=&gt;"fra1",
         :address=&gt;"10.2.3.4"},
       21=&gt;
        {:name=&gt;"storage2.fra1",
         :type=&gt;"storage",
         :macaddresses=&gt;["BB:BB:BB:BB:BB:BB", "AA:AA:AA:AA:AA:AA"],
         :logcal_location=&gt;"fra1",
         :address=&gt;"10.2.3.5"}
      }
    }

Thanks in advance and mybe you beings have some sweet tips.  


Codesnipet for readability: [https://gitlab.com/-/snippets/2009689](https://gitlab.com/-/snippets/2009689)
## [9][The super keyword in Ruby](https://www.reddit.com/r/ruby/comments/ihs9k0/the_super_keyword_in_ruby/)
- url: https://medium.com/rubycademy/the-super-keyword-a75b67f46f05
---

## [10][Multipart file upload to Active Storage using GraphQL-ruby and Apollo](https://www.reddit.com/r/ruby/comments/ihqqrg/multipart_file_upload_to_active_storage_using/)
- url: https://www.abhaynikam.me/posts/active-storage-multipart-file-upload-graphql-ruby/
---

