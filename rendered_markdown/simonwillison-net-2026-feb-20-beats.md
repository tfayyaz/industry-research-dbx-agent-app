# [Simon Willison’s Weblog](/)

[Subscribe](/about/#subscribe)

**Sponsored by:** Teleport — Secure, Govern, and Operate AI at Engineering Scale. [Learn more](https://fandf.co/4u0HQAD)

## Adding TILs, releases, museums, tools and research to my blog

20th February 2026

I’ve been wanting to add indications of my various other online activities to my blog for a while now. I just turned on a new feature I’m calling “beats” (after story beats, naming this was hard!) which adds five new types of content to my site, all corresponding to activity elsewhere.

Here’s what beats look like:

![Screenshot of a fragment of a page showing three entries from 30th Dec 2025. First: [RELEASE] "datasette-turnstile 0.1a0 — Configurable CAPTCHAs for Datasette paths usin…" at 7:23 pm. Second: [TOOL] "Software Heritage Repository Retriever — Download archived Git repositories f…" at 11:41 pm. Third: [TIL] "Downloading archived Git repositories from archive.softwareheritage.org — …" at 11:43 pm.](https://static.simonwillison.net/static/2026/three-beats.jpg)

Those three are from [the 30th December 2025](https://simonwillison.net/2025/Dec/30/) archive page.

Beats are little inline links with badges that fit into different content timeline views around my site, including the homepage, search and archive pages.

There are currently five types of beats:

* [Releases](https://simonwillison.net/elsewhere/release/) are GitHub releases of my many different open source projects, imported from [this JSON file](https://github.com/simonw/simonw/blob/main/releases_cache.json) that was constructed [by GitHub Actions](https://simonwillison.net/2020/Jul/10/self-updating-profile-readme/).
* [TILs](https://simonwillison.net/elsewhere/til/) are the posts from my [TIL blog](https://til.simonwillison.net/), imported using [a SQL query over JSON and HTTP](https://github.com/simonw/simonwillisonblog/blob/f883b92be23892d082de39dbada571e406f5cfbf/blog/views.py#L1169) against the Datasette instance powering that site.
* [Museums](https://simonwillison.net/elsewhere/museum/) are new posts on my [niche-museums.com](https://www.niche-museums.com/) blog, imported from [this custom JSON feed](https://github.com/simonw/museums/blob/909bef71cc8d336bf4ac1f13574db67a6e1b3166/plugins/export.py).
* [Tools](https://simonwillison.net/elsewhere/tool/) are HTML and JavaScript tools I’ve vibe-coded on my [tools.simonwillison.net](https://tools.simonwillison.net/) site, as described in [Useful patterns for building HTML tools](https://simonwillison.net/2025/Dec/10/html-tools/).
* [Research](https://simonwillison.net/elsewhere/research/) is for AI-generated research projects, hosted in my [simonw/research repo](https://github.com/simonw/research) and described in [Code research projects with async coding agents like Claude Code and Codex](https://simonwillison.net/2025/Nov/6/async-code-research/).

That’s five different custom integrations to pull in all of that data. The good news is that this kind of integration project is the kind of thing that coding agents *really* excel at. I knocked most of the feature out in a single morning while working in parallel on various other things.

I didn’t have a useful structured feed of my Research projects, and it didn’t matter because I gave Claude Code a link to [the raw Markdown README](https://raw.githubusercontent.com/simonw/research/refs/heads/main/README.md) that lists them all and it [spun up a parser regex](https://github.com/simonw/simonwillisonblog/blob/f883b92be23892d082de39dbada571e406f5cfbf/blog/importers.py#L77-L80). Since I’m responsible for both the source and the destination I’m fine with a brittle solution that would be too risky against a source that I don’t control myself.

Claude also handled all of the potentially tedious UI integration work with my site, making sure the new content worked on all of my different page types and was handled correctly by my [faceted search engine](https://simonwillison.net/2017/Oct/5/django-postgresql-faceted-search/).

#### Prototyping with Claude Artifacts [#](/2026/Feb/20/beats/#prototyping-with-claude-artifacts)

I actually prototyped the initial concept for beats in regular Claude—not Claude Code—taking advantage of the fact that it can clone public repos from GitHub these days. I started with:

> `Clone simonw/simonwillisonblog and tell me about the models and views`

And then later in the brainstorming session said:

> `use the templates and CSS in this repo to create a new artifact with all HTML and CSS inline that shows me my homepage with some of those inline content types mixed in`

After some iteration we got to [this artifact mockup](https://gisthost.github.io/?c3f443cc4451cf8ce03a2715a43581a4/preview.html), which was enough to convince me that the concept had legs and was worth handing over to full [Claude Code for web](https://code.claude.com/docs/en/claude-code-on-the-web) to implement.

If you want to see how the rest of the build played out the most interesting PRs are [Beats #592](https://github.com/simonw/simonwillisonblog/pull/592) which implemented the core feature and [Add Museums Beat importer #595](https://github.com/simonw/simonwillisonblog/pull/595/changes) which added the Museums content type.

Posted [20th February 2026](/2026/Feb/20/) at 11:47 pm · Follow me on [Mastodon](https://fedi.simonwillison.net/%40simon), [Bluesky](https://bsky.app/profile/simonwillison.net), [Twitter](https://twitter.com/simonw) or [subscribe to my newsletter](https://simonwillison.net/about/#subscribe)

## More recent articles

* [Two new Showboat tools: Chartroom and datasette-showboat](/2026/Feb/17/chartroom-and-datasette-showboat/) - 17th February 2026
* [Deep Blue](/2026/Feb/15/deep-blue/) - 15th February 2026

This is **Adding TILs, releases, museums, tools and research to my blog** by Simon Willison, posted on [20th February 2026](/2026/Feb/20/).

Part of series **[How I blog](/series/blogging/)**

9. [A homepage redesign for my blog's 22nd birthday](/2024/Jun/12/homepage-redesign/) - June 12, 2024, 7:59 p.m.
10. [My approach to running a link blog](/2024/Dec/22/link-blog/) - Dec. 22, 2024, 6:37 p.m.
11. [How I automate my Substack newsletter with content from my blog](/2025/Nov/19/how-i-automate-my-substack-newsletter/) - Nov. 19, 2025, 10 p.m.
12. **Adding TILs, releases, museums, tools and research to my blog** - Feb. 20, 2026, 11:47 p.m.

[blogging
118](/tags/blogging/)
[museums
27](/tags/museums/)
[ai
1862](/tags/ai/)
[til
25](/tags/til/)
[generative-ai
1649](/tags/generative-ai/)
[llms
1614](/tags/llms/)
[ai-assisted-programming
340](/tags/ai-assisted-programming/)
[claude-artifacts
36](/tags/claude-artifacts/)
[claude-code
95](/tags/claude-code/)

**Previous:** [Two new Showboat tools: Chartroom and datasette-showboat](/2026/Feb/17/chartroom-and-datasette-showboat/)

### Monthly briefing

Sponsor me for **$10/month** and get a curated email digest of the month's most important LLM developments.

Pay me to send you less!

[Sponsor & subscribe](https://github.com/sponsors/simonw/)

* [Disclosures](/about/#disclosures)
* [Colophon](/about/#about-site)
* ©
* [2002](/2002/)
* [2003](/2003/)
* [2004](/2004/)
* [2005](/2005/)
* [2006](/2006/)
* [2007](/2007/)
* [2008](/2008/)
* [2009](/2009/)
* [2010](/2010/)
* [2011](/2011/)
* [2012](/2012/)
* [2013](/2013/)
* [2014](/2014/)
* [2015](/2015/)
* [2016](/2016/)
* [2017](/2017/)
* [2018](/2018/)
* [2019](/2019/)
* [2020](/2020/)
* [2021](/2021/)
* [2022](/2022/)
* [2023](/2023/)
* [2024](/2024/)
* [2025](/2025/)
* [2026](/2026/)