from pathlib import Path
import re
from markitdown import MarkItDown

HTML_DIR = Path("rendered_html")
MD_DIR = Path("rendered_markdown")
SUMMARY_PATH = Path("article_paragraph_summaries.md")

NOISE_PATTERNS = [
    "privacy",
    "legal",
    "copyright",
    "©",
    "youtube",
    "github",
    "meetup",
    "sponsor me",
    "stripe.com",
    "docs]",
    "blog]",
]


def summarize(text: str, max_words: int = 30) -> str:
    words = text.split()
    if len(words) <= max_words:
        return text
    return " ".join(words[:max_words]) + "..."


def is_noise(block: str) -> bool:
    lowered = block.lower()
    if any(pattern in lowered for pattern in NOISE_PATTERNS):
        return True
    if lowered.count("http") >= 2:
        return True
    if lowered.count("](") >= 2:
        return True
    return False


def paragraph_candidates(markdown_text: str):
    blocks = re.split(r"\n\s*\n", markdown_text)
    candidates = []
    for block in blocks:
        lines = [ln.strip() for ln in block.splitlines() if ln.strip()]
        if not lines:
            continue

        joined = " ".join(lines).strip()
        if joined.startswith(("#", "-", "*", ">", "```", "|", "![")):
            continue
        if re.match(r"^\[[^\]]+\]\([^\)]+\)$", joined):
            continue
        if len(joined) < 100:
            continue
        if is_noise(joined):
            continue
        candidates.append(joined)
    return candidates


def main():
    if not HTML_DIR.exists():
        raise SystemExit(f"Missing HTML directory: {HTML_DIR}")

    MD_DIR.mkdir(parents=True, exist_ok=True)
    md = MarkItDown()
    html_files = sorted(HTML_DIR.glob("*.html"))
    rows = []

    for html_file in html_files:
        result = md.convert(str(html_file))
        markdown_text = result.text_content
        md_file = MD_DIR / f"{html_file.stem}.md"
        md_file.write_text(markdown_text, encoding="utf-8")

        paragraphs = paragraph_candidates(markdown_text)
        first_par = paragraphs[0] if paragraphs else "N/A"
        last_par = paragraphs[-1] if paragraphs else "N/A"

        rows.append(
            {
                "slug": html_file.stem,
                "markdown_path": str(md_file),
                "first_summary": summarize(first_par),
                "last_summary": summarize(last_par),
            }
        )

    lines = [
        "# Article paragraph summary validation",
        "",
        "This file confirms MarkItDown conversion succeeded by summarizing the first and last substantial article paragraph found in each converted markdown file.",
        "",
    ]

    for row in rows:
        lines.extend(
            [
                f"## {row['slug']}",
                f"- markdown file: `{row['markdown_path']}`",
                f"- first paragraph summary: {row['first_summary']}",
                f"- last paragraph summary: {row['last_summary']}",
                "",
            ]
        )

    SUMMARY_PATH.write_text("\n".join(lines), encoding="utf-8")


if __name__ == "__main__":
    main()
