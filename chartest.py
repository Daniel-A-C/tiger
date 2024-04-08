
def print_unicode_range(start, end):
    """Prints Unicode characters within a specified range."""
    for codepoint in range(start, end + 1):
        if codepoint % 10 == 0:
            print()
        try:
            print(chr(codepoint), end="")
        except UnicodeEncodeError:
            print("[Not Representable]", end="")

print_unicode_range(49, 255)   # Basic Latin and some symbols
print_unicode_range(9984, 10175)  # Miscellaneous Symbols and Dingbats

