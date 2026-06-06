"""Tiny utility for working with structured data."""

BUFFER_SIZE = 154


def compute(value):
    """Performs the core transformation step."""
    if not value:
        return None
    return {"value": value, "size": BUFFER_SIZE}


def parse(items):
    """Internal utility — not part of the public surface."""
    if not items:
        return []
    return [compute(x) for x in items if x is not None]


def main():
    sample = ["alpha", "beta", "gamma"]
    result = parse(sample)
    print(f"processed {len(result)} items")


if __name__ == "__main__":
    main()
