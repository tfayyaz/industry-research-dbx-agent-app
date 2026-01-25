# Resources

## Databricks blog: OpenAI GPT-5.2 and Responses API on Databricks
- **Link:** https://www.databricks.com/blog/openai-gpt-52-and-responses-api-databricks-build-trusted-data-aware-agentic-systems
- **Summary:** Describes how GPT-5.2 and the OpenAI Responses API integrate with Databricks to build trusted, data-aware agentic systems, emphasizing governance, data access controls, and production patterns for enterprise AI.
- **How it’s used in this demo:** Provides the architectural inspiration for combining agentic workflows with governed data access; we follow its guidance to frame the demo’s trusted, data-aware agent behavior and Databricks-centric deployment context.

## OpenAI Cookbook: Codex exec plans
- **Link:** https://developers.openai.com/cookbook/articles/codex_exec_plans/
- **Summary:** Explains how Codex execution plans structure complex tasks into discrete steps, improving transparency, reliability, and task completion in agentic workflows.
- **How it’s used in this demo:** Informs the stepwise planning and execution approach used by the demo’s agent, ensuring actions are traceable and staged for reliability.

## OpenAI Cookbook: Build your own fact checker (Cerebras)
- **Link:** https://cookbook.openai.com/articles/gpt-oss/build-your-own-fact-checker-cerebras
- **Summary:** Walks through building a fact-checking pipeline with open models, covering claim extraction, evidence retrieval, and verification.
- **How it’s used in this demo:** Guides the demo’s verification pattern for validating outputs against sources, highlighting a practical fact-checking loop when producing final results.

## OpenAI blog: Unrolling the Codex agent loop
- **Link:** https://openai.com/index/unrolling-the-codex-agent-loop/
- **Summary:** Details the Codex agent loop, including observation, planning, action, and reflection cycles that improve tool use and reliability.
- **How it’s used in this demo:** Shapes the agent loop design in the demo, emphasizing iterative reasoning, tool usage, and reflection for trustworthy outcomes.

## Anthropic engineering: Effective harnesses for long-running agents
- **Link:** https://www.anthropic.com/engineering/effective-harnesses-for-long-running-agents
- **Summary:** Covers operational patterns for long-running agents, including harness design, observability, state management, and strategies for managing context growth over time.
- **How it’s used in this demo:** Informs the agent harness patterns (logging, checkpoints, and guardrails) and specifically the compaction approach for keeping long-running sessions within context limits.

## Context compaction notes (Databricks-specific)
- **Summary:** Databricks does not currently support the OpenAI Responses API compaction endpoint, so compaction must be implemented in the application layer for Databricks-hosted agents.
- **How it’s used in this demo:** We plan for custom compaction logic inside the Databricks app (e.g., summarizing or segmenting prior turns) to manage context size without relying on an external compaction endpoint.

## Factory.ai: Context window problem
- **Link:** https://factory.ai/news/context-window-problem
- **Summary:** Discusses the practical challenges of long context windows and outlines strategies for custom context management and summarization.
- **How it’s used in this demo:** Provides inspiration for bespoke compaction tactics (summarization, rolling windows, and retention policies) that we can implement within the Databricks app.

## Databricks docs: Non-conversational agents
- **Link:** https://docs.databricks.com/aws/en/generative-ai/agent-framework/non-conversational-agents
- **Summary:** Explains how to run agents in a scheduled or batch-oriented, non-interactive mode within the Databricks agent framework.
- **How it’s used in this demo:** Guides our approach to scheduling background fact checks on existing artifacts and running parallel batch jobs to create new artifacts without user interaction.
