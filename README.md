# humanizecompact

**humanizecompact** is a Go library that converts numeric strings into more human-friendly representations (e.g., `1.2K`, `9,99 만`, `1 bilhão`) based on [CLDR](https://cldr.unicode.org/) data and locale rules.

## Features

- **Locale-aware**: Relies on per-locale data to determine how to abbreviate numbers (thousand, million,万,亿,만,억, etc.) and which plural forms to use.
- **Long or Short**: Offers long-form strings (`1 thousand`) or short-form strings (`1K`), configurable via `OptionLong` or `OptionShort`.
- **Fallback mechanism**: When a number cannot be humanized (e.g., it’s not an integer or out of range), the user-supplied fallback function is called.
- **Easy integration**: Simply implement the `Locale` interface and provide `CldrData` for custom languages or variants.