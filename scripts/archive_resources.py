#!/usr/bin/env python3
from __future__ import annotations

import hashlib
import os
import re
import subprocess
import sys
from pathlib import Path
from urllib.parse import urlparse

from markitdown import MarkItDown
from playwright.sync_api import sync_playwright

URL_PATTERN = re.compile(r"https?://[^\s)>\"]+")


def run_git(*args: str) -> str:
    result = subprocess.run(
        ["git", *args],
        check=True,
        capture_output=True,
        text=True,
    )
    return result.stdout


def extract_urls_from_text(text: str) -> set[str]:
    return set(match.group(0).rstrip(".,") for match in URL_PATTERN.finditer(text))


def diff_added_lines(base_sha: str, head_sha: str) -> str:
    if base_sha == "0" * 40:
        return ""
    return run_git("diff", "--unified=0", base_sha, head_sha, "--", "RESOURCES.md")


def urls_from_diff(base_sha: str, head_sha: str) -> set[str]:
    diff_text = diff_added_lines(base_sha, head_sha)
    if not diff_text:
        return set()
    added_lines = []
    for line in diff_text.splitlines():
        if line.startswith("+++") or line.startswith("---"):
            continue
        if line.startswith("+"):
            added_lines.append(line[1:])
    return extract_urls_from_text("\n".join(added_lines))


def urls_from_file() -> set[str]:
    content = Path("RESOURCES.md").read_text(encoding="utf-8")
    return extract_urls_from_text(content)


def slugify_url(url: str) -> str:
    parsed = urlparse(url)
    base = f"{parsed.netloc}{parsed.path}"
    base = base.strip("/") or "index"
    base = re.sub(r"[^a-zA-Z0-9]+", "-", base).strip("-")
    suffix = hashlib.sha256(url.encode("utf-8")).hexdigest()[:8]
    return f"{base}-{suffix}"


def save_html_and_markdown(urls: set[str], output_dir: Path) -> None:
    output_dir.mkdir(parents=True, exist_ok=True)
    converter = MarkItDown()
    with sync_playwright() as playwright:
        browser = playwright.chromium.launch()
        context = browser.new_context()
        page = context.new_page()
        for url in sorted(urls):
            slug = slugify_url(url)
            html_path = output_dir / f"{slug}.html"
            md_path = output_dir / f"{slug}.md"
            page.goto(url, wait_until="networkidle", timeout=60000)
            html_path.write_text(page.content(), encoding="utf-8")
            markdown = converter.convert(str(html_path)).text_content
            md_path.write_text(markdown, encoding="utf-8")
        context.close()
        browser.close()


def main() -> int:
    if len(sys.argv) < 3:
        print("Usage: archive_resources.py <base_sha> <head_sha>", file=sys.stderr)
        return 2
    base_sha, head_sha = sys.argv[1], sys.argv[2]
    if base_sha == "0" * 40:
        urls = urls_from_file()
    else:
        urls = urls_from_diff(base_sha, head_sha)
    if not urls:
        print("No new URLs detected.")
        return 0
    output_dir = Path("resources-archive")
    save_html_and_markdown(urls, output_dir)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
