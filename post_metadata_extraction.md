# RWAP-style metadata extraction from the four posts

The fields below are mapped from your sketch (company, project type, agent stack, integrations, storage, MCP/CLI/evals, etc.) and populated only with details explicitly stated in the converted post content.

## 1) Ramp — Ramp Research (Builders blog)

- **Company**: Ramp
- **Company type (private/public)**: Private (not explicitly stated in post; inferred from company context)
- **Project**: Ramp Research
- **Project type**: Internal
- **Primary goal**: Internal data analyst agent in Slack to reduce analytics bottlenecks
- **Model / model version**: Not specified
- **OSS / closed**: Closed/internal
- **Agent framework / harness / SDK**: In-house agent, Python mini-framework for evaluations in dbt project
- **Runtime**: Internal service (runtime details not explicitly specified)
- **Sandbox / isolation**: Not explicitly described
- **Tools**: SQL/data inspection tools; table/column inspection, branching/backtracking in agent flow
- **Integrations**: Slack, dbt, Looker, Snowflake, Redash
- **Storage / data systems**: Analytics warehouse with thousands of tables/views; file-system docs for domain context
- **Memory / state**: Multi-turn thread state in Slack
- **Database / data warehouse**: Snowflake warehouse
- **Search**: Metadata indexing and retrieval over dbt/Looker/Snowflake context
- **MCP**: Not mentioned
- **CLI**: Not mentioned
- **Logging / trajectory**: Intermediate-step assertions in evals (tool calls, table refs, query shape)
- **Evals**: End-to-end tests + Python mini-framework asserting final and intermediate steps
- **Durable workflows**: Not explicitly stated
- **Temporal**: Not mentioned
- **Browser use**: Not mentioned
- **Computer use**: Not mentioned

## 2) Ramp — Inspect (Modal customer story)

- **Company**: Ramp
- **Company type (private/public)**: Private (not explicitly stated in post; inferred from company context)
- **Project**: Ramp Inspect (background coding agent)
- **Project type**: Internal
- **Primary goal**: Give all builders cloud-parallel coding capacity with full internal context
- **Model / model version**: Not specified
- **OSS / closed**: Closed/internal
- **Agent framework / harness / SDK**: OpenCode agent running in sandbox
- **Runtime**: Modal Sandboxes + Modal Functions/Dicts/Queues/Cron + snapshot-based startup
- **Sandbox / isolation**: Per-session full-stack sandbox environments
- **Tools**: VS Code server, web terminal, VNC stack with Chromium, screenshot and browser navigation checks
- **Integrations**: GitHub, Slack, Buildkite, Sentry, Datadog, LaunchDarkly, Temporal
- **Storage / data systems**: Filesystem snapshots, Dicts for lock/image metadata, Queues for prompt routing
- **Memory / state**: Shared session coordination through Dicts + synchronized multi-client session state
- **Database / data warehouse**: Postgres, Redis, RabbitMQ (plus Temporal service)
- **Search**: Not explicitly stated
- **MCP**: Not mentioned
- **CLI**: Not called out directly (clients called out: Slack, web, Chrome extension)
- **Logging / trajectory**: Observability stack via Sentry/Datadog
- **Evals**: Visual verification (before/after screenshots, real browser checks)
- **Durable workflows**: Cron + queue/dict coordination patterns
- **Temporal**: Explicitly used in stack
- **Browser use**: Yes (Chromium in VNC stack)
- **Computer use**: Effectively yes (full remote dev environment, interactive tooling)

## 3) Simon Willison — “beats” feature post

- **Company**: Simon Willison (individual project)
- **Company type (private/public)**: Public/open personal project
- **Project**: “beats” feature for simonwillison.net
- **Project type**: Public personal product/project
- **Primary goal**: Aggregate external activity types (releases, TILs, museums, tools, research) into blog timelines
- **Model / model version**: Claude / Claude Code mentioned, no model version specified
- **OSS / closed**: Open/public repos and integrations
- **Agent framework / harness / SDK**: Claude Code + Claude artifacts workflow
- **Runtime**: Existing blog app/runtime (Django stack implied by references)
- **Sandbox / isolation**: Not explicitly described
- **Tools**: GitHub Actions, Datasette-based query/import, parser regex for README extraction
- **Integrations**: GitHub releases JSON, TIL site feed/query, museums JSON feed, tools site, research repo README
- **Storage / data systems**: JSON feeds/files + blog database/content system
- **Memory / state**: Not discussed
- **Database / data warehouse**: Not explicit; Datasette/query layer referenced
- **Search**: Faceted search integration mentioned
- **MCP**: Not mentioned
- **CLI**: Claude Code for web mentioned (workflow/tooling level)
- **Logging / trajectory**: Not discussed
- **Evals**: Not formalized; rapid iterative prototyping + PR workflow
- **Durable workflows**: GitHub Actions-based feed generation
- **Temporal**: Not mentioned
- **Browser use**: Not an agent-browser stack discussion
- **Computer use**: Not discussed

## 4) Stripe — Minions

- **Company**: Stripe
- **Company type (private/public)**: Private
- **Project**: Minions (one-shot unattended coding agents)
- **Project type**: Internal
- **Primary goal**: End-to-end unattended code changes from request to CI-ready PR
- **Model / model version**: Not specified
- **OSS / closed**: Closed/internal system, with fork of OSS Goose
- **Agent framework / harness / SDK**: Custom harness over forked Block Goose + deterministic orchestration steps
- **Runtime**: Isolated pre-warmed devboxes with Stripe repos/services preloaded
- **Sandbox / isolation**: Devboxes isolated from internet and production resources
- **Tools**: Git ops, linters, local executable checks, CI, Sourcegraph search, docs/tickets/build status tools
- **Integrations**: Slack, CLI, web UI, internal docs platform, feature-flag platform, ticketing UI, CI
- **Storage / data systems**: Standard Stripe internal systems (not itemized as DB/storage products in this post)
- **Memory / state**: Context hydration from thread links + MCP tools
- **Database / data warehouse**: Not discussed
- **Search**: Sourcegraph search explicitly used
- **MCP**: Yes; internal central MCP server “Toolshed” with 400+ tools
- **CLI**: Yes (explicitly listed as entry point)
- **Logging / trajectory**: Web UI shows decisions/actions of minion runs
- **Evals**: Multi-layer tests (local lint heuristics + selective CI + autofix loops, max two CI rounds)
- **Durable workflows**: Branch/push/CI/PR pipeline orchestration
- **Temporal**: Not mentioned
- **Browser use**: Not highlighted
- **Computer use**: Yes in spirit (unattended coding agent acting in dev environments)

